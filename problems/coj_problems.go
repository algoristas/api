package problems

import (
        "fmt"
        "log"
        "net"
        "time"
        "net/http"
        "encoding/json"
        "io/ioutil"
)

// generate_request_url generates the URL to make a GET Resquest.
// Generates the format as follows:
//    http://URL_REQUESTED?param1=value1&param2=value2&...&paramn=valuen
// It returns a String that is the well format URL to make the request.
func generate_request_url(method, url_requested string, parameters map[string]string) (string, error){
    request, err := http.NewRequest(method, url_requested, nil)
    if err != nil {
        log.Panic(err)
        return "", err
    }

    q := request.URL.Query()
    for key, value := range parameters { 
        q.Add(key, value)
    }
    request.URL.RawQuery = q.Encode()
    return request.URL.String(), nil
}

// COJJudgement is the struct to represent a COJ judment for a submition.
type COJJudgement struct {
    Id int
    Date string
    User string
    Prob int
    Judgment string
    Errortestcase int
    Time int
    Memory string
    Tam string
    Lang string
}

// COJJudgmentResponse is the struct to represent a collection of COJ Judgments.
type COJJudgmentResponse struct {
    Collection [] COJJudgement
}


// SubmitionsProblemInfo represents the important information of the submitions to a problem
// The values for has_solved && has_tried are "YES" && "NO"
type SubmitionsProblemInfo struct {
    has_solved  string
    has_tried   string
}


// getHasSolvedInCOJ returns the important information that said if a user "USERNAME" has solve or has tried
// the problem "PID" in the COJ.
func getHasSolvedInCOJ(pid string, username string) (SubmitionsProblemInfo, error){
    submitionsProblemInfo := SubmitionsProblemInfo{"NO", "NO"}

    // Parameters sent to the API
    api_parameters := map[string]string{ "username": username, "pid": pid};
    // Generate the URL for make the request
    api_url, url_err := generate_request_url("GET", "http://coj.uci.cu/api/judgment", api_parameters)
    
    if url_err !=nil {
        log.Panic(url_err)
        return submitionsProblemInfo, url_err
    }

    // Generate the http.Client to make the request. We do the request in this way so we can controll
    // the time out, otherwise we may have troubles if COJ API fails.
    netTransport := &http.Transport{
      Dial: (&net.Dialer{
        Timeout: 5 * time.Second,
      }).Dial,
      TLSHandshakeTimeout: 5 * time.Second,
    }
    netClient := &http.Client{
      Timeout: time.Second * 10,
      Transport: netTransport,
    }

    response, response_error := netClient.Get(api_url)
    if response_error != nil {
        log.Panic(response_error)
        return submitionsProblemInfo, response_error
    }


    buffer,_ := ioutil.ReadAll(response.Body)
    // Note: The Judge response is a JSON array, we should parse each element as a COJJudgment
    // so we can use it as a struct. We can use other approach that is parse the JSON as a map
    // See the example below:
    // 		data := make([] map[string]interface{}, 0)

    data := make([] COJJudgement, 0)
    err := json.Unmarshal(buffer, &data)
    if err != nil {
        log.Panic(err)
    }

    // Iterate over all submitions to check if the problem has been solve by the user.
    // We can make another call to the API adding an parameter status=ac but we already have all the
    // submitions that is a waste of time.
    for _, sumbition := range data {
        if submitionsProblemInfo.has_tried == "" || submitionsProblemInfo.has_tried == "NO" {
            submitionsProblemInfo.has_tried = "YES"
        }
        if sumbition.Judgment == "Accepted" {
            submitionsProblemInfo.has_solved = "YES"
        }
    }
    return submitionsProblemInfo, nil
}
