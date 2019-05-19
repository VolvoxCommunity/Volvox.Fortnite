package api

import (
	"encoding/json"
	"errors"
	"github.com/ewohltman/pool"
	"github.com/volvoxcommunity/volvox.fortnite/src/logging"
	"io/ioutil"
	"net/http"
	"time"
)

/**

 * Created by cxnky on 25/04/2019 at 00:13
 * api
 * https://github.com/cxnky/

**/

type Stats struct {
	LifeTimeStats []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"lifeTimeStats"`
}

var APIKey = ""
var pooledClient = pool.NewPClient(&http.Client{Timeout: 5 * time.Second}, 1, 1)

func FetchLifetimeWins(username, platform string) (wins string, err error) {
	urlString := "https://api.fortnitetracker.com/v1/profile/" + platform + "/" + username
	req, err := http.NewRequest("GET", urlString, nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("TRN-Api-Key", APIKey)
	resp, err := pooledClient.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", nil
	}

	var stats Stats
	err = json.Unmarshal(bytes, &stats)

	if err != nil {
		return "", err
	}

	logging.Log.Debug(username, "wins", stats)

	for _, v := range stats.LifeTimeStats {
		if v.Key == "Wins" {
			return v.Value, nil
		}
	}

	return "", errors.New("unable to find any stats for that user")
}
