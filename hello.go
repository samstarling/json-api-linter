package main

import "fmt"
import "net/http"
import "io/ioutil"
import "os"
import "encoding/json"

func perror(err error) {
	if err != nil {
		panic(err)
	}
}
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
	fmt.Println(dat)
}
