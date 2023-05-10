package vkAPI

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"vkBot/clients/botsLP"
	"vkBot/e"
	"vkBot/events"
)

const (
	getLongPollServerMethod = "groups.getLongPollServer"
	groupId                 = "220376788"
	messageSend             = "messages.send"
)

type VkAPI struct {
	host     string
	token    string
	basePath string
	client   http.Client
}

func New(host string, token string) *VkAPI {
	return &VkAPI{
		host:     host,
		token:    token,
		basePath: "method",
		client:   http.Client{},
	}
}

func (v *VkAPI) GetLastMessage(userId string) (*botsLP.HistoryResponse, error) {
	var q = mandatoryParams(v.token)
	q = makeQueryParams(q, "user_id", userId)
	q = makeQueryParams(q, "count", "3")
	q = makeQueryParams(q, "offset", "1")
	q = makeQueryParams(q, "peer_id", userId)
	q = makeQueryParams(q, "start_message_id", "-1")

	data, err := v.doRequest("messages.getHistory", q)

	if err != nil {
		return nil, err
	}

	var res botsLP.HistoryResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (v *VkAPI) SendMessage(text string, userId string, keyboard events.ViewKeyBoard) (err error) {
	defer func() { err = e.WrapIfError("can't send message", err) }()

	randId := rand.Int()

	keyboardJSON, _ := json.Marshal(keyboard)

	var q = mandatoryParams(v.token)

	q = makeQueryParams(q, "user_id", userId)
	q = makeQueryParams(q, "random_id", strconv.Itoa(randId))
	q = makeQueryParams(q, "keyboard", string(keyboardJSON[:]))

	_, err = v.doRequest(messageSend, makeQueryParams(q, "message", text))

	if err != nil {
		return err
	}

	return nil
}

func (v *VkAPI) DataForBotsLongPoll() (lp *botsLP.BotsLongPollAPI, err error) {
	defer func() { err = e.WrapIfError("can't connect to LongPoll", err) }()

	var q = mandatoryParams(v.token)

	dataLP, err := v.doRequest(getLongPollServerMethod, makeQueryParams(q, "group_id", groupId))

	if err != nil {
		return nil, err
	}

	var res botsLP.BotLongPollResp

	if err := json.Unmarshal(dataLP, &res); err != nil {
		return nil, err
	}

	return res.Response, nil

}

func makeQueryParams(queryParams url.Values, key string, value string) url.Values {
	q := queryParams
	q.Add(key, value)
	return q
}

func mandatoryParams(token string) url.Values {
	queryParams := url.Values{}

	queryParams.Add("v", "5.131")
	queryParams.Add("access_token", token)

	return queryParams
}

func (v *VkAPI) doRequest(method string, queryParams url.Values) (data []byte, err error) {
	defer func() { err = e.WrapIfError("can't do request to VK API", err) }()

	u := url.URL{
		Scheme: "https",
		Host:   v.host,
		Path:   path.Join(v.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = queryParams.Encode()

	resp, err := v.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	log.Println("connection to vk api was successful")

	return body, nil
}
