package utils

import "strings"

func GetSubdomain(hostname string) string {
	parts := strings.Split(hostname, ".")

	if len(parts) >= 2 {
		return parts[0]
	}
	return ""
}
