package shared

import (
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/clients/rest/httptransport"
)

type SaladCloudImdsSdkError struct {
	Err      error
	Body     []byte
	Metadata SaladCloudImdsSdkErrorMetadata
}

type SaladCloudImdsSdkErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewSaladCloudImdsSdkError[T any](transportError *httptransport.ErrorResponse[T]) *SaladCloudImdsSdkError {
	return &SaladCloudImdsSdkError{
		Err:  transportError.GetError(),
		Body: transportError.GetBody(),
		Metadata: SaladCloudImdsSdkErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

func (e *SaladCloudImdsSdkError) Error() string {
	return e.Err.Error()
}
