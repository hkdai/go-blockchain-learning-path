### 写一个查询地址最近交易的函数
以太坊本身不支持“按地址查交易”解决方法：用区块范围扫描+遍历交易，从中筛选涉及目标地址的交易

1.获取各个参数，包括rpcURL,targetAddress
2.处理以太坊地址的时候要把string转换成以太坊的地址对象，使用到targetAddress := common.HexToAddress(os.Args[3])
3.连接以太坊网络，使用defer关闭资源
4.获取最新区块高度，然后递减500个（可配置），把每一个里面的from和to的地址拿出来与targetAddress比对，符合条件的打印出来
5.区块高度转换成int64类型
6.使用适当的签名算法从交易中恢复发送者地址 from, err := types.Sender(types.LatestSignerForChainID(chainID), tx)
7.交易接收者直接使用to = *tx.To()
8.做好异常处理，有链ID可能为0或空


注意事项：
1.Args[0] 是main程序所在路径， 所以执行命令后面的参数一般都从1开始

调用命令：
go run main.go querytx https://mainnet.infura.io/v3/5e90e80eb50f4d888db6614644bdc875 0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5
go run main.go querytx http://127.0.0.1:7545 0xEbB1795B2F5d8dDF227096a50955D3D8e7C20Aee


### 写一个发起ETH转账的函数
1.获取各个参数 rpcURL，privateKey，to，amount
2.用私钥字符串调用crypto.HexToECDSA(privKey)创建私钥对象
3.从私钥对象中获取公钥，并调用方法转换成from地址
4.nonce
5.定义gasPrice
6.调用 tx := types.NewTransaction(nonce,to,amount,21000,gasPrice,nil) 创建转账
7.调用 client.NetworkID(context.Background()) 获取chainID
8. 使用私钥和chainID签名，获取signedTx
9. 使用signedTx发送交易
10.集成到 main路由里去，使用sendtx作为路由标识



go run main.go sendtx http://127.0.0.1:7545 0x61008a160d039111f9e4a66ff23af74fcd4fcdfdbc55112b4dee1178794e9322 0x1920C713B37f026b2A03755aa2113e26dBD9e413 1000000000000000000

遇到的问题：
1.私钥没有处理，不能识别带0x的，手动处理掉。
2. Invalid signature v value  通常表示在使用给定的私钥签名交易时出现了问题，这可能与区块链网络的chainID设置有关。对于本地Ganache，我们需要专门处理其特殊的签名要求。
不再从网络获取chainID，而是直接使用Ganache的默认chainID (1337)
