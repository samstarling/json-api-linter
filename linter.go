package main

import "fmt"
import "net/http"
import "io/ioutil"
import "os"
import "encoding/json"

func main() {

	args := os.Args[1:]

	url := args[0]
	res := fetch(url)
	dat := parse(res)
	validate(dat)
}

func fetch(url string) *http.Response {
	res, err := http.Get(url)
	perror(err)
	return res
}

func parse(res *http.Response) map[string]interface{} {
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	perror(err)

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}
	return dat
}

func validate(dat map[string]interface{}) {
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
