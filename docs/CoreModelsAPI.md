# \CoreModelsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1CoreGetAny**](CoreModelsAPI.md#ApiV1CoreGetAny) | **Get** /api/v1/core/any/{abid} | Get Any
[**ApiV1CoreGetArchiveresult**](CoreModelsAPI.md#ApiV1CoreGetArchiveresult) | **Get** /api/v1/core/archiveresult/{archiveresult_id} | Get Archiveresult
[**ApiV1CoreGetArchiveresults**](CoreModelsAPI.md#ApiV1CoreGetArchiveresults) | **Get** /api/v1/core/archiveresults | Get Archiveresults
[**ApiV1CoreGetSnapshot**](CoreModelsAPI.md#ApiV1CoreGetSnapshot) | **Get** /api/v1/core/snapshot/{snapshot_id} | Get Snapshot
[**ApiV1CoreGetSnapshots**](CoreModelsAPI.md#ApiV1CoreGetSnapshots) | **Get** /api/v1/core/snapshots | Get Snapshots
[**ApiV1CoreGetTag**](CoreModelsAPI.md#ApiV1CoreGetTag) | **Get** /api/v1/core/tag/{tag_id} | Get Tag
[**ApiV1CoreGetTags**](CoreModelsAPI.md#ApiV1CoreGetTags) | **Get** /api/v1/core/tags | Get Tags



## ApiV1CoreGetAny

> Response ApiV1CoreGetAny(ctx, abid).Execute()

Get Any

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
	abid := "abid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreModelsAPI.ApiV1CoreGetAny(context.Background(), abid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreModelsAPI.ApiV1CoreGetAny``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CoreGetAny`: Response
	fmt.Fprintf(os.Stdout, "Response from `CoreModelsAPI.ApiV1CoreGetAny`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**abid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CoreGetAnyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Response**](Response.md)

### Authorization

[HeaderTokenAuth](../README.md#HeaderTokenAuth), [QueryParamTokenAuth](../README.md#QueryParamTokenAuth), [UsernameAndPasswordAuth](../README.md#UsernameAndPasswordAuth), [BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1CoreGetArchiveresult

> ArchiveResultSchema ApiV1CoreGetArchiveresult(ctx, archiveresultId).Execute()

Get Archiveresult



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
	archiveresultId := "archiveresultId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreModelsAPI.ApiV1CoreGetArchiveresult(context.Background(), archiveresultId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreModelsAPI.ApiV1CoreGetArchiveresult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CoreGetArchiveresult`: ArchiveResultSchema
	fmt.Fprintf(os.Stdout, "Response from `CoreModelsAPI.ApiV1CoreGetArchiveresult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**archiveresultId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CoreGetArchiveresultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArchiveResultSchema**](ArchiveResultSchema.md)

### Authorization

[HeaderTokenAuth](../README.md#HeaderTokenAuth), [QueryParamTokenAuth](../README.md#QueryParamTokenAuth), [UsernameAndPasswordAuth](../README.md#UsernameAndPasswordAuth), [BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1CoreGetArchiveresults

> PagedArchiveResultSchema ApiV1CoreGetArchiveresults(ctx).Id(id).Search(search).SnapshotId(snapshotId).SnapshotUrl(snapshotUrl).SnapshotTag(snapshotTag).Status(status).Output(output).Extractor(extractor).Cmd(cmd).Pwd(pwd).CmdVersion(cmdVersion).CreatedAt(createdAt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).Limit(limit).Offset(offset).Page(page).Execute()

Get Archiveresults



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	snapshotId := "snapshotId_example" // string |  (optional)
	snapshotUrl := "snapshotUrl_example" // string |  (optional)
	snapshotTag := "snapshotTag_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	output := "output_example" // string |  (optional)
	extractor := "extractor_example" // string |  (optional)
	cmd := "cmd_example" // string |  (optional)
	pwd := "pwd_example" // string |  (optional)
	cmdVersion := "cmdVersion_example" // string |  (optional)
	createdAt := time.Now() // time.Time |  (optional)
	createdAtGte := time.Now() // time.Time |  (optional)
	createdAtLt := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 200)
	offset := int32(56) // int32 |  (optional) (default to 0)
	page := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreModelsAPI.ApiV1CoreGetArchiveresults(context.Background()).Id(id).Search(search).SnapshotId(snapshotId).SnapshotUrl(snapshotUrl).SnapshotTag(snapshotTag).Status(status).Output(output).Extractor(extractor).Cmd(cmd).Pwd(pwd).CmdVersion(cmdVersion).CreatedAt(createdAt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).Limit(limit).Offset(offset).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreModelsAPI.ApiV1CoreGetArchiveresults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CoreGetArchiveresults`: PagedArchiveResultSchema
	fmt.Fprintf(os.Stdout, "Response from `CoreModelsAPI.ApiV1CoreGetArchiveresults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CoreGetArchiveresultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **search** | **string** |  | 
 **snapshotId** | **string** |  | 
 **snapshotUrl** | **string** |  | 
 **snapshotTag** | **string** |  | 
 **status** | **string** |  | 
 **output** | **string** |  | 
 **extractor** | **string** |  | 
 **cmd** | **string** |  | 
 **pwd** | **string** |  | 
 **cmdVersion** | **string** |  | 
 **createdAt** | **time.Time** |  | 
 **createdAtGte** | **time.Time** |  | 
 **createdAtLt** | **time.Time** |  | 
 **limit** | **int32** |  | [default to 200]
 **offset** | **int32** |  | [default to 0]
 **page** | **int32** |  | [default to 0]

### Return type

[**PagedArchiveResultSchema**](PagedArchiveResultSchema.md)

### Authorization

[HeaderTokenAuth](../README.md#HeaderTokenAuth), [QueryParamTokenAuth](../README.md#QueryParamTokenAuth), [UsernameAndPasswordAuth](../README.md#UsernameAndPasswordAuth), [BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1CoreGetSnapshot

> SnapshotSchema ApiV1CoreGetSnapshot(ctx, snapshotId).WithArchiveresults(withArchiveresults).Execute()

Get Snapshot



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
	snapshotId := "snapshotId_example" // string | 
	withArchiveresults := true // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreModelsAPI.ApiV1CoreGetSnapshot(context.Background(), snapshotId).WithArchiveresults(withArchiveresults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreModelsAPI.ApiV1CoreGetSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CoreGetSnapshot`: SnapshotSchema
	fmt.Fprintf(os.Stdout, "Response from `CoreModelsAPI.ApiV1CoreGetSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CoreGetSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withArchiveresults** | **bool** |  | [default to true]

### Return type

[**SnapshotSchema**](SnapshotSchema.md)

### Authorization

[HeaderTokenAuth](../README.md#HeaderTokenAuth), [QueryParamTokenAuth](../README.md#QueryParamTokenAuth), [UsernameAndPasswordAuth](../README.md#UsernameAndPasswordAuth), [BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1CoreGetSnapshots

> PagedSnapshotSchema ApiV1CoreGetSnapshots(ctx).Id(id).Abid(abid).CreatedById(createdById).CreatedByUsername(createdByUsername).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAt(createdAt).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLt(modifiedAtLt).Search(search).Url(url).Tag(tag).Title(title).Timestamp(timestamp).BookmarkedAtGte(bookmarkedAtGte).BookmarkedAtLt(bookmarkedAtLt).WithArchiveresults(withArchiveresults).Limit(limit).Offset(offset).Page(page).Execute()

Get Snapshots



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string |  (optional)
	abid := "abid_example" // string |  (optional)
	createdById := "createdById_example" // string |  (optional)
	createdByUsername := "createdByUsername_example" // string |  (optional)
	createdAtGte := time.Now() // time.Time |  (optional)
	createdAtLt := time.Now() // time.Time |  (optional)
	createdAt := time.Now() // time.Time |  (optional)
	modifiedAt := time.Now() // time.Time |  (optional)
	modifiedAtGte := time.Now() // time.Time |  (optional)
	modifiedAtLt := time.Now() // time.Time |  (optional)
	search := "search_example" // string |  (optional)
	url := "url_example" // string |  (optional)
	tag := "tag_example" // string |  (optional)
	title := "title_example" // string |  (optional)
	timestamp := "timestamp_example" // string |  (optional)
	bookmarkedAtGte := time.Now() // time.Time |  (optional)
	bookmarkedAtLt := time.Now() // time.Time |  (optional)
	withArchiveresults := true // bool |  (optional) (default to false)
	limit := int32(56) // int32 |  (optional) (default to 200)
	offset := int32(56) // int32 |  (optional) (default to 0)
	page := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreModelsAPI.ApiV1CoreGetSnapshots(context.Background()).Id(id).Abid(abid).CreatedById(createdById).CreatedByUsername(createdByUsername).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAt(createdAt).ModifiedAt(modifiedAt).ModifiedAtGte(modifiedAtGte).ModifiedAtLt(modifiedAtLt).Search(search).Url(url).Tag(tag).Title(title).Timestamp(timestamp).BookmarkedAtGte(bookmarkedAtGte).BookmarkedAtLt(bookmarkedAtLt).WithArchiveresults(withArchiveresults).Limit(limit).Offset(offset).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreModelsAPI.ApiV1CoreGetSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CoreGetSnapshots`: PagedSnapshotSchema
	fmt.Fprintf(os.Stdout, "Response from `CoreModelsAPI.ApiV1CoreGetSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CoreGetSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **abid** | **string** |  | 
 **createdById** | **string** |  | 
 **createdByUsername** | **string** |  | 
 **createdAtGte** | **time.Time** |  | 
 **createdAtLt** | **time.Time** |  | 
 **createdAt** | **time.Time** |  | 
 **modifiedAt** | **time.Time** |  | 
 **modifiedAtGte** | **time.Time** |  | 
 **modifiedAtLt** | **time.Time** |  | 
 **search** | **string** |  | 
 **url** | **string** |  | 
 **tag** | **string** |  | 
 **title** | **string** |  | 
 **timestamp** | **string** |  | 
 **bookmarkedAtGte** | **time.Time** |  | 
 **bookmarkedAtLt** | **time.Time** |  | 
 **withArchiveresults** | **bool** |  | [default to false]
 **limit** | **int32** |  | [default to 200]
 **offset** | **int32** |  | [default to 0]
 **page** | **int32** |  | [default to 0]

### Return type

[**PagedSnapshotSchema**](PagedSnapshotSchema.md)

### Authorization

[HeaderTokenAuth](../README.md#HeaderTokenAuth), [QueryParamTokenAuth](../README.md#QueryParamTokenAuth), [UsernameAndPasswordAuth](../README.md#UsernameAndPasswordAuth), [BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1CoreGetTag

> TagSchema ApiV1CoreGetTag(ctx, tagId).WithSnapshots(withSnapshots).Execute()

Get Tag

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
	tagId := "tagId_example" // string | 
	withSnapshots := true // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreModelsAPI.ApiV1CoreGetTag(context.Background(), tagId).WithSnapshots(withSnapshots).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreModelsAPI.ApiV1CoreGetTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CoreGetTag`: TagSchema
	fmt.Fprintf(os.Stdout, "Response from `CoreModelsAPI.ApiV1CoreGetTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CoreGetTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withSnapshots** | **bool** |  | [default to true]

### Return type

[**TagSchema**](TagSchema.md)

### Authorization

[HeaderTokenAuth](../README.md#HeaderTokenAuth), [QueryParamTokenAuth](../README.md#QueryParamTokenAuth), [UsernameAndPasswordAuth](../README.md#UsernameAndPasswordAuth), [BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1CoreGetTags

> PagedTagSchema ApiV1CoreGetTags(ctx).Limit(limit).Offset(offset).Page(page).Execute()

Get Tags

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
	limit := int32(56) // int32 |  (optional) (default to 200)
	offset := int32(56) // int32 |  (optional) (default to 0)
	page := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreModelsAPI.ApiV1CoreGetTags(context.Background()).Limit(limit).Offset(offset).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreModelsAPI.ApiV1CoreGetTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CoreGetTags`: PagedTagSchema
	fmt.Fprintf(os.Stdout, "Response from `CoreModelsAPI.ApiV1CoreGetTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CoreGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 200]
 **offset** | **int32** |  | [default to 0]
 **page** | **int32** |  | [default to 0]

### Return type

[**PagedTagSchema**](PagedTagSchema.md)

### Authorization

[HeaderTokenAuth](../README.md#HeaderTokenAuth), [QueryParamTokenAuth](../README.md#QueryParamTokenAuth), [UsernameAndPasswordAuth](../README.md#UsernameAndPasswordAuth), [BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

