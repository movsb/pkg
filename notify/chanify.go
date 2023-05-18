package notify

import (
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

func (c *Chanify) Send(title string, content string, sound bool) {
	u := c.base.JoinPath(`/v1/sender`, c.token, content)
	v := u.Query()
	if title != `` {
		v.Set(`title`, title)
	}
	if sound {
		v.Set(`sound`, `1`)
	}
	u.RawQuery = v.Encode()
	rsp, err := http.Get(u.String())
	if err != nil {
		log.Println(err)
	}
	defer rsp.Body.Close()
}
