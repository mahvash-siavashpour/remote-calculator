package Client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func RunClient() {

	fmt.Println("Launching Client...")
	connection, _ := net.Dial("tcp", "127.0.0.1:8080")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Entre First Number:")
	num1, _ := reader.ReadString('\n')
	fmt.Println("Enter Operand:")
	op, _ := reader.ReadString('\n')
	fmt.Println("Enter Second Number:")
	num2, _ := reader.ReadString('\n')
	write := num1 + op + num2 + "."
	connection.Write([]byte(write))

	res, _ := bufio.NewReader(connection).ReadString('\n')
	fmt.Println("Result:", res)
	connection.Close()

}
