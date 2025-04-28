rpc 节点地址：https://mainnet.infura.io/v3/5e90e80eb50f4d888db6614644bdc875

本地Ganache http://127.0.0.1:7545

cd .\DAY1_to_DAY4_ETH_CLI\
go mod init ethcli  # 初始化Go module，项目名是ethcli，告诉Go，我要建立一个新的项目module,通俗点：告诉Go，我这里有一个叫ethcli的新项目，请帮我管理依赖。

go get github.com/ethereum/go-ethereum # 使用go get拉去go-ethereum 模块代码

go mod tidy  # 自动补齐所有依赖


context.Background() : 创建一个空的根上下文（最顶层的Context），一般用在main函数、初始化时
context.WithTimeout(parent, timeout) : 基于传入的Context，创建一个带超时控制的子Context。如果超过设定时间（5秒），自动取消这个context
ctx : 新生成的带超时控制的Context，后续网络请求、数据库操作等都基于它
cancel : 返回一个取消函数（cancel()），必须调用，负责释放内部资源，防止内存泄漏
defer cancel() : Go惯例：一旦main函数退出或者提前返回，不管有没有超时，都会调用cancel释放资源


调用： go run main.go http://127.0.0.1:7545