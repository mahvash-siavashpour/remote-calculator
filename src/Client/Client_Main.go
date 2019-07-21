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
	fmt.Println("Enter Second Number:")
	num2, _ := reader.ReadString('\n')

	connection.Write([]byte(num1))
	connection.Write([]byte(num2))

	//res, err := bufio.NewReader(connection).ReadString('\n')
	//fmt.Println(err)
	//fmt.Println(res)
}