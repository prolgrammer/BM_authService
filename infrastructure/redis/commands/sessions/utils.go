package sessions

import "fmt"

func getKey(accessToken string) string {
	return fmt.Sprintf("session_%v", accessToken)
}
