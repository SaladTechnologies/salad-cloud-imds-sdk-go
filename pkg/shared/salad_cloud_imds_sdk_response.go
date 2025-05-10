package shared

import (
	"encoding/json"

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

func (r *SaladCloudImdsSdkResponse[T]) GetData() T {
	return r.Data
}

func (r SaladCloudImdsSdkResponse[T]) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: SaladCloudImdsSdkResponse to string"
	}
	return string(jsonData)
}
