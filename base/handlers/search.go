package handlers

import (
	"net/http"
	"time"

	"github.com/patrickdappollonio/division-lfg/base/helpers/render"
	"github.com/patrickdappollonio/division-lfg/base/models"

	"golang.org/x/net/context"
)

type SearchResponse struct {
	Results int             `json:"results"`
	Agents  []*models.Agent `json:"agents"`
}

func Search(_ context.Context, w http.ResponseWriter, r *http.Request) {

	var response SearchResponse

	a1 := &models.Agent{
		Username:    "Agnizamaka",
		Platform:    "pc",
		Microphone:  true,
		StoryLevel:  30,
		DZLevel:     75,
		GearLevel:   182,
		Activity:    "Misiones Diarias",
		Description: "need one more for daily challenge quick run",
		Firearms:    187234,
		Stamina:     56783,
		Electronics: 6485,
		Registered:  time.Now(),
	}

	a2 := &models.Agent{
		Username:    "Oceanfabreze",
		Platform:    "ps4",
		StoryLevel:  30,
		DZLevel:     99,
		GearLevel:   204,
		Activity:    "Zona Oscura",
		Description: "Looking for a group of adults that have a good route to run to farm exp. 170k60k20k. Send inv",
		Firearms:    170000,
		Stamina:     60000,
		Electronics: 20000,
		Registered:  time.Now(),
	}

	response.Agents = append(response.Agents, a1, a2)
	response.Results = len(response.Agents)

	render.Template.JSON(w, http.StatusOK, response)
}
