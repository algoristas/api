package problems

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/algoristas/api/util"
)

// COJJudgement is the struct to represent a COJ judment for a submition.
type COJJudgement struct {
	ID        int    `json:"id"`
	Date      string `json:"date"`
	Username  string `json:"user"`
	ProblemID int    `json:"problem"`
	Status    string `json:"judgment"`
	Language  string `json:"lang"`
}

// COJJudgmentResponse is the struct to represent a collection of COJ Judgments.
type COJJudgmentResponse struct {
	Collection []COJJudgement
}

// SubmissionProblemInfo represents the important information of the submitions to a problem
// The values for HasSolved && HasTried are "YES" && "NO"
type SubmissionProblemInfo struct {
	HasSolved bool
	HasTried  bool
}

// getHasSolvedInCOJ returns the important information that said if a user "USERNAME" has solve or has tried
// the problem "PID" in the COJ.
func getHasSolvedInCOJ(pid string, username string) (SubmissionProblemInfo, error) {
	submissionInfo := SubmissionProblemInfo{false, false}

	// Parameters sent to the API
	apiParameters := map[string]string{"username": username, "pid": pid}
	// Generate the URL for make the request
	req, err := util.BuildRequest("GET", "http://coj.uci.cu/api/judgment", apiParameters)

	if err != nil {
		return submissionInfo, err
	}

	// Generate the http.Client to make the request. We do the request in this way so we can control
	// the time out, otherwise we may have troubles if COJ API fails.
	netClient := util.NewHTTPClient()

	response, errResponse := netClient.Do(req)
	if errResponse != nil {
		log.Println("GET Resquest failed: ", errResponse)
		return submissionInfo, errResponse
	}

	buffer, errBuffer := ioutil.ReadAll(response.Body)
	if errBuffer != nil {
		log.Println("Parse response to buffer failed: ", errBuffer)
		return submissionInfo, errBuffer
	}

	// Note: The Judge response is a JSON array, we should parse each element as a COJJudgment
	// so we can use it as a struct. We can use other approach that is parse the JSON as a map
	// See the example below:
	//      data := make([] map[string]interface{}, 0)

	data := make([]COJJudgement, 0)
	errParse := json.Unmarshal(buffer, &data)
	if errParse != nil {
		log.Println("Parse response to COJJudgement failed: ", errParse)
		return submissionInfo, errParse
	}

	// Iterate over all submitions to check if the problem has been solve by the user.
	// We can make another call to the API adding an parameter status=ac but we already have all the
	// submissions that is a waste of time.
	if len(data) > 0 {
		submissionInfo.HasTried = true
	}

	for _, submission := range data {
		if submission.Status == "Accepted" {
			submissionInfo.HasSolved = true
		}
	}
	return submissionInfo, nil
}
