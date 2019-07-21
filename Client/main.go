package Client

import (
	"bufio"
	"net"
	"os"
)

func main() {

	connection, _ := net.Dial("tcp", "127.0.0.1:8080")

	for {
		reader := bufio.NewReader(os.Stdin)

	}
}
