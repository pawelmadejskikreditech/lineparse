// Package lineparse helps to extract elements from our internal log structure
package lineparse

import (
	"regexp"
)

// IPParser parses log line in search for IP address specified in RFC791 represented by 32-bit integer value
// in the 'dot-decimal' notation (eg. 192.168.10.1)
type IPParser struct {}

// Parse returns sting containing dot-decimal notation of IPv4 address found in given line
func (IPParser) Parse(line string) string {
	r := regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`)
	return r.FindStringSubmatch(line)[0]
}

// Extract extracts searched log values from given log line.
// log line structure is predefined and should follow schema:
// 		<application_name> | <time in ISO8601> <log_lvl> <requester_ip> <http_method> <path> <response_status> <message>
// example line:
// 		campaign_manager | 2018-11-14T10:45:31Z INFO 127.0.0.1 GET /campaigns/1 403 user "guest" is not allowed to show campaign with id: '1'
func Extract(line string, parser IPParser) string {
	return parser.Parse(line)
}
