package main

import (
	"net"
	"fmt"
	"encoding/json"
	"../bruker"
)



func main() {
	//ny struct av type bruker
	ocj := &bruker.Bruker {
		Navn: "Ole Christian",
		Epost: "ocj@uia.no" }
	
	//strA, _ := json.Marshal(map[string]string{"Navn": "Ole", "Epost": "oc@uia.no"})
	//fmt.Println(string(strA))
	
	//gjør om til json struktur
	strB, _ := json.Marshal(ocj)
	fmt.Printf("JSON: %s \n", strB)
	
	res := bruker.Bruker{}
	json.Unmarshal(strB, &res)

	fmt.Println("Navn: ", res.Navn)
	fmt.Println("Epost: ", res.Epost)

	conn,err := net.Dial("tcp", "127.0.0.1:5001")
	if err != nil {
		panic(err)
	}

	conn.Write([]byte("hello"))
	fmt.Println("Sent hello")

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		panic(err)
	}

	fmt.Println("Reply: ", string(reply))
	 
}