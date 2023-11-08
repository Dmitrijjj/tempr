package servis

import (
	"std/github.com/tempr/repository"
)

func PogodaMainTemp(CName string) float64 {
	resp := repository.PogodaMain(CName)

	return resp.Temp
}
