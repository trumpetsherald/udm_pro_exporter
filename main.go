package main

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	unifi_api "github.com/trumpetsherald/udm_pro_exporter/unifi_api"
)

func main() {
	var username string = os.Getenv("UDM_EXPORTER_USER")
	var password string = os.Getenv("UDM_EXPORTER_PASS")
	var device_url string = os.Getenv("UDM_DEVICE_URL")
	missing_param := false

	if len(username) < 1 {
		log.Println("Required environment variable UDM_EXPORTER_USER was missing.")
		missing_param = true
	}
	if len(password) < 1 {
		log.Println("Required environment variable UDM_EXPORTER_PASS was missing.")
		missing_param = true
	}
	if len(device_url) < 1 {
		log.Println("Required environment variable UDM_DEVICE_URL was missing.")
		missing_param = true
	}
	if missing_param {
		log.Fatal("One or more required environment variables were missing, exiting.")
	}

	api_broker := unifi_api.NewAPIBroker(device_url, username, password)

	if api_broker.VerifyConnectivity() {
		collector := unifi_api.NewUDMProCollector(api_broker)

		prometheus.MustRegister(collector)
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":9182", nil)
	} else {
		log.Fatal("Couldn't connect to device, exiting.")
	}
}
