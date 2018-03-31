package main

import "fmt"

type RequestData struct {
	url       string
	browser   string
	sessionId string
	id        int
}

var queueRequest []RequestData

func initNewRequestData(browser, url, sessionId string) RequestData {
	var req RequestData
	req.browser = browser
	req.url = url
	req.sessionId = sessionId
	req.id = len(queueRequest) + 1
	return req
}

func main() {

	var aaa = initNewRequestData("chrome", "localhost:5455/", "session some new")
	queueRequest[1] = aaa
	fmt.Println("test", aaa, queueRequest)
}
