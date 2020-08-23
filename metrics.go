package main

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	commonLabels = []string{"activity_type", "character"}

	kills = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "d2_total_kills",
			Help: "Total count for kills",
		},
		commonLabels,
	)

	deaths = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "d2_total_deaths",
			Help: "Total count for deaths",
		},
		commonLabels,
	)

	assists = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "d2_total_assists",
			Help: "Total count for assists",
		},
		commonLabels,
	)

	killDeathRatio = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "d2_kd",
			Help: "Kill/Death ratio",
		},
		commonLabels,
	)

	killDeathAssists = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "d2_kda",
			Help: "Kill/Death/Assists",
		},
		commonLabels,
	)

	secondsPlayed = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "d2_seconds_played",
			Help: "Total seconds played",
		},
		commonLabels,
	)

	activitiesCleared = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "d2_activities_cleared",
			Help: "Total PvE activities cleared",
		},
	)

	activitiesEntered = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "d2_activities_entered",
			Help: "Total PvE activities entered",
		},
	)

	killsWith = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "d2_kills_with",
			Help: "Kill statistics based on weapon",
		},
		append(commonLabels, "weapon"),
	)

	precisionKills = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "d2_precision_kills",
			Help: "Total precision kills",
		},
		commonLabels,
	)

	resurrections = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "d2_resurrections",
			Help: "Total resurrections",
		},
		append(commonLabels, "action"),
	)

	suicides = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "d2_suicides",
			Help: "Total suicides",
		},
		commonLabels,
	)

	orbs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "d2_orbs",
			Help: "Total orbs dropped or gathered",
		},
		append(commonLabels, "action"),
	)
)

func updateMetrics() {
	mergedAllPvEDataAllTime := historicalStatsData.
		MergedAllCharacters.
		Results.
		AllPvE.
		AllTime

	mergedAllPvPDataAllTime := historicalStatsData.
		MergedAllCharacters.
		Results.
		AllPvP.
		AllTime

	//
	//  Total Kills
	//
	// Merged
	kills.WithLabelValues("pve", "merged").Set(mergedAllPvEDataAllTime.Kills.Basic.Value)
	kills.WithLabelValues("pvp", "merged").Set(mergedAllPvPDataAllTime.Kills.Basic.Value)

	// Per character
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		kills.WithLabelValues("pve", class).Set(allTimeResults.Kills.Basic.Value)
	}

	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		kills.WithLabelValues("pvp", class).Set(allTimeResults.Kills.Basic.Value)
	}

	//
	//  Total Deaths
	//
	// Merged
	deaths.WithLabelValues("pve", "merged").Set(mergedAllPvEDataAllTime.Deaths.Basic.Value)
	deaths.WithLabelValues("pvp", "merged").Set(mergedAllPvPDataAllTime.Deaths.Basic.Value)

	// Per character
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		deaths.WithLabelValues("pve", class).Set(allTimeResults.Deaths.Basic.Value)
	}

	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		deaths.WithLabelValues("pvp", class).Set(allTimeResults.Deaths.Basic.Value)
	}

	//
	//  Total Assists
	//
	// Merged
	assists.WithLabelValues("pve", "merged").Set(mergedAllPvEDataAllTime.Assists.Basic.Value)
	assists.WithLabelValues("pvp", "merged").Set(mergedAllPvPDataAllTime.Assists.Basic.Value)

	// Per character
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		assists.WithLabelValues("pve", class).Set(allTimeResults.Assists.Basic.Value)
	}

	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		assists.WithLabelValues("pvp", class).Set(allTimeResults.Assists.Basic.Value)
	}

	//
	//  KD & KDA
	//
	// Merged
	killDeathRatio.WithLabelValues("pve", "merged").Set(mergedAllPvEDataAllTime.KillsDeathsRatio.Basic.Value)
	killDeathRatio.WithLabelValues("pvp", "merged").Set(mergedAllPvPDataAllTime.KillsDeathsRatio.Basic.Value)
	killDeathAssists.WithLabelValues("pve", "merged").Set(mergedAllPvEDataAllTime.KillsDeathsAssists.Basic.Value)
	killDeathAssists.WithLabelValues("pvp", "merged").Set(mergedAllPvPDataAllTime.KillsDeathsAssists.Basic.Value)

	// Per character
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killDeathRatio.WithLabelValues("pve", class).Set(allTimeResults.KillsDeathsRatio.Basic.Value)
	}

	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killDeathRatio.WithLabelValues("pvp", class).Set(allTimeResults.KillsDeathsRatio.Basic.Value)
	}

	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killDeathAssists.WithLabelValues("pve", class).Set(allTimeResults.KillsDeathsAssists.Basic.Value)
	}

	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killDeathAssists.WithLabelValues("pvp", class).Set(allTimeResults.KillsDeathsAssists.Basic.Value)
	}

	//
	//  Seconds Played
	//
	// Merged
	// TODO: Fix bug here where time displays weird
	secondsPlayed.WithLabelValues("pve", "merged").Set(mergedAllPvEDataAllTime.SecondsPlayed.Basic.Value)
	secondsPlayed.WithLabelValues("pvp", "merged").Set(mergedAllPvPDataAllTime.SecondsPlayed.Basic.Value)

	// Per character
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		secondsPlayed.WithLabelValues("pve", class).Set(allTimeResults.SecondsPlayed.Basic.Value)
	}

	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		secondsPlayed.WithLabelValues("pvp", class).Set(allTimeResults.SecondsPlayed.Basic.Value)
	}

	//
	//  Activities
	//
	activitiesCleared.Set(mergedAllPvEDataAllTime.ActivitiesCleared.Basic.Value)
	activitiesEntered.Set(mergedAllPvEDataAllTime.ActivitiesEntered.Basic.Value)

	//
	//  Kills With (PvE)
	//
	killsWith.WithLabelValues("pve", "merged", "super").Set(mergedAllPvEDataAllTime.WeaponKillsSuper.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pve", class, "super").Set(allTimeResults.WeaponKillsSuper.Basic.Value)
	}

	killsWith.WithLabelValues("pve", "merged", "melee").Set(mergedAllPvEDataAllTime.WeaponKillsMelee.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pve", class, "melee").Set(allTimeResults.WeaponKillsMelee.Basic.Value)
	}

	killsWith.WithLabelValues("pve", "merged", "grenade").Set(mergedAllPvEDataAllTime.WeaponKillsGrenade.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pve", class, "grenade").Set(allTimeResults.WeaponKillsGrenade.Basic.Value)
	}

	killsWith.WithLabelValues("pve", "merged", "ability").Set(mergedAllPvEDataAllTime.WeaponKillsAbility.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pve", class, "ability").Set(allTimeResults.WeaponKillsAbility.Basic.Value)
	}

	killsWith.WithLabelValues("pve", "merged", "autorilfe").Set(mergedAllPvEDataAllTime.WeaponKillsAutoRifle.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pve", class, "autorifle").Set(allTimeResults.WeaponKillsAutoRifle.Basic.Value)
	}

	killsWith.WithLabelValues("pve", "merged", "fusionrilfe").Set(mergedAllPvEDataAllTime.WeaponKillsFusionRifle.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pve", class, "fusionrifle").Set(allTimeResults.WeaponKillsFusionRifle.Basic.Value)
	}

	killsWith.WithLabelValues("pve", "merged", "handcannon").Set(mergedAllPvEDataAllTime.WeaponKillsHandCannon.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pve", class, "handcannon").Set(allTimeResults.WeaponKillsHandCannon.Basic.Value)
	}

	killsWith.WithLabelValues("pve", "merged", "machinegun").Set(mergedAllPvEDataAllTime.WeaponKillsMachinegun.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pve", class, "machinegun").Set(allTimeResults.WeaponKillsMachinegun.Basic.Value)
	}

	killsWith.WithLabelValues("pve", "merged", "pulserifle").Set(mergedAllPvEDataAllTime.WeaponKillsPulseRifle.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pve", class, "pulserifle").Set(allTimeResults.WeaponKillsPulseRifle.Basic.Value)
	}

	killsWith.WithLabelValues("pve", "merged", "rocketlauncher").Set(mergedAllPvEDataAllTime.WeaponKillsRocketLauncher.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pve", class, "rocketlauncher").Set(allTimeResults.WeaponKillsRocketLauncher.Basic.Value)
	}

	killsWith.WithLabelValues("pve", "merged", "scoutrifle").Set(mergedAllPvEDataAllTime.WeaponKillsScoutRifle.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pve", class, "scoutrifle").Set(allTimeResults.WeaponKillsScoutRifle.Basic.Value)
	}

	killsWith.WithLabelValues("pve", "merged", "shotgun").Set(mergedAllPvEDataAllTime.WeaponKillsShotgun.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pve", class, "shotgun").Set(allTimeResults.WeaponKillsShotgun.Basic.Value)
	}

	killsWith.WithLabelValues("pve", "merged", "sniper").Set(mergedAllPvEDataAllTime.WeaponKillsSniper.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pve", class, "sniper").Set(allTimeResults.WeaponKillsSniper.Basic.Value)
	}

	killsWith.WithLabelValues("pve", "merged", "submachinegun").Set(mergedAllPvEDataAllTime.WeaponKillsSubmachinegun.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pve", class, "submachinegun").Set(allTimeResults.WeaponKillsSubmachinegun.Basic.Value)
	}

	killsWith.WithLabelValues("pve", "merged", "relic").Set(mergedAllPvEDataAllTime.WeaponKillsRelic.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pve", class, "relic").Set(allTimeResults.WeaponKillsRelic.Basic.Value)
	}

	killsWith.WithLabelValues("pve", "merged", "sidearm").Set(mergedAllPvEDataAllTime.WeaponKillsSideArm.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pve", class, "sidearm").Set(allTimeResults.WeaponKillsSideArm.Basic.Value)
	}

	killsWith.WithLabelValues("pve", "merged", "sword").Set(mergedAllPvEDataAllTime.WeaponKillsSword.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pve", class, "sword").Set(allTimeResults.WeaponKillsSword.Basic.Value)
	}

	//
	//  Kills With (PvP)
	//
	killsWith.WithLabelValues("pvp", "merged", "super").Set(mergedAllPvPDataAllTime.WeaponKillsSuper.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pvp", class, "super").Set(allTimeResults.WeaponKillsSuper.Basic.Value)
	}

	killsWith.WithLabelValues("pvp", "merged", "melee").Set(mergedAllPvPDataAllTime.WeaponKillsMelee.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pvp", class, "melee").Set(allTimeResults.WeaponKillsMelee.Basic.Value)
	}

	killsWith.WithLabelValues("pvp", "merged", "grenade").Set(mergedAllPvPDataAllTime.WeaponKillsGrenade.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pvp", class, "grenade").Set(allTimeResults.WeaponKillsGrenade.Basic.Value)
	}

	killsWith.WithLabelValues("pvp", "merged", "ability").Set(mergedAllPvPDataAllTime.WeaponKillsAbility.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pvp", class, "ability").Set(allTimeResults.WeaponKillsAbility.Basic.Value)
	}

	killsWith.WithLabelValues("pvp", "merged", "autorilfe").Set(mergedAllPvPDataAllTime.WeaponKillsAutoRifle.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pvp", class, "autorifle").Set(allTimeResults.WeaponKillsAutoRifle.Basic.Value)
	}

	killsWith.WithLabelValues("pvp", "merged", "fusionrilfe").Set(mergedAllPvPDataAllTime.WeaponKillsFusionRifle.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pvp", class, "fusionrifle").Set(allTimeResults.WeaponKillsFusionRifle.Basic.Value)
	}

	killsWith.WithLabelValues("pvp", "merged", "handcannon").Set(mergedAllPvPDataAllTime.WeaponKillsHandCannon.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pvp", class, "handcannon").Set(allTimeResults.WeaponKillsHandCannon.Basic.Value)
	}

	killsWith.WithLabelValues("pvp", "merged", "machinegun").Set(mergedAllPvPDataAllTime.WeaponKillsMachinegun.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pvp", class, "machinegun").Set(allTimeResults.WeaponKillsMachinegun.Basic.Value)
	}

	killsWith.WithLabelValues("pvp", "merged", "pulserifle").Set(mergedAllPvPDataAllTime.WeaponKillsPulseRifle.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pvp", class, "pulserifle").Set(allTimeResults.WeaponKillsPulseRifle.Basic.Value)
	}

	killsWith.WithLabelValues("pvp", "merged", "rocketlauncher").Set(mergedAllPvPDataAllTime.WeaponKillsRocketLauncher.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pvp", class, "rocketlauncher").Set(allTimeResults.WeaponKillsRocketLauncher.Basic.Value)
	}

	killsWith.WithLabelValues("pvp", "merged", "scoutrifle").Set(mergedAllPvPDataAllTime.WeaponKillsScoutRifle.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pvp", class, "scoutrifle").Set(allTimeResults.WeaponKillsScoutRifle.Basic.Value)
	}

	killsWith.WithLabelValues("pvp", "merged", "shotgun").Set(mergedAllPvPDataAllTime.WeaponKillsShotgun.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pvp", class, "shotgun").Set(allTimeResults.WeaponKillsShotgun.Basic.Value)
	}

	killsWith.WithLabelValues("pvp", "merged", "sniper").Set(mergedAllPvPDataAllTime.WeaponKillsSniper.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pvp", class, "sniper").Set(allTimeResults.WeaponKillsSniper.Basic.Value)
	}

	killsWith.WithLabelValues("pvp", "merged", "submachinegun").Set(mergedAllPvPDataAllTime.WeaponKillsSubmachinegun.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pvp", class, "submachinegun").Set(allTimeResults.WeaponKillsSubmachinegun.Basic.Value)
	}

	killsWith.WithLabelValues("pvp", "merged", "relic").Set(mergedAllPvPDataAllTime.WeaponKillsRelic.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pvp", class, "relic").Set(allTimeResults.WeaponKillsRelic.Basic.Value)
	}

	killsWith.WithLabelValues("pvp", "merged", "sidearm").Set(mergedAllPvPDataAllTime.WeaponKillsSideArm.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pvp", class, "sidearm").Set(allTimeResults.WeaponKillsSideArm.Basic.Value)
	}

	killsWith.WithLabelValues("pvp", "merged", "sword").Set(mergedAllPvPDataAllTime.WeaponKillsSword.Basic.Value)
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		killsWith.WithLabelValues("pvp", class, "sword").Set(allTimeResults.WeaponKillsSword.Basic.Value)
	}

	//
	//  Precision Kills
	//
	// Merged
	precisionKills.WithLabelValues("pve", "merged").Set(mergedAllPvEDataAllTime.PrecisionKills.Basic.Value)
	precisionKills.WithLabelValues("pvp", "merged").Set(mergedAllPvPDataAllTime.PrecisionKills.Basic.Value)

	// Per character
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		precisionKills.WithLabelValues("pve", class).Set(allTimeResults.PrecisionKills.Basic.Value)
	}

	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		precisionKills.WithLabelValues("pvp", class).Set(allTimeResults.PrecisionKills.Basic.Value)
	}

	//
	//  Resurrections
	//
	// Merged
	resurrections.WithLabelValues("pve", "merged", "performed").Set(mergedAllPvEDataAllTime.ResurrectionsPerformed.Basic.Value)
	resurrections.WithLabelValues("pvp", "merged", "performed").Set(mergedAllPvPDataAllTime.ResurrectionsPerformed.Basic.Value)
	resurrections.WithLabelValues("pve", "merged", "received").Set(mergedAllPvEDataAllTime.ResurrectionsReceived.Basic.Value)
	resurrections.WithLabelValues("pvp", "merged", "received").Set(mergedAllPvPDataAllTime.ResurrectionsReceived.Basic.Value)

	// Per character
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		resurrections.WithLabelValues("pve", class, "performed").Set(allTimeResults.ResurrectionsPerformed.Basic.Value)
	}

	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		resurrections.WithLabelValues("pvp", class, "performed").Set(allTimeResults.ResurrectionsPerformed.Basic.Value)
	}

	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		resurrections.WithLabelValues("pve", class, "received").Set(allTimeResults.ResurrectionsReceived.Basic.Value)
	}

	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		resurrections.WithLabelValues("pvp", class, "received").Set(allTimeResults.ResurrectionsReceived.Basic.Value)
	}

	//
	//  Suicides
	//
	// Merged
	suicides.WithLabelValues("pve", "merged").Set(mergedAllPvEDataAllTime.Suicides.Basic.Value)
	suicides.WithLabelValues("pvp", "merged").Set(mergedAllPvPDataAllTime.Suicides.Basic.Value)

	// Per character
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		suicides.WithLabelValues("pve", class).Set(allTimeResults.Suicides.Basic.Value)
	}

	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		suicides.WithLabelValues("pvp", class).Set(allTimeResults.Suicides.Basic.Value)
	}

	// Orbs
	orbs.WithLabelValues("pve", "merged", "dropped").Set(mergedAllPvEDataAllTime.OrbsDropped.Basic.Value)
	orbs.WithLabelValues("pvp", "merged", "dropped").Set(mergedAllPvPDataAllTime.OrbsDropped.Basic.Value)
	orbs.WithLabelValues("pve", "merged", "gathered").Set(mergedAllPvEDataAllTime.OrbsGathered.Basic.Value)
	orbs.WithLabelValues("pvp", "merged", "gathered").Set(mergedAllPvPDataAllTime.OrbsGathered.Basic.Value)

	// Per character
	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		orbs.WithLabelValues("pve", class, "dropped").Set(allTimeResults.OrbsDropped.Basic.Value)
	}

	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		orbs.WithLabelValues("pvp", class, "dropped").Set(allTimeResults.OrbsDropped.Basic.Value)
	}

	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvE.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		orbs.WithLabelValues("pve", class, "gathered").Set(allTimeResults.OrbsGathered.Basic.Value)
	}

	for _, c := range historicalStatsData.Characters {
		allTimeResults := c.Results.AllPvP.AllTime
		class := strings.ToLower(classIdentifiers[characterProfile.Characters.Data[c.CharacterID].ClassType])
		orbs.WithLabelValues("pvp", class, "gathered").Set(allTimeResults.OrbsGathered.Basic.Value)
	}
}
