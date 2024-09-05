package metadata

// Represents a request to reallocate a container.
type ReallocateContainer struct {
	// The reason for reallocating the container. This value is reported to SaladCloud support for quality assurance of Salad Nodes.
	Reason *string `json:"reason,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (r *ReallocateContainer) SetReason(reason string) {
	r.Reason = &reason
}

func (r *ReallocateContainer) GetReason() *string {
	if r == nil {
		return nil
	}
	return r.Reason
}
