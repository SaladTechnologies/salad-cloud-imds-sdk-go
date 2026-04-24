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
