package date

import "time"

// GetNow gets datetime now
func GetNow() string {
  return time.Now().UTC().Local().Format(time.RFC3339)
}