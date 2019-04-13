package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	statusCalled := promauto.NewCounter(prometheus.CounterOpts{
		Name: "simple_web_server_status_called_total",
		Help: "The total number of request made to status endpoint",
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!\n"))
	})

	http.HandleFunc("/_status", func(w http.ResponseWriter, r *http.Request) {
		statusCalled.Inc()
		w.Write([]byte("Ok\n"))
	})

	http.HandleFunc("/fmt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello formater!")
	})

	http.HandleFunc("/param", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Give param is: ")
		fmt.Fprintln(w, r.URL.Query().Get("p"))
	})

	http.Handle("/metrics", promhttp.Handler())

	log.Fatalln(http.ListenAndServe(":4000", nil)) // yield to package default
}
