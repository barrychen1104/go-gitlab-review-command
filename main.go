package main

import (
	"fmt"
	"go-gitlab-review-command/common"
	"go-gitlab-review-command/service"
	"os"
	"strconv"

	util "github.com/restuwahyu13/gin-rest-api/utils"
)

func main() {
	if len(os.Args) < 5 {
		panic("參數不足")
	}

	openaiKey := os.Args[1]
	gitlabToken := os.Args[2]
	projectId, _ := strconv.Atoi(os.Args[3])
	mrId, _ := strconv.Atoi(os.Args[4])
	gitlabUrl := util.GodotEnv("GITLAB_SERVER_URL")
	openaiUrl := util.GodotEnv("OPENAI_API_URL")

	common.NewAppConfig(openaiUrl, openaiKey, gitlabUrl, gitlabToken, projectId, mrId)
	fmt.Println("Config: ", common.AppConfig)

	diffs, err := service.GetChanges(common.AppConfig.ProjectId, common.AppConfig.MrId)
	if err != nil {
		fmt.Println("Error getting changes:", err)
		return
	}

	comments, err := service.ReviewCode(diffs)
	if err != nil {
		fmt.Println("Error reviewing code:", err)
		return
	}

	if err := service.WriteComments(projectId, mrId, comments); err != nil {
		fmt.Println("Error writing comments:", err)
		return
	}

	fmt.Println("GPT Code Review Done!")
}
