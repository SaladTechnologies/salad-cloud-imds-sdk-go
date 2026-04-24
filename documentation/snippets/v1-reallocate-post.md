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
