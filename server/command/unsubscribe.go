package command

import (
	"strings"

	"github.com/mattermost/mattermost-server/model"
	"github.com/techbot/server/techbuzz"
)

func commandUnsubscribeTopics() *Config {
	return &Config{
		Command: &model.Command{
			Trigger:          "unsubscribe",
			AutoComplete:     true,
			AutoCompleteDesc: "unsubscribe to tech post.",
			AutoCompleteHint: "<tags...>",
		},
		HelpText: "",
		Validate: validate,
		Execute:  unsubscribe,
	}
}

func validate(args []string, context Context) (*model.CommandResponse, *model.AppError) {
	tagsNotFound := []string{}
	tags := []string{}
	for _, arg := range args {
		if techbuzz.TechTag[strings.ToLower(arg)] {
			tags = append(tags, strings.ToLower(arg))
		} else {
			tagsNotFound = append(tagsNotFound, strings.ToLower(arg))
		}
	}
	context.Props["tags"] = tags
	context.Props["tagsNotFound"] = tagsNotFound

	return nil, nil
}

func unsubscribe(args []string, context Context) (*model.CommandResponse, *model.AppError) {
	userID := context.CommandArgs.UserId
	if len(args) == 0 {
		userConfig := techbuzz.GetUserConfig(userID)
		userConfig.Enabled = false

		techbuzz.SaveConfig(userID, userConfig)
		techbuzz.Unsubscribe(userID, techbuzz.TechList)
		return &model.CommandResponse{
			Type: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text: "Successfully unsubscribed from all tech post",
		}, nil
	}

	var tt, te string
	tags := context.Props["tags"].([]string)
	tagsNotFound := context.Props["tagsNotFound"].([]string)
	techbuzz.Unsubscribe(userID, tags)
	for _, val := range tags {
		te = te + " " + val
	}
	if len(tagsNotFound) == 0 {
		return &model.CommandResponse{
			Type: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text: "Successfully unsubscribed to post having the tags:" + te,
		}, nil
	}
	if len(tags) != 0 {
		tt = "Successfully unsubscribed to post having tags:"
		tt = tt + te + "\n"
	}

	tt = tt + "Invalid tags :"
	for _, val := range tagsNotFound {
		tt = tt + " " + val
	}

	return &model.CommandResponse{
		Type: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
		Text: tt,
	}, nil
}
