package main

import consumer "domain-event/internal/app/consumer_app"

func main() {
	app, err := consumer.NewApp()
	if err != nil {
		panic("can't run")
	}
	app.Waiter().Wait()
}
