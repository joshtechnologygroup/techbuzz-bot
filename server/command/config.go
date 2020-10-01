package command

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/techbot/server/techbuzz"
)

func commandGetConfig() *Config {
	return &Config{
		Command: &model.Command{
			Trigger:          "config",
			AutoCompleteDesc: "config of tech post",
			AutoComplete:     true,
		},
		HelpText: "",
		Validate: validateGetConfig,
		Execute:  getConfig,
	}
}

func validateGetConfig(args []string, context Context) (*model.CommandResponse, *model.AppError) {
	return nil, nil
}

func getConfig(args []string, context Context) (*model.CommandResponse, *model.AppError) {
	userID := context.CommandArgs.UserId
	config := techbuzz.GetUserConfig(userID)
	if config.Enabled == false {
		return &model.CommandResponse{
			Type: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text: "You have not subscribed yet.",
		}, nil
	}
	text := "You have subscribed these tags:"
	Tags := config.Tags
	for key, val := range Tags {
		if val.Enabled {
			text = text + " " + key
		}
	}
	return &model.CommandResponse{
		Type: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
		Text: text,
	}, nil
}
