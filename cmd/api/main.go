package main

import (
	"flag"
	"log"

	"github.com/wDRxxx/yandex_calculator_api/cmd/internal/app"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8080", "The port to listen on")
}

func main() {
	flag.Parse()

	app := app.NewApp(port)
	err := app.Start()
	if err != nil {
		log.Fatal(err)
	}

}
