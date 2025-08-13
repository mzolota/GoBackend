package main

func main() {
	hub := newHub()
	go hub.run()

	StartWebserver(hub)
}
