# 使用 golang:latest 作为基础镜像
FROM golang:latest

# 在容器内创建一个工作目录
WORKDIR /app

EXPOSE 8888

# 设置容器启动时的默认命令，即运行 Go 进程
CMD ["go", "run", "main.go"]
