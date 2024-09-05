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
