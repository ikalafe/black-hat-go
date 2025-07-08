package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	host := os.Getenv("MSFHOST")
	pass := os.Getenv("MSFPASS")
	user := "msf"

	fmt.Println("MSFHOST:", host)
	fmt.Println("MSFPASS:", pass)

	if host == "" || pass == "" {
		log.Fatalln("Missing required enviroment variable MSFHOST or MSFPASS")
	}
	msf, err := New(host, user, pass)
	if err != nil {
		log.Panicln(err)
	}
	defer msf.Logout()

	sessions, err := msf.SessionList()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Sessions:")
	for _, session := range sessions {
		fmt.Printf("%5d %s\n", session.ID, session.Info)
	}
}
