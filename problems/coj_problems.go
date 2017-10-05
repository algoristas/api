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

// COJJudgement is the struct to represent a COJ judment for a submition.
type COJJudgement struct {
    Id        int     `json:"id"`
    Date      string  `json:"date"`
    Username  string  `json:"user"`
    ProblemId int     `json:"problem"`
    Status    string  `json:"judgment"`
    Language  string  `json:"lang"`
}


// COJJudgmentResponse is the struct to represent a collection of COJ Judgments.
type COJJudgmentResponse struct {
    Collection [] COJJudgement
}


// SubmitionsProblemInfo represents the important information of the submitions to a problem
// The values for HasSolved && HasTried are "YES" && "NO"
type SubmitionsProblemInfo struct {
    HasSolved  bool
    HasTried   bool
}

// generateRequestURL generates the URL to make a GET Resquest.
// 
// Generates the format as follows:
//    http://URL_REQUESTED?param1=value1&param2=value2&...&paramn=valuen
// It returns a String that is the well format URL to make the request.
func generateRequestURL(method, url_requested string, parameters map[string]string) (string, error){
    request, err := http.NewRequest(method, url_requested, nil)
    if err != nil {
        log.Println("GenerateResquestUrl failed: ", err)
        return "", err
    }

    q := request.URL.Query()
    for key, value := range parameters { 
        q.Add(key, value)
    }
    request.URL.RawQuery = q.Encode()
    return request.URL.String(), nil
}

// createNetClient generates a http client to make request.
// 
// It returns a pointer of type http.Client to make request to servers.
// We do the request in this way so we can control the timeout, otherwise we may have troubles
// if the server we are requesting for fails.
func createNetClient() (*http.Client){
    const timeoutRequest = 10 * time.Second // Requests timeouts
    const timeoutDial = 180 * time.Second   // limits the time spent establishing a TCP connection,
                                            //   often around 3 minutes.
    const timeoutTSL = 30 * time.Second     // limits the time spent performing the TLS handshake.

    // cap the TCP connect and TLS handshake timeouts
    // as well as establishing an end-to-end request timeout.
    netTransport := &http.Transport{
      Dial: (&net.Dialer{
        Timeout: timeoutDial,
      }).Dial,
      TLSHandshakeTimeout: timeoutTSL,
    }
    netClient := &http.Client{
      Timeout: time.Second,
      Transport: netTransport,
    }

    return netClient
}

// getHasSolvedInCOJ returns the important information that said if a user "USERNAME" has solve or has tried
// the problem "PID" in the COJ.
func getHasSolvedInCOJ(pid string, username string) (SubmitionsProblemInfo, error){
    sumbitionsInfo := SubmitionsProblemInfo{false, false}

    // Parameters sent to the API
    apiParameters := map[string]string{ "username": username, "pid": pid};
    // Generate the URL for make the request
    apiURL, err := generateRequestURL("GET", "http://coj.uci.cu/api/judgment", apiParameters)
    
    if err != nil {
        log.Println("GenerateResquestUrl failed: ", err)
        return sumbitionsInfo, err
    }

    // Generate the http.Client to make the request. We do the request in this way so we can control
    // the time out, otherwise we may have troubles if COJ API fails.
    netClient := createNetClient()

    response, errResponse := netClient.Get(apiURL)
    if errResponse != nil {
        log.Println("GET Resquest failed: ", errResponse)
        return sumbitionsInfo, errResponse
    }

    buffer, errBuffer := ioutil.ReadAll(response.Body)
    if errBuffer != nil {
        log.Println("Parse response to buffer failed: ", errBuffer)
        return sumbitionsInfo, errBuffer
    }

    // Note: The Judge response is a JSON array, we should parse each element as a COJJudgment
    // so we can use it as a struct. We can use other approach that is parse the JSON as a map
    // See the example below:
    //      data := make([] map[string]interface{}, 0)

    data := make([] COJJudgement, 0)
    errParse := json.Unmarshal(buffer, &data)
    if errParse != nil {
        log.Println("Parse response to COJJudgement failed: ", errParse)
        return sumbitionsInfo, errParse
    }

    // Iterate over all submitions to check if the problem has been solve by the user.
    // We can make another call to the API adding an parameter status=ac but we already have all the
    // submitions that is a waste of time.
    for _, sumbition := range data {
        if sumbitionsInfo.HasTried == false {
            sumbitionsInfo.HasTried = true
        }
        if sumbition.Status == "Accepted" {
            sumbitionsInfo.HasSolved = true
        }
    }
    return sumbitionsInfo, nil
}
