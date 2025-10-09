package main

import (
	"fmt"

	"github.com/prometheus/common/promslog"
)

func main() {
	fmt.Println("Hello World")

	promslogConfig := &promslog.Config{}
	fmt.Println(promslogConfig)
}
