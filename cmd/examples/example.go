package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/metadata"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdk"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
)

func main() {
	loadEnv()

	config := saladcloudimdssdkconfig.NewConfig()
	client := saladcloudimdssdk.NewSaladCloudImdsSdk(config)

	request := metadata.ReallocateContainer{}
	request.SetReason("Reason")

	response, err := client.Metadata.ReallocateContainer(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", response)
}

func loadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
