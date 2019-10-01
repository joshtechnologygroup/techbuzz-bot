package config

import (
	"github.com/pkg/errors"
	"strings"
	"time"

	"github.com/mattermost/mattermost-server/plugin"
	"go.uber.org/atomic"
)

const (
	CommandPrefix       = "techbot"
	TechMembers         = "tech_members1"
	UserConfig          = "user1_config"
	TechData            = "tech_data1"
	TechQuestions       = "tech_question1"
	URLMappingKeyPrefix = "url_"
	BotUsername         = "techbot"
	BotDisplayName      = "TechBot"

	URLPluginBase  = "/plugins/" + "techbot"
	URLStaticBase  = URLPluginBase + "/static"
	RunnerInterval = 60 * time.Second

	HeaderMattermostUserID = "Mattermost-User-Id"
)

var (
	config     atomic.Value
	Mattermost plugin.API
)

type Configuration struct {
	SiteURL   string `json:"SiteURL"`
	BotUserID string `json:"botUserId"`
	Apikey    string `json:"Apikey"`
	TechBuzzChannel string `json:"TechBuzzChannel"`
	AskJtgChannel string `json:"AskJtgChannel"`
}

func GetConfig() *Configuration {
	return config.Load().(*Configuration)
}

func SetConfig(c *Configuration) {
	config.Store(c)
}

func (c *Configuration) ProcessConfiguration() error {
	c.Apikey = strings.TrimSpace(c.Apikey)
	c.TechBuzzChannel = strings.TrimSpace(c.TechBuzzChannel)

	return nil
}

func (c *Configuration) IsValid() error {
	if c.TechBuzzChannel == "" {
		return errors.New("TechBuzz Channel ID cannot be empty")
	}

	if _, appErr := Mattermost.GetChannel(c.TechBuzzChannel); appErr != nil {
		return errors.Wrap(errors.New(appErr.Error()), "invalid TechBuzz Channel ID")
	}

	return nil
}
