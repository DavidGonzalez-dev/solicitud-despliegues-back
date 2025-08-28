package pkgHttp

import (
	"encoding/json"
	"errors"
	"net/http"
)

func DoHttpRequest(request *http.Request, v any) error {
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to parse response")
	}
	return json.NewDecoder(resp.Body).Decode(v)
}