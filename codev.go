package main

import (
	E2M "codev/eco2mix"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	os.Setenv("DBUSER", "b42c4d4f0f1d03")
	os.Setenv("DBPASS", "4d227bbc")
	os.Setenv("NET", "tcp")
	os.Setenv("ADDR", "eu-cdbr-west-02.cleardb.net")
	os.Setenv("API_KEY", "751d9da6351bcb1d0a3710004096de3c5b0cb94d36e3264cb6d1d5f4")

	ticker := time.NewTicker(1 * time.Hour)
	done := make(chan bool)
	timesUp()
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				timesUp()
			}
		}
	}()

	httpServer()

	fmt.Println("This shouldn't be printed")
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

}

func timesUp() {
	t := time.Now()
	timeBeg := t.Add(-time.Hour * 4)
	min := (timeBeg.Minute() + (15 - (timeBeg.Minute() % 15)))
	dateBeg := time.Date(timeBeg.Year(), timeBeg.Month(), timeBeg.Day(), timeBeg.Hour(), min, 0, 0, timeBeg.Location())

	timeEnd := dateBeg.Add(time.Minute * 45)

	strBeg := strconv.Itoa(dateBeg.Year()) + "-" + strconv.Itoa(int(dateBeg.Month())) + "-" + strconv.Itoa(dateBeg.Day()) + "T"
	strBeg += strconv.Itoa(dateBeg.Hour()) + "%3A" + strconv.Itoa(dateBeg.Minute()) + "%3A00"

	strEnd := strconv.Itoa(timeEnd.Year()) + "-" + strconv.Itoa(int(timeEnd.Month())) + "-" + strconv.Itoa(timeEnd.Day()) + "T"
	strEnd += strconv.Itoa(timeEnd.Hour()) + "%3A" + strconv.Itoa(timeEnd.Minute()) + "%3A00"

	strDate := strBeg + "Z+TO+" + strEnd

	callE2M(strDate)
}

func callE2M(date string) {
	E2M.Maineco2mix(date)
}

func httpServer() {
	port := os.Getenv("PORT")
	//List all handlers
	http.HandleFunc("/eco2mix", eco2mixHandler)
	http.HandleFunc("/eco2mix/24h", dayHandler)

	log.Fatal(http.ListenAndServe(port, nil))
}

// HTTP HANDLERS
func eco2mixHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	timesUp()
	consos, err := E2M.RecentData()
	if err != nil {
		log.Println(err)
	}
	consosJson, errjson := json.Marshal(consos)
	if err != nil {
		fmt.Println(errjson)
		return
	}
	fmt.Println("LAST HOUR ECO2MIX CALLED")
	fmt.Fprintf(w, string(consosJson))
}

func dayHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	consos, err := E2M.DayRecentData()
	if err != nil {
		log.Println(err)
	}
	consosJson, errjson := json.Marshal(consos)
	if err != nil {
		fmt.Println(errjson)
		return
	}
	fmt.Println("LAST DAY ECO2MIX CALLED")
	fmt.Fprintf(w, string(consosJson))
}
