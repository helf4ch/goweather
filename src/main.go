package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	res, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=50.2796&longitude=127.5405&current=relative_humidity_2m,temperature_2m&timezone=auto")
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.StatusCode)
	fmt.Printf("%s", body)
}
