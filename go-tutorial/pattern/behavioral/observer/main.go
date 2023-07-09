package main

import (
	item2 "tutorial/go-tutorial/pattern/behavioral/observer/item"
	weather2 "tutorial/go-tutorial/pattern/behavioral/observer/weather"
)

func main() {
	weather := weather2.Weather{}
	phone := weather2.Phone{}
	desktop := weather2.Desktop{}
	weather.Add(phone)
	weather.Add(desktop)
	weather.Notify()

	//----------------------------------------
	item := item2.Item{}
	item.Name = "testItem"
	customer := item2.Customer{}

	item.Subscribe(customer)
	item.NotifyAll()
}
