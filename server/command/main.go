package command

import (
	"fmt"

	"github.com/mattermost/mattermost-server/model"
)

type Context struct {
	CommandArgs *model.CommandArgs
	Props       map[string]interface{}
}

type Config struct {
	Command  *model.Command
	HelpText string
	Execute  func([]string, Context) (*model.CommandResponse, *model.AppError)
	Validate func([]string, Context) (*model.CommandResponse, *model.AppError)
}

func (c *Config) Syntax() string {
	return fmt.Sprintf("/%s %s", c.Command.Trigger, c.Command.AutoCompleteHint)
}

var commands = map[string]*Config{
	commandSubscribeTopics().Command.Trigger:   commandSubscribeTopics(),
	commandUnsubscribeTopics().Command.Trigger: commandUnsubscribeTopics(),
	commandGetConfig().Command.Trigger:         commandGetConfig(),
	commandInsertData().Command.Trigger:        commandInsertData(),
	commandAskQuestion().Command.Trigger:       commandAskQuestion(),
}
