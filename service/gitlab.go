package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-gitlab-review-command/common"
	"net/http"

	util "github.com/restuwahyu13/gin-rest-api/utils"
)

func GetChanges(projectId, mrId int) ([]string, error) {
	url := fmt.Sprintf("%s/api/v4/projects/%d/merge_requests/%d/changes", util.GodotEnv("GITLAB_SERVER_URL"), common.AppConfig.ProjectId, common.AppConfig.MrId)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Private-Token", common.AppConfig.GITLAB_PRIVATE_TOKEN)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("API returned an error:", res.Status)
		return nil, err
	}

	// Read response body into a buffer
	var responseBodyBuffer bytes.Buffer
	_, err = responseBodyBuffer.ReadFrom(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var mrChanges map[string]interface{}
	err = json.Unmarshal(responseBodyBuffer.Bytes(), &mrChanges)
	if err != nil {
		fmt.Println("Error parsing response body:", err)
		return nil, err
	}

	var diffs []string
	changes := mrChanges["changes"].([]interface{})
	for _, change := range changes {
		diff := change.(map[string]interface{})["diff"].(string)
		diffs = append(diffs, diff)
	}

	return diffs, nil
}

func WriteComments(projectId, mrId int, content string) error {
	url := fmt.Sprintf("%s/api/v4/projects/%d/merge_requests/%d/notes", util.GodotEnv("GITLAB_SERVER_URL"), common.AppConfig.ProjectId, common.AppConfig.MrId)
	requestBody := map[string]string{
		"body": content,
	}
	jsonData, _ := json.Marshal(requestBody)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("建立請求時發生錯誤：", err)
		return err
	}
	req.Header.Set("Private-Token", common.AppConfig.GITLAB_PRIVATE_TOKEN)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("傳送請求時發生錯誤：", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Println("請求不成功，狀態碼：", resp.StatusCode)
		return err
	}

	var responseBody []byte
	_, err = resp.Body.Read(responseBody)
	if err != nil {
		fmt.Println("讀取回應時發生錯誤：", err)
		return err
	}

	return nil
}

// func filterDiffContent(diffContent string) string {
// 	// filteredContent := re.sub(r'(^-.*\n)|(^@@.*\n)', '', diffContent, flags=re.MULTILINE)
// 	filteredContent := regexp.MustCompile(`(^-.*\n)|(^@@.*\n)`).ReplaceAllString(diffContent, "")
// 	splittedContent := strings.Split(filteredContent, "\n")
// 	var processedCode []string
// 	for _, line := range splittedContent {
// 		if strings.HasPrefix(line, "+") {
// 			processedCode = append(processedCode, line[1:])
// 		} else {
// 			processedCode = append(processedCode, line)
// 		}
// 	}
// 	return strings.Join(processedCode, "\n")
// }
