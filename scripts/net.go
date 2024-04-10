package scripts

import (
	"log"
	"net/http"
)

func Web() {

	log.Printf("Server on ")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "scripts/index.html")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Error:", err)
		return
	}

}
