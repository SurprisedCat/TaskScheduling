package main

import (
	"fmt"
	"net/http"
)

//scheduler gets request from clients and add it into mapstring struct
//scheduler gets infomation about task servers
//scheduler assigns the task servers to clients according to different algorithm
//scheduler retrun the results to task clients
type textHandler struct {
	responseText string
}

func (th *textHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, th.responseText)
}

type indexHandler struct{}

func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	html := `<doctype html>
        <html>
        <head>
          <title>Hello World</title>
        </head>
        <body>
        <p>
          <a href="/welcome">Welcome</a> |  <a href="/message">Message</a>
        </p>
        </body>
</html>`
	fmt.Fprintln(w, html)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &indexHandler{})

	thWelcome := &textHandler{"TextHandler !"}
	mux.Handle("/text", thWelcome)

	http.ListenAndServe(":8000", mux)
}
