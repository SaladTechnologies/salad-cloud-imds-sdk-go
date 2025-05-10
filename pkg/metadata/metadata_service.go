package metadata

import (
	"context"
	restClient "github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/clients/rest"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/clients/rest/httptransport"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/shared"
	"time"
)

type MetadataService struct {
	manager *configmanager.ConfigManager
}

func NewMetadataService() *MetadataService {
	return &MetadataService{
		manager: configmanager.NewConfigManager(saladcloudimdssdkconfig.Config{}),
	}
}

func (api *MetadataService) WithConfigManager(manager *configmanager.ConfigManager) *MetadataService {
	api.manager = manager
	return api
}

func (api *MetadataService) getConfig() *saladcloudimdssdkconfig.Config {
	return api.manager.GetMetadata()
}

func (api *MetadataService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *MetadataService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

// Gets the deletion cost of the current container instance
func (api *MetadataService) GetDeletionCost(ctx context.Context) (*shared.SaladCloudImdsSdkResponse[DeletionCost], *shared.SaladCloudImdsSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/v1/deletion-cost").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DeletionCost](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudImdsSdkError[DeletionCost](err)
	}

	return shared.NewSaladCloudImdsSdkResponse[DeletionCost](resp), nil
}

// Replaces the deletion cost of the current container instance
func (api *MetadataService) ReplaceDeletionCost(ctx context.Context, deletionCost DeletionCost) (*shared.SaladCloudImdsSdkResponse[DeletionCost], *shared.SaladCloudImdsSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/v1/deletion-cost").
		WithConfig(config).
		WithBody(deletionCost).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DeletionCost](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudImdsSdkError[DeletionCost](err)
	}

	return shared.NewSaladCloudImdsSdkResponse[DeletionCost](resp), nil
}

// Reallocates the current container instance to another SaladCloud node
func (api *MetadataService) Reallocate(ctx context.Context, reallocatePrototype ReallocatePrototype) (*shared.SaladCloudImdsSdkResponse[any], *shared.SaladCloudImdsSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/v1/reallocate").
		WithConfig(config).
		WithBody(reallocatePrototype).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudImdsSdkError[any](err)
	}

	return shared.NewSaladCloudImdsSdkResponse[any](resp), nil
}

// Recreates the current container instance on the same SaladCloud node
func (api *MetadataService) Recreate(ctx context.Context) (*shared.SaladCloudImdsSdkResponse[any], *shared.SaladCloudImdsSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/v1/recreate").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudImdsSdkError[any](err)
	}

	return shared.NewSaladCloudImdsSdkResponse[any](resp), nil
}

// Restarts the current container instance on the same SaladCloud node
func (api *MetadataService) Restart(ctx context.Context) (*shared.SaladCloudImdsSdkResponse[any], *shared.SaladCloudImdsSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/v1/restart").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudImdsSdkError[any](err)
	}

	return shared.NewSaladCloudImdsSdkResponse[any](resp), nil
}

// Gets the health statuses of the current container instance
func (api *MetadataService) GetStatus(ctx context.Context) (*shared.SaladCloudImdsSdkResponse[Status], *shared.SaladCloudImdsSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/v1/status").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Status](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudImdsSdkError[Status](err)
	}

	return shared.NewSaladCloudImdsSdkResponse[Status](resp), nil
}

// Gets the identity token of the current container instance
func (api *MetadataService) GetToken(ctx context.Context) (*shared.SaladCloudImdsSdkResponse[Token], *shared.SaladCloudImdsSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/v1/token").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Token](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudImdsSdkError[Token](err)
	}

	return shared.NewSaladCloudImdsSdkResponse[Token](resp), nil
}
