package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/stenstromen/rancher-renewer/types"
)

func GetRancherTokenInfo(apiKey string, rancherURL string, token string) (types.RancherAPIResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", rancherURL+"/v3/token/"+token, nil)
	if err != nil {
		return types.RancherAPIResponse{}, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return types.RancherAPIResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.RancherAPIResponse{}, err
	}

	var tokenInfo types.RancherAPIResponse
	json.Unmarshal(body, &tokenInfo)

	return tokenInfo, nil
}

func TokenIsExpiringSoon(expiration time.Time) bool {
	oneWeekLater := time.Now().Add(7 * 24 * time.Hour)
	return expiration.Before(oneWeekLater)
}
