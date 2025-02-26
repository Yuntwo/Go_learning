module blockchain

go 1.23.4

require (
	github.com/davecgh/go-spew v1.1.1 // 在 console 中直接查看 struct 和 slice 这两种数据结构
	github.com/gorilla/mux v1.8.1 // 写 web handler
	github.com/joho/godotenv v1.5.1 // 读取项目根目录中的 .env 配置文件，避免 http port 之类的配置硬编码
)

