package eco2mixstruct

import (
	"time"
)

func NewEco2MixAPI() *Eco2mixAPI {
	p := &Eco2mixAPI{}
	return p
}

type Eco2mixAPI struct {
	Nhits      int `json:"nhits"`
	Parameters struct {
		Dataset  string   `json:"dataset"`
		Rows     int      `json:"rows"`
		Start    int      `json:"start"`
		Facet    []string `json:"facet"`
		Format   string   `json:"format"`
		Timezone string   `json:"timezone"`
	} `json:"parameters"`
	Records []struct {
		Datasetid string `json:"datasetid"`
		Recordid  string `json:"recordid"`
		Fields    struct {
			Charbon                  int       `json:"charbon"`
			Bioenergies              int       `json:"bioenergies"`
			EchCommSuisse            int       `json:"ech_comm_suisse"`
			HydrauliqueStepTurbinage int       `json:"hydraulique_step_turbinage"`
			BioenergiesBiogaz        int       `json:"bioenergies_biogaz"`
			Nature                   string    `json:"nature"`
			GazAutres                int       `json:"gaz_autres"`
			HydrauliqueLacs          int       `json:"hydraulique_lacs"`
			BioenergiesBiomasse      int       `json:"bioenergies_biomasse"`
			TauxCo2                  int       `json:"taux_co2"`
			Nucleaire                int       `json:"nucleaire"`
			Eolien                   int       `json:"eolien"`
			EchPhysiques             int       `json:"ech_physiques"`
			Gaz                      int       `json:"gaz"`
			EchCommEspagne           string    `json:"ech_comm_espagne"`
			Perimetre                string    `json:"perimetre"`
			EchCommAngleterre        string    `json:"ech_comm_angleterre"`
			BioenergiesDechets       int       `json:"bioenergies_dechets"`
			EchCommAllemagneBelgique int       `json:"ech_comm_allemagne_belgique"`
			FioulCogen               int       `json:"fioul_cogen"`
			EchCommItalie            int       `json:"ech_comm_italie"`
			Solaire                  int       `json:"solaire"`
			Hydraulique              string    `json:"hydraulique"`
			GazTac                   int       `json:"gaz_tac"`
			Consommation             int       `json:"consommation"`
			PrevisionJ               int       `json:"prevision_j"`
			FioulAutres              int       `json:"fioul_autres"`
			GazCcg                   int       `json:"gaz_ccg"`
			Fioul                    int       `json:"fioul"`
			Pompage                  int       `json:"pompage"`
			DateHeure                time.Time `json:"date_heure"`
			FioulTac                 int       `json:"fioul_tac"`
			Date                     string    `json:"date"`
			Heure                    string    `json:"heure"`
			HydrauliqueFilEauEclusee string    `json:"hydraulique_fil_eau_eclusee"`
			GazCogen                 int       `json:"gaz_cogen"`
			PrevisionJ1              int       `json:"prevision_j1"`
		} `json:"fields"`
		RecordTimestamp time.Time `json:"record_timestamp"`
	} `json:"records"`
	FacetGroups []struct {
		Name   string `json:"name"`
		Facets []struct {
			Name  string `json:"name"`
			Count int    `json:"count"`
			State string `json:"state"`
			Path  string `json:"path"`
		} `json:"facets"`
	} `json:"facet_groups"`
}
