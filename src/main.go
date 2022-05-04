package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

func newSolution() solution {
	return solution{}
}

func (app *solution) parseConfig() error {
	// parse config file
	if _, err := os.Stat(configJSONPath); os.IsNotExist(err) {
		return errors.New("failed to find config file")
	}

	jsonBytes, err := ioutil.ReadFile(configJSONPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonBytes, &app.Config)
}

func (app *solution) runInBackground() {
	forever := make(chan struct{})
	// run in background
	<-forever
}

func main() {
	app := newSolution()

	err := checkErrors(
		app.parseConfig,
		app.runDaemon,
		app.blockchainConnect,
		app.setupFront,
		app.openFront,
	)
	if err != nil {
		log.Fatalln(err)
	}

	app.runInBackground()
}
