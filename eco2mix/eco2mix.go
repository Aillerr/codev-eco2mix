package eco2mix

import (
	"codev/eco2mix/eco2mixstruct"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func MakeRequest(key string) {
	address := "https://opendata.reseaux-energies.fr/api/records/1.0/search/?dataset=eco2mix-national-tr&q=&facet=nature&facet=date_heure"
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

}
