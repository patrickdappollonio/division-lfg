package models

var Platforms = map[string]string{
	"pc":    "PC (UPlay / Steam)",
	"ps4":   "PlayStation 4",
	"xbone": "XBox One",
}

var Activities = map[string]string{
	"incursions":  "Incursiones",
	"dailies":     "Misiones Diarias",
	"assignments": "Asignaciones",
	"darkzone":    "Zona Oscura",
	"freemode":    "Modo Libre",
	"storymode":   "Modo Historia",
}

func GetPlatformKeys() []string {
	return getKeys(Platforms)
}

func GetActivitiesKeys() []string {
	return getKeys(Activities)
}

func getKeys(origin map[string]string) []string {
	var keys []string

	for key := range origin {
		keys = append(keys, key)
	}

	return keys
}
