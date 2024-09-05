package metadata

// Represents the health statuses of the running container.
type ContainerStatus struct {
	// `true` if the running container is ready. If a readiness probe is defined, this returns the latest result of the probe. If a readiness probe is not defined but a startup probe is defined, this returns the same value as the `started` property. If neither a readiness probe nor a startup probe are defined, returns `true`.
	Ready *bool `json:"ready,omitempty" required:"true"`
	// `true` if the running container is started. If a startup probe is defined, this returns the latest result of the probe. If a startup probe is not defined, returns `true`.
	Started *bool `json:"started,omitempty" required:"true"`
}

func (c *ContainerStatus) SetReady(ready bool) {
	c.Ready = &ready
}

func (c *ContainerStatus) GetReady() *bool {
	if c == nil {
		return nil
	}
	return c.Ready
}

func (c *ContainerStatus) SetStarted(started bool) {
	c.Started = &started
}

func (c *ContainerStatus) GetStarted() *bool {
	if c == nil {
		return nil
	}
	return c.Started
}
