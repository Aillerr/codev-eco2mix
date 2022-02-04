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
var dateFormat string

// MAIN FUNC
func Maineco2mix(apikey string) {
	dateFormat = "2006-01-02 15:04:05.000000"
	table = "codev"
	initDB()

	//makeRequest(apikey)
	var consosDB []eco2mixstruct.ConsoDB = fillSlice(makeRequest(apikey))
	/*slice := make([]eco2mixstruct.ConsoDB, 2)
	slice[0] = eco2mixstruct.ConsoDB{
		Région:    "ARA",
		DateHeure: "1999-01-02 00:00",
		Total:     5000,
		Pompage:   2500,
		Nucléaire: 2500,
	}
	slice[1] = eco2mixstruct.ConsoDB{
		Région:    "ARV",
		DateHeure: "1999-01-02 00:00",
		Total:     5000,
		Pompage:   2500,
		Nucléaire: 2500,
	}*/

	addMultipletoDB(consosDB)

	consos, err := recentData()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Consos found: %v\n", consos)
}

// API CALLS TO FETCH DATA
func makeRequest(key string) eco2mixstruct.Eco2mixAPI {
	address := "https://opendata.reseaux-energies.fr/api/records/1.0/search/?dataset=eco2mix-regional-tr&q=&sort=-date_heure&facet=libelle_region&facet=nature&facet=date_heure&refine.date_heure=2022%2F02%2F04"
	apiURL := address + "?key=" + key

	resAPI := eco2mixstruct.Eco2mixAPI{}

	reqClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("User-Agent", "spacecount-tutorial")

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
		conso.DateHeure = fields.DateHeure
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

	fmt.Println(dsn)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected")
}

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

	rows, err := db.Query("SELECT * FROM eco2mixconso LIMIT 12")
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
