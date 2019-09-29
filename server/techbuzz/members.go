package techbuzz

import (
	"strings"
)

func GetTagMemberIDs(tag string) []string {
	memberIDs := []string{}
	usersIDs := GetTechMembers()
	for _, userID := range usersIDs {
		userConfig := GetUserConfig(userID)
		if userConfig.Enabled == false {
			continue
		} else {
			for key, value := range userConfig.Tags {
				if key == strings.ToLower(tag) && value.Enabled == true {
					memberIDs = append(memberIDs, userID)
				}
			}
		}
	}
	return memberIDs
}
