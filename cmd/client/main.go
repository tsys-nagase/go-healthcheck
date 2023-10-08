package main

import (
	// (一部抜粋)
	"bufio"
	"context"
	"errors"
	hellopb "example/hello/pkg/grpc"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	scanner *bufio.Scanner
	client  hellopb.GreetingServiceClient
)

func main() {
	fmt.Println("start gRPC Client.")

	// 1. 標準入力から文字列を受け取るスキャナを用意
	scanner = bufio.NewScanner(os.Stdin)

	// 2. gRPCサーバーとのコネクションを確立
	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	// 3. gRPCクライアントを生成
	client = hellopb.NewGreetingServiceClient(conn)

	for {
		fmt.Println("1: send Request")
		fmt.Println("2: HelloServerStream")
		fmt.Println("3: exit")
		fmt.Print("please enter >")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			Hello()
		case "2":
			HelloServerStream()
		case "3":
			fmt.Println("bye.")
			goto M
		}
	}
M:
}

func Hello() {
	fmt.Println("🌽Please enter your name.")
	scanner.Scan()
	name := scanner.Text()

	req := &hellopb.HelloRequest{
		Name: name,
	}
	res, err := client.Hello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetMessage())
	}
}

func HelloServerStream() {
	fmt.Println("😘Please enter your name.")
	scanner.Scan()
	name := scanner.Text()

	req := &hellopb.HelloRequest{
		Name: name,
	}
	stream, err := client.HelloServerStream(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		res, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("all the responses have already received.")
			break
		}

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}
