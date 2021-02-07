package entity

type Reputation struct {
	Level int `json:"level"`
	Stars int `json:"stars"`
}

func (rep *Reputation) UpLevel(scoreboard Scoreboard, levelThresholdInfo []LevelThresholdInfo) (shouldIncreaseLevel bool, shouldIncreaseStar bool) {

	var currentLevelThreshold LevelThresholdInfo
	for _, levelInfo := range levelThresholdInfo {
		if levelInfo.Level == rep.Level {
			currentLevelThreshold = levelInfo
			break
		}
	}

	shouldIncreaseStar = (scoreboard.HitsAmount % currentLevelThreshold.HitsPerStar) == 0

	if shouldIncreaseStar {
		rep.addNewStar()

		isNotLastLevel := currentLevelThreshold.Level < len(levelThresholdInfo)
		shouldIncreaseLevel = rep.Stars >= currentLevelThreshold.TotalStarsToNextLevel && isNotLastLevel

		if shouldIncreaseLevel {
			rep.upgradeToNextLevel()
		}
	}

	return
}

func (rep *Reputation) addNewStar() {
	rep.Stars++
}

func (rep *Reputation) upgradeToNextLevel() {
	rep.Level++
	rep.Stars = 0
}
