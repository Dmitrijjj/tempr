package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"std/github.com/tempr/models"
)

func main() {
	url := "https://api.openweathermap.org/data/2.5/weather?q=Novosibirsk&appid=62bea516ac233fbb63a5d21ce6985458"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	result := &models.Memury

	err = json.Unmarshal(body, result)
	if err != nil {
		fmt.Println(err)
	}
	r := result.Main.grads()
	r2 := result.Main.hum()
	fmt.Println(r, r2)
}
