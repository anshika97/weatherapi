package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second * 5,
}

func GetData(city string) ([]byte, error) {
	var apiURL = fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&APPID=19e8588cb3d7d0623e3a5a8ec529232f", city)
	req, _ := http.NewRequest("GET", apiURL, nil)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Status code not 200")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
