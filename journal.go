package edsm

import (
	"encoding/json"
	"net/http"
)

const (
	pathJournal = "api-journal-v1"
	pathDiscard = "api-journal-v1/discard"
)

func (srv *Service) Discard(events []string) error {
	rq, _ := http.NewRequest("GET", srv.url(pathDiscard), nil)
	q := rq.URL.Query()
	rq.URL.RawQuery = q.Encode()
	rq.Header.Set("Accept", "application/json")
	resp, err := srv.HttpClient.Do(rq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&events)
	return err
}

func (srv *Service) Journal(cmdr string, entry string) error {
	rq, _ := http.NewRequest("POST", srv.url(pathJournal), nil)
	q := rq.URL.Query()
	q.Set("commanderName", cmdr)
	if srv.Creds != nil {
		q.Set("apiKey", srv.Creds.ApiKey)
	}
	q.Set("fromSoftware", Software)
	q.Set("fromSoftwareVersion", VersionStr())
	q.Set("message", entry)
	rq.URL.RawQuery = q.Encode()
	rq.Header.Set("Accept", "application/json")
	_, err := srv.HttpClient.Do(rq)
	return err
}
