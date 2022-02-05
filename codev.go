package main

import (
	E2M "codev/eco2mix"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var port int = 8000
var portStr = ":" + strconv.Itoa(port)

func main() {
	os.Setenv("DBUSER", "root")
	os.Setenv("DBPASS", "")
	os.Setenv("NET", "tcp")
	os.Setenv("ADDR", "127.0.0.1")
	os.Setenv("API_KEY", "751d9da6351bcb1d0a3710004096de3c5b0cb94d36e3264cb6d1d5f4")

	ticker := time.NewTicker(30 * time.Minute)
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
	fmt.Println("http started")

	time.Sleep(999999999999)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

}

func timesUp() {
	t := time.Now()
	newt := t.Add(-time.Hour * 1)
	newt2 := newt.Add(-time.Minute * 45)
	fmt.Println(newt2)
	min := (newt2.Minute() + (15 - (newt2.Minute() % 15)))
	newt3 := time.Date(newt2.Year(), newt2.Month(), newt2.Day(), newt2.Hour(), min, 0, 0, newt2.Location())
	fmt.Println(newt3)
	for i := 0; i < 4; i++ {
		evolvt := newt3.Add(time.Minute * time.Duration(15*i))
		date := strconv.Itoa(evolvt.Year()) + "%2F" + strconv.Itoa(int(evolvt.Month())) + "%2F" + strconv.Itoa(evolvt.Day())
		hour := strconv.Itoa(evolvt.Hour()) + "%3A" + strconv.Itoa(evolvt.Minute())
		refDate := "&refine.date_heure=" + date
		refHeure := "&refine.heure=" + hour
		callE2M(refDate, refHeure)
	}
}

func callE2M(date string, hour string) {
	E2M.Maineco2mix(date, hour)
}

func httpServer() {
	//List all handlers
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(portStr, nil))
}

// Handler type function for a http call
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "E2M")
	//Special case, depends on what the server gets
	timesUp()
}
