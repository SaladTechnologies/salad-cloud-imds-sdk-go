package metadata

import "encoding/json"

// The health statuses of the current container instance.
type Status struct {
	// `true` if the running container is ready. If a readiness probe is defined, this returns the latest result of the probe. If a readiness probe is not defined but a startup probe is defined, this returns the same value as the `started` property. If neither a readiness probe nor a startup probe are defined, returns `true`.
	Ready *bool `json:"ready,omitempty" required:"true"`
	// `true` if the running container is started. If a startup probe is defined, this returns the latest result of the probe. If a startup probe is not defined, returns `true`.
	Started *bool `json:"started,omitempty" required:"true"`
}

func (s *Status) GetReady() *bool {
	if s == nil {
		return nil
	}
	return s.Ready
}

func (s *Status) SetReady(ready bool) {
	s.Ready = &ready
}

func (s *Status) GetStarted() *bool {
	if s == nil {
		return nil
	}
	return s.Started
}

func (s *Status) SetStarted(started bool) {
	s.Started = &started
}

func (s Status) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: Status to string"
	}
	return string(jsonData)
}
