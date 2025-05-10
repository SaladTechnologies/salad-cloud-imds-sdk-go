package metadata

import "encoding/json"

// The identity token of the current container instance.
type Token struct {
	// The JSON Web Token (JWT) that may be used to identify the running container. The JWT may be verified using the JSON Web Key Set (JWKS) available at https://matrix-rest-api.salad.com/.well-known/workload-jwks.json.
	Jwt *string `json:"jwt,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (t *Token) GetJwt() *string {
	if t == nil {
		return nil
	}
	return t.Jwt
}

func (t *Token) SetJwt(jwt string) {
	t.Jwt = &jwt
}

func (t Token) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: Token to string"
	}
	return string(jsonData)
}
