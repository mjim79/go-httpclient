package gomime

import "testing"

func TestHeaders(t *testing.T) {
	if ContentTypeJson != "application/json" {
		t.Error("invalid content type json")
	}
	if ContentTypeXml != "application/xml" {
		t.Error("invalid content type xml")
	}
}
