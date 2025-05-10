# DeletionCost

The deletion cost of the current container instance.

**Properties**

| Name         | Type  | Required | Description                                                                                                                                                                                                                                                                                                                                                                        |
| :----------- | :---- | :------- | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| DeletionCost | int64 | ✅       | An integer value that identifies the relative cost to the application running across the container group if the current container instance is deleted. A higher value indicates a higher cost, and a lower value indicates a lower cost. If the container group is scaled down, the scheduler will attempt to delete the container instances with the lowest deletion costs first. |
