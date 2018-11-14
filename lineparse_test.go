package lineparse

import "testing"

const (
	testLogLine = `campaign_manager | 2018-11-14T10:45:31Z INFO 127.0.0.1 GET /campaigns/1 403 user "guest" is not allowed to show campaign with id: '1'`
)

func TestExtract(t *testing.T) {
	ip := Extract(testLogLine, IPParser{})
	if ip != "127.0.0.1" {
		t.Errorf("ip mismatch: expected \"127.0.0.1\" got %q", ip)
	}
}