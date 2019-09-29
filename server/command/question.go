package command

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/techbot/server/techbuzz"
	"github.com/techbot/server/util"
)

func commandAskQuestion() *Config {
	return &Config{
		Command: &model.Command{
			Trigger:          "question",
			AutoComplete:     true,
			AutoCompleteDesc: "Ask a question.",
		},
		HelpText: "",
		Validate: validatequestion,
		Execute:  askQuestion,
	}
}

func validatequestion(args []string, context Context) (*model.CommandResponse, *model.AppError) {
	if len(args) < 2 {
		return util.SendEphemeralText("Please specify both tag and question")
	}

	context.Props["tag"] = args[0]
	context.Props["question"] = args[1]

	return nil, nil
}

func askQuestion(args []string, context Context) (*model.CommandResponse, *model.AppError) {
	userID := context.CommandArgs.UserId
	tag := context.Props["tag"].(string)
	question := context.Props["question"].(string)
	questionID := techbuzz.AddQuestion(question)
	memberIDs := techbuzz.GetTagMemberIDs(tag)
	techbuzz.PostQuestion(memberIDs, question, userID, questionID)
	return &model.CommandResponse{
		Type: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
		Text: "Creativity flows when curiosity is stoked :smile:.",
	}, nil
}
