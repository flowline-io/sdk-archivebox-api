# \ArchiveBoxCLISubCommandsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1CliCliAdd**](ArchiveBoxCLISubCommandsAPI.md#ApiV1CliCliAdd) | **Post** /api/v1/cli/add | archivebox add [args] [urls]
[**ApiV1CliCliList**](ArchiveBoxCLISubCommandsAPI.md#ApiV1CliCliList) | **Post** /api/v1/cli/list | archivebox list [args] [filter_patterns]
[**ApiV1CliCliRemove**](ArchiveBoxCLISubCommandsAPI.md#ApiV1CliCliRemove) | **Post** /api/v1/cli/remove | archivebox remove [args] [filter_patterns]
[**ApiV1CliCliSchedule**](ArchiveBoxCLISubCommandsAPI.md#ApiV1CliCliSchedule) | **Post** /api/v1/cli/schedule | archivebox schedule [args] [import_path]
[**ApiV1CliCliUpdate**](ArchiveBoxCLISubCommandsAPI.md#ApiV1CliCliUpdate) | **Post** /api/v1/cli/update | archivebox update [args] [filter_patterns]



## ApiV1CliCliAdd

> CLICommandResponseSchema ApiV1CliCliAdd(ctx).AddCommandSchema(addCommandSchema).Execute()

archivebox add [args] [urls]

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	addCommandSchema := *openapiclient.NewAddCommandSchema([]string{"Urls_example"}) // AddCommandSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArchiveBoxCLISubCommandsAPI.ApiV1CliCliAdd(context.Background()).AddCommandSchema(addCommandSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArchiveBoxCLISubCommandsAPI.ApiV1CliCliAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CliCliAdd`: CLICommandResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `ArchiveBoxCLISubCommandsAPI.ApiV1CliCliAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CliCliAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addCommandSchema** | [**AddCommandSchema**](AddCommandSchema.md) |  | 

### Return type

[**CLICommandResponseSchema**](CLICommandResponseSchema.md)

### Authorization

[HeaderTokenAuth](../README.md#HeaderTokenAuth), [QueryParamTokenAuth](../README.md#QueryParamTokenAuth), [UsernameAndPasswordAuth](../README.md#UsernameAndPasswordAuth), [BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1CliCliList

> CLICommandResponseSchema ApiV1CliCliList(ctx).ListCommandSchema(listCommandSchema).Execute()

archivebox list [args] [filter_patterns]

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	listCommandSchema := *openapiclient.NewListCommandSchema() // ListCommandSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArchiveBoxCLISubCommandsAPI.ApiV1CliCliList(context.Background()).ListCommandSchema(listCommandSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArchiveBoxCLISubCommandsAPI.ApiV1CliCliList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CliCliList`: CLICommandResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `ArchiveBoxCLISubCommandsAPI.ApiV1CliCliList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CliCliListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listCommandSchema** | [**ListCommandSchema**](ListCommandSchema.md) |  | 

### Return type

[**CLICommandResponseSchema**](CLICommandResponseSchema.md)

### Authorization

[HeaderTokenAuth](../README.md#HeaderTokenAuth), [QueryParamTokenAuth](../README.md#QueryParamTokenAuth), [UsernameAndPasswordAuth](../README.md#UsernameAndPasswordAuth), [BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1CliCliRemove

> CLICommandResponseSchema ApiV1CliCliRemove(ctx).RemoveCommandSchema(removeCommandSchema).Execute()

archivebox remove [args] [filter_patterns]

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	removeCommandSchema := *openapiclient.NewRemoveCommandSchema() // RemoveCommandSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArchiveBoxCLISubCommandsAPI.ApiV1CliCliRemove(context.Background()).RemoveCommandSchema(removeCommandSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArchiveBoxCLISubCommandsAPI.ApiV1CliCliRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CliCliRemove`: CLICommandResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `ArchiveBoxCLISubCommandsAPI.ApiV1CliCliRemove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CliCliRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeCommandSchema** | [**RemoveCommandSchema**](RemoveCommandSchema.md) |  | 

### Return type

[**CLICommandResponseSchema**](CLICommandResponseSchema.md)

### Authorization

[HeaderTokenAuth](../README.md#HeaderTokenAuth), [QueryParamTokenAuth](../README.md#QueryParamTokenAuth), [UsernameAndPasswordAuth](../README.md#UsernameAndPasswordAuth), [BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1CliCliSchedule

> CLICommandResponseSchema ApiV1CliCliSchedule(ctx).ScheduleCommandSchema(scheduleCommandSchema).Execute()

archivebox schedule [args] [import_path]

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	scheduleCommandSchema := *openapiclient.NewScheduleCommandSchema() // ScheduleCommandSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArchiveBoxCLISubCommandsAPI.ApiV1CliCliSchedule(context.Background()).ScheduleCommandSchema(scheduleCommandSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArchiveBoxCLISubCommandsAPI.ApiV1CliCliSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CliCliSchedule`: CLICommandResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `ArchiveBoxCLISubCommandsAPI.ApiV1CliCliSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CliCliScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scheduleCommandSchema** | [**ScheduleCommandSchema**](ScheduleCommandSchema.md) |  | 

### Return type

[**CLICommandResponseSchema**](CLICommandResponseSchema.md)

### Authorization

[HeaderTokenAuth](../README.md#HeaderTokenAuth), [QueryParamTokenAuth](../README.md#QueryParamTokenAuth), [UsernameAndPasswordAuth](../README.md#UsernameAndPasswordAuth), [BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1CliCliUpdate

> CLICommandResponseSchema ApiV1CliCliUpdate(ctx).UpdateCommandSchema(updateCommandSchema).Execute()

archivebox update [args] [filter_patterns]

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateCommandSchema := *openapiclient.NewUpdateCommandSchema() // UpdateCommandSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArchiveBoxCLISubCommandsAPI.ApiV1CliCliUpdate(context.Background()).UpdateCommandSchema(updateCommandSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArchiveBoxCLISubCommandsAPI.ApiV1CliCliUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CliCliUpdate`: CLICommandResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `ArchiveBoxCLISubCommandsAPI.ApiV1CliCliUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CliCliUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateCommandSchema** | [**UpdateCommandSchema**](UpdateCommandSchema.md) |  | 

### Return type

[**CLICommandResponseSchema**](CLICommandResponseSchema.md)

### Authorization

[HeaderTokenAuth](../README.md#HeaderTokenAuth), [QueryParamTokenAuth](../README.md#QueryParamTokenAuth), [UsernameAndPasswordAuth](../README.md#UsernameAndPasswordAuth), [BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

