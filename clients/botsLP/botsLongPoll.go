package botsLP

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"vkBot/e"
)

type BotLongPollResp struct {
	Response *BotsLongPollAPI `json:"response"`
}

type BotsLongPollAPI struct {
	Host   string `json:"server"`
	Key    string `json:"key"`
	Ts     string `json:"ts"`
	Client http.Client
}

func New(b *BotsLongPollAPI) *BotsLongPollAPI {
	return b
}

func (lp BotsLongPollAPI) makeQueryParams(ts string) url.Values {
	q := url.Values{}

	q.Add("act", "a_check")
	q.Add("key", lp.Key)
	q.Add("ts", ts)
	q.Add("wait", "25")

	return q
}

func (lp BotsLongPollAPI) Update(ts string) (updates []Update, err error) {
	defer func() { err = e.WrapIfError("can't get updates", err) }()
	helper := strings.Split(lp.Host, "/")
	u := url.URL{
		Scheme: "https",
		Host:   helper[2],
		Path:   helper[3],
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = lp.makeQueryParams(ts).Encode()

	resp, err := lp.Client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var res UpdateResponse

	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	return res.Updates, nil
}
