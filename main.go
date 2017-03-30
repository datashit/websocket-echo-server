package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":80", "Websocket (server adress:port)")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func socketHandle(w http.ResponseWriter, r *http.Request) {

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {

		messageType, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Read error !")
			return
		}
		log.Printf("MessageType %v Recived: %s ", messageType, message)

		text := []byte("Server message \r\n")

		err = c.WriteMessage(messageType, text)
		if err != nil {
			log.Println("Write error 2 !")
			return
		}
		err = c.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Write error !")
			return
		}
		log.Println("Message sending.")

	}

}

func main() {
	fmt.Println("Websocket working!")
	http.HandleFunc("/", socketHandle)
	log.Fatal(http.ListenAndServe(*addr, nil))

}
