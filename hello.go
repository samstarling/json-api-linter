package main

import "fmt"
import "net/http"
import "io/ioutil"
import "os"
import "encoding/json"

func main() {

	args := os.Args[1:]

	url := args[0]
	res, err := http.Get(url)
	perror(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	perror(err)

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}
  if _, ok := dat["name"]; ok {
    fmt.Println("Name element found")
  } else {
    fmt.Println("Name element not found")
  }
}

func perror(err error) {
	if err != nil {
		panic(err)
	}
}
