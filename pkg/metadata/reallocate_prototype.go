package metadata

import "encoding/json"

// Represents a request to reallocate the current container instance to another SaladCloud node.
type ReallocatePrototype struct {
	// The reason for reallocating the current container instance. This value is reported to SaladCloud support for quality assurance purposes of SaladCloud nodes.
	Reason *string `json:"reason,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (r *ReallocatePrototype) GetReason() *string {
	if r == nil {
		return nil
	}
	return r.Reason
}

func (r *ReallocatePrototype) SetReason(reason string) {
	r.Reason = &reason
}

func (r ReallocatePrototype) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReallocatePrototype to string"
	}
	return string(jsonData)
}
