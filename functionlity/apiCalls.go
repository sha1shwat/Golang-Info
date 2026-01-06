package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Country struct {
	Team1      string
	Team2      string
	Team1goals string
	Team2goals string
}

type ApiResponse struct {
	Data       []Country `json:"data"`
	Page       int       `json:"page"`
	PerPage    int       `json:"per_page"`
	TotalPages int       `json:"total_pages"`
	Total      int       `json:"total"`
}

const url = "https://jsonmock.hackerrank.com/api/football_matches?year=%d&team1goals=%d&page=%d"

func getCapitalByCountry() []Country {

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := client.Get(fmt.Sprintf(url, 2011, 1, 1))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer resp.Body.Close()

	var apiResponse ApiResponse

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		fmt.Println(err)
		return nil
	}

	ans := make([]Country, 0)
	ans = append(ans, apiResponse.Data...)

	for page := 0; ; page++ {
		resp, err := client.Get(fmt.Sprintf(url, 2011, 1, page))

		if err != nil {
			fmt.Println(err)
			return ans
		}

		if err = json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
			fmt.Println(err)
			return ans
		}

		if len(apiResponse.Data) == 0 {
			return ans
		}

		ans = append(ans, apiResponse.Data...)

	}

	return ans

}

func clientWithHeaders() []Country {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(url, 2011, 1, 1),
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer <token>")
	req.Header.Set("X-Custom-Header", "value")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	var apiResponse ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		fmt.Println(err)
		return nil
	}
	return apiResponse.Data

}
