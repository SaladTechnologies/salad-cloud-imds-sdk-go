package shared

import (
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/clients/rest/httptransport"
)

type SaladCloudImdsSdkResponse[T any] struct {
	Data     T
	Metadata SaladCloudImdsSdkResponseMetadata
}

type SaladCloudImdsSdkResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewSaladCloudImdsSdkResponse[T any](resp *httptransport.Response[T]) *SaladCloudImdsSdkResponse[T] {
	return &SaladCloudImdsSdkResponse[T]{
		Data: resp.Data,
		Metadata: SaladCloudImdsSdkResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}
