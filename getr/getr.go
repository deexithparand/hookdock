package getr

import (
	"fmt"
	"io"
	"net/http"
)

func GetRequestRootURL(w http.ResponseWriter, r *http.Request) {
	var writer io.Writer = w
	n, err := writer.Write([]byte("HELLO WORLD!"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Number of bytes sent in response %d\n", n)
}
