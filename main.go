package main

import (
	"fmt"
	"encoding/json"
	"os"
	"net"
	"bufio"
	"log"
	"time"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

type mining_notify struct {
	id   int      `json:"id"`
	method []string `json:"fruits"`
	params []interface{} `json:"fruits"`
}


const (
	CONN_ADDR = "btc.viabtc.com:443"
	CONN_TYPE = "tcp"
	MSG = `{"id": null, "method": "mining.notify", "params": ["74f3", "3fe9a4faffbc78b7965f10df55a8b3c59d0456310017deb20000000000000000", "01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff5a03bd3608172f5669614254432f4d696e6564206279206c786a686b2f2cfabe6d6dd4807c2beb9813393c72b1714afe755866d531bc6271c98b461838e9a8b85dfa040000000000000010f374370862685029", "ffffffff0228a2f84a000000001976a914536ffa992491508dca0354e52f32a3a7a679a53a88ac0000000000000000266a24aa21a9ed8ecc629135396d0104bdfab4847467bffa37cf268d295708c455c6c38409897200000000", ["b3cd8df76fece8022d57246897be0956ee0b539422f4e8bf52ca91af5eb72e28", "e2277b92b9ff0ae5db8fdf0f28b362aca772a4c13133d9693038d77801826cf7", "21f231463db335c45043b3519cc16b7ea67c356d9f9a87ac67b7db6607ccfcd1", "7997e50226d297379c1e46a8bffec84f30ba37f70126f9219d514b6e355f3961", "c05efb3d4a74b035fd52e75e7a4eb3d39002d900b49612545b9f76634a3fc462", "499c44d38767df2f79140ecdcd093290a67ebc21fc0b7070393681be2cdd0a2d", "4449e9e22f20defb2fbcc67c946c722c435566b80e449df93a9db280a26a21c7", "8eb9071c0fb51a14899bb2dc3079266f46d72d505fa3849a16c33e7ba94c93ff", "0eeca878d8ca23d68fce827535b9ad78ece3ea7d71d311e34138289c30f29e0e", "fe050724335beef17551dd551c55d0157bf1850c20711acb3e2103bf977f231b", "07d969027f572e958d046255c8c2ff0d2020fa4d3efb0ff0a950a2db4e50a7a1"], "20000000", "1729d72d", "5b807185", true]}`
)



func listenAndPrint(conn *net.Conn){
	connbuf := bufio.NewReader(*conn)
	for {
		// read in input from stdin
		//str, err := connbuf.ReadString('\n')
		str, err := connbuf.ReadString('\n')
		if err != nil{
			(*conn).Close()
			log.Println("Connection Closed due to ",err)
			return
		}
		if len(str)>0 {
			fmt.Print("Message from server: " + str)
		}
	}
}

func sendText(conn *net.Conn){
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		_, err := fmt.Fprintf(*conn, text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}
	}
}

func main(){

	conn, _ := net.Dial(CONN_TYPE, CONN_ADDR)
	err := conn.(*net.TCPConn).SetKeepAlive(true)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	go listenAndPrint(&conn)
	go sendText(&conn)
	time.Sleep(time.Second * 1000)



	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []interface{}{[]string{"1","2"}, "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)


}
