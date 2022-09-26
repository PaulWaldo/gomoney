package main

import (
	"flag"

	"github.com/PaulWaldo/gomoney/internal/application"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

func main() {
	var sample, inMem bool
	flag.BoolVar(&sample, "sampleData", false, "Whether to load sample data or not")
	flag.BoolVar(&inMem, "inMem", false, "Whether to use In-Memory database or not")
	flag.Parse()

	application.RunApp(&application.AppData{Accounts: []models.Account{}, LoadSampleData: sample, InMemDatabase: inMem})
}
