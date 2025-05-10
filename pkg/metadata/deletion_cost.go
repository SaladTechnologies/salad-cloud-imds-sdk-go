package metadata

import "encoding/json"

// The deletion cost of the current container instance.
type DeletionCost struct {
	// An integer value that identifies the relative cost to the application running across the container group if the current container instance is deleted. A higher value indicates a higher cost, and a lower value indicates a lower cost. If the container group is scaled down, the scheduler will attempt to delete the container instances with the lowest deletion costs first.
	DeletionCost *int64 `json:"deletion_cost,omitempty" required:"true" min:"-2147483648" max:"2147483647"`
}

func (d *DeletionCost) GetDeletionCost() *int64 {
	if d == nil {
		return nil
	}
	return d.DeletionCost
}

func (d *DeletionCost) SetDeletionCost(deletionCost int64) {
	d.DeletionCost = &deletionCost
}

func (d DeletionCost) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeletionCost to string"
	}
	return string(jsonData)
}
