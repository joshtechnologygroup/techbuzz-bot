package techbuzz

import (
	"encoding/json"
	"github.com/thoas/go-funk"

	"github.com/standup-raven/standup-raven/server/logger"
	"github.com/standup-raven/standup-raven/server/util"
	"github.com/techbot/server/config"
)

var TechTag = map[string]bool{
	"python":     true,
	"javascript": true,
	"java":       true,
	"ruby":       true,
	"php":        true,
	"other":      true,
}

var TechList = []string {
	"python",
	"javascript",
	"java",
	"ruby",
	"php",
	"other",
}

type Tag struct {
	SequenceNumber int  `json:"sno"`
	Enabled        bool `json:"enabled"`
}

type UserConfig struct {
	Enabled bool           `json:"enabled"`
	Tags    map[string]Tag `json:"tags"`
}

func Unsubscribe(userID string, tags []string) error {
	userConfig := GetUserConfig(userID)
	if userConfig != nil {
		for _, tag := range tags {
			if val, ok := userConfig.Tags[tag]; ok {
				val.Enabled = false
				userConfig.Tags[tag] = val
			}
		}
	}
	if err := SaveConfig(userID, userConfig); err != nil {
		return err
	}

	return nil
}

func SaveUserConfig(userID string, tags []string) error {
	userConfig := GetUserConfig(userID)
	if userConfig == nil {
		if err := AddTechMembers(userID); err != nil {
			return err
		}
		tags = funk.UniqString(tags)
		var techTags = make(map[string]Tag)
		for _, tag := range tags {
			techpost := Tag{
				SequenceNumber: 0,
				Enabled:        true,
			}
			techTags[tag] = techpost
		}
		userConfig = &UserConfig{
			Enabled: true,
			Tags:    techTags,
		}
	} else {
		userConfig.Enabled = true
		if userConfig.Tags == nil {
			userConfig.Tags = make(map[string]Tag)
		}
		for _, tag := range tags {
			if val, ok := userConfig.Tags[tag]; ok {
				val.Enabled = true
				userConfig.Tags[tag] = val
			} else {
				userConfig.Tags[tag] = Tag{
					Enabled:        true,
					SequenceNumber: 0,
				}
			}
		}
	}
	if err := SaveConfig(userID, userConfig); err != nil {
		return err
	}

	return nil
}

func GetTechMembers() []string {
	data, _ := config.Mattermost.KVGet(util.GetKeyHash(config.TechMembers))
	var users []string
	json.Unmarshal(data, &users)
	return users
}

func AddTechMembers(userID string) error {
	users := GetTechMembers()
	users = append(users, userID)

	serilizedData, err := json.Marshal(users)
	if err != nil {
		logger.Error("Couldn't marshal config", err, nil)
		return err
	}

	if err := config.Mattermost.KVSet(util.GetKeyHash(config.TechMembers), serilizedData); err != nil {
		return err
	}
	return nil
}

func GetUserConfig(userID string) *UserConfig {
	data, _ := config.Mattermost.KVGet(util.GetKeyHash(config.UserConfig + "_" + userID))
	if len(data) == 0 {
		return nil
	}
	var userConfig *UserConfig
	if len(data) > 0 {
		userConfig = &UserConfig{}
		json.Unmarshal(data, &userConfig)
	}
	return userConfig
}

func SaveConfig(userID string, userConfig *UserConfig) error {
	serilizedData, err := json.Marshal(userConfig)
	if err != nil {
		logger.Error("Couldn't marshal config", err, nil)
		return err
	}

	if err := config.Mattermost.KVSet(util.GetKeyHash(config.UserConfig+"_"+userID), serilizedData); err != nil {
		return err
	}
	return nil
}

func GetData(tag string) []string {
	data, _ := config.Mattermost.KVGet(util.GetKeyHash(config.TechData + "_" + tag))
	var techData []string
	json.Unmarshal(data, &techData)
	return techData
}

func InsertData(tag string, text string) error {
	techData := GetData(tag)
	techData = append(techData, text)
	serilizedData, err := json.Marshal(techData)
	if err != nil {
		logger.Error("Couldn't marshal tech data", err, nil)
		return err
	}

	if err := config.Mattermost.KVSet(util.GetKeyHash(config.TechData+"_"+tag), serilizedData); err != nil {
		return err
	}
	return nil
}
func GetQuestionByID(questionID int) string {
	techQuestions := Getquestions()
	return techQuestions[questionID-1]
}

func Getquestions() []string {
	data, _ := config.Mattermost.KVGet(util.GetKeyHash(config.TechQuestions))
	var techQuestions []string
	json.Unmarshal(data, &techQuestions)
	return techQuestions
}

func AddQuestion(text string) int {
	techQuestions := Getquestions()
	techQuestions = append(techQuestions, text)
	serilizedData, _ := json.Marshal(techQuestions)

	config.Mattermost.KVSet(util.GetKeyHash(config.TechQuestions), serilizedData)
	return len(techQuestions)
}
