# RemoveCommandSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delete** | Pointer to **bool** |  | [optional] [default to true]
**After** | Pointer to **NullableFloat32** |  | [optional] 
**Before** | Pointer to **NullableFloat32** |  | [optional] 
**FilterType** | Pointer to **string** |  | [optional] [default to "exact"]
**FilterPatterns** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRemoveCommandSchema

`func NewRemoveCommandSchema() *RemoveCommandSchema`

NewRemoveCommandSchema instantiates a new RemoveCommandSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveCommandSchemaWithDefaults

`func NewRemoveCommandSchemaWithDefaults() *RemoveCommandSchema`

NewRemoveCommandSchemaWithDefaults instantiates a new RemoveCommandSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelete

`func (o *RemoveCommandSchema) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *RemoveCommandSchema) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *RemoveCommandSchema) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *RemoveCommandSchema) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetAfter

`func (o *RemoveCommandSchema) GetAfter() float32`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *RemoveCommandSchema) GetAfterOk() (*float32, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *RemoveCommandSchema) SetAfter(v float32)`

SetAfter sets After field to given value.

### HasAfter

`func (o *RemoveCommandSchema) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### SetAfterNil

`func (o *RemoveCommandSchema) SetAfterNil(b bool)`

 SetAfterNil sets the value for After to be an explicit nil

### UnsetAfter
`func (o *RemoveCommandSchema) UnsetAfter()`

UnsetAfter ensures that no value is present for After, not even an explicit nil
### GetBefore

`func (o *RemoveCommandSchema) GetBefore() float32`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *RemoveCommandSchema) GetBeforeOk() (*float32, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *RemoveCommandSchema) SetBefore(v float32)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *RemoveCommandSchema) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### SetBeforeNil

`func (o *RemoveCommandSchema) SetBeforeNil(b bool)`

 SetBeforeNil sets the value for Before to be an explicit nil

### UnsetBefore
`func (o *RemoveCommandSchema) UnsetBefore()`

UnsetBefore ensures that no value is present for Before, not even an explicit nil
### GetFilterType

`func (o *RemoveCommandSchema) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *RemoveCommandSchema) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *RemoveCommandSchema) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *RemoveCommandSchema) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetFilterPatterns

`func (o *RemoveCommandSchema) GetFilterPatterns() []string`

GetFilterPatterns returns the FilterPatterns field if non-nil, zero value otherwise.

### GetFilterPatternsOk

`func (o *RemoveCommandSchema) GetFilterPatternsOk() (*[]string, bool)`

GetFilterPatternsOk returns a tuple with the FilterPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPatterns

`func (o *RemoveCommandSchema) SetFilterPatterns(v []string)`

SetFilterPatterns sets FilterPatterns field to given value.

### HasFilterPatterns

`func (o *RemoveCommandSchema) HasFilterPatterns() bool`

HasFilterPatterns returns a boolean if a field has been set.

### SetFilterPatternsNil

`func (o *RemoveCommandSchema) SetFilterPatternsNil(b bool)`

 SetFilterPatternsNil sets the value for FilterPatterns to be an explicit nil

### UnsetFilterPatterns
`func (o *RemoveCommandSchema) UnsetFilterPatterns()`

UnsetFilterPatterns ensures that no value is present for FilterPatterns, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


