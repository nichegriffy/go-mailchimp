package mailchimp

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
)

// Unsubscribe ...
func (c *Client) Unsubscribe(listID string, email string) error {
	// Hash email
	emailMD5 := fmt.Sprintf("%x", md5.Sum([]byte(email)))
	// Make request
	resp, err := c.do(
		"DELETE",
		fmt.Sprintf("/lists/%s/members/%s", listID, emailMD5),
		nil,
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 204 {
		return nil
	}

	// Read the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Request failed
	errorResponse, err := extractError(data)
	if err != nil {
		return err
	}
	return errorResponse
}
