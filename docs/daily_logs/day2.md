实现了查询区块信息的小工具，使用命令：
go run main.go queryblock http://127.0.0.1:7545 latest

go run main.go queryblock https://mainnet.infura.io/v3/5e90e80eb50f4d888db6614644bdc875 latest

实现步骤为：
1.创建上下文对象
2.利用上下文对象创建eth节点连接对象
3.解析区块高度（原始是16进制的）
4.同时支持10进制的区块高度查询
5.针对infura的访问，使用了代理，本地的ganache直接访问

