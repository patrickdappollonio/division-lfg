package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/patrickdappollonio/division-lfg/base/helpers/render"
	"github.com/patrickdappollonio/division-lfg/base/models"
	"github.com/patrickdappollonio/division-lfg/base/models/binding"

	"golang.org/x/net/context"
)

func AddUser(_ context.Context, w http.ResponseWriter, r *http.Request) {
	// Create a reader from the request body that contains JSON
	decoder := json.NewDecoder(r.Body)

	// Create a temporary placeholder, this is needed because the
	// original `model.Agent` structure doesn't match the strings we get
	// from the POST values
	var values struct {
		Username    string `json:"username"`
		Platform    string `json:"platform"`
		Activity    string `json:"activity"`
		Microphone  string `json:"microphone"`
		StoryLevel  string `json:"storylevel"`
		DZLevel     string `json:"dzlevel"`
		GearLevel   string `json:"gearlevel"`
		Description string `json:"description"`
		Firearms    string `json:"firearms"`
		Stamina     string `json:"stamina"`
		Electronics string `json:"electronics"`
	}

	// Send the request body to the variable placeholder
	if err := decoder.Decode(&values); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert string values to integer
	var (
		storyLevel  = intify(values.StoryLevel)
		dzLevel     = intify(values.DZLevel)
		gearLevel   = intify(values.GearLevel)
		firearms    = intify(values.Firearms)
		stamina     = intify(values.Stamina)
		electronics = intify(values.Electronics)
	)

	// Convert string value to boolean
	var hasMicrophone bool
	if values.Microphone == "true" {
		hasMicrophone = true
	}

	// Create an agent
	agent := &models.Agent{
		Username:    values.Username,
		Platform:    values.Platform,
		Activity:    values.Activity,
		Microphone:  hasMicrophone,
		StoryLevel:  storyLevel,
		DZLevel:     dzLevel,
		GearLevel:   gearLevel,
		Description: values.Description,
		Firearms:    firearms,
		Stamina:     stamina,
		Electronics: electronics,
		Registered:  time.Now(),
	}

	// Try validating the agent
	if err := agent.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO:
	// 1. Create an appengine key
	// 2. Try looking for the user in the given platform
	// 3. Check wether if exists or not, if it does exists, warn; if not add
	// 4. Return an okay!

	render.Template.JSON(w, http.StatusOK, binding.KV{
		"agent": agent,
	})
}

func intify(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}
