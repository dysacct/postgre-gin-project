-- 清理数据库中的空记录
-- 删除 zbx_id 为空的记录
DELETE FROM idc_info WHERE zbx_id = '' OR zbx_id IS NULL;
DELETE FROM machine_info WHERE zbx_id = '' OR zbx_id IS NULL;
DELETE FROM business_info WHERE zbx_id = '' OR zbx_id IS NULL;
DELETE FROM network_info WHERE zbx_id = '' OR zbx_id IS NULL;

-- 查看当前数据
SELECT 'idc_info' as table_name, count(*) as count FROM idc_info
UNION ALL
SELECT 'machine_info' as table_name, count(*) as count FROM machine_info
UNION ALL
SELECT 'business_info' as table_name, count(*) as count FROM business_info
UNION ALL
SELECT 'network_info' as table_name, count(*) as count FROM network_info;
