# MetadataService

A list of all methods in the `MetadataService` service. Click on the method name to view detailed information about that method.

| Methods                                     | Description                                             |
| :------------------------------------------ | :------------------------------------------------------ |
| [ReallocateContainer](#reallocatecontainer) | Reallocates the running container to another Salad Node |
| [GetContainerStatus](#getcontainerstatus)   | Gets the health statuses of the running container       |
| [GetContainerToken](#getcontainertoken)     | Gets the identity token of the running container        |

## ReallocateContainer

Reallocates the running container to another Salad Node

- HTTP Method: `POST`
- Endpoint: `/v1/reallocate`

**Parameters**

| Name                | Type                | Required | Description                 |
| :------------------ | :------------------ | :------- | :-------------------------- |
| ctx                 | Context             | ✅       | Default go language context |
| reallocateContainer | ReallocateContainer | ✅       |                             |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdk"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/metadata"
)

config := saladcloudimdssdkconfig.NewConfig()
client := saladcloudimdssdk.NewSaladCloudImdsSdk(config)


request := metadata.ReallocateContainer{}
request.SetReason("Reason")

response, err := client.Metadata.ReallocateContainer(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## GetContainerStatus

Gets the health statuses of the running container

- HTTP Method: `GET`
- Endpoint: `/v1/status`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`ContainerStatus`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdk"
)

config := saladcloudimdssdkconfig.NewConfig()
client := saladcloudimdssdk.NewSaladCloudImdsSdk(config)

response, err := client.Metadata.GetContainerStatus(context.Background())
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## GetContainerToken

Gets the identity token of the running container

- HTTP Method: `GET`
- Endpoint: `/v1/token`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`ContainerToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdk"
)

config := saladcloudimdssdkconfig.NewConfig()
client := saladcloudimdssdk.NewSaladCloudImdsSdk(config)

response, err := client.Metadata.GetContainerToken(context.Background())
if err != nil {
  panic(err)
}

fmt.Print(response)
```
