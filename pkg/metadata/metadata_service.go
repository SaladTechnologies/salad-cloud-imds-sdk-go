package metadata

import (
	"context"
	restClient "github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/clients/rest"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/clients/rest/hooks"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/clients/rest/httptransport"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/shared"
	"time"
)

// MetadataService provides methods to interact with MetadataService-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type MetadataService struct {
	manager *configmanager.ConfigManager
	hook    hooks.Hook
}

func NewMetadataService() *MetadataService {
	return &MetadataService{
		manager: configmanager.NewConfigManager(saladcloudimdssdkconfig.Config{}),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *MetadataService) WithConfigManager(manager *configmanager.ConfigManager) *MetadataService {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *MetadataService) WithHook(hook hooks.Hook) *MetadataService {
	api.hook = hook
	return api
}

func (api *MetadataService) getConfig() *saladcloudimdssdkconfig.Config {
	return api.manager.GetMetadata()
}

func (api *MetadataService) getHook() hooks.Hook {
	return api.hook
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
func (api *MetadataService) GetDeletionCost(ctx context.Context, params GetDeletionCostRequestParams) (*shared.SaladCloudImdsSdkResponse[DeletionCost], *shared.SaladCloudImdsSdkError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/v1/deletion-cost").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DeletionCost, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudImdsSdkError[[]byte](err)
	}

	return shared.NewSaladCloudImdsSdkResponse[DeletionCost](resp), nil
}

// Replaces the deletion cost of the current container instance
func (api *MetadataService) ReplaceDeletionCost(ctx context.Context, deletionCost DeletionCost, params ReplaceDeletionCostRequestParams) (*shared.SaladCloudImdsSdkResponse[any], *shared.SaladCloudImdsSdkError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/v1/deletion-cost").
		WithConfig(config).
		WithBody(deletionCost).
		AddHeader("CONTENT-TYPE", "application/json").
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudImdsSdkError[[]byte](err)
	}

	return shared.NewSaladCloudImdsSdkResponse[any](resp), nil
}

// Reallocates the current container instance to another SaladCloud node
func (api *MetadataService) Reallocate(ctx context.Context, reallocatePrototype ReallocatePrototype, params ReallocateRequestParams) (*shared.SaladCloudImdsSdkResponse[any], *shared.SaladCloudImdsSdkError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/v1/reallocate").
		WithConfig(config).
		WithBody(reallocatePrototype).
		AddHeader("CONTENT-TYPE", "application/json").
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudImdsSdkError[[]byte](err)
	}

	return shared.NewSaladCloudImdsSdkResponse[any](resp), nil
}

// Recreates the current container instance on the same SaladCloud node
func (api *MetadataService) Recreate(ctx context.Context, params RecreateRequestParams) (*shared.SaladCloudImdsSdkResponse[any], *shared.SaladCloudImdsSdkError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/v1/recreate").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudImdsSdkError[[]byte](err)
	}

	return shared.NewSaladCloudImdsSdkResponse[any](resp), nil
}

// Restarts the current container instance on the same SaladCloud node
func (api *MetadataService) Restart(ctx context.Context, params RestartRequestParams) (*shared.SaladCloudImdsSdkResponse[any], *shared.SaladCloudImdsSdkError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/v1/restart").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudImdsSdkError[[]byte](err)
	}

	return shared.NewSaladCloudImdsSdkResponse[any](resp), nil
}

// Gets the health statuses of the current container instance
func (api *MetadataService) GetStatus(ctx context.Context, params GetStatusRequestParams) (*shared.SaladCloudImdsSdkResponse[Status], *shared.SaladCloudImdsSdkError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/v1/status").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Status, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudImdsSdkError[[]byte](err)
	}

	return shared.NewSaladCloudImdsSdkResponse[Status](resp), nil
}

// Gets the identity token of the current container instance
func (api *MetadataService) GetToken(ctx context.Context, params GetTokenRequestParams) (*shared.SaladCloudImdsSdkResponse[Token], *shared.SaladCloudImdsSdkError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/v1/token").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Token, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudImdsSdkError[[]byte](err)
	}

	return shared.NewSaladCloudImdsSdkResponse[Token](resp), nil
}
