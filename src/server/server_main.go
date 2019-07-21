package server

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func RunServer() {
	fmt.Println("Launching Server...")

	client, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	connection, err := client.Accept()
	if err != nil {
		fmt.Println(err)
	}

	receive, err := bufio.NewReader(connection).ReadString('.')
	st := strings.Split(receive, "\n")

	n1, _ := strconv.Atoi(st[0])
	op := st[1]
	n2, _ := strconv.Atoi(st[2])

	var res int
	if op == "+" {
		res = n1 + n2
	} else if op == "-" {
		res = n1 - n2
	} else if op == "*" {
		res = n1 * n2
	} else if op == "/" {
		res = n1 / n2
	}

	_, _ = connection.Write([]byte(strconv.Itoa(res)))
	connection.Close()
}
