
# gRPCサンプルコード

## 概要
gRPCの簡単なサンプルコードです。
以下のエンドポイントがそれぞれ実装されています。

- /hello
- /hello/{name}
- /hello/{name}/age/{age}
- /status

## 起動方法




### client 

docker compose up


- 開発用  
`air`  

- 起動  
`go run client/main.go`

### server

docker compose up

もしくは

- 開発用  
`air`  
- 起動  
`go run server/main.go`



