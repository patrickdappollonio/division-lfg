package models

import "time"

type Agent struct {
	Username    string    `json:"username"`
	Platform    string    `json:"platform"`
	Microphone  bool      `json:"microphone"`
	StoryLevel  int       `json:"storylevel"`
	DZLevel     int       `json:"dzlevel"`
	GearLevel   int       `json:"gearlevel"`
	Activity    string    `json:"activity"`
	Description string    `json:"description"`
	Firearms    int       `json:"firearms"`
	Stamina     int       `json:"stamina"`
	Electronics int       `json:"electronics"`
	Registered  time.Time `json:"registered"`
}
