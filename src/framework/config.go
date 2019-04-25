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
	TRNAPIKey  string `json:"trn_api_key"`
	PCRole     string `json:"pc_role"`
	PS4Role    string `json:"ps4_role"`
	XboxRole   string `json:"xbox_role"`
	SwitchRole string `json:"switch_role"`
	MobileRole string `json:"mobile_role"`
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
