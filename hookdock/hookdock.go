package hookdock

import (
	"fmt"
	"io"
	"net/http"
)

type CustomHandler struct{}

func (c *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request method : ", r.Method)
	// fmt.Fprintf(w, "Hello, World!\n")

	var writer io.Writer = w
	n, err := writer.Write([]byte("HELLO WORLD!"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Number of bytes sent in response %d\n", n)
}

func StartServer() error {

	host, port := "localhost:", "5431"

	custom_handler := &CustomHandler{}

	server := &http.Server{
		Addr:    host + port,
		Handler: custom_handler,
	}

	err := http.ListenAndServe(server.Addr, server.Handler)
	if err != nil {
		return err
	}
	return nil

}

func ExitServer() {
	fmt.Println("Server Sleeps ...")
}
