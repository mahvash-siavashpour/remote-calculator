package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func RunClient() {

	fmt.Println("Launching Client...")
	connection, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}

	//get the operator and the operands from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Entre First Number:")
	num1, _ := reader.ReadString('\n')
	fmt.Println("Enter Operator:")
	op, _ := reader.ReadString('\n')
	fmt.Println("Enter Second Number:")
	num2, _ := reader.ReadString('\n')

	//prepare the string to be sent
	write := num1 + op + num2 + "."

	//send the string
	_, err = connection.Write([]byte(write))
	if err != nil {
		fmt.Println(err)
	}

	//receive the result
	res, err := bufio.NewReader(connection).ReadString('\n')

	fmt.Println("Result:", res)
	connection.Close()
}
