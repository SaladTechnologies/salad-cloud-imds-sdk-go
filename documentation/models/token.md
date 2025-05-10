# Token

The identity token of the current container instance.

**Properties**

| Name | Type   | Required | Description                                                                                                                                                                                                           |
| :--- | :----- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Jwt  | string | ✅       | The JSON Web Token (JWT) that may be used to identify the running container. The JWT may be verified using the JSON Web Key Set (JWKS) available at https://matrix-rest-api.salad.com/.well-known/workload-jwks.json. |
