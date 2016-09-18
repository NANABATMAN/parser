package main

import (
	"bytes"
	"html"
	"io/ioutil"
	"net/http"
	"strings"
)

type MessageContent struct {
	Mentions  []string `json:"mentions"`
	Emoticons []string `json:"emoticons"`
	Links     []Link   `json:"links"`
}

var (
	mentions  = make(map[string]string)
	emoticons = make(map[string]string)
	urls      = make(map[string]string)
)

func (messageContent *MessageContent) AddMention(mention string) {
	if _, ok := mentions[mention]; !ok {
		mentions[mention] = mention
		messageContent.Mentions = append(messageContent.Mentions, mention)
	}
}

func (messageContent *MessageContent) AddEmoticon(emoticon string) {
	if _, ok := emoticons[emoticon]; !ok {
		emoticons[emoticon] = emoticon
		messageContent.Emoticons = append(messageContent.Emoticons, emoticon)
	}
}

func (messageContent *MessageContent) AddLink(url string) {
	if _, ok := urls[url]; !ok {
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			html, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				titleChan := make(chan string, 1)
				go getTitle(html, titleChan)

				link := Link{}
				link.setUrl(url)
				link.setTitle(<-titleChan)
				urls[url] = url
				messageContent.Links = append(messageContent.Links, link)
			}
		}
	}
}

func getTitle(htmlCode []byte, titleChan chan<- string) {
	rows := strings.Split(string(htmlCode), NEW_LINE)
	var title bytes.Buffer
	for i := 0; i < len(rows); i++ {
		if strings.Contains(rows[i], TITLE_OPEN_TEG) && strings.Contains(rows[i], TITLE_CLOSE_TEG) {
			title.WriteString(clearTitle(rows[i]))
		}
	}
	titleChan <- title.String()
}

func clearTitle(htmlTitle string) string {
	htmlTitle = strings.Replace(htmlTitle, TITLE_OPEN_TEG, EMPTY_STRING, 1)
	htmlTitle = strings.Replace(htmlTitle, TITLE_CLOSE_TEG, EMPTY_STRING, 1)
	for string(htmlTitle[0]) == SPACE {
		htmlTitle = strings.Replace(htmlTitle, SPACE, EMPTY_STRING, 1)
	}
	return html.UnescapeString(htmlTitle)
}
