package main

var which string

func main() {
	if which == "client" {
		clientMain()
	} else {
		serverMain()
	}
}
