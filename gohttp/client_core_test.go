package gohttp

import (
	"github.com/mjim79/go-httpclient/gomime"
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {

	//Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", gomime.ContentTypeJson)
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.builder = &clientBuilder{
		headers:   commonHeaders,
		userAgent: "cool-user-agent",
	}

	//Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")
	finalHeaders := client.getRequestHeaders(requestHeaders)

	//Validation
	if len(finalHeaders) != 3 {
		t.Error("we expect 3 Headers")
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("invalid request received")
	}

	if finalHeaders.Get("Content-Type") != gomime.ContentTypeJson {
		t.Error("invalid content type received")
	}

	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("invalid user agent received")
	}
}

func TestGetRequestBody(t *testing.T) {
	// Initialization
	client := httpClient{}

	t.Run("noBodyNilResponse", func(t *testing.T) {
		// Execution
		body, err := client.getRequestBody("", nil)

		// Validation
		if err != nil {
			t.Error("no error expected when passing a nil String")
		}

		if body != nil {
			t.Error("no String expected when passing a nil String")
		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody(gomime.ContentTypeJson, requestBody)

		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json String obtained")
		}
	})

	t.Run("BodyWithXml", func(t *testing.T) {
		requestBody := []string{"<field>value</field>"}
		body, err := client.getRequestBody("application/xml", requestBody)

		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}

		if string(body) != `<string>&lt;field&gt;value&lt;/field&gt;</string>` {
			t.Error("invalid xml String obtained")
		}
	})

	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/another", requestBody)

		if err != nil {
			t.Error("no error expected when marshaling slice as default")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json as default String obtained")
		}
	})

}
