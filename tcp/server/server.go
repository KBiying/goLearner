package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func process(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("server process: %T\n", conn)

	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("receive data from client: ", recvStr)
		inputReader := bufio.NewReader(os.Stdin)
		s, _ := inputReader.ReadString('\n') //?
		t := strings.Trim(s, "\r\n")         //?
		conn.Write([]byte(t))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8877")
	fmt.Printf("server: %T======\n", listen)
	if err != nil {
		fmt.Println("listen failed, err", err)
		return

	}
	for {
		conn, err := listen.Accept()
		fmt.Println("当前建立tcp连接")
		if err != nil {
			fmt.Println("accpect failed, err: ", err)
			continue
		}
		go process(conn)
	}

}
