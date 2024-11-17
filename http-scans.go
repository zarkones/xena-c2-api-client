package c2api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func GetHttpScans(page int) (scans []HttpScan, err error) {
	req, err := http.NewRequest(http.MethodGet, *BaseURL+"/v1/scans/http", nil)
	if err != nil {
		return nil, err
	}

	setAuth(req)

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return []HttpScan{}, nil
	}

	return scans, json.NewDecoder(resp.Body).Decode(&scans)
}

func InsertHttpScan(requestID string, agentIDs []string) (err error) {
	jsonBody, err := json.Marshal(&HttpScan{
		ReqID:    requestID,
		AgentIDs: strings.Join(agentIDs, ","),
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, *BaseURL+"/v1/scans/http", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	setAuth(req)

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return errors.Join(ErrUnexpectedStatusCode, errors.New(resp.Status))
	}

	return nil
}
