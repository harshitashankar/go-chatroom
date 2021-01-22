package main

import(
	"flag"
	"log"

	"github.com/harshitashankar/go-chatroom/client"
	"github.com/harshitashankar/go-chatroom/tui"
)

func main() {
	address := flag.String("server", "", "which server to connect to" )

	flag.Parse()

	client := client.NewClient()
	err := client.Dial(*address)
	log.Printf("after client dialing")

	if err != nil {
		log.Printf("err: after client dialing")
		log.Fatal(err)
	}

	defer client.Close()

	go client.Start()

	tui.StartUi(client)
}