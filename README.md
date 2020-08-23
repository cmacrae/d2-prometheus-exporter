# Destiny 2 Prometheus Exporter
A [Prometheus](https://prometheus.io/) exporter for the [Destiny 2](https://bungie.net/) API: because "why not"?

## About
Bungie provide a rich, treasure trove of data via their public API for Destiny 2. Exposing this data via Prometheus metrics allows
for really easy dashboard building with [Grafana](https://grafana.com/), complex mathematical functions over the data, and
just anything else you can do with metrics... you could even set up alerts when certain conditions are observed

![](https://i.imgur.com/MDAPtKC.gif)

## Usage
There are 3 environment variables the `d2-prometheus-exporter` requires to function:
| Variable          | Description                                                                            |
|:------------------|:---------------------------------------------------------------------------------------|
| `BNET_API_KEY`    | Your bungie.net API key                                                                |
| `PLAYER_PLATFORM` | The platform to use for player lookup (stats are account-wide, if you use cross-save)  |
| `PLAYER_NAME`     | The player name (gamertag, PSN ID, etc.) to lookup                                     |

_**Note** `PLAYER_PLATFORM` must be one of "xbox", "playstation", "steam", "blizzard", "stadia"_

## Currently Implemented Metrics
| Metric                | Labels                           | Description                                                |
|:----------------------|:---------------------------------|:-----------------------------------------------------------|
| d2_total_kills        | activity_type, character         | Total kill count                                           |
| d2_total_deaths       | activity_type, character         | Total death count                                          |
| d2_total_assists      | activity_type, character         | Total assist count                                         |
| d2_kd                 | activity_type, character         | Kill/Death ratio                                           |
| d2_kda                | activity_type, character         | Kill/Death/Assist ratio                                    |
| d2_seconds_played     | activity_type, character         | Total seconds played                                       |
| d2_activities_cleared | None                             | Total activities entered                                   |
| d2_activities_entered | None                             | Total activities cleared                                   |
|  d2_kills_with        | activity_type, character, weapon | Total kills, partitioned by weapon                         |
|  d2_precision_kills   | activity_type, character         | Total precision kills                                      |
|  d2_resurrections     | activity_type, character, action | Total resurrections, partitioned by "performed"/"received" |
|  d2_suicides          | activity_type, character         | Total suicides                                             |
|  d2_orbs              | activity_type, character, action | Total orbs, partitioned by "gathered"/"dropped"            |

_**Note** The `activity_type` label is either `pve` or `pvp`_  
_**Note** The `character` label is currently denoted by class (`hunter`, `warlock`, or `titan`). This may prove troublesome if you have more than one character of the same class. In the future, this may be denoted by a combination of class, race, and gender_  
_**Note** The `character` label also has a special value: `merged`, which represents all statistics across characters_  

## Example Prometheus Queries
#### Activity clear rate
Represents the your overall activity clear rate as a percentage:  
```
round((d2_activities_cleared / d2_activities_entered) * 100)
```

#### Percent of time played as character
Hunter in PvP, for example:  
```
(d2_seconds_played{activity_type="pvp", character="hunter"} / ignoring (character) group_left d2_seconds_played{activity_type="pvp", character="merged"}) * 100
```

## Local Development/Sandbox
Included in the `sandbox/` directory is a [Docker Compose](https://docs.docker.com/compose/) environment, which will spin up:  
- d2-prometheus-exporter: `localhost:3000/metrics`
- [Prometheus](https://prometheus.io/): `localhost:9090`
- [Alertmanager](https://prometheus.io/docs/alerting/latest/alertmanager/): `localhost:9093`
- [Grafana](https://grafana.com/): `localhost:3001`

This will build the exporter and its container, so can be used for local development. Or, simply as an easy way to play around.  
You'll need to ensure the [required environment variables](#usage) are set. Docker Compose will inherit them.

You can log into Grafana with **username** `admin`, **password** `destiny`.  
An example dashboard is included - you can copy this and play around to build your own - feel free to share any!

## Docker & Kubernetes
A [Docker](https://docker.com/) image is available at [cmacrae/d2-prometheus-exporter](https://hub.docker.com/repository/docker/cmacrae/d2-prometheus-exporter) on Docker Hub.  

You can also find a [Helm](https://helm.sh/) Chart (`d2-prometheus-exporter`) for deployment to Kubernetes in my Helm repository: https://charts.cmacr.ae
