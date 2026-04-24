package saladcloudimdssdk

import (
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/clients/rest/hooks"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/metadata"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
	"time"
)

// SaladCloudImdsSdk is the main SDK client that provides access to all service endpoints.
// It manages configuration, authentication, and service instances with centralized settings.
type SaladCloudImdsSdk struct {
	Metadata *metadata.MetadataService
	manager  *configmanager.ConfigManager
}

func NewSaladCloudImdsSdk(config saladcloudimdssdkconfig.Config) *SaladCloudImdsSdk {
	metadata := metadata.NewMetadataService()

	manager := configmanager.NewConfigManager(config)
	hook := hooks.NewDefaultHook()
	metadata.WithConfigManager(manager)
	metadata.WithHook(hook)

	return &SaladCloudImdsSdk{
		Metadata: metadata,
		manager:  manager,
	}
}

func (s *SaladCloudImdsSdk) SetBaseUrl(baseUrl string) {
	s.manager.SetBaseUrl(baseUrl)
}

func (s *SaladCloudImdsSdk) SetTimeout(timeout time.Duration) {
	s.manager.SetTimeout(timeout)
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
