package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (server *Server) MakeBet(token string, puuid string, betsize float64) (err error) {
	req, err := http.NewRequest("POST", `https://provider.expsoftdev.com/api/v0/users/bet`, nil)
	q := req.URL.Query()
	q.Add("token", token)
	q.Add("puuid", puuid)
	q.Add("game_uuid", nodeConfig.Server.GameUID)
	q.Add("txn_id", server.TxnId())
	q.Add("bet", fmt.Sprintf("%.2f", betsize))
	req.URL.RawQuery = q.Encode()
	req.Header.Set("X-Auth-Token", token)

	//token, puuid, game_uuid, txn_id, bet
	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Status code not 200:" + resp.Status)
	}

	return
}

func (server *Server) MakeWin(token string, puuid string, winsize float64) (err error) {
	req, err := http.NewRequest("POST", `https://provider.expsoftdev.com/api/v0/users/win`, nil)
	q := req.URL.Query()
	q.Add("token", token)
	q.Add("puuid", puuid)
	q.Add("game_uuid", nodeConfig.Server.GameUID)
	q.Add("txn_id", server.TxnId())
	q.Add("win", fmt.Sprintf("%.2f", winsize))
	req.URL.RawQuery = q.Encode()
	req.Header.Set("X-Auth-Token", token)

	//token, puuid, game_uuid, txn_id, bet
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Status code not 200:" + resp.Status)
	}

	return
}

func (server *Server) UserLogin(token string, puuid string) (err error) {
	req, err := http.NewRequest("GET", "http://provider.expsoftdev.com/api/users/login?token="+token+"&puuid="+puuid, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Auth-Token", token)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("Can't login user: " + resp.Status)
	}

	return
}

func (server *Server) GetUserInfo(token string, puuid string) (info *UserInfo, err error) {
	req, err := http.NewRequest("GET", "https://provider.expsoftdev.com/api/v0/users/userInfo?token="+token+"&puuid="+puuid + "&game_uuid=" + nodeConfig.Server.GameUID, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Auth-Token", token)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Can't get user info: " + resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	info = &UserInfo{}
	err = json.Unmarshal(body, info)
	if err != nil {
		return
	}

	return
}
