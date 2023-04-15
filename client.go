package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get(url + "/todos/1")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	defer resp.Body.Close() //body를 닫아야 소켓이 닫힘
	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(-1)
		}
		var item todo
		err = json.Unmarshal(body, &item)
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(-1)
		}
		fmt.Printf("%#v\n", item)
	}
}
