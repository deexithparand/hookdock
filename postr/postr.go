package postr

import (
	"fmt"
	"io"
	"net/http"
)

func PostRequestRootURL(w http.ResponseWriter, r *http.Request) {
	var writer io.Writer = w
	n, err := writer.Write([]byte("Received Log"))
	if err != nil {
		panic(err)
	}

	request_body := r.Body
	defer request_body.Close()

	var data []byte
	data, err = io.ReadAll(request_body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Parsed Data from request body : %s and sent response of %d bytes\n", string(data), n)

}
