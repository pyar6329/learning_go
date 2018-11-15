package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	switch rand.Intn(6) + 1 {
	case 1:
		fmt.Println("凶")
	case 2, 3:
		fmt.Println("吉")
	case 4, 5:
		fmt.Println("中吉")
	case 6:
		fmt.Println("大吉")
	}
}
