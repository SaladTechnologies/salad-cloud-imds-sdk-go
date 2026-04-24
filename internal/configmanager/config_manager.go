package configmanager

import (
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
	"time"
)

// ConfigManager manages configuration across all services with synchronized updates.
// Provides centralized configuration management and OAuth token handling for multiple services.
type ConfigManager struct {
	Metadata saladcloudimdssdkconfig.Config
}

// NewConfigManager creates a new configuration manager with the provided config and optional OAuth token service.
// Initializes service-specific configs and sets up OAuth token management if enabled.
func NewConfigManager(config saladcloudimdssdkconfig.Config) *ConfigManager {
	return &ConfigManager{
		Metadata: config,
	}
}

// SetBaseUrl updates the BaseUrl configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetBaseUrl(baseUrl string) {
	c.Metadata.SetBaseUrl(baseUrl)
}

// SetTimeout updates the Timeout configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetTimeout(timeout time.Duration) {
	c.Metadata.SetTimeout(timeout)
}

// GetMetadata returns the configuration for the Metadata service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetMetadata() *saladcloudimdssdkconfig.Config {
	return &c.Metadata
}
