package hookdock

import (
	"fmt"
	"net/http"

	"github.com/deexithparand/hookdock/getr"
	"github.com/deexithparand/hookdock/postr"
)

type CustomHandler struct{}

func (c *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request method : ", r.Method)

	// single route - not required to handle routes

	switch r.Method {
	case "GET":
		getr.GetRequestRootURL(w, r)
	case "POST":
		postr.PostRequestRootURL(w, r)
	default:
		fmt.Println("other than GET  & POST")
	}
}

func StartServer() error {

	host, port := "localhost:", "5431"

	custom_handler := &CustomHandler{}

	server := &http.Server{
		Addr:    host + port,
		Handler: custom_handler,
	}

	fmt.Println("Server started and running ....")
	err := http.ListenAndServe(server.Addr, server.Handler)
	if err != nil {
		return err
	}
	return nil

}

func ExitServer() {
	fmt.Println("Server Sleeps ...")
}
