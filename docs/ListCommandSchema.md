# ListCommandSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterPatterns** | Pointer to **[]string** |  | [optional] 
**FilterType** | Pointer to **string** |  | [optional] [default to "substring"]
**Status** | Pointer to [**NullableStatusChoices**](StatusChoices.md) |  | [optional] 
**After** | Pointer to **NullableFloat32** |  | [optional] 
**Before** | Pointer to **NullableFloat32** |  | [optional] 
**Sort** | Pointer to **string** |  | [optional] [default to "bookmarked_at"]
**AsJson** | Pointer to **bool** |  | [optional] [default to true]
**AsHtml** | Pointer to **bool** |  | [optional] [default to false]
**AsCsv** | Pointer to [**AsCsv**](AsCsv.md) |  | [optional] [default to timestamp,url]
**WithHeaders** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewListCommandSchema

`func NewListCommandSchema() *ListCommandSchema`

NewListCommandSchema instantiates a new ListCommandSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCommandSchemaWithDefaults

`func NewListCommandSchemaWithDefaults() *ListCommandSchema`

NewListCommandSchemaWithDefaults instantiates a new ListCommandSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterPatterns

`func (o *ListCommandSchema) GetFilterPatterns() []string`

GetFilterPatterns returns the FilterPatterns field if non-nil, zero value otherwise.

### GetFilterPatternsOk

`func (o *ListCommandSchema) GetFilterPatternsOk() (*[]string, bool)`

GetFilterPatternsOk returns a tuple with the FilterPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPatterns

`func (o *ListCommandSchema) SetFilterPatterns(v []string)`

SetFilterPatterns sets FilterPatterns field to given value.

### HasFilterPatterns

`func (o *ListCommandSchema) HasFilterPatterns() bool`

HasFilterPatterns returns a boolean if a field has been set.

### SetFilterPatternsNil

`func (o *ListCommandSchema) SetFilterPatternsNil(b bool)`

 SetFilterPatternsNil sets the value for FilterPatterns to be an explicit nil

### UnsetFilterPatterns
`func (o *ListCommandSchema) UnsetFilterPatterns()`

UnsetFilterPatterns ensures that no value is present for FilterPatterns, not even an explicit nil
### GetFilterType

`func (o *ListCommandSchema) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *ListCommandSchema) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *ListCommandSchema) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *ListCommandSchema) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetStatus

`func (o *ListCommandSchema) GetStatus() StatusChoices`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListCommandSchema) GetStatusOk() (*StatusChoices, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListCommandSchema) SetStatus(v StatusChoices)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListCommandSchema) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ListCommandSchema) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ListCommandSchema) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAfter

`func (o *ListCommandSchema) GetAfter() float32`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *ListCommandSchema) GetAfterOk() (*float32, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *ListCommandSchema) SetAfter(v float32)`

SetAfter sets After field to given value.

### HasAfter

`func (o *ListCommandSchema) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### SetAfterNil

`func (o *ListCommandSchema) SetAfterNil(b bool)`

 SetAfterNil sets the value for After to be an explicit nil

### UnsetAfter
`func (o *ListCommandSchema) UnsetAfter()`

UnsetAfter ensures that no value is present for After, not even an explicit nil
### GetBefore

`func (o *ListCommandSchema) GetBefore() float32`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *ListCommandSchema) GetBeforeOk() (*float32, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *ListCommandSchema) SetBefore(v float32)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *ListCommandSchema) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### SetBeforeNil

`func (o *ListCommandSchema) SetBeforeNil(b bool)`

 SetBeforeNil sets the value for Before to be an explicit nil

### UnsetBefore
`func (o *ListCommandSchema) UnsetBefore()`

UnsetBefore ensures that no value is present for Before, not even an explicit nil
### GetSort

`func (o *ListCommandSchema) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ListCommandSchema) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ListCommandSchema) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ListCommandSchema) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetAsJson

`func (o *ListCommandSchema) GetAsJson() bool`

GetAsJson returns the AsJson field if non-nil, zero value otherwise.

### GetAsJsonOk

`func (o *ListCommandSchema) GetAsJsonOk() (*bool, bool)`

GetAsJsonOk returns a tuple with the AsJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsJson

`func (o *ListCommandSchema) SetAsJson(v bool)`

SetAsJson sets AsJson field to given value.

### HasAsJson

`func (o *ListCommandSchema) HasAsJson() bool`

HasAsJson returns a boolean if a field has been set.

### GetAsHtml

`func (o *ListCommandSchema) GetAsHtml() bool`

GetAsHtml returns the AsHtml field if non-nil, zero value otherwise.

### GetAsHtmlOk

`func (o *ListCommandSchema) GetAsHtmlOk() (*bool, bool)`

GetAsHtmlOk returns a tuple with the AsHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsHtml

`func (o *ListCommandSchema) SetAsHtml(v bool)`

SetAsHtml sets AsHtml field to given value.

### HasAsHtml

`func (o *ListCommandSchema) HasAsHtml() bool`

HasAsHtml returns a boolean if a field has been set.

### GetAsCsv

`func (o *ListCommandSchema) GetAsCsv() AsCsv`

GetAsCsv returns the AsCsv field if non-nil, zero value otherwise.

### GetAsCsvOk

`func (o *ListCommandSchema) GetAsCsvOk() (*AsCsv, bool)`

GetAsCsvOk returns a tuple with the AsCsv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsCsv

`func (o *ListCommandSchema) SetAsCsv(v AsCsv)`

SetAsCsv sets AsCsv field to given value.

### HasAsCsv

`func (o *ListCommandSchema) HasAsCsv() bool`

HasAsCsv returns a boolean if a field has been set.

### GetWithHeaders

`func (o *ListCommandSchema) GetWithHeaders() bool`

GetWithHeaders returns the WithHeaders field if non-nil, zero value otherwise.

### GetWithHeadersOk

`func (o *ListCommandSchema) GetWithHeadersOk() (*bool, bool)`

GetWithHeadersOk returns a tuple with the WithHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithHeaders

`func (o *ListCommandSchema) SetWithHeaders(v bool)`

SetWithHeaders sets WithHeaders field to given value.

### HasWithHeaders

`func (o *ListCommandSchema) HasWithHeaders() bool`

HasWithHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


