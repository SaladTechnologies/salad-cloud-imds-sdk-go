package metadata

import (
	"context"
	restClient "github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/clients/rest"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/clients/rest/httptransport"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/shared"
)

type MetadataService struct {
	manager *configmanager.ConfigManager
}

func NewMetadataService(manager *configmanager.ConfigManager) *MetadataService {
	return &MetadataService{
		manager: manager,
	}
}

func (api *MetadataService) getConfig() *saladcloudimdssdkconfig.Config {
	return api.manager.GetMetadata()
}

func (api *MetadataService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

// Reallocates the running container to another Salad Node
func (api *MetadataService) ReallocateContainer(ctx context.Context, reallocateContainer ReallocateContainer) (*shared.SaladCloudImdsSdkResponse[any], *shared.SaladCloudImdsSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "POST", "/v1/reallocate", config)
	request.Headers["Content-Type"] = "application/json"

	request.Body = reallocateContainer

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudImdsSdkError[any](err)
	}

	return shared.NewSaladCloudImdsSdkResponse[any](resp), nil
}

// Gets the health statuses of the running container
func (api *MetadataService) GetContainerStatus(ctx context.Context) (*shared.SaladCloudImdsSdkResponse[ContainerStatus], *shared.SaladCloudImdsSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[ContainerStatus](config)

	request := httptransport.NewRequest(ctx, "GET", "/v1/status", config)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudImdsSdkError[ContainerStatus](err)
	}

	return shared.NewSaladCloudImdsSdkResponse[ContainerStatus](resp), nil
}

// Gets the identity token of the running container
func (api *MetadataService) GetContainerToken(ctx context.Context) (*shared.SaladCloudImdsSdkResponse[ContainerToken], *shared.SaladCloudImdsSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[ContainerToken](config)

	request := httptransport.NewRequest(ctx, "GET", "/v1/token", config)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudImdsSdkError[ContainerToken](err)
	}

	return shared.NewSaladCloudImdsSdkResponse[ContainerToken](resp), nil
}
