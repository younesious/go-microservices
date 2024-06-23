package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	// "os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	fmt.Println("Starting front end service on port 8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Panic(err)
	}
}

//go:embed templates
var templateFS embed.FS

func render(w http.ResponseWriter, t string) {
	partials := []string{
		"templates/base.layout.gohtml",
		"templates/header.partial.gohtml",
		"templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("templates/%s", t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	tmpl, err := template.ParseFS(templateFS, templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data struct {
		BrokerURL string
	}

	/*
		// comment out if you wanna use docker swarm
		data.BrokerURL = os.Getenv("BOROKER_URL")
		if data.BrokerURL == "" {
			data.BrokerURL = "http://backend"
		}
	*/

	/*
		// comment out if you wanna use k8s
		data.BrokerURL = os.Getenv("BROKER_URL")
		if data.BrokerURL == "" {
			data.BrokerURL = "http://broker-service.info" // replcae this address with address of broker-service svc in minikube
		}
	*/

	// comment out if you wanna run on local or with docker-compose
	data.BrokerURL = "http://localhost:8080"

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
