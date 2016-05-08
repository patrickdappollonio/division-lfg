package binding

import (
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/patrickdappollonio/division-lfg/base/models"
)

type KV map[string]interface{}

func GetDefault(r *http.Request) KV {
	return KV{
		"CSRFToken": nosurf.Token(r),

		"AvailableActivities": models.Activities,
		"AvailablePlatforms":  models.Platforms,

		"MinUsernameLength": models.MinUsernameLength,
		"MaxUsernameLength": models.MaxUsernameLength,

		"MinStoryLevel": models.MinStoryLevel,
		"MaxStoryLevel": models.MaxStoryLevel,

		"MinDZLevel": models.MinDZLevel,
		"MaxDZLevel": models.MaxDZLevel,

		"MinGearLevel": models.MinGearLevel,
		"MaxGearLevel": models.MaxGearLevel,

		"MinDescriptionLength": models.MinDescriptionLength,
		"MaxDescriptionLength": models.MaxDescriptionLength,
	}
}
