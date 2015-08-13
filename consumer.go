
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://api.github.com/repos/smahawar/myApp/commits/a086e7ff7a6dc53ff5cbc1f5ed0e600d2243865a/comments"

	res, err := http.Get(url)

	if err != nil{
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil{
		panic(err)
	}

	err = json.Unmarshal(body, &p)

	if err != nil{
		panic(err)
	}

	fmt.Println(p)
	
}