package aoc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

const (
	aocBaseUrl          = "https://adventofcode.com"
	dayInputUrlTemplate = aocBaseUrl + "/2019/day/%d/input"
	filenameTemplate    = "day%d-input.txt"
)

type aocClient struct {
	httpClient http.Client
	sessionID  string
}

func GetInput(day int) {
	filename := fmt.Sprintf(filenameTemplate, day)
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		fmt.Printf("File %s exists, quitting\n", filename)
		return
	}
	c := newClient()
	input := c.downloadInput(day)
	saveInput(day, input)
}

func newClient() aocClient {
	aocUrl, _ := url.Parse(aocBaseUrl)

	cookies := []*http.Cookie{{
		Name:   "session",
		Value:  os.Getenv("AOC_SESSION_ID"),
		Path:   "/",
		Domain: ".adventofcode.com",
	}}

	jar, _ := cookiejar.New(nil)
	jar.SetCookies(aocUrl, cookies)

	return aocClient{
		httpClient: http.Client{Jar: jar},
	}
}

func (c aocClient) downloadInput(day int) []byte {
	url := fmt.Sprintf(dayInputUrlTemplate, day)
	fmt.Printf("Downloading input from %s\n", url)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		msg := fmt.Sprintf("Failed to download input for day %d", day)
		panic(msg)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return body
}

func saveInput(day int, input []byte) {
	filename := fmt.Sprintf(filenameTemplate, day)
	if err := ioutil.WriteFile(filename, input, 0644); err != nil {
		msg := fmt.Sprintf("Couldn't save file for day %d", day)
		panic(msg)
	}
}
