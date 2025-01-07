package utils

import (
	"fmt"
)

func HttpCode(code int) string {
	var message string
	switch code {
	case 401:
		message = "Unauthorized"
	case 403:
		message = "Forbiden"
	case 200:
		message = "OK"
	case 500:
		message = "Internal Server Error"
	case 400:
		message = "Bad Request"
	default:
		message = "Not Found"
	}
	return fmt.Sprintf("%d %s", code, message)
}
