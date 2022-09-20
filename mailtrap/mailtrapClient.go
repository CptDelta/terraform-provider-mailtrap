package mailtrap

import (
	"fmt"
	"net/http"
	"time"
)

func mailTrapClient(api string) *http.Response {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", "https://mailtrap.io", api), nil)
	req.Header.Add("Api-Token", "6313a012dbee479ca4b038fba237ba4c")

	if err != nil {
		err.Error()
	}

	resp, err := client.Do(req)

	return resp
}
