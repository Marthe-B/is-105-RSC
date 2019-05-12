package main

import (
	"encoding/json"
	"fmt"
	"net"

	"../bruker"
)

func handler(c net.Conn, b []string) {
	//c.Write([]byte("ok"))
	c.Write([]byte(b[0]))
	fmt.Println("Motatt en forespørsel.")
	fmt.Println("Remote address: " + c.RemoteAddr().String())
	fmt.Println("Local address: " + c.LocalAddr().String())
	c.Close()
}
func main() {
	ocj := &bruker.Bruker{
		Navn:  "Ole Christian",
		Epost: "olecj12@uia.no"}

	var b []string
	str, _ := json.Marshal(ocj)
	b = append(b, string(str))

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
