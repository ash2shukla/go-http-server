package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"os"

	"github.com/ash2shukla/go-http-server/lib/core"
	"github.com/ash2shukla/go-http-server/lib/handlers"
)

func ProcessClient(connection net.Conn) {
	handlers := map[string]func(core.HttpBasicHeaderInfo) []byte{
		"GET": handlers.GetStaticHandler,
	}

	fmt.Printf(
		"Connected Client: %s\n",
		connection.RemoteAddr(),
	)

	defer connection.Close()

	reader := bufio.NewReader(connection)
	tp := textproto.NewReader(reader)
	reqInitLine, _ := tp.ReadLine()
	basicInfo, err := core.GetHttpBasicInfo(reqInitLine)

	if err != nil {
		log.Panicln(err)
	}

	response := handlers[basicInfo.Method](basicInfo)
	connection.Write([]byte(response))

	fmt.Printf("Closed client : %s\n", connection.RemoteAddr())
}

func ServerLoop(host string, port string) {
	hostPortStr := fmt.Sprintf("%s:%s", host, port)

	server, err := net.Listen("tcp", hostPortStr)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer server.Close()

	fmt.Printf("Serving at Host: %s Port: %s\n", host, port)
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go ProcessClient(connection)
	}
}
