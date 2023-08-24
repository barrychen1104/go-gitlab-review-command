package common

type Config struct {
	OPENAI_API_URL       string
	OPENAI_API_KEY       string
	GITLAB_SERVER_URL    string
	GITLAB_PRIVATE_TOKEN string
	ProjectId            int
	MrId                 int
}

var AppConfig *Config

func NewAppConfig(openaiUrl, openaiKey, gitlabUrl, gitlabToken string, projectId, mrId int) {
	AppConfig = &Config{
		OPENAI_API_URL:       openaiUrl,
		OPENAI_API_KEY:       openaiKey,
		GITLAB_SERVER_URL:    gitlabUrl,
		GITLAB_PRIVATE_TOKEN: gitlabToken,
		ProjectId:            projectId,
		MrId:                 mrId,
	}
}
