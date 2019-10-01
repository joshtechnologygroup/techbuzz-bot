package command

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/techbot/server/techbuzz"
)

func commandInsertData() *Config {
	return &Config{
		Command: &model.Command{
			Trigger:          "data",
			AutoCompleteDesc: "config of tech post",
			AutoComplete:     true,
		},
		HelpText: "",
		Validate: validatedata,
		Execute:  insertData,
	}
}

func validatedata(args []string, context Context) (*model.CommandResponse, *model.AppError) {
	return nil, nil
}

func insertData(args []string, context Context) (*model.CommandResponse, *model.AppError) {
	techbuzz.InsertData(args[0], args[1])
	return &model.CommandResponse{
		Type: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
		Text: "-",
	}, nil
}
