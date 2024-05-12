package main

import (
	"github.com/gorilla/websocket"
	"log"
	"math"
	"math/rand"
	"time"
	"tolling/types"
)

const wsEndpoint = "ws://127.0.0.1:3000/ws"

var sendInterval = time.Second

func getLatLong() (float64, float64) {
	return getCoord(), getCoord()
}

func getCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	obuIDS := generateOBUIDS(20)
	for {
		for i := 0; i < len(obuIDS); i++ {
			lat, long := getLatLong()
			data := types.OBUData{
				OBUID: obuIDS[i],
				Lat:   lat,
				Long:  long,
			}
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(sendInterval)
	}
}

func generateOBUIDS(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
