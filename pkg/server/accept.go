package server

import (
	"net/http"
	"strings"
)

func CanAccept(r *http.Request, values ...string) bool {
	accepts := parseAcceptValues(r.Header)

	if len(accepts) == 0 {
		return false
	}

	lowered := make([]string, 0, len(values))
	for _, v := range values {
		lowered = append(lowered, strings.ToLower(v))
	}

	for _, spec := range accepts {
		if spec == "*/*" {
			return true
		}
		for _, want := range lowered {
			if spec == want {
				return true
			}
		}
	}

	return false
}

// parseAcceptValues parses the Accept header values into a list of media types
// without parameters (e.g., "application/json"). It performs a minimal parse
// sufficient for equality checks and the */* wildcard and is case-insensitive.
func parseAcceptValues(h http.Header) []string {
	var out []string
	for _, headerVal := range h.Values("Accept") {
		for _, part := range strings.Split(headerVal, ",") {
			token := strings.TrimSpace(part)
			if token == "" {
				continue
			}
			if i := strings.IndexByte(token, ';'); i >= 0 {
				token = token[:i]
			}
			token = strings.ToLower(strings.TrimSpace(token))
			if token != "" {
				out = append(out, token)
			}
		}
	}
	return out
}
