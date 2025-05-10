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
