package main

import (
	"github.com/deexithparand/hookdock/hookdock"
)

func main() {

	// need a server running at port (5431)
	err := hookdock.StartServer()
	if err != nil {
		panic(err)
	}

	defer hookdock.ExitServer()

}
