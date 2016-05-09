package models

import (
	"fmt"
	"regexp"
	"time"
	"unicode/utf8"
)

const (
	MinUsernameLength = 3
	MaxUsernameLength = 20

	MinStoryLevel = 1
	MaxStoryLevel = 30

	MinDZLevel = 1
	MaxDZLevel = 99

	MinGearLevel = 150
	MaxGearLevel = 240

	MinDescriptionLength = 3
	MaxDescriptionLength = 150
)

const AgentGroup = "agents"

type Agent struct {
	Username    string    `json:"username"`
	Platform    string    `json:"platform"`
	Activity    string    `json:"activity"`
	Microphone  bool      `json:"microphone"`
	StoryLevel  int       `json:"storylevel"`
	DZLevel     int       `json:"dzlevel"`
	GearLevel   int       `json:"gearlevel"`
	Description string    `json:"description"`
	Firearms    int       `json:"firearms"`
	Stamina     int       `json:"stamina"`
	Electronics int       `json:"electronics"`
	Registered  time.Time `json:"registered"`
}

var (
	strLevelOverflow = "el nivel %v está fuera de los niveles disponibles"
	strYouveSelected = "has seleccionado %v"
)

var (
	ErrUsernameTooShort    = fmt.Errorf("el nombre de usuario debe tener entre %v y %v caracteres", MinUsernameLength, MaxUsernameLength)
	ErrDescriptionTooShort = fmt.Errorf("la descripción debe tener entre %v y %v caracteres", MinDescriptionLength, MaxDescriptionLength)
	ErrUsernameInvalid     = fmt.Errorf("el nonbre de usuario no parece ser válido")
	ErrPlatformIncorrect   = fmt.Errorf(strYouveSelected, "una plataforma inválida")
	ErrActivityIncorrect   = fmt.Errorf(strYouveSelected, "una actividad inválida")
	ErrStoryLevelOverflow  = fmt.Errorf(strLevelOverflow, "en el modo historia")
	ErrDZLevelOverflow     = fmt.Errorf(strLevelOverflow, "de la zona oscura")
	ErrGearLevelOverflow   = fmt.Errorf(strLevelOverflow, "de equipo")
)

var (
	reUsername = regexp.MustCompile(`([\w \-\_\.]){3,20}`)
)

func (a Agent) Validate() error {
	// Check if username has required length
	if !hasLength(a.Username, MinUsernameLength, MaxUsernameLength) {
		return ErrUsernameTooShort
	}

	// Check if username meets the standards
	if !reUsername.MatchString(a.Username) {
		return ErrUsernameInvalid
	}

	// Check if the platform exists
	if !PlatformExists(a.Platform) {
		return ErrPlatformIncorrect
	}

	// Check if the activity exists
	if !ActivityExists(a.Activity) {
		return ErrActivityIncorrect
	}

	// Check if story level is between the allowed values
	if !isBetweenMinMax(a.StoryLevel, MinStoryLevel, MaxStoryLevel) {
		return ErrStoryLevelOverflow
	}

	// Check if darkzone level is between the allowed values
	if !isBetweenMinMax(a.DZLevel, MinDZLevel, MaxDZLevel) {
		return ErrDZLevelOverflow
	}

	// Check if gear score is between the allowed values
	if !isBetweenMinMax(a.GearLevel, MinGearLevel, MaxGearLevel) {
		return ErrGearLevelOverflow
	}

	// Check if description has the required length
	if !hasLength(a.Description, MinDescriptionLength, MaxDescriptionLength) {
		return ErrDescriptionTooShort
	}

	return nil
}

func isBetweenMinMax(value, min, max int) bool {
	return value >= min && value <= max
}

func hasLength(value string, min, max int) bool {
	return utf8.RuneCountInString(value) >= min && utf8.RuneCountInString(value) <= max
}
