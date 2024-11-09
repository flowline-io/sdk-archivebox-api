# \AuthenticationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1AuthCheckApiToken**](AuthenticationAPI.md#ApiV1AuthCheckApiToken) | **Post** /api/v1/auth/check_api_token | Validate an API token to make sure its valid and non-expired
[**ApiV1AuthGetApiToken**](AuthenticationAPI.md#ApiV1AuthGetApiToken) | **Post** /api/v1/auth/get_api_token | Generate an API token for a given username &amp; password (or currently logged-in user)



## ApiV1AuthCheckApiToken

> ApiV1AuthCheckApiToken(ctx).TokenAuthSchema(tokenAuthSchema).Execute()

Validate an API token to make sure its valid and non-expired

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
	tokenAuthSchema := *openapiclient.NewTokenAuthSchema("Token_example") // TokenAuthSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.ApiV1AuthCheckApiToken(context.Background()).TokenAuthSchema(tokenAuthSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.ApiV1AuthCheckApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1AuthCheckApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenAuthSchema** | [**TokenAuthSchema**](TokenAuthSchema.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1AuthGetApiToken

> ApiV1AuthGetApiToken(ctx).PasswordAuthSchema(passwordAuthSchema).Execute()

Generate an API token for a given username & password (or currently logged-in user)

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
	passwordAuthSchema := *openapiclient.NewPasswordAuthSchema() // PasswordAuthSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.ApiV1AuthGetApiToken(context.Background()).PasswordAuthSchema(passwordAuthSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.ApiV1AuthGetApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1AuthGetApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordAuthSchema** | [**PasswordAuthSchema**](PasswordAuthSchema.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

