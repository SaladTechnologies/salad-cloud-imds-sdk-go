# SaladCloudImdsSdk Go SDK 0.9.0-alpha.3

Welcome to the SaladCloudImdsSdk SDK documentation. This guide will help you get started with integrating and using the SaladCloudImdsSdk SDK in your project.

## Versions

- API version: `0.9.0-alpha.1`
- SDK version: `0.9.0-alpha.3`

## About the API

The SaladCloud Instance Metadata Service (IMDS). Please refer to the [SaladCloud API Documentation](https://docs.salad.com/api-reference) for more details.

## Table of Contents

- [Setup & Configuration](#setup--configuration)
  - [Supported Language Versions](#supported-language-versions)
  - [Installation](#installation)
- [Services](#services)
  - [Response Wrappers](#response-wrappers)
- [Models](#models)
- [License](#license)

# Setup & Configuration

## Supported Language Versions

This SDK is compatible with the following versions: `Go >= 1.19.0`

## Services

The SDK provides various services to interact with the API.

<details> 
<summary>Below is a list of all available services with links to their detailed documentation:</summary>

| Name                                                          |
| :------------------------------------------------------------ |
| [MetadataService](documentation/services/metadata_service.md) |

</details>

### Response Wrappers

All services use response wrappers to provide a consistent interface to return the responses from the API.

The response wrapper itself is a generic struct that contains the response data and metadata.

<details>
<summary>Below are the response wrappers used in the SDK:</summary>

#### `SaladCloudImdsSdkResponse[T]`

This response wrapper is used to return the response data from the API. It contains the following fields:

| Name     | Type                                | Description                                 |
| :------- | :---------------------------------- | :------------------------------------------ |
| Data     | `T`                                 | The body of the API response                |
| Metadata | `SaladCloudImdsSdkResponseMetadata` | Status code and headers returned by the API |

#### `SaladCloudImdsSdkError`

This response wrapper is used to return an error. It contains the following fields:

| Name     | Type                                | Description                                 |
| :------- | :---------------------------------- | :------------------------------------------ |
| Err      | `error`                             | The error that occurred                     |
| Body     | `T`                                 | The body of the API response                |
| Metadata | `SaladCloudImdsSdkResponseMetadata` | Status code and headers returned by the API |

#### `SaladCloudImdsSdkResponseMetadata`

This struct is shared by both response wrappers and contains the following fields:

| Name       | Type                | Description                                      |
| :--------- | :------------------ | :----------------------------------------------- |
| Headers    | `map[string]string` | A map containing the headers returned by the API |
| StatusCode | `int`               | The status code returned by the API              |

</details>

## Models

The SDK includes several models that represent the data structures used in API requests and responses. These models help in organizing and managing the data efficiently.

<details> 
<summary>Below is a list of all available models with links to their detailed documentation:</summary>

| Name                                                                | Description                                              |
| :------------------------------------------------------------------ | :------------------------------------------------------- |
| [ReallocateContainer](documentation/models/reallocate_container.md) | Represents a request to reallocate a container.          |
| [ContainerStatus](documentation/models/container_status.md)         | Represents the health statuses of the running container. |
| [ContainerToken](documentation/models/container_token.md)           | Represents the identity token of the running container.  |

</details>

## License

This SDK is licensed under the MIT License.

See the [LICENSE](LICENSE) file for more details.
