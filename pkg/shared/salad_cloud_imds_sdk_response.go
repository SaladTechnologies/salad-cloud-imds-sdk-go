package shared

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/clients/rest/httptransport"
	"net/http"
)

// SaladCloudImdsSdkResponse is the user-facing wrapper for API responses.
// It contains the deserialized data, raw HTTP response, and metadata like headers and status code.
type SaladCloudImdsSdkResponse[T any] struct {
	Data     T
	Raw      *http.Response
	Metadata SaladCloudImdsSdkResponseMetadata
}

// SaladCloudImdsSdkResponseMetadata contains HTTP metadata from the API response.
// Includes status code and headers for inspection and debugging.
type SaladCloudImdsSdkResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

// NewSaladCloudImdsSdkResponse creates a new response wrapper from an internal transport response.
// Extracts data and metadata into a user-facing structure.
func NewSaladCloudImdsSdkResponse[T any](resp *httptransport.Response[T]) *SaladCloudImdsSdkResponse[T] {
	return &SaladCloudImdsSdkResponse[T]{
		Data: resp.Data,
		Raw:  resp.Raw,
		Metadata: SaladCloudImdsSdkResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}

// GetData returns the deserialized response data.
func (r *SaladCloudImdsSdkResponse[T]) GetData() T {
	return r.Data
}

// String returns a JSON representation of the response for debugging.
// Returns an error message if JSON marshaling fails.
func (r SaladCloudImdsSdkResponse[T]) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: SaladCloudImdsSdkResponse to string"
	}
	return string(jsonData)
}
