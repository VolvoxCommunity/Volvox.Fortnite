package main

import (
	"encoding/json"
	"io/ioutil"
)

/**

 * Created by cxnky on 24/04/2019 at 15:13
 * volvox_fortnite
 * https://github.com/cxnky/

**/

type Configuration struct {
	Token    string `json:"token"`
	ClientID string `json:"client_id"`
}

func ReadConfig() Configuration {
	bytes, err := ioutil.ReadFile("config.json")

	if err != nil {
		panic("could not read config.json. does it exist?")
	}

	var config Configuration
	err = json.Unmarshal(bytes, &config)

	if err != nil {
		panic("unable to parse the JSON in config.json: " + err.Error())
	}

	return config

}
