package metadata

// Represents the identity token of the running container.
type ContainerToken struct {
	// The JSON Web Token (JWT) that may be used to identify the running container. The JWT may be verified using the JSON Web Key Set (JWKS) available at https://matrix-rest-api.salad.com/.well-known/stash-jwks.json.
	Jwt *string `json:"jwt,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (c *ContainerToken) SetJwt(jwt string) {
	c.Jwt = &jwt
}

func (c *ContainerToken) GetJwt() *string {
	if c == nil {
		return nil
	}
	return c.Jwt
}
