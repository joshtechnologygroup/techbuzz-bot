package techbuzz

import (
	"fmt"

	"github.com/mattermost/mattermost-server/model"
	"github.com/techbot/server/config"
)

func SendPost() {
	usersIDs := GetTechMembers()
	fmt.Println("length=", len(usersIDs))
	for _, userID := range usersIDs {
		userConfig := GetUserConfig(userID)
		fmt.Println("userConfig=", userConfig)
		if userConfig.Enabled == false {
			continue
		} else {
			var newTags = make(map[string]Tag)
			for key, value := range userConfig.Tags {
				if value.Enabled == true {
					if sendTechPost(key,userID,value.SequenceNumber) {
						value.SequenceNumber++
					}
				}
				newTags[key]=value
			}
			userConfig.Tags= newTags
			SaveConfig(userID, userConfig)
		}
	}
}

func sendTechPost(tag ,userID string, sequenceNumber int) bool {
	techData := GetData(tag)
	if len(techData) <= sequenceNumber {
		return false
	}
	channel, _ := config.Mattermost.GetDirectChannel(config.GetConfig().BotUserID, userID)
	post := &model.Post{
		ChannelId: channel.Id,
		UserId:    config.GetConfig().BotUserID,
		Message:   techData[sequenceNumber],
	}
	_, err := config.Mattermost.CreatePost(post)
	if err != nil {
		return false
	}
	return true
}

func PostQuestion(userIDs []string, text string, userID string, questionID int) {
	for _, id := range userIDs {
		if id == userID {
			continue
		}
		channel, _ := config.Mattermost.GetDirectChannel(config.GetConfig().BotUserID, id)
		post := &model.Post{
			ChannelId: channel.Id,
			UserId:    config.GetConfig().BotUserID,
			Message:   "Hi one of our friend's need our help",
		}
		actions := []*model.PostAction{}

		actions = append(actions, &model.PostAction{
			Type: "button",
			Name: "Submit Answer",
			Integration: &model.PostActionIntegration{
				URL: fmt.Sprintf("%s/plugins/%s/%s?id=%d&user_id=%s", *config.Mattermost.GetConfig().ServiceSettings.SiteURL, config.PluginName, "submit-answer", questionID, userID),
			},
		})

		post.AddProp("attachments", []*model.SlackAttachment{
			{
				Text:    fmt.Sprintf("**Q.** %s \n", text),
				Actions: actions,
			},
		})
		config.Mattermost.CreatePost(post)
	}
}
