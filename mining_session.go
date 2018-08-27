package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	CONN_ADDR = "btc.viabtc.com:443"
	CONN_TYPE = "tcp"
	MSG1      = `{"id": 1, "method": "mining.subscribe", "params": []}`
	MSG2      = `{"params": ["lxjhk.01", "password"], "id": 2, "method": "mining.authorize"}`
)

type miningSession struct {
	workerName         string
	tcpConn            *net.Conn
	incomingDifficulty int
	extranonce1        string //Hex-encoded, per-connection unique string which will be used for coinbase serialization later. Keep it safe!
	extranonce2Size    int
	//jobQueue
}

func (self *miningSession) start() {
	conn, _ := net.Dial(CONN_TYPE, CONN_ADDR)
	err := conn.(*net.TCPConn).SetKeepAlive(true)
	if err != nil {
		fmt.Println(err)
		return
	}

	self.tcpConn = &conn

	go listenAndPrint(&conn, self)
	time.Sleep(time.Second * 2)
	sendText(&conn, MSG1)
	sendText(&conn, MSG2)

	time.Sleep(time.Second * 2000)
}

func (self *miningSession) processMsg(m string) {
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(m), &dat); err != nil {
		panic(err)
	}

	// For initialization only
	if v, ok := dat["id"]; ok && v != nil {
		switch int(v.(float64)) {
		case 1:
			d := dat["result"].([]interface{})
			self.extranonce1 = d[1].(string)
			self.extranonce2Size = int(d[2].(float64))
			log.Printf("extranonce1 set to %s and extranonce2Size set to %d", self.extranonce1, self.extranonce2Size)
			return
		case 2:
			if dat["result"] == true {
				log.Println("Successfully authorised!")
			} else {
				panic("Authorisation with mining pool failed")
			}
			return
		}
	}

	// Switch for message handlers
	if _, ok := dat["method"]; ok {
		switch dat["method"].(string) {
		case "mining.set_difficulty":
			newDiff := int(dat["params"].([]interface{})[0].(float64))
			self.incomingDifficulty = newDiff
			log.Printf("Current Difficulty is set to %d", self.incomingDifficulty)
		case "mining.notify":
			self.processMiningNotify(dat)
		default:
			fmt.Fprintf(os.Stderr, "No such method %s is implemented!\n", dat["method"])
		}
	} else {
		fmt.Fprintf(os.Stderr, "method is not a key in the json payload! %s\n", dat)
	}

}
