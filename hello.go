package main

import "fmt"
import "net/http"
import "io/ioutil"

func perror(err error) {
    if err != nil {
        panic(err)
    }
}
func main() {
    url := "http://samstarling.co.uk"
    res, err := http.Get(url)
    perror(err)
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    perror(err)
    fmt.Printf("%s", body)
}