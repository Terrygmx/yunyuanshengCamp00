# 构建本地镜像
cd week2
docker build  -t go-http2 .

#上传到docker仓库
docker tag go-http2 terrygmx/go-httpserv:latest
docker push terrygmx/go-httpserv:latest 
# 通过 Docker 命令本地启动 httpserver
docker run terrygmx/go-httpserv


# 查询容器IP配置
docker inspect 4c572cdfe00d
#通过 nsenter 进入容器查看 IP 配置
docker inspect 4c572cdfe00d | grep Pid #找到pid
nsenter -t 30404 -n /bin/sh #进入容器
ifconfig # 查到ip为172.17.0.2

# 测试http服务
curl -v 172.17.0.2:8080/
curl -v 172.17.0.2:8080/healthz

