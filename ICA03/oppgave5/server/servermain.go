package main

import (
	"net"
	"fmt"
	"encoding/json"
	"../bruker"
)



func handler(c net.Conn, b []string) {
	//c.Write([]byte("ok"))
	c.Write([]byte(b[0]))
	fmt.Println("Got something?")
    c.Close()
}
func main() {
	ocj := &bruker.Bruker {
		Navn: "Ole Christian",
		Epost: "ocj@uia.no" }
	
	var b []string
	str, _ := json.Marshal(ocj)
	b = append(b,string(str))

	l, err := net.Listen("tcp", ":5001")
	if err != nil {
		panic(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}

		go handler(c, b)
	}
}