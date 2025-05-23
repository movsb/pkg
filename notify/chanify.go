package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// https://github.com/chanify/chanify#http-api
type Chanify struct {
	endpoint string
	token    string
	base     *url.URL
}

func NewOfficialChanify(token string) *Chanify {
	return NewChanify(`https://api.chanify.net`, token)
}

func NewChanify(endpoint string, token string) *Chanify {
	u, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}
	c := &Chanify{
		endpoint: endpoint,
		token:    token,
		base:     u,
	}
	return c
}

type ChanifyMessage struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Sound int    `json:"sound"`
}

func (c *Chanify) Send(title string, content string, sound bool) error {
	u := c.base.JoinPath(`/v1/sender`, c.token)

	m := ChanifyMessage{
		Title: title,
		Text:  content,
		Sound: 1,
	}

	b := bytes.NewBuffer(nil)
	json.NewEncoder(b).Encode(m)

	rsp, err := http.Post(u.String(), `application/json`, b)
	if err != nil {
		log.Println(err)
		return err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode < 200 || rsp.StatusCode >= 300 {
		log.Println("Chanify error:", rsp.Status)
		return fmt.Errorf(`chanify error: %d %s`, rsp.StatusCode, u.String())
	}

	return nil
}
