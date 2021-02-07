package testing

import (
	"testing"

	"github.com/sabidos/core/entity"
)

var levelThresholdInfo = []entity.LevelThresholdInfo{
	{
		Level:                 1,
		HitsPerStar:           30,
		TotalStarsToNextLevel: 5,
	},
	{
		Level:                 2,
		HitsPerStar:           65,
		TotalStarsToNextLevel: 5,
	},
	{
		Level:                 3,
		HitsPerStar:           95,
		TotalStarsToNextLevel: 5,
	},
	{
		Level:                 4,
		HitsPerStar:           120,
		TotalStarsToNextLevel: 5,
	},
}

func TestUpdatePlayerLevelWithNotUpgradeLevel(t *testing.T) {

	scoreboard := entity.Scoreboard{
		Nickname:          "",
		HitsAmount:        12,
		ScoreEndTimestamp: 123456,
	}

	reputation := entity.Reputation{
		Level: 1,
		Stars: 0,
	}

	levelHasBeenUp, starHasBeenUp := reputation.UpLevel(scoreboard, levelThresholdInfo)

	if reputation.Level != 1 {
		t.Errorf("TestUpdatePlayerLevelWithNotUpgradeLevel() failed, expected %v, got %v", 1, reputation.Level)
	}

	if reputation.Stars != 0 {
		t.Errorf("TestUpdatePlayerLevelWithNotUpgradeLevel() failed, expected %v, got %v", 0, reputation.Stars)
	}

	if levelHasBeenUp || starHasBeenUp {
		t.Errorf("TestUpdatePlayerLevelUpgradingToNextLevel() failed, level shouldn`t be raised")
	}

}

func TestUpdatePlayerLevelUpgradingFirstStar(t *testing.T) {

	scoreboard := entity.Scoreboard{
		Nickname:          "",
		HitsAmount:        30,
		ScoreEndTimestamp: 123456,
	}

	reputation := entity.Reputation{
		Level: 1,
		Stars: 0,
	}

	levelHasBeenUp, starHasBeenUp := reputation.UpLevel(scoreboard, levelThresholdInfo)

	if reputation.Level != 1 {
		t.Errorf("TestUpdatePlayerLevelWithNotUpgradeLevel() failed, expected %v, got %v", 1, reputation.Level)
	}

	if reputation.Stars != 1 {
		t.Errorf("TestUpdatePlayerLevelWithNotUpgradeLevel() failed, expected %v, got %v", 0, reputation.Stars)
	}

	if levelHasBeenUp || !starHasBeenUp {
		t.Errorf("TestUpdatePlayerLevelUpgradingToNextLevel() failed, star should be raised")
	}

}

func TestUpdatePlayerLevelUpgradingToNextLevel(t *testing.T) {

	scoreboard := entity.Scoreboard{
		Nickname:          "",
		HitsAmount:        30,
		ScoreEndTimestamp: 123456,
	}

	reputation := entity.Reputation{
		Level: 1,
		Stars: 4,
	}

	levelHasBeenUp, starHasBeenUp := reputation.UpLevel(scoreboard, levelThresholdInfo)

	if reputation.Level != 2 {
		t.Errorf("TestUpdatePlayerLevelUpgradingToNextLevel() failed, expected %v, got %v", 1, reputation.Level)
	}

	if reputation.Stars != 0 {
		t.Errorf("TestUpdatePlayerLevelUpgradingToNextLevel() failed, expected %v, got %v", 0, reputation.Stars)
	}

	if !levelHasBeenUp || !starHasBeenUp {
		t.Errorf("TestUpdatePlayerLevelUpgradingToNextLevel() failed, star should be raised")
	}

}

func TestUpdatePlayerLevelUpgradingLastStar(t *testing.T) {

	scoreboard := entity.Scoreboard{
		Nickname:          "",
		HitsAmount:        120,
		ScoreEndTimestamp: 123456,
	}

	reputation := entity.Reputation{
		Level: 4,
		Stars: 4,
	}

	levelHasBeenUp, starHasBeenUp := reputation.UpLevel(scoreboard, levelThresholdInfo)

	if reputation.Level != 4 {
		t.Errorf("TestUpdatePlayerLevelUpgradingToNextLevel() failed, expected %v, got %v", 4, reputation.Level)
	}

	if reputation.Stars != 5 {
		t.Errorf("TestUpdatePlayerLevelUpgradingToNextLevel() failed, expected %v, got %v", 5, reputation.Stars)
	}

	if levelHasBeenUp || !starHasBeenUp {
		t.Errorf("TestUpdatePlayerLevelUpgradingToNextLevel() failed, level and star should be raised")
	}

}

func TestUpdatePlayerLevelWithoutThreshold(t *testing.T) {

	scoreboard := entity.Scoreboard{
		Nickname:          "",
		HitsAmount:        120,
		ScoreEndTimestamp: 123456,
	}

	reputation := entity.Reputation{
		Level: 4,
		Stars: 5,
	}

	levelHasBeenUp, starHasBeenUp := reputation.UpLevel(scoreboard, levelThresholdInfo)

	if reputation.Level != 4 {
		t.Errorf("TestUpdatePlayerLevelUpgradingToNextLevel() failed, expected %v, got %v", 4, reputation.Level)
	}

	if reputation.Stars != 6 {
		t.Errorf("TestUpdatePlayerLevelUpgradingToNextLevel() failed, expected %v, got %v", 6, reputation.Stars)
	}

	if levelHasBeenUp || !starHasBeenUp {
		t.Errorf("TestUpdatePlayerLevelUpgradingToNextLevel() failed, star should be raised")
	}

}
