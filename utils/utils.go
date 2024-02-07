package utils

import (
	"fmt"
	"strings"
)

/*
    Urls returned can be one of two kinds, one that starts with http or https and one that starts with /.
	If the url starts with http or https then we can check if the url is part of the domain by checking if the url starts with the domain.
	If the url starts with / then we can assume that it is part of the domain and we can append the domain to the url.
*/

// CheckDomainMatch checks if the url is part of the domain
func CheckDomainMatch(domain string, url string) (string, bool) {
	if strings.HasPrefix(url, "http") {
		if strings.HasPrefix(url, domain) {
			return url, true
		}
		return "", false
	}
	if len(url) > 2 && url[0] == '/' && url[len(url)-1] == '/' {
		return fmt.Sprintf(domain+"%s", url), true
	}
	return "", false
}
