package pkgHttp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DoHttpRequest(request *http.Request, v any) error {
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("failed to parse response, %w", err)
	}

	return nil
}
