package network

import (
	"flag"
	"github.com/gorilla/websocket"
	"net/http"
	"log"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var upgrader = websocket.Upgrader{}

func Upgrade(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("There was an issue upgrading a request: ", err)
	}
	NewClient(c);
}

func StartServer() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/ws", Upgrade)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}