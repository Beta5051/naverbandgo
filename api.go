package naverbandgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const ENDPOINT = "https://openapi.band.us"

// Client 클라이언트
type Client struct {
	AccessToken string
	HttpClient *http.Client
}

// NewClient 클라이언트 생성
func NewClient(accessToken string, client *http.Client) *Client {
	if client == nil {
		client = &http.Client{}
	}
	return &Client{
		accessToken,
		client,
	}
}

// APIResult API 응답
type APIResult struct {
	ResultCode uint `json:"result_code"`
	ResultData interface{} `json:"result_data"`
}

// To API 응답 변환
func (result *APIResult) To(v interface{}) error {
	bs, err := json.Marshal(result.ResultData)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, v)
}

// CallAPI API 요청
func (client *Client) CallAPI(method string, path string, data map[string]interface{}) (*APIResult, error) {
	method = strings.ToUpper(method)
	if method != "GET" && method != "POST" {
		return nil, errors.New("api method must be GET or POST")
	}
	req, err := http.NewRequest(method, ENDPOINT + path, nil)
	query := &url.Values{}
	query.Add("access_token", client.AccessToken)
	if len(data) != 0 {
		for key, value := range data {
			query.Add(key, fmt.Sprint(value))
		}
	}
	req.URL.RawQuery = query.Encode()
	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := new(APIResult)
	err = json.Unmarshal(bs, result)
	if err != nil {
		return nil, err
	}
	if result.ResultCode != 1 {
		return nil, errors.New(string(bs))
	}
	return result, err
}