package scripts

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Client() {

	log.Printf("Client on")
	client := &http.Client{}
	url := "http://localhost:8080/"
	resp, err := client.Get(url)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	if err != nil {
		log.Println("Error:", err)
		return
	}
	log.Println("Response:", string(body))
}
