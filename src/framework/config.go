package framework

import (
	"encoding/json"
	"io/ioutil"
)

/**

 * Created by cxnky on 24/04/2019 at 15:13
 * volvox_fortnite
 * https://github.com/cxnky/

**/

// Configuration is the object into which the config.json file will be read
type Configuration struct {
	Token      string `json:"token"`
	ClientID   string `json:"client_id"`
	Prefix     string `json:"prefix"`
	LogChannel string `json:"log_channel"`
	TRNAPIKey  string `json:"trn_api_key"`
	GuildID    string `json:"guild_id"`
	Roles      Roles  `json:"roles"`
}

type Roles struct {
	Pc    string `json:"pc"`
	Ps4   string `json:"ps4"`
	Xbox  string `json:"xbox"`
	Tier1 string `json:"tier1"`
	Tier2 string `json:"tier2"`
	Tier3 string `json:"tier3"`
	Tier4 string `json:"tier4"`
	Tier5 string `json:"tier5"`
}

// ReadConfig attempts to parse the JSON in config.json
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
