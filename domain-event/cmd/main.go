package main

import (
	"domain-event/internal/app"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		panic("can't run")
	}
	app.Waiter().Wait()
}
