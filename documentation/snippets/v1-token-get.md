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
