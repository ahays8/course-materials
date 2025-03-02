package shodan

// Do not use this file directly, do not attemp to compile this source file directly
// Go To lab/3/shodan/main/main.go

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"	//to verify page number given as argument
)

type HostLocation struct {
	City         string  `json:"city"`
	RegionCode   string  `json:"region_code"`
	AreaCode     int     `json:"area_code"`
	Longitude    float32 `json:"longitude"`
	CountryCode3 string  `json:"country_code3"`
	CountryName  string  `json:"country_name"`
	PostalCode   string  `json:"postal_code"`
	DMACode      int     `json:"dma_code"`
	CountryCode  string  `json:"country_code"`
	Latitude     float32 `json:"latitude"`
}

type Host struct {
	OS        string       `json:"os"`
	Timestamp string       `json:"timestamp"`
	ISP       string       `json:"isp"`
	ASN       string       `json:"asn"`
	Hostnames []string     `json:"hostnames"`
	Location  HostLocation `json:"location"`
	IP        int64        `json:"ip"`
	Domains   []string     `json:"domains"`
	Org       string       `json:"org"`
	Data      string       `json:"data"`
	Port      int          `json:"port"`
	IPString  string       `json:"ip_str"`
}

type HostSearch struct {
	Matches []Host `json:"matches"`
}

func (s *Client) HostSearch(q string, page string) (*HostSearch, error) {
	pageNumber, err :=  strconv.Atoi(page);
	if err != nil {
		return nil, err
	}
	res, err := http.Get(
		fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s&page=%d", BaseURL, s.apiKey, q, pageNumber),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret HostSearch
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}