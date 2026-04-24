package shared

import (
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/clients/rest/httptransport"
	"net/http"
)

// SaladCloudImdsSdkError wraps API errors with detailed metadata including status code, headers, and raw response.
// It implements the error interface and provides structured access to error information.
type SaladCloudImdsSdkError[T any] struct {
	Err      error
	Data     *T
	Body     []byte
	Raw      *http.Response
	Metadata SaladCloudImdsSdkErrorMetadata
}

// SaladCloudImdsSdkErrorMetadata contains HTTP metadata associated with an error response.
type SaladCloudImdsSdkErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

// NewSaladCloudImdsSdkError creates a new SaladCloudImdsSdkError from an internal transport error.
// It extracts error details, body, status code, and headers into a user-facing error structure.
func NewSaladCloudImdsSdkError[T any](transportError *httptransport.ErrorResponse[T]) *SaladCloudImdsSdkError[T] {
	return &SaladCloudImdsSdkError[T]{
		Err:  transportError.GetError(),
		Data: transportError.Data,
		Body: transportError.GetBody(),
		Raw:  transportError.Raw,
		Metadata: SaladCloudImdsSdkErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

// Error implements the error interface, returning the error message string.
func (e *SaladCloudImdsSdkError[T]) Error() string {
	return e.Err.Error()
}

// GetData returns the deserialized error response data.
// Returns nil if unmarshaling failed or the response body was empty.
func (e *SaladCloudImdsSdkError[T]) GetData() *T {
	return e.Data
}
