package service

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func genID() string {
	return fmt.Sprintf("%x", rand.Int())
}
