package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/sipeed/picoclaw/pkg/config"
)

type HTTPProvider struct {
	apiKey     string
	apiBase    string
	student    string
	password   string
	httpClient *http.Client
	lastUpdate int64
}

type HTTPProviderConfig struct {
	apiKey     string
	apiBase    string
	proxy      string
	student    string
	password   string
	lastUpdate int64
}

func refreshJWTBeforeChat(p *HTTPProvider) error {
	if strings.Contains(p.apiBase, "chat.shou.edu.cn") {
		if (time.Now().Unix() - p.lastUpdate) > (24-1)*60*60 {
			//fmt.Printf("\n\u001B[36m[DEBUG]\u001B[0m Original Config: %+v\n", p)
			apiKey, _right := authShouMiddle(p.apiBase, p.student, p.password)
			if _, ok := _right.(error); !ok {
				p.apiKey = apiKey.(string)
				p.lastUpdate = time.Now().Unix()
			} else {
				return _right.(error)
			}
		}
	}
	return nil
}

func authShouMiddle(apiBase string, student string, password string) (interface{}, interface{}) {
	var apiKey string
	data := map[string]string{
		"user":     student,
		"password": password,
	}
	//fmt.Printf("\n\u001B[36m[DEBUG]\u001B[0m Auth Data: %s\n", data)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	req, err := http.NewRequest("POST", apiBase+"/api/v1/auths/ldap", bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to update jwt: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send post request: %w", err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var respJson map[string]interface{}
	if err := json.Unmarshal(respBody, &respJson); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	token, ok := respJson["token"].(string)
	//fmt.Printf("\n\u001B[36m[DEBUG]\u001B[0m Original Response: %s\n", respJson)
	if !ok {
		return nil, fmt.Errorf("token not found or not a string in response")
	}
	apiKey = token
	fmt.Println("ApiKey Updated")
	return apiKey, nil
}

func authShou(cfg *config.Config) (interface{}, interface{}) {
	return authShouMiddle(cfg.Providers.SHOU.APIBase, cfg.Providers.SHOU.Student, cfg.Providers.SHOU.PASSWORD)
}
