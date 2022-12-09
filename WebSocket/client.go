package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"time"
)

type Demo struct {
    Name string `json:"name"`
    Age  int `json:"age"`
}

func main(){
	var headers http.Header = make(map[string][]string)
	wconn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", headers)
	if err != nil {
		log.Printf("ERROR: %s. cannot connect to RED", err)
		return 
	}
	x := Demo {
		Name: "test_ws_user",
		Age: 10,
	}
	for {
		err = wconn.WriteJSON(x)
		if err != nil {
			log.Printf("ERROR: %s. cannot authenticate with RED.", err)
			return 
		}

		msgType, msg, err := wconn.ReadMessage()
		if err != nil {
			log.Printf("%v", err)
		}
		log.Printf("%s sent:Msgtype[%v] Msg[%s]\n",
			   wconn.RemoteAddr(), msgType, string(msg))
		time.Sleep(2 * time.Second)
	}
}
