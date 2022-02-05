package eco2mix

import (
	"codev/eco2mix/eco2mixstruct"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var table string

// MAIN FUNC
func Maineco2mix(date string, hour string) {

	table = "codev"
	initDB()

	var consosDB []eco2mixstruct.ConsoDB = fillSlice(makeRequest(date, hour))

	addMultipletoDB(consosDB)

	consos, err := recentData()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Consos found: %v\n", consos)
}

// API CALLS TO FETCH DATA
func makeRequest(date string, hour string) eco2mixstruct.Eco2mixAPI {
	address := "https://opendata.reseaux-energies.fr/api/records/1.0/search/?dataset=eco2mix-regional-tr&q=&rows=20&sort=date_heure&facet=libelle_region&facet=nature&facet=date_heure&facet=heure" + date + hour
	fmt.Println(address)

	resAPI := eco2mixstruct.Eco2mixAPI{}

	reqClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, address, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("User-Agent", "spacecount-tutorial")
	req.Header.Add("apikey", os.Getenv("API_KEY"))

	res, getErr := reqClient.Do(req)

	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &resAPI)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(resAPI.Nhits)
	return resAPI
}

func fillSlice(apiData eco2mixstruct.Eco2mixAPI) []eco2mixstruct.ConsoDB {
	var consos []eco2mixstruct.ConsoDB
	for _, element := range apiData.Records {
		var conso eco2mixstruct.ConsoDB
		var fields = element.Fields

		conso.Région = fields.LibelleRegion
		conso.DateHeure = fields.DateHeure.Add(time.Hour * 1)
		conso.Total = fields.Consommation
		conso.Thermique = fields.Thermique
		conso.Nucléaire = fields.Nucleaire
		conso.Éolien = fields.Eolien
		conso.Solaire = fields.Solaire
		conso.Hydraulique = fields.Hydraulique
		conso.Pompage = fields.Pompage
		conso.Bioénergies = fields.Bioenergies

		consos = append(consos, conso)
	}
	return consos

}

// SQL AND DB RELATED STUFF
func initDB() {
	dsn := os.Getenv("DBUSER") + ":" + os.Getenv("DBPASS") + "@" + os.Getenv("NET") + "(" + os.Getenv("ADDR") + ")/"
	dsn += table
	dsn += "?parseTime=true"
	//fmt.Println(dsn)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	//fmt.Println("Connected")
}

//QUERIES
func addToDB(conso eco2mixstruct.ConsoDB) error {
	_, err := db.Exec("INSERT INTO eco2mixconso (région, dateHeure, total, thermique, nucléaire, éolien, solaire, hydraulique, pompage, bioénergies) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", conso.Région, conso.DateHeure, conso.Total, conso.Thermique, conso.Nucléaire, conso.Éolien, conso.Solaire, conso.Hydraulique, conso.Pompage, conso.Bioénergies)
	if err != nil {
		return fmt.Errorf("addToDB: %v", err)
	}
	if err != nil {
		return fmt.Errorf("addToDB: %v", err)
	}
	return nil
}

func addMultipletoDB(consos []eco2mixstruct.ConsoDB) {
	for _, element := range consos {
		err := addToDB(element)
		if err != nil {
			log.Println(err)
		}
	}
}

//Most recent data
func recentData() ([]eco2mixstruct.ConsoDB, error) {
	var consos []eco2mixstruct.ConsoDB

	rows, err := db.Query("SELECT * FROM `eco2mixconso` ORDER BY `dateHeure` DESC LIMIT 12")
	if err != nil {
		return nil, fmt.Errorf("recentData : %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var conso eco2mixstruct.ConsoDB
		if err := rows.Scan(&conso.Région, &conso.DateHeure, &conso.Total, &conso.Thermique, &conso.Nucléaire, &conso.Éolien, &conso.Solaire, &conso.Hydraulique, &conso.Pompage, &conso.Bioénergies); err != nil {
			return nil, fmt.Errorf("recentData : %v", err)
		}
		consos = append(consos, conso)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("recentData : %v", err)
	}
	return consos, nil
}
