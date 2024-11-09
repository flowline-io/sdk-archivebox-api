# AddCommandSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Urls** | **[]string** |  | 
**Tag** | Pointer to **string** |  | [optional] [default to ""]
**Depth** | Pointer to **int32** |  | [optional] [default to 0]
**Update** | Pointer to **bool** |  | [optional] [default to false]
**UpdateAll** | Pointer to **bool** |  | [optional] [default to false]
**IndexOnly** | Pointer to **bool** |  | [optional] [default to false]
**Overwrite** | Pointer to **bool** |  | [optional] [default to false]
**Init** | Pointer to **bool** |  | [optional] [default to false]
**Extractors** | Pointer to **string** |  | [optional] [default to ""]
**Parser** | Pointer to **string** |  | [optional] [default to "auto"]

## Methods

### NewAddCommandSchema

`func NewAddCommandSchema(urls []string, ) *AddCommandSchema`

NewAddCommandSchema instantiates a new AddCommandSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCommandSchemaWithDefaults

`func NewAddCommandSchemaWithDefaults() *AddCommandSchema`

NewAddCommandSchemaWithDefaults instantiates a new AddCommandSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrls

`func (o *AddCommandSchema) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *AddCommandSchema) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *AddCommandSchema) SetUrls(v []string)`

SetUrls sets Urls field to given value.


### GetTag

`func (o *AddCommandSchema) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AddCommandSchema) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AddCommandSchema) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AddCommandSchema) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetDepth

`func (o *AddCommandSchema) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *AddCommandSchema) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *AddCommandSchema) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *AddCommandSchema) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetUpdate

`func (o *AddCommandSchema) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *AddCommandSchema) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *AddCommandSchema) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *AddCommandSchema) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUpdateAll

`func (o *AddCommandSchema) GetUpdateAll() bool`

GetUpdateAll returns the UpdateAll field if non-nil, zero value otherwise.

### GetUpdateAllOk

`func (o *AddCommandSchema) GetUpdateAllOk() (*bool, bool)`

GetUpdateAllOk returns a tuple with the UpdateAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAll

`func (o *AddCommandSchema) SetUpdateAll(v bool)`

SetUpdateAll sets UpdateAll field to given value.

### HasUpdateAll

`func (o *AddCommandSchema) HasUpdateAll() bool`

HasUpdateAll returns a boolean if a field has been set.

### GetIndexOnly

`func (o *AddCommandSchema) GetIndexOnly() bool`

GetIndexOnly returns the IndexOnly field if non-nil, zero value otherwise.

### GetIndexOnlyOk

`func (o *AddCommandSchema) GetIndexOnlyOk() (*bool, bool)`

GetIndexOnlyOk returns a tuple with the IndexOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexOnly

`func (o *AddCommandSchema) SetIndexOnly(v bool)`

SetIndexOnly sets IndexOnly field to given value.

### HasIndexOnly

`func (o *AddCommandSchema) HasIndexOnly() bool`

HasIndexOnly returns a boolean if a field has been set.

### GetOverwrite

`func (o *AddCommandSchema) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *AddCommandSchema) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *AddCommandSchema) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *AddCommandSchema) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetInit

`func (o *AddCommandSchema) GetInit() bool`

GetInit returns the Init field if non-nil, zero value otherwise.

### GetInitOk

`func (o *AddCommandSchema) GetInitOk() (*bool, bool)`

GetInitOk returns a tuple with the Init field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInit

`func (o *AddCommandSchema) SetInit(v bool)`

SetInit sets Init field to given value.

### HasInit

`func (o *AddCommandSchema) HasInit() bool`

HasInit returns a boolean if a field has been set.

### GetExtractors

`func (o *AddCommandSchema) GetExtractors() string`

GetExtractors returns the Extractors field if non-nil, zero value otherwise.

### GetExtractorsOk

`func (o *AddCommandSchema) GetExtractorsOk() (*string, bool)`

GetExtractorsOk returns a tuple with the Extractors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractors

`func (o *AddCommandSchema) SetExtractors(v string)`

SetExtractors sets Extractors field to given value.

### HasExtractors

`func (o *AddCommandSchema) HasExtractors() bool`

HasExtractors returns a boolean if a field has been set.

### GetParser

`func (o *AddCommandSchema) GetParser() string`

GetParser returns the Parser field if non-nil, zero value otherwise.

### GetParserOk

`func (o *AddCommandSchema) GetParserOk() (*string, bool)`

GetParserOk returns a tuple with the Parser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParser

`func (o *AddCommandSchema) SetParser(v string)`

SetParser sets Parser field to given value.

### HasParser

`func (o *AddCommandSchema) HasParser() bool`

HasParser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


