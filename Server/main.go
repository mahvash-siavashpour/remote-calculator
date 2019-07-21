package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

func main() {
	fmt.Println("Launching Server...")

	client, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	connection, err := client.Accept()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected")

	for {
		number1, err := bufio.NewReader(connection).ReadString('\n')
		fmt.Println(err)
		number2, err := bufio.NewReader(connection).ReadString('\n')
		fmt.Println(err)

		//fmt.Println(strconv.Atoi("1123"))
		n1, _ := strconv.Atoi(number1)
		n2, _ := strconv.Atoi(number2)
		sum := n1 + n2

		fmt.Println(sum)
		_, _ = connection.Write([]byte(strconv.Itoa(sum)))
	}
}
