package handlers

import (
	"encoding/json"
	"gin-postgre-project/database"
	"gin-postgre-project/models"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

const cacheTTL = 30 * time.Minute // 缓存过期时间30分钟

// 同一的redis缓存key格式
func cacheKey(zbxID string) string {
	return "cache:machine:" + zbxID
}

// GetMachine godoc
// @Summary      获取单个机器信息
// @Description  根据zbx_id获取机器的详细信息，包括IDC、机器、业务、网络信息
// @Tags         machines
// @Accept       json
// @Produce      json
// @Param        zbx_id path string true "机器唯一标识"
// @Success      200 {object} models.Response
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Security     ApiKeyAuth
// @Router       /machine/{zbx_id} [get]
func GetMachine(c *gin.Context) {
	zbxID := c.Param("zbx_id") // 从URL路径参数获取zbx_id
	if zbxID == "" {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Message: "zbx_id不能为空"})
		return
	}

	// Step 1 : 先查Redis缓存
	ctx := c.Request.Context() // Request是Gin框架的请求对象, Context()是获取请求的上下文
	cacheResult := database.RedisClient.Get(ctx, cacheKey(zbxID))

	if cacheResult.Err() == nil { // .Err() 是redis.Nil错误, 表示缓存不存在
		// 缓存命中
		var result map[string]interface{}
		json.Unmarshal([]byte(cacheResult.Val()), &result)
		c.JSON(http.StatusOK, models.Response{
			Code:    200,
			Message: "缓存命中",
			Data:    result,
		})
		return
	}

	// redis.Nil 是redis的错误, 表示缓存不存在
	if cacheResult.Err() != redis.Nil {
		slog.Warn("Redis 查询失败, 将直接查库", "err", cacheResult.Err()) // slog.Warn 是日志记录器, 用于记录警告信息
	}

	// Step 2 : 缓存 miss，查 PostgreSQL (联表查询)
	var idc models.IDCInfo
	err := database.DB.First(&idc, "zbx_id = ?", zbxID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Code:    404,
			Message: "IDC信息不存在",
		})
		return
	}
	var machine models.MachineInfo
	database.DB.First(&machine, "zbx_id = ?", zbxID)

	var business models.BusinessInfo
	database.DB.First(&business, "zbx_id = ?", zbxID)

	var networks []models.NetworkInfo
	database.DB.Find(&networks, "zbx_id = ?", zbxID)

	// 组装最终返回数据
	result := gin.H{
		"idc_info":      idc,
		"machine_info":  machine,
		"business_info": business,
		"network_info":  networks,
	}

	// Step 3 : 写入Redis缓存
	jsonData, _ := json.Marshal(result)
	database.RedisClient.Set(ctx, cacheKey(zbxID), jsonData, cacheTTL)
	c.JSON(http.StatusOK, models.Response{
		Code:    200,
		Message: "查询成功",
		Data:    result,
	})
	slog.Info("数据库查询并已经缓存", "zbx_id", zbxID)
}

// CreateMachine godoc
// @Summary      创建机器
// @Description  创建新的机器记录
// @Tags         machines
// @Accept       json
// @Produce      json
// @Param        body body models.IDCInfo true "机器信息"
// @Success      200 {object} models.Response
// @Failure      400 {object} models.Response
// @Failure      500 {object} models.Response
// @Security     ApiKeyAuth
// @Router       /machine [post]
func CreateMachine(c *gin.Context) {
	var idc models.IDCInfo
	// ShouldBindJSON 是Gin框架的函数, 用于将请求体中的JSON数据绑定到结构体中
	if err := c.ShouldBindJSON(&idc); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	// 插入主表
	// 失败情况: 1. 数据库插入失败 2. 主键冲突
	if err := database.DB.Create(&idc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:    500,
			Message: "创建失败",
		})
		return
	}

	// 清除缓存(防止别人之前查过但没这台机器)
	database.RedisClient.Del(c.Request.Context(), cacheKey(idc.ZbxID)) // ZbxID 是IDCInfo的唯一标识

	c.JSON(http.StatusOK, models.Response{
		Code:    200,
		Message: "创建成功",
		Data:    idc,
	})
}

// UpdateMachine godoc
// @Summary      更新机器信息
// @Description  根据zbx_id更新机器信息
// @Tags         machines
// @Accept       json
// @Produce      json
// @Param        zbx_id path string true "机器唯一标识"
// @Param        body body models.IDCInfo true "更新的机器信息"
// @Success      200 {object} models.Response
// @Failure      400 {object} models.Response
// @Security     ApiKeyAuth
// @Router       /machine/{zbx_id} [put]
func UpdateMachine(c *gin.Context) {
	zbxID := c.Param("zbx_id") // 从URL中获取zbx_id
	var idc models.IDCInfo
	// Error 是Gorm的错误, 表示查询失败
	if err := database.DB.First(&idc, "zbx_id = ?", zbxID).Error; err != nil {
		// http.StatusBadRequest 是HTTP状态码, 表示请求错误
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    400,
			Message: "机器不存在",
		})
		return
	}
	// Save 是Gorm的函数, 用于更新数据
	database.DB.Save(&idc)
	// 这一步是清除缓存, 防止别人之前查过但没这台机器
	database.RedisClient.Del(c.Request.Context(), cacheKey(zbxID))

	c.JSON(http.StatusOK, models.Response{
		Code:    200,
		Message: "更新成功",
		Data:    idc,
	})
}

// DeleteMachine godoc
// @Summary      删除机器
// @Description  根据zbx_id删除机器记录
// @Tags         machines
// @Accept       json
// @Produce      json
// @Param        zbx_id path string true "机器唯一标识"
// @Success      200 {object} models.Response
// @Failure      500 {object} models.Response
// @Security     ApiKeyAuth
// @Router       /machine/{zbx_id} [delete]
func DeleteMachine(c *gin.Context) {
	zbxID := c.Param("zbx_id")
	// http.StatusInternalServerError 是HTTP状态码, 表示服务器内部错误
	if err := database.DB.Delete(&models.IDCInfo{}, "zbx_id = ?", zbxID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:    500,
			Message: "删除失败",
		})
		return
	}

	database.RedisClient.Del(c.Request.Context(), cacheKey(zbxID))
	c.JSON(http.StatusOK, models.Response{
		Code:    200,
		Message: "删除成功",
	})
}
