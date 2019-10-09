package main

import (
    "fmt"
    "net/http"
    "net/http/httputil"
)

const DEFAULT_PORT = 6666

func handler(w http.ResponseWriter, r *http.Request) {
    dump, err := httputil.DumpRequest(r, true)
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        fmt.Println(err)
        return
    }
    fmt.Println(string(dump))
}

func main() {
    var port int = DEFAULT_PORT

    fmt.Println("Listening on ", port)
    http.HandleFunc("/", handler)
    http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

