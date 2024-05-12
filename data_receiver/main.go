package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"tolling/types"
)

func main() {
	recv := NewDataReceiver()
	http.HandleFunc("/ws", recv.handleWS)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return
	}
	fmt.Println("Data receiver work")
}

type DataReceiver struct {
	msgChan chan types.OBUData
	conn    *websocket.Conn
}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgChan: make(chan types.OBUData, 128),
	}
}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err)
	}
	dr.conn = conn

	go dr.wsReceiveLoop()
}

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("client connected to server")
	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println(err)
			continue
		}
		dr.msgChan <- data
	}
}
