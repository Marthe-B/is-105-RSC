package main

import (
	"encoding/json"
	"fmt"
	"net"

	"../bruker"
)

func main() {
	//test av json lokalt (ikke del av ICA5, kun til referanse)
	/*ocj := &bruker.Bruker{
		Navn:  "Ole Christian",
		Epost: "ocj@uia.no"}
	strA, _ := json.Marshal(map[string]string{"Navn": "Ole", "Epost": "oc@uia.no"})
	fmt.Println(string(strA))
	//gjør om til json struktur
	strB, _ := json.Marshal(ocj)
	fmt.Printf("JSON: %q \n", strB)
	*/

	conn, err := net.Dial("tcp", "84.48.236.40:5001")
	//conn, err := net.Dial("tcp", "10.0.0.7:5001")
	//conn, err := net.Dial("tcp", "127.0.0.1:5001")
	if err != nil {
		panic(err)
	}

	conn.Write([]byte("hello"))
	fmt.Println("Sent hello")

	reply := make([]byte, 1024)
	n, err := conn.Read(reply)
	if err != nil {
		panic(err)
	}
	//lag en ny byteslice som har lengde (int n) som tilsvarer meldingen vi mottok:
	reply2 := make([]byte, 0, 1024)
	reply2 = append(reply2, reply[:n]...)

	//Fra oppgave 5a, der vi kunne printe ut medlingen direkte som en streng:
	fmt.Println("Reply (string printout): ", string(reply2))

	//lag en Bruker struct som vi kan unmarshalle JSON svaret i
	res := bruker.Bruker{}
	err2 := json.Unmarshal(reply2, &res)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println("Navn: ", res.Navn)
	fmt.Println("Epost: ", res.Epost)
}
