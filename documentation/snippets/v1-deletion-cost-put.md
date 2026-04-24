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
