package utils_test

import (
	"encoding/json"
	"github.com/cheebo/go-config/utils"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Race struct {
	Alias      string  `json:"alias"`
	Number     int     `json:"number"`
	Friendly   bool    `json:"friendly"`
	Population float64 `json:"population"`
}

type TestJSON struct {
	Name         string  `json:"name"`
	Age          int     `json:"age"`
	Alian        bool    `json:"alian"`
	PlanetNumber float64 `json:"planet_number"`
	Race         Race    `json:"race"`
}

func TestParse(t *testing.T) {
	assert := assert.New(t)

	cfg := TestJSON{
		Name:         "who",
		Age:          2000,
		Alian:        true,
		PlanetNumber: 1.2,
		Race: Race{
			Alias:      "timelord",
			Number:     2,
			Friendly:   true,
			Population: 2,
		},
	}

	data, err := json.Marshal(cfg)
	assert.NoError(err)

	m := utils.JsonParse(data)
	spew.Dump(m)
	assert.Equal(cfg.Name, m["name"].(string), "Incorrect name")
	assert.Equal(float64(cfg.Age), m["age"].(float64), "Incorrect age")
	assert.Equal(cfg.Alian, m["alian"].(bool), "Incorrect alian status")
	assert.Equal(cfg.PlanetNumber, m["planet_number"].(float64), "Incorrect plunet number")
	assert.Equal(cfg.Race.Alias, m["race.alias"].(string), "Incorrect race alias")
	assert.Equal(float64(cfg.Race.Number), m["race.number"].(float64), "Incorrect race number")
	assert.Equal(cfg.Race.Friendly, m["race.friendly"].(bool), "Incorrect friendly status")
	assert.Equal(cfg.Race.Population, m["race.population"].(float64), "Incorrect population")
}
