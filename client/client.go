package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Digite uma mensagem ('sair' para sair): ")
		message, _ := reader.ReadString('\n')

		// Envia a mensagem para o servidor
		fmt.Fprint(conn, message)

		if message == "sair\n" {
			break
		}
	}
}
