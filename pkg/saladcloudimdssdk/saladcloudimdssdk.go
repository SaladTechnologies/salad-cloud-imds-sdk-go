package saladcloudimdssdk

import (
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/metadata"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
)

type SaladCloudImdsSdk struct {
	Metadata *metadata.MetadataService
	manager  *configmanager.ConfigManager
}

func NewSaladCloudImdsSdk(config saladcloudimdssdkconfig.Config) *SaladCloudImdsSdk {
	manager := configmanager.NewConfigManager(config)
	return &SaladCloudImdsSdk{
		Metadata: metadata.NewMetadataService(manager),
		manager:  manager,
	}
}

func (s *SaladCloudImdsSdk) SetBaseUrl(baseUrl string) {
	s.manager.SetBaseUrl(baseUrl)
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
