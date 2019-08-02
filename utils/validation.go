package utils

import (
	"regexp"
	"strings"
)

func IsOnion(identifier string) bool {
	// TODO: At some point we will want to support i2p
	if len(identifier) >= 22 && strings.HasSuffix(identifier, ".onion") {
		matched, _ := regexp.MatchString(`(^|\.)[a-z0-9]{16}\.onion$|(^|\.)[a-z0-9]{56}\.onion$`, identifier)
		return matched
	}
	return false
}
