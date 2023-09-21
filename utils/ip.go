package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type IPAddressInfo struct {
	IP            string `json:"ip"`
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	Type          string `json:"type"`
	Continent     string `json:"continent"`
	ContinentCode string `json:"continent_code"`
	Country       string `json:"country"`
	CountryCode   string `json:"country_code"`
	Region        string `json:"region"`
	RegionCode    string `json:"region_code"`
	City          string `json:"city"`
}

func GetIPAddress(ip string) (string, error) {
	address := ""
	if ip == "127.0.0.1" {
		address = "本地"
	} else {
		url := fmt.Sprintf("https://ipwho.is/%s?lang=zh-CN", ip[0])
		header := map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36",
		}
		body, status, err := HttpGet(url, header, "application/json", 0)
		if err != nil {
			return "", err
		}
		if status != http.StatusOK {
			return "", fmt.Errorf("response status:%d", status)
		}
		var ipAddressInfo IPAddressInfo
		if err = json.Unmarshal(body, &ipAddressInfo); err != nil {
			return "", err
		}
		if ipAddressInfo.Success {
			address = ipAddressInfo.Country + "," + ipAddressInfo.Region + "," + ipAddressInfo.City
		} else {
			return "", errors.New(ipAddressInfo.Message)
		}
	}
	return address, nil
}
