package gorge

import (
	"encoding/base64"
)
func auth(username, password string) string {
  auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}