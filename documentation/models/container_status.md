# ContainerStatus

Represents the health statuses of the running container.

**Properties**

| Name    | Type   | Required | Description                                                                                                                                                                                                                                                                                                                      |
| :------ | :----- | :------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Ready   | `bool` | ✅       | `true` if the running container is ready. If a readiness probe is defined, this returns the latest result of the probe. If a readiness probe is not defined but a startup probe is defined, this returns the same value as the `started` property. If neither a readiness probe nor a startup probe are defined, returns `true`. |
| Started | `bool` | ✅       | `true` if the running container is started. If a startup probe is defined, this returns the latest result of the probe. If a startup probe is not defined, returns `true`.                                                                                                                                                       |
