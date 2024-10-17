package services

import (
	"api/models/dto"
	util "api/utils"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type TimeService struct{}

func NewTimeService() *TimeService {
	return &TimeService{}
}

func (s *TimeService) GetCurrentTime(timezone string) (dto.TimeResponse, error) {
	var timeResponse dto.TimeResponse
	url, err := url.Parse("https://timeapi.io/api/time/current/zone")
	if err != nil {
		log.Fatal(err)
	}
	query := url.Query()
	query.Add("timeZone", timezone)
	url.RawQuery = query.Encode()

	resp, err := http.Get(url.String())
	if err != nil {
		return timeResponse, util.InternalServerError
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return timeResponse, util.NotFound
	}

	if err := json.NewDecoder(resp.Body).Decode(&timeResponse); err != nil {
		return timeResponse, util.InternalServerError
	}

	return timeResponse, nil
}
