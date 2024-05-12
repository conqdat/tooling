package main

import (
	"fmt"
	"math/rand"
	"time"
)


var sendInterval = time.Second

type OBUData struct {
	OBUID int `json:"obuID"`
	Lat float64 `json:"lat"`
	Long float64 `json:"long"`
}

func getCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}

func main() {
	for {
		fmt.Println(getCoord())
		time.Sleep(sendInterval)
	}
}