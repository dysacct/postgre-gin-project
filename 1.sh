#!/bin/bash
# 1. 先登录拿到 token（admin）
TOKEN=$(curl -s -X POST http://localhost:8080/api/login -d '{"username":"admin","password":"abcd001002"}' -H "Content-Type: application/json" | jq -r .data.token)

# 创建机器
# curl -s -X POST http://localhost:8080/api/machine \
#   -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" \
#   -d '{"zbx_id":"ipmi-bj-999","idc_code":"BJ01","idc_name":"北京亦庄","ipmi_ip":"10.0.1.999","ssh_ip":"192.168.1.999"}' | jq

# 查询
# curl -s -H "Authorization: Bearer $TOKEN" http://localhost:8080/api/machine/ipmi-bj-999 | jq

curl -s -H "Authorization: Bearer $TOKEN" "http://localhost:8080/api/machines?page=1&size=10&search=北京" | jq