package utils

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// checks if the text is inside the patch request json
func DoesPatchIncludeTextInJSON(r *http.Response, text string) bool {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	return strings.Contains(buf.String(), text)
}

func CalculateETA(startTime time.Time, currentIteration, totalIterations int) string {
	elapsedTime := time.Since(startTime)
	averageTimePerIteration := elapsedTime / time.Duration(currentIteration)
	remainingIterations := totalIterations - currentIteration
	estimatedRemainingTime := averageTimePerIteration * time.Duration(remainingIterations)

	eta := elapsedTime + estimatedRemainingTime
	days := int(eta.Hours()) / 24
	hours := int(eta.Hours()) % 24
	minutes := int(eta.Minutes()) % 60
	seconds := int(eta.Seconds()) % 60

	return fmt.Sprintf("%dd:%02dh:%02dm:%02ds", days, hours, minutes, seconds)
}
