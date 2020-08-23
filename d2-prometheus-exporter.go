package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cmacrae/d2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	historicalStatsData d2.CommonAccountStats
	characterProfile    d2.GetProfileCharactersResponse
	userAPIKey          string
	playerPlatform      string
	playerName          string

	raceIdentifiers = map[int]string{
		0: "Human",
		1: "Awoken",
		2: "Exo",
		3: "Unknown",
	}

	genderIdentifiers = map[int]string{
		0: "Male",
		1: "Female",
		2: "Unknown",
	}

	classIdentifiers = map[int]string{
		0: "Titan",
		1: "Hunter",
		2: "Warlock",
		3: "Unknown",
	}

	platformIdentifiers = map[string]string{
		"all":         "All",
		"none":        "None",
		"xbox":        "TigerXbox",
		"playstation": "TigerPsn",
		"steam":       "TigerSteam",
		"blizzard":    "TigerBlizzard",
		"stadia":      "TigerStadia",
		"demon":       "TigerDemon",
		"bungienext":  "BungieNext",
	}
)

func init() {
	if err := getConfig(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	go func() {
		if err := serveMetrics(3000); err != nil {
			log.Printf("Unable to serve metric: %v\n", err)
		}
	}()
}

func main() {
	api := d2.NewClient(nil, userAPIKey)
	accountData, err := api.Player.SearchDestinyPlayer(playerPlatform, playerName)
	if err != nil {
		log.Println(err)
	}

	// TODO: Rather than run every minute, just run on scrape?
	for range time.NewTicker(1 * time.Minute).C {
		characterProfile, err = api.Player.GetProfileCharacters(playerPlatform, accountData[0].MembershipID)
		if err != nil {
			log.Println(err)
		}

		historicalStatsData, err = api.Player.GetHistoricalStatsForAccount(playerPlatform, accountData[0].MembershipID)
		if err != nil {
			log.Println(err)
		}

		updateMetrics()
	}
}

func serveMetrics(port int) error {
	prometheus.MustRegister(kills)
	prometheus.MustRegister(deaths)
	prometheus.MustRegister(assists)
	prometheus.MustRegister(killDeathRatio)
	prometheus.MustRegister(killDeathAssists)
	prometheus.MustRegister(secondsPlayed)
	prometheus.MustRegister(activitiesCleared)
	prometheus.MustRegister(activitiesEntered)
	prometheus.MustRegister(killsWith)
	prometheus.MustRegister(precisionKills)
	prometheus.MustRegister(resurrections)
	prometheus.MustRegister(suicides)
	prometheus.MustRegister(orbs)

	http.HandleFunc("/healthz", healthz)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	return nil
}

func getEnv(env string) (string, error) {
	v := os.Getenv(env)
	if v == "" {
		return "", fmt.Errorf("%s variable is empty", env)
	}
	return v, nil
}

func getConfig() error {
	var err error
	userAPIKey, err = getEnv("BNET_API_KEY")
	if err != nil {
		return err
	}

	plPlat, err := getEnv("PLAYER_PLATFORM")
	if err != nil {
		return err
	}

	playerPlatform, err = playerPlatformValidator(strings.ToLower(plPlat))
	if err != nil {
		return err
	}

	playerName, err = getEnv("PLAYER_NAME")
	if err != nil {
		return err
	}

	return nil
}

func playerPlatformValidator(p string) (string, error) {
	for k, v := range platformIdentifiers {
		if k == p {
			return v, nil
		}
	}
	return "", fmt.Errorf("%s is not a valid platform", p)
}

func healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
