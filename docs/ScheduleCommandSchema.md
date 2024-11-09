# ScheduleCommandSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportPath** | Pointer to **NullableString** |  | [optional] 
**Add** | Pointer to **bool** |  | [optional] [default to false]
**Every** | Pointer to **NullableString** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] [default to ""]
**Depth** | Pointer to **int32** |  | [optional] [default to 0]
**Overwrite** | Pointer to **bool** |  | [optional] [default to false]
**Update** | Pointer to **bool** |  | [optional] [default to false]
**Clear** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewScheduleCommandSchema

`func NewScheduleCommandSchema() *ScheduleCommandSchema`

NewScheduleCommandSchema instantiates a new ScheduleCommandSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleCommandSchemaWithDefaults

`func NewScheduleCommandSchemaWithDefaults() *ScheduleCommandSchema`

NewScheduleCommandSchemaWithDefaults instantiates a new ScheduleCommandSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportPath

`func (o *ScheduleCommandSchema) GetImportPath() string`

GetImportPath returns the ImportPath field if non-nil, zero value otherwise.

### GetImportPathOk

`func (o *ScheduleCommandSchema) GetImportPathOk() (*string, bool)`

GetImportPathOk returns a tuple with the ImportPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPath

`func (o *ScheduleCommandSchema) SetImportPath(v string)`

SetImportPath sets ImportPath field to given value.

### HasImportPath

`func (o *ScheduleCommandSchema) HasImportPath() bool`

HasImportPath returns a boolean if a field has been set.

### SetImportPathNil

`func (o *ScheduleCommandSchema) SetImportPathNil(b bool)`

 SetImportPathNil sets the value for ImportPath to be an explicit nil

### UnsetImportPath
`func (o *ScheduleCommandSchema) UnsetImportPath()`

UnsetImportPath ensures that no value is present for ImportPath, not even an explicit nil
### GetAdd

`func (o *ScheduleCommandSchema) GetAdd() bool`

GetAdd returns the Add field if non-nil, zero value otherwise.

### GetAddOk

`func (o *ScheduleCommandSchema) GetAddOk() (*bool, bool)`

GetAddOk returns a tuple with the Add field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdd

`func (o *ScheduleCommandSchema) SetAdd(v bool)`

SetAdd sets Add field to given value.

### HasAdd

`func (o *ScheduleCommandSchema) HasAdd() bool`

HasAdd returns a boolean if a field has been set.

### GetEvery

`func (o *ScheduleCommandSchema) GetEvery() string`

GetEvery returns the Every field if non-nil, zero value otherwise.

### GetEveryOk

`func (o *ScheduleCommandSchema) GetEveryOk() (*string, bool)`

GetEveryOk returns a tuple with the Every field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvery

`func (o *ScheduleCommandSchema) SetEvery(v string)`

SetEvery sets Every field to given value.

### HasEvery

`func (o *ScheduleCommandSchema) HasEvery() bool`

HasEvery returns a boolean if a field has been set.

### SetEveryNil

`func (o *ScheduleCommandSchema) SetEveryNil(b bool)`

 SetEveryNil sets the value for Every to be an explicit nil

### UnsetEvery
`func (o *ScheduleCommandSchema) UnsetEvery()`

UnsetEvery ensures that no value is present for Every, not even an explicit nil
### GetTag

`func (o *ScheduleCommandSchema) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ScheduleCommandSchema) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ScheduleCommandSchema) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ScheduleCommandSchema) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetDepth

`func (o *ScheduleCommandSchema) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *ScheduleCommandSchema) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *ScheduleCommandSchema) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *ScheduleCommandSchema) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetOverwrite

`func (o *ScheduleCommandSchema) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *ScheduleCommandSchema) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *ScheduleCommandSchema) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *ScheduleCommandSchema) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetUpdate

`func (o *ScheduleCommandSchema) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ScheduleCommandSchema) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ScheduleCommandSchema) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ScheduleCommandSchema) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetClear

`func (o *ScheduleCommandSchema) GetClear() bool`

GetClear returns the Clear field if non-nil, zero value otherwise.

### GetClearOk

`func (o *ScheduleCommandSchema) GetClearOk() (*bool, bool)`

GetClearOk returns a tuple with the Clear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClear

`func (o *ScheduleCommandSchema) SetClear(v bool)`

SetClear sets Clear field to given value.

### HasClear

`func (o *ScheduleCommandSchema) HasClear() bool`

HasClear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


