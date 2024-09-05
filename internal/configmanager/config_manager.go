package configmanager

import "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"

type ConfigManager struct {
	Metadata saladcloudimdssdkconfig.Config
}

func NewConfigManager(config saladcloudimdssdkconfig.Config) *ConfigManager {
	return &ConfigManager{
		Metadata: config,
	}
}

func (c *ConfigManager) SetBaseUrl(baseUrl string) {
	c.Metadata.SetBaseUrl(baseUrl)
}

func (c *ConfigManager) GetMetadata() *saladcloudimdssdkconfig.Config {
	return &c.Metadata
}
