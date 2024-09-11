# ContainerToken

Represents the identity token of the running container.

**Properties**

| Name | Type     | Required | Description                                                                                                                                                                                                        |
| :--- | :------- | :------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Jwt  | `string` | ✅       | The JSON Web Token (JWT) that may be used to identify the running container. The JWT may be verified using the JSON Web Key Set (JWKS) available at https://matrix-rest-api.salad.com/.well-known/stash-jwks.json. |
