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

| Name   | Type                         | Required | Description                   |
| :----- | :--------------------------- | :------- | :---------------------------- |
| ctx    | Context                      | ✅       | Default go language context   |
| params | GetDeletionCostRequestParams | ✅       | Additional request parameters |

**Return Type**

`DeletionCost`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdk"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/metadata"
)

config := saladcloudimdssdkconfig.NewConfig()

client := saladcloudimdssdk.NewSaladCloudImdsSdk(config)

metadata := metadata.METADATA_TRUE

params := metadata.GetDeletionCostRequestParams{
  Metadata: &metadata,
}

response, err := client.Metadata.GetDeletionCost(context.Background(), params)
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

| Name         | Type                             | Required | Description                   |
| :----------- | :------------------------------- | :------- | :---------------------------- |
| ctx          | Context                          | ✅       | Default go language context   |
| deletionCost | DeletionCost                     | ✅       |                               |
| params       | ReplaceDeletionCostRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdk"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/metadata"
)

config := saladcloudimdssdkconfig.NewConfig()

client := saladcloudimdssdk.NewSaladCloudImdsSdk(config)

metadata := metadata.METADATA_TRUE

params := metadata.ReplaceDeletionCostRequestParams{
  Metadata: &metadata,
}


request := metadata.DeletionCost{
  DeletionCost: util.ToPointer(int64(100)),
}

response, err := client.Metadata.ReplaceDeletionCost(context.Background(), request, params)
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

| Name                | Type                    | Required | Description                   |
| :------------------ | :---------------------- | :------- | :---------------------------- |
| ctx                 | Context                 | ✅       | Default go language context   |
| reallocatePrototype | ReallocatePrototype     | ✅       |                               |
| params              | ReallocateRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdk"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/metadata"
)

config := saladcloudimdssdkconfig.NewConfig()

client := saladcloudimdssdk.NewSaladCloudImdsSdk(config)

metadata := metadata.METADATA_TRUE

params := metadata.ReallocateRequestParams{
  Metadata: &metadata,
}


request := metadata.ReallocatePrototype{
  Reason: util.ToPointer("Insufficient VRAM"),
}

response, err := client.Metadata.Reallocate(context.Background(), request, params)
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

| Name   | Type                  | Required | Description                   |
| :----- | :-------------------- | :------- | :---------------------------- |
| ctx    | Context               | ✅       | Default go language context   |
| params | RecreateRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdk"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/metadata"
)

config := saladcloudimdssdkconfig.NewConfig()

client := saladcloudimdssdk.NewSaladCloudImdsSdk(config)

metadata := metadata.METADATA_TRUE

params := metadata.RecreateRequestParams{
  Metadata: &metadata,
}

response, err := client.Metadata.Recreate(context.Background(), params)
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

| Name   | Type                 | Required | Description                   |
| :----- | :------------------- | :------- | :---------------------------- |
| ctx    | Context              | ✅       | Default go language context   |
| params | RestartRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdk"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/metadata"
)

config := saladcloudimdssdkconfig.NewConfig()

client := saladcloudimdssdk.NewSaladCloudImdsSdk(config)

metadata := metadata.METADATA_TRUE

params := metadata.RestartRequestParams{
  Metadata: &metadata,
}

response, err := client.Metadata.Restart(context.Background(), params)
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

| Name   | Type                   | Required | Description                   |
| :----- | :--------------------- | :------- | :---------------------------- |
| ctx    | Context                | ✅       | Default go language context   |
| params | GetStatusRequestParams | ✅       | Additional request parameters |

**Return Type**

`Status`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdk"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/metadata"
)

config := saladcloudimdssdkconfig.NewConfig()

client := saladcloudimdssdk.NewSaladCloudImdsSdk(config)

metadata := metadata.METADATA_TRUE

params := metadata.GetStatusRequestParams{
  Metadata: &metadata,
}

response, err := client.Metadata.GetStatus(context.Background(), params)
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

| Name   | Type                  | Required | Description                   |
| :----- | :-------------------- | :------- | :---------------------------- |
| ctx    | Context               | ✅       | Default go language context   |
| params | GetTokenRequestParams | ✅       | Additional request parameters |

**Return Type**

`Token`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdk"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/metadata"
)

config := saladcloudimdssdkconfig.NewConfig()

client := saladcloudimdssdk.NewSaladCloudImdsSdk(config)

metadata := metadata.METADATA_TRUE

params := metadata.GetTokenRequestParams{
  Metadata: &metadata,
}

response, err := client.Metadata.GetToken(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
