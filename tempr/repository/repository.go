package repository

import (
	"std/github.com/tempr/models"
)

var Memury = map[string]models.Main{}

func PogodaMain(key string) models.Main {
	return Memury[key]
}
