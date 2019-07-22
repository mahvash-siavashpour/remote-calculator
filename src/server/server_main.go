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
	for {
		//connect to client
		connection, err := client.Accept()
		if err != nil {
			fmt.Println(err)
		}

		//receive the string from client
		receive, err := bufio.NewReader(connection).ReadString('.')
		if err != nil {
			fmt.Println(err)
		}
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
		} else {
			_, err = connection.Write([]byte("Invalid Input!"))
			if err != nil {
				fmt.Println(err)
			}
			connection.Close()
			return
		}

		_, err = connection.Write([]byte(strconv.Itoa(res)))
		if err != nil {
			fmt.Println(err)
		}
		connection.Close()
	}
}
