package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func listenAndPrint(conn *net.Conn, ms *miningSession) {
	connbuf := bufio.NewReader(*conn)
	log.Println("start to listing on tcp ....")
	for {
		str, err := connbuf.ReadString('\n')
		if err != nil {
			(*conn).Close()
			log.Println("Connection Closed due to ", err)
			panic(err)
		}
		if len(str) > 0 {
			str = strings.TrimSuffix(str, "\n")
			log.Println("Message from server: " + str)
			ms.processMsg(str)
		}
	}
}

func sendText(conn *net.Conn, msg string) {
	log.Print("Text to send: ", msg, "\n")
	_, err := fmt.Fprintf(*conn, msg+"\n")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		panic(err)
	}
}
