package metadata

// GetDeletionCostRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type GetDeletionCostRequestParams struct {
	Metadata *Metadata `explode:"false" serializationStyle:"simple" headerParam:"Metadata" required:"true"`
}

// SetMetadata sets the Metadata parameter.
func (params *GetDeletionCostRequestParams) SetMetadata(metadata Metadata) {
	params.Metadata = &metadata
}

// ReplaceDeletionCostRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type ReplaceDeletionCostRequestParams struct {
	Metadata *Metadata `explode:"false" serializationStyle:"simple" headerParam:"Metadata" required:"true"`
}

// SetMetadata sets the Metadata parameter.
func (params *ReplaceDeletionCostRequestParams) SetMetadata(metadata Metadata) {
	params.Metadata = &metadata
}

// ReallocateRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type ReallocateRequestParams struct {
	Metadata *Metadata `explode:"false" serializationStyle:"simple" headerParam:"Metadata" required:"true"`
}

// SetMetadata sets the Metadata parameter.
func (params *ReallocateRequestParams) SetMetadata(metadata Metadata) {
	params.Metadata = &metadata
}

// RecreateRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type RecreateRequestParams struct {
	Metadata *Metadata `explode:"false" serializationStyle:"simple" headerParam:"Metadata" required:"true"`
}

// SetMetadata sets the Metadata parameter.
func (params *RecreateRequestParams) SetMetadata(metadata Metadata) {
	params.Metadata = &metadata
}

// RestartRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type RestartRequestParams struct {
	Metadata *Metadata `explode:"false" serializationStyle:"simple" headerParam:"Metadata" required:"true"`
}

// SetMetadata sets the Metadata parameter.
func (params *RestartRequestParams) SetMetadata(metadata Metadata) {
	params.Metadata = &metadata
}

// GetStatusRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type GetStatusRequestParams struct {
	Metadata *Metadata `explode:"false" serializationStyle:"simple" headerParam:"Metadata" required:"true"`
}

// SetMetadata sets the Metadata parameter.
func (params *GetStatusRequestParams) SetMetadata(metadata Metadata) {
	params.Metadata = &metadata
}

// GetTokenRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type GetTokenRequestParams struct {
	Metadata *Metadata `explode:"false" serializationStyle:"simple" headerParam:"Metadata" required:"true"`
}

// SetMetadata sets the Metadata parameter.
func (params *GetTokenRequestParams) SetMetadata(metadata Metadata) {
	params.Metadata = &metadata
}
