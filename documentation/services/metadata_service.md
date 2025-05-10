# MetadataService

A list of all methods in the `MetadataService` service. Click on the method name to view detailed information about that method.

| Methods                                     | Description                                                           |
| :------------------------------------------ | :-------------------------------------------------------------------- |
| [GetDeletionCost](#getdeletioncost)         | Gets the deletion cost of the current container instance              |
| [ReplaceDeletionCost](#replacedeletioncost) | Replaces the deletion cost of the current container instance          |
| [Reallocate](#reallocate)                   | Reallocates the current container instance to another SaladCloud node |
| [Recreate](#recreate)                       | Recreates the current container instance on the same SaladCloud node  |
| [Restart](#restart)                         | Restarts the current container instance on the same SaladCloud node   |
| [GetStatus](#getstatus)                     | Gets the health statuses of the current container instance            |
| [GetToken](#gettoken)                       | Gets the identity token of the current container instance             |

## GetDeletionCost

Gets the deletion cost of the current container instance

- HTTP Method: `GET`
- Endpoint: `/v1/deletion-cost`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`DeletionCost`

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

response, err := client.Metadata.GetDeletionCost(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ReplaceDeletionCost

Replaces the deletion cost of the current container instance

- HTTP Method: `PUT`
- Endpoint: `/v1/deletion-cost`

**Parameters**

| Name         | Type         | Required | Description                 |
| :----------- | :----------- | :------- | :-------------------------- |
| ctx          | Context      | ✅       | Default go language context |
| deletionCost | DeletionCost | ✅       |                             |

**Return Type**

`DeletionCost`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdk"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/metadata"
)

config := saladcloudimdssdkconfig.NewConfig()
client := saladcloudimdssdk.NewSaladCloudImdsSdk(config)


request := metadata.DeletionCost{
  DeletionCost: util.ToPointer(int64(123)),
}

response, err := client.Metadata.ReplaceDeletionCost(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## Reallocate

Reallocates the current container instance to another SaladCloud node

- HTTP Method: `POST`
- Endpoint: `/v1/reallocate`

**Parameters**

| Name                | Type                | Required | Description                 |
| :------------------ | :------------------ | :------- | :-------------------------- |
| ctx                 | Context             | ✅       | Default go language context |
| reallocatePrototype | ReallocatePrototype | ✅       |                             |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdk"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/metadata"
)

config := saladcloudimdssdkconfig.NewConfig()
client := saladcloudimdssdk.NewSaladCloudImdsSdk(config)


request := metadata.ReallocatePrototype{
  Reason: util.ToPointer("Reason"),
}

response, err := client.Metadata.Reallocate(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## Recreate

Recreates the current container instance on the same SaladCloud node

- HTTP Method: `POST`
- Endpoint: `/v1/recreate`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`any`

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

response, err := client.Metadata.Recreate(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## Restart

Restarts the current container instance on the same SaladCloud node

- HTTP Method: `POST`
- Endpoint: `/v1/restart`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`any`

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

response, err := client.Metadata.Restart(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetStatus

Gets the health statuses of the current container instance

- HTTP Method: `GET`
- Endpoint: `/v1/status`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`Status`

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

response, err := client.Metadata.GetStatus(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetToken

Gets the identity token of the current container instance

- HTTP Method: `GET`
- Endpoint: `/v1/token`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`Token`

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

response, err := client.Metadata.GetToken(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```
