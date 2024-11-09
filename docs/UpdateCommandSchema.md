# UpdateCommandSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resume** | Pointer to **NullableFloat32** |  | [optional] 
**OnlyNew** | Pointer to **bool** |  | [optional] [default to true]
**IndexOnly** | Pointer to **bool** |  | [optional] [default to false]
**Overwrite** | Pointer to **bool** |  | [optional] [default to false]
**After** | Pointer to **NullableFloat32** |  | [optional] 
**Before** | Pointer to **NullableFloat32** |  | [optional] 
**Status** | Pointer to [**NullableStatusChoices**](StatusChoices.md) |  | [optional] 
**FilterType** | Pointer to **NullableString** |  | [optional] 
**FilterPatterns** | Pointer to **[]string** |  | [optional] 
**Extractors** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateCommandSchema

`func NewUpdateCommandSchema() *UpdateCommandSchema`

NewUpdateCommandSchema instantiates a new UpdateCommandSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCommandSchemaWithDefaults

`func NewUpdateCommandSchemaWithDefaults() *UpdateCommandSchema`

NewUpdateCommandSchemaWithDefaults instantiates a new UpdateCommandSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResume

`func (o *UpdateCommandSchema) GetResume() float32`

GetResume returns the Resume field if non-nil, zero value otherwise.

### GetResumeOk

`func (o *UpdateCommandSchema) GetResumeOk() (*float32, bool)`

GetResumeOk returns a tuple with the Resume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResume

`func (o *UpdateCommandSchema) SetResume(v float32)`

SetResume sets Resume field to given value.

### HasResume

`func (o *UpdateCommandSchema) HasResume() bool`

HasResume returns a boolean if a field has been set.

### SetResumeNil

`func (o *UpdateCommandSchema) SetResumeNil(b bool)`

 SetResumeNil sets the value for Resume to be an explicit nil

### UnsetResume
`func (o *UpdateCommandSchema) UnsetResume()`

UnsetResume ensures that no value is present for Resume, not even an explicit nil
### GetOnlyNew

`func (o *UpdateCommandSchema) GetOnlyNew() bool`

GetOnlyNew returns the OnlyNew field if non-nil, zero value otherwise.

### GetOnlyNewOk

`func (o *UpdateCommandSchema) GetOnlyNewOk() (*bool, bool)`

GetOnlyNewOk returns a tuple with the OnlyNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyNew

`func (o *UpdateCommandSchema) SetOnlyNew(v bool)`

SetOnlyNew sets OnlyNew field to given value.

### HasOnlyNew

`func (o *UpdateCommandSchema) HasOnlyNew() bool`

HasOnlyNew returns a boolean if a field has been set.

### GetIndexOnly

`func (o *UpdateCommandSchema) GetIndexOnly() bool`

GetIndexOnly returns the IndexOnly field if non-nil, zero value otherwise.

### GetIndexOnlyOk

`func (o *UpdateCommandSchema) GetIndexOnlyOk() (*bool, bool)`

GetIndexOnlyOk returns a tuple with the IndexOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexOnly

`func (o *UpdateCommandSchema) SetIndexOnly(v bool)`

SetIndexOnly sets IndexOnly field to given value.

### HasIndexOnly

`func (o *UpdateCommandSchema) HasIndexOnly() bool`

HasIndexOnly returns a boolean if a field has been set.

### GetOverwrite

`func (o *UpdateCommandSchema) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *UpdateCommandSchema) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *UpdateCommandSchema) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *UpdateCommandSchema) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetAfter

`func (o *UpdateCommandSchema) GetAfter() float32`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *UpdateCommandSchema) GetAfterOk() (*float32, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *UpdateCommandSchema) SetAfter(v float32)`

SetAfter sets After field to given value.

### HasAfter

`func (o *UpdateCommandSchema) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### SetAfterNil

`func (o *UpdateCommandSchema) SetAfterNil(b bool)`

 SetAfterNil sets the value for After to be an explicit nil

### UnsetAfter
`func (o *UpdateCommandSchema) UnsetAfter()`

UnsetAfter ensures that no value is present for After, not even an explicit nil
### GetBefore

`func (o *UpdateCommandSchema) GetBefore() float32`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *UpdateCommandSchema) GetBeforeOk() (*float32, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *UpdateCommandSchema) SetBefore(v float32)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *UpdateCommandSchema) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### SetBeforeNil

`func (o *UpdateCommandSchema) SetBeforeNil(b bool)`

 SetBeforeNil sets the value for Before to be an explicit nil

### UnsetBefore
`func (o *UpdateCommandSchema) UnsetBefore()`

UnsetBefore ensures that no value is present for Before, not even an explicit nil
### GetStatus

`func (o *UpdateCommandSchema) GetStatus() StatusChoices`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateCommandSchema) GetStatusOk() (*StatusChoices, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateCommandSchema) SetStatus(v StatusChoices)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateCommandSchema) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *UpdateCommandSchema) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *UpdateCommandSchema) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetFilterType

`func (o *UpdateCommandSchema) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *UpdateCommandSchema) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *UpdateCommandSchema) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *UpdateCommandSchema) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### SetFilterTypeNil

`func (o *UpdateCommandSchema) SetFilterTypeNil(b bool)`

 SetFilterTypeNil sets the value for FilterType to be an explicit nil

### UnsetFilterType
`func (o *UpdateCommandSchema) UnsetFilterType()`

UnsetFilterType ensures that no value is present for FilterType, not even an explicit nil
### GetFilterPatterns

`func (o *UpdateCommandSchema) GetFilterPatterns() []string`

GetFilterPatterns returns the FilterPatterns field if non-nil, zero value otherwise.

### GetFilterPatternsOk

`func (o *UpdateCommandSchema) GetFilterPatternsOk() (*[]string, bool)`

GetFilterPatternsOk returns a tuple with the FilterPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPatterns

`func (o *UpdateCommandSchema) SetFilterPatterns(v []string)`

SetFilterPatterns sets FilterPatterns field to given value.

### HasFilterPatterns

`func (o *UpdateCommandSchema) HasFilterPatterns() bool`

HasFilterPatterns returns a boolean if a field has been set.

### SetFilterPatternsNil

`func (o *UpdateCommandSchema) SetFilterPatternsNil(b bool)`

 SetFilterPatternsNil sets the value for FilterPatterns to be an explicit nil

### UnsetFilterPatterns
`func (o *UpdateCommandSchema) UnsetFilterPatterns()`

UnsetFilterPatterns ensures that no value is present for FilterPatterns, not even an explicit nil
### GetExtractors

`func (o *UpdateCommandSchema) GetExtractors() string`

GetExtractors returns the Extractors field if non-nil, zero value otherwise.

### GetExtractorsOk

`func (o *UpdateCommandSchema) GetExtractorsOk() (*string, bool)`

GetExtractorsOk returns a tuple with the Extractors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractors

`func (o *UpdateCommandSchema) SetExtractors(v string)`

SetExtractors sets Extractors field to given value.

### HasExtractors

`func (o *UpdateCommandSchema) HasExtractors() bool`

HasExtractors returns a boolean if a field has been set.

### SetExtractorsNil

`func (o *UpdateCommandSchema) SetExtractorsNil(b bool)`

 SetExtractorsNil sets the value for Extractors to be an explicit nil

### UnsetExtractors
`func (o *UpdateCommandSchema) UnsetExtractors()`

UnsetExtractors ensures that no value is present for Extractors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


