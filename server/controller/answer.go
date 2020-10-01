package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/mattermost/mattermost-server/model"
	"github.com/techbot/server/config"
	"github.com/techbot/server/techbuzz"
)

var postAnswer = &Endpoint{
	Path:    "/submit-answer",
	Execute: submitAnswer,
}

var sendAnswer = &Endpoint{
	Path:    "/send-answer",
	Execute: sentAnswerToUser,
}

func sentAnswerToUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		config.Mattermost.LogError("Unable to read the request body.", "Error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var feedback map[string]interface{}
	if err = json.Unmarshal(b, &feedback); err != nil {
		config.Mattermost.LogError("Unable to unmarshal response.", "Error", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	questionText := feedback["state"].(string)
	responseBy := feedback["user_id"].(string)
	answer := feedback["submission"].(map[string]interface{})["Answer"].(string)
	user, _ := config.Mattermost.GetUser(responseBy)
	channel, _ := config.Mattermost.GetDirectChannel(config.GetConfig().BotUserID, userID)
	post := &model.Post{
		ChannelId: channel.Id,
		UserId:    config.GetConfig().BotUserID,
		Message:   "Hi @" + user.Username + " responded to your query :smile:",
	}
	post.AddProp("attachments", []*model.SlackAttachment{
		{
			Text: fmt.Sprintf("**Q.** %s \n\n**A.** %s", questionText, answer),
		},
	})
	config.Mattermost.CreatePost(post)

	post1 := &model.Post{
		ChannelId: config.GetConfig().AskJtgChannel,
		UserId:    config.GetConfig().BotUserID,
		Message:   "Hi @" + user.Username + " responded to a query",
	}
	post1.AddProp("attachments", []*model.SlackAttachment{
		{
			Text: fmt.Sprintf("**Q.** %s \n\n**A.** %s", questionText, answer),
		},
	})
	config.Mattermost.CreatePost(post1)
	channelDM, _ := config.Mattermost.GetDirectChannel(config.GetConfig().BotUserID, responseBy)
	config.Mattermost.SendEphemeralPost(responseBy, &model.Post{
		ChannelId: channelDM.Id,
		Message:   " Keeping knowledge erodes power. Sharing is the fuel to your growth engine :wink:.",
	})
}
func submitAnswer(w http.ResponseWriter, r *http.Request) {
	questionID := r.URL.Query().Get("id")
	userID := r.URL.Query().Get("user_id")
	id, _ := strconv.Atoi(questionID)
	question := techbuzz.GetQuestionByID(id)

	decoder := json.NewDecoder(r.Body)
	params := &model.PostActionIntegrationRequest{}

	if err := decoder.Decode(&params); err != nil {
		config.Mattermost.LogError("Error decoding PostActionIntegrationRequest params: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dialogRequest := model.OpenDialogRequest{
		TriggerId: params.TriggerId,
		URL:       fmt.Sprintf("%s/plugins/%s/%s?user_id=%s", *config.Mattermost.GetConfig().ServiceSettings.SiteURL, config.PluginName, "send-answer", userID),
		Dialog: model.Dialog{
			Title: "Share Your Answer",
			Elements: []model.DialogElement{{
				DisplayName: "Answer",
				Name:        "Answer",
				Type:        "textarea",
				Placeholder: "Please enter your answer.",
			}},
			State: question,
		},
	}
	config.Mattermost.OpenInteractiveDialog(dialogRequest)

	response := &model.PostActionIntegrationResponse{}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(response.ToJson()); err != nil {
		config.Mattermost.LogWarn("failed to write PostActionIntegrationResponse", "Error", err.Error())
	}
}
