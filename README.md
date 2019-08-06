# 远程工作

:newspaper: 抓取最新的远程工作机会

## Dokku 部署

1. `go.mod` 中定义 `go install` 路径
2. `main.go` 中定义 `+build tag1,tag2`
3. install 后的二进制存放在 `/app/bin`
4. `Procfile` 中定义执行的二进制文件
5. 执行文件夹在 `/app`
