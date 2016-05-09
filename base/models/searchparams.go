package models

import (
	"fmt"
	"strconv"
)

type SearchParams struct {
	Platform       string `json:"platform"`
	Activity       string `json:"activity"`
	Microphone     string `json:"microphone"`
	BoolMicrophone bool   `json:"_microphone"`
	StoryLevel     string `json:"storylevel"`
	IntStoryLevel  int    `json:"_storylevel"`
	DZLevel        string `json:"dzlevel"`
	IntDZLevel     int    `json:"_dzlevel"`
	GearLevel      string `json:"gearlevel"`
	IntGearLevel   int    `json:"_gearlevel"`
	Search         string `json:"search"`
}

const ANY = "any"

var strOptionInvalid = "la opción de %v no es válida"

var (
	ErrMicrophoneOptionInvalid = fmt.Errorf(strOptionInvalid, "micrófono seleccionado")
	ErrStoryLevelOptionInvalid = fmt.Errorf(strOptionInvalid, "nivel en el modo historia")
	ErrDZLevelOptionInvalid    = fmt.Errorf(strOptionInvalid, "nivel en la zona oscura")
	ErrGearLevelOptionInvalid  = fmt.Errorf(strOptionInvalid, "nivel de equipo")
)

func (a *SearchParams) Validate() error {
	// Check if the platform exists
	if !PlatformExists(a.Platform) && a.Platform != ANY {
		return ErrPlatformIncorrect
	}

	// Check if the activity exists
	if !ActivityExists(a.Activity) && a.Activity != ANY {
		return ErrActivityIncorrect
	}

	// Check if the microphone value is valid
	if a.Microphone != ANY {
		if a.Microphone != "true" && a.Microphone != "false" {
			return ErrMicrophoneOptionInvalid
		}

		// Convert to the proper value, false by default
		// so we only need to check for true
		if a.Microphone == "true" {
			a.BoolMicrophone = true
		}
	}

	// Check if int story level is invalid
	if a.StoryLevel != ANY {
		// Convert string to int
		val, err := strconv.Atoi(a.StoryLevel)

		// Check if there was an error in the conversion
		if err != nil {
			return ErrStoryLevelOptionInvalid
		}

		// Check if value is in between allowed levels
		if !isBetweenMinMax(val, MinStoryLevel, MaxStoryLevel) {
			return ErrStoryLevelOverflow
		}

		// If not, we can add it as the proper value
		a.IntStoryLevel = val
	}

	// Check if int darkzone level is invalid
	if a.DZLevel != ANY {
		// Convert string to int
		val, err := strconv.Atoi(a.DZLevel)

		// Check if there was an error in the conversion
		if err != nil {
			return ErrDZLevelOptionInvalid
		}

		// Check if value is in between allowed levels
		if !isBetweenMinMax(val, MinDZLevel, MaxDZLevel) {
			return ErrDZLevelOverflow
		}

		// If not, we can add it as the proper value
		a.IntDZLevel = val
	}

	// Check if int darkzone level is invalid
	if a.GearLevel != ANY {
		// Convert string to int
		val, err := strconv.Atoi(a.GearLevel)

		// Check if there was an error in the conversion
		if err != nil {
			return ErrGearLevelOptionInvalid
		}

		// Check if value is in between allowed levels
		if !isBetweenMinMax(val, MinGearLevel, MaxGearLevel) {
			return ErrGearLevelOverflow
		}

		// If not, we can add it as the proper value
		a.IntGearLevel = val
	}

	return nil
}
