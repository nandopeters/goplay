package main

import (
    "fmt"
    "net/http"
	"log"
   	"github.com/tmaniaci/mtrax1/emr/clinician"
   	"github.com/tmaniaci/mtrax1/emr/schedule"
)

func exit_handler(w http.ResponseWriter, r *http.Request) {
	log.Fatal("/exit handler called")
	return
}

func clinician_handler(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "%s",clinician.GetClinican()) 
}


func schedule_handler(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "%s", schedule.GetSchedule())
}

func main() {
	http.HandleFunc("/schedule", schedule_handler)
	http.HandleFunc("/clinician", clinician_handler)
	http.HandleFunc("/exit", exit_handler)

	http.ListenAndServe(":9025", nil)
}
