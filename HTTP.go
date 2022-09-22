package discordhooks

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// forward will enforce and send the http request to the webhook server.
func (H *Hook) forward(b []byte, timeout time.Duration) error {
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("https://discord.com/api/webhooks/%d/%s", H.ID, H.Key), bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	results, err := client.Do(request)
	if err != nil {
		return err
	}

	contents, err := io.ReadAll(results.Body)
	if err == nil && len(contents) > 0 {
		return errors.New("usually the api responses with an empty body, this time it returned a value")
	}

	if results.StatusCode >= 205 {
		return errors.New("bad status code returned, possibly sent")
	}

	return nil
}