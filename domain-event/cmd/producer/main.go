package main

import producer "domain-event/internal/app/producer_app"

func main() {
	app, err := producer.NewApp()
	if err != nil {
		panic("can't run")
	}
	app.Waiter().Wait()
}
