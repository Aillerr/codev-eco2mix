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
		Sort     []string `json:"sort"`
		Facet    []string `json:"facet"`
		Format   string   `json:"format"`
		Timezone string   `json:"timezone"`
	} `json:"parameters"`
	Records []struct {
		Datasetid string `json:"datasetid"`
		Recordid  string `json:"recordid"`
		Fields    struct {
			Pompage                                      string    `json:"pompage"`
			TchEolien                                    float64   `json:"tch_eolien"`
			Thermique                                    int       `json:"thermique"`
			Nucleaire                                    int       `json:"nucleaire"`
			TcoBioenergies                               float64   `json:"tco_bioenergies"`
			FluxPhysiquesDeBretagneVersNouvelleAquitaine string    `json:"flux_physiques_de_bretagne_vers_nouvelle_aquitaine"`
			TcoSolaire                                   float64   `json:"tco_solaire"`
			Date                                         string    `json:"date"`
			TcoEolien                                    float64   `json:"tco_eolien"`
			LibelleRegion                                string    `json:"libelle_region"`
			FluxPhysiquesDeNouvelleAquitaineVersBretagne string    `json:"flux_physiques_de_nouvelle_aquitaine_vers_bretagne"`
			DateHeure                                    time.Time `json:"date_heure"`
			Heure                                        string    `json:"heure"`
			TchSolaire                                   float64   `json:"tch_solaire"`
			TchBioenergies                               float64   `json:"tch_bioenergies"`
			TcoThermique                                 float64   `json:"tco_thermique"`
			TchHydraulique                               float64   `json:"tch_hydraulique"`
			Consommation                                 int       `json:"consommation"`
			Solaire                                      int       `json:"solaire"`
			Hydraulique                                  int       `json:"hydraulique"`
			TcoHydraulique                               float64   `json:"tco_hydraulique"`
			EchPhysiques                                 int       `json:"ech_physiques"`
			Bioenergies                                  int       `json:"bioenergies"`
			CodeInseeRegion                              string    `json:"code_insee_region"`
			Nature                                       string    `json:"nature"`
			Eolien                                       int       `json:"eolien"`
			TchThermique                                 float64   `json:"tch_thermique"`
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

type RegionDB struct {
	id         int
	code_INSEE int
	nom        string
}

type ConsoDB struct {
	Région      string
	DateHeure   time.Time
	Total       int
	Thermique   int
	Nucléaire   int
	Éolien      int
	Solaire     int
	Hydraulique int
	Pompage     string
	Bioénergies int
}
