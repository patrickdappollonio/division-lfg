package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"

	"github.com/patrickdappollonio/division-lfg/base/helpers/render"
	"github.com/patrickdappollonio/division-lfg/base/models"
	"github.com/patrickdappollonio/division-lfg/base/models/binding"

	"golang.org/x/net/context"
)

var ErrCantProcessSearch = fmt.Errorf("no pudimos procesar la búsqueda, intenta más tarde")

func Search(_ context.Context, w http.ResponseWriter, r *http.Request) {
	// Create a new appengine ctx
	ctx := appengine.NewContext(r)

	// Create a reader from the request body that contains JSON
	decoder := json.NewDecoder(r.Body)

	// Create a placeholder to place all the request info
	var searchparams models.SearchParams

	// Send the request body to the variable placeholder
	if err := decoder.Decode(&searchparams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Validate the fields
	if err := searchparams.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create an empty slice of agents
	var agents []*models.Agent

	// Create a search object
	q := datastore.NewQuery(models.AgentGroup)

	// If there's a Platform, add it
	if searchparams.Platform != models.ANY {
		q = q.Filter("Platform =", searchparams.Platform)
	}

	// If there's a Activity, add it
	if searchparams.Activity != models.ANY {
		q = q.Filter("Activity =", searchparams.Activity)
	}

	// If there's a Microphone, add it
	if searchparams.Microphone != models.ANY {
		q = q.Filter("Microphone =", searchparams.BoolMicrophone)
	}

	// If there's a StoryLevel, add it
	if searchparams.StoryLevel != models.ANY {
		q = q.Filter("StoryLevel =", searchparams.IntStoryLevel)
	}

	// If there's a DZLevel, add it
	if searchparams.DZLevel != models.ANY {
		q = q.Filter("DZLevel =", searchparams.IntDZLevel)
	}

	// If there's a GearLevel, add it
	if searchparams.GearLevel != models.ANY {
		q = q.Filter("GearLevel =", searchparams.IntGearLevel)
	}

	// TODO once Google AppEngine kind-of supports
	// text searches, then add `searchparams.Search`
	// as part of the parameters of the search.

	// Perform the search
	if _, err := q.GetAll(ctx, &agents); err != nil {
		http.Error(w, ErrCantProcessSearch.Error(), http.StatusInternalServerError)
		return
	}

	// =====================
	agents = append(agents, &models.Agent{
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
	}, &models.Agent{
		Username:    "Oceanfabreze",
		Platform:    "ps4",
		StoryLevel:  30,
		DZLevel:     99,
		GearLevel:   204,
		Activity:    "Zona Oscura",
		Description: "Looking for a group of adults that have a good route to run to farm exp. 170k60k20k. Send inv",
		Firearms:    0,
		Stamina:     0,
		Electronics: 0,
		Registered:  time.Now(),
	})
	// =====================

	// Send the JSON response
	render.Template.JSON(w, http.StatusOK, binding.KV{
		"agents": agents,
		"count":  len(agents),
		"query":  searchparams,
	})
}

func validateSearchParams(r *http.Request) error {
	return nil
}
