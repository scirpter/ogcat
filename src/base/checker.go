package base

import (
	"bytes"
	"fmt"
	"net/http"
	"ogcat/src/common"
	"ogcat/src/enums"
	"ogcat/src/utils"
	"ogcat/src/utils/cnsl"
	"ogcat/src/utils/gen"
	"time"
)

func RunChecker(as *string, idleTime time.Duration, usernames *[]string) {
	currentTime := time.Now()
	totalIterations := len(*usernames)
	startTime := time.Now()

	for i, username := range *usernames {
	REDO:
		jsonStr := []byte(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, "password"))
		client := &http.Client{}
		req, err := http.NewRequest("PATCH", common.DISCORD_ENDPOINT, bytes.NewBuffer(jsonStr))
		if err != nil {
			cnsl.Error(fmt.Sprintf("failed to create request: %s", err.Error()))
			time.Sleep(idleTime * time.Second)
			goto REDO
		}
		req.Header.Set("authorization", *as)
		req.Header.Set("content-type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			cnsl.Error(fmt.Sprintf("failed to send request: %s", err.Error()))
			time.Sleep(idleTime * time.Second)
			goto REDO
		}

		elapsedTime := time.Since(startTime)
		iterationsCompleted := i + 1
		iterationsRemaining := totalIterations - iterationsCompleted
		avgTimePerIteration := elapsedTime / time.Duration(iterationsCompleted)
		estimatedTimeRemaining := avgTimePerIteration * time.Duration(iterationsRemaining)

		eta := currentTime.Add(estimatedTimeRemaining).Format("02 Jan 2006 15:04:05")

		if utils.DoesPatchIncludeTextInJSON(resp, "USERNAME_ALREADY_TAKEN") {
			cnsl.LogLowPrio(fmt.Sprintf("%s is taken (%d/%d) ETA: %s", username, i+1, len(*usernames), eta))
			gen.AddUsernameToTxt(&username, enums.RetTaken)
		} else if resp.StatusCode != 400 {
			cnsl.Error(fmt.Sprintf("%s failed, unauthorized or ratelimit. we'll retry. (%d/%d) ETA: %s", username, i+1, len(*usernames), eta))
			time.Sleep(12 * time.Second)
			goto REDO
		} else {
			cnsl.Ok(fmt.Sprintf("%s is available (%d/%d) ETA: %s", username, i+1, len(*usernames), eta))
			gen.AddUsernameToTxt(&username, enums.RetAvailable)
		}

		time.Sleep(idleTime * time.Second)
	}
}
