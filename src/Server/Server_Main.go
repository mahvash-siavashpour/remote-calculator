package Server

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
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

	number1, err := bufio.NewReader(connection).ReadString('\n')
	//fmt.Println(err)
	number2, err := bufio.NewReader(connection).ReadString('\n')
	//fmt.Println(err)

	//omitting the "\n"
	a := []rune(number1)
	b := []rune(number2)
	//converting to int
	n1, _ := strconv.Atoi(string(a[0 : len(a)-1]))
	n2, _ := strconv.Atoi(string(b[0 : len(b)-1]))
	sum := n1 + n2

	fmt.Println(sum)
	_, _ = connection.Write([]byte(strconv.Itoa(sum)))

}
