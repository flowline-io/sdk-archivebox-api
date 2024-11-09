# ArchiveResultFilterSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Search** | Pointer to **NullableString** |  | [optional] 
**SnapshotId** | Pointer to **NullableString** |  | [optional] 
**SnapshotUrl** | Pointer to **NullableString** |  | [optional] 
**SnapshotTag** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Output** | Pointer to **NullableString** |  | [optional] 
**Extractor** | Pointer to **NullableString** |  | [optional] 
**Cmd** | Pointer to **NullableString** |  | [optional] 
**Pwd** | Pointer to **NullableString** |  | [optional] 
**CmdVersion** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAtGte** | Pointer to **NullableTime** |  | [optional] 
**CreatedAtLt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewArchiveResultFilterSchema

`func NewArchiveResultFilterSchema() *ArchiveResultFilterSchema`

NewArchiveResultFilterSchema instantiates a new ArchiveResultFilterSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveResultFilterSchemaWithDefaults

`func NewArchiveResultFilterSchemaWithDefaults() *ArchiveResultFilterSchema`

NewArchiveResultFilterSchemaWithDefaults instantiates a new ArchiveResultFilterSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArchiveResultFilterSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArchiveResultFilterSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArchiveResultFilterSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArchiveResultFilterSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ArchiveResultFilterSchema) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ArchiveResultFilterSchema) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSearch

`func (o *ArchiveResultFilterSchema) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *ArchiveResultFilterSchema) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *ArchiveResultFilterSchema) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *ArchiveResultFilterSchema) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *ArchiveResultFilterSchema) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *ArchiveResultFilterSchema) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetSnapshotId

`func (o *ArchiveResultFilterSchema) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ArchiveResultFilterSchema) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ArchiveResultFilterSchema) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ArchiveResultFilterSchema) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *ArchiveResultFilterSchema) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *ArchiveResultFilterSchema) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetSnapshotUrl

`func (o *ArchiveResultFilterSchema) GetSnapshotUrl() string`

GetSnapshotUrl returns the SnapshotUrl field if non-nil, zero value otherwise.

### GetSnapshotUrlOk

`func (o *ArchiveResultFilterSchema) GetSnapshotUrlOk() (*string, bool)`

GetSnapshotUrlOk returns a tuple with the SnapshotUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotUrl

`func (o *ArchiveResultFilterSchema) SetSnapshotUrl(v string)`

SetSnapshotUrl sets SnapshotUrl field to given value.

### HasSnapshotUrl

`func (o *ArchiveResultFilterSchema) HasSnapshotUrl() bool`

HasSnapshotUrl returns a boolean if a field has been set.

### SetSnapshotUrlNil

`func (o *ArchiveResultFilterSchema) SetSnapshotUrlNil(b bool)`

 SetSnapshotUrlNil sets the value for SnapshotUrl to be an explicit nil

### UnsetSnapshotUrl
`func (o *ArchiveResultFilterSchema) UnsetSnapshotUrl()`

UnsetSnapshotUrl ensures that no value is present for SnapshotUrl, not even an explicit nil
### GetSnapshotTag

`func (o *ArchiveResultFilterSchema) GetSnapshotTag() string`

GetSnapshotTag returns the SnapshotTag field if non-nil, zero value otherwise.

### GetSnapshotTagOk

`func (o *ArchiveResultFilterSchema) GetSnapshotTagOk() (*string, bool)`

GetSnapshotTagOk returns a tuple with the SnapshotTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTag

`func (o *ArchiveResultFilterSchema) SetSnapshotTag(v string)`

SetSnapshotTag sets SnapshotTag field to given value.

### HasSnapshotTag

`func (o *ArchiveResultFilterSchema) HasSnapshotTag() bool`

HasSnapshotTag returns a boolean if a field has been set.

### SetSnapshotTagNil

`func (o *ArchiveResultFilterSchema) SetSnapshotTagNil(b bool)`

 SetSnapshotTagNil sets the value for SnapshotTag to be an explicit nil

### UnsetSnapshotTag
`func (o *ArchiveResultFilterSchema) UnsetSnapshotTag()`

UnsetSnapshotTag ensures that no value is present for SnapshotTag, not even an explicit nil
### GetStatus

`func (o *ArchiveResultFilterSchema) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArchiveResultFilterSchema) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArchiveResultFilterSchema) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArchiveResultFilterSchema) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ArchiveResultFilterSchema) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ArchiveResultFilterSchema) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetOutput

`func (o *ArchiveResultFilterSchema) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ArchiveResultFilterSchema) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ArchiveResultFilterSchema) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ArchiveResultFilterSchema) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *ArchiveResultFilterSchema) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *ArchiveResultFilterSchema) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetExtractor

`func (o *ArchiveResultFilterSchema) GetExtractor() string`

GetExtractor returns the Extractor field if non-nil, zero value otherwise.

### GetExtractorOk

`func (o *ArchiveResultFilterSchema) GetExtractorOk() (*string, bool)`

GetExtractorOk returns a tuple with the Extractor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractor

`func (o *ArchiveResultFilterSchema) SetExtractor(v string)`

SetExtractor sets Extractor field to given value.

### HasExtractor

`func (o *ArchiveResultFilterSchema) HasExtractor() bool`

HasExtractor returns a boolean if a field has been set.

### SetExtractorNil

`func (o *ArchiveResultFilterSchema) SetExtractorNil(b bool)`

 SetExtractorNil sets the value for Extractor to be an explicit nil

### UnsetExtractor
`func (o *ArchiveResultFilterSchema) UnsetExtractor()`

UnsetExtractor ensures that no value is present for Extractor, not even an explicit nil
### GetCmd

`func (o *ArchiveResultFilterSchema) GetCmd() string`

GetCmd returns the Cmd field if non-nil, zero value otherwise.

### GetCmdOk

`func (o *ArchiveResultFilterSchema) GetCmdOk() (*string, bool)`

GetCmdOk returns a tuple with the Cmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmd

`func (o *ArchiveResultFilterSchema) SetCmd(v string)`

SetCmd sets Cmd field to given value.

### HasCmd

`func (o *ArchiveResultFilterSchema) HasCmd() bool`

HasCmd returns a boolean if a field has been set.

### SetCmdNil

`func (o *ArchiveResultFilterSchema) SetCmdNil(b bool)`

 SetCmdNil sets the value for Cmd to be an explicit nil

### UnsetCmd
`func (o *ArchiveResultFilterSchema) UnsetCmd()`

UnsetCmd ensures that no value is present for Cmd, not even an explicit nil
### GetPwd

`func (o *ArchiveResultFilterSchema) GetPwd() string`

GetPwd returns the Pwd field if non-nil, zero value otherwise.

### GetPwdOk

`func (o *ArchiveResultFilterSchema) GetPwdOk() (*string, bool)`

GetPwdOk returns a tuple with the Pwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwd

`func (o *ArchiveResultFilterSchema) SetPwd(v string)`

SetPwd sets Pwd field to given value.

### HasPwd

`func (o *ArchiveResultFilterSchema) HasPwd() bool`

HasPwd returns a boolean if a field has been set.

### SetPwdNil

`func (o *ArchiveResultFilterSchema) SetPwdNil(b bool)`

 SetPwdNil sets the value for Pwd to be an explicit nil

### UnsetPwd
`func (o *ArchiveResultFilterSchema) UnsetPwd()`

UnsetPwd ensures that no value is present for Pwd, not even an explicit nil
### GetCmdVersion

`func (o *ArchiveResultFilterSchema) GetCmdVersion() string`

GetCmdVersion returns the CmdVersion field if non-nil, zero value otherwise.

### GetCmdVersionOk

`func (o *ArchiveResultFilterSchema) GetCmdVersionOk() (*string, bool)`

GetCmdVersionOk returns a tuple with the CmdVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdVersion

`func (o *ArchiveResultFilterSchema) SetCmdVersion(v string)`

SetCmdVersion sets CmdVersion field to given value.

### HasCmdVersion

`func (o *ArchiveResultFilterSchema) HasCmdVersion() bool`

HasCmdVersion returns a boolean if a field has been set.

### SetCmdVersionNil

`func (o *ArchiveResultFilterSchema) SetCmdVersionNil(b bool)`

 SetCmdVersionNil sets the value for CmdVersion to be an explicit nil

### UnsetCmdVersion
`func (o *ArchiveResultFilterSchema) UnsetCmdVersion()`

UnsetCmdVersion ensures that no value is present for CmdVersion, not even an explicit nil
### GetCreatedAt

`func (o *ArchiveResultFilterSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArchiveResultFilterSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArchiveResultFilterSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArchiveResultFilterSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ArchiveResultFilterSchema) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ArchiveResultFilterSchema) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedAtGte

`func (o *ArchiveResultFilterSchema) GetCreatedAtGte() time.Time`

GetCreatedAtGte returns the CreatedAtGte field if non-nil, zero value otherwise.

### GetCreatedAtGteOk

`func (o *ArchiveResultFilterSchema) GetCreatedAtGteOk() (*time.Time, bool)`

GetCreatedAtGteOk returns a tuple with the CreatedAtGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtGte

`func (o *ArchiveResultFilterSchema) SetCreatedAtGte(v time.Time)`

SetCreatedAtGte sets CreatedAtGte field to given value.

### HasCreatedAtGte

`func (o *ArchiveResultFilterSchema) HasCreatedAtGte() bool`

HasCreatedAtGte returns a boolean if a field has been set.

### SetCreatedAtGteNil

`func (o *ArchiveResultFilterSchema) SetCreatedAtGteNil(b bool)`

 SetCreatedAtGteNil sets the value for CreatedAtGte to be an explicit nil

### UnsetCreatedAtGte
`func (o *ArchiveResultFilterSchema) UnsetCreatedAtGte()`

UnsetCreatedAtGte ensures that no value is present for CreatedAtGte, not even an explicit nil
### GetCreatedAtLt

`func (o *ArchiveResultFilterSchema) GetCreatedAtLt() time.Time`

GetCreatedAtLt returns the CreatedAtLt field if non-nil, zero value otherwise.

### GetCreatedAtLtOk

`func (o *ArchiveResultFilterSchema) GetCreatedAtLtOk() (*time.Time, bool)`

GetCreatedAtLtOk returns a tuple with the CreatedAtLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtLt

`func (o *ArchiveResultFilterSchema) SetCreatedAtLt(v time.Time)`

SetCreatedAtLt sets CreatedAtLt field to given value.

### HasCreatedAtLt

`func (o *ArchiveResultFilterSchema) HasCreatedAtLt() bool`

HasCreatedAtLt returns a boolean if a field has been set.

### SetCreatedAtLtNil

`func (o *ArchiveResultFilterSchema) SetCreatedAtLtNil(b bool)`

 SetCreatedAtLtNil sets the value for CreatedAtLt to be an explicit nil

### UnsetCreatedAtLt
`func (o *ArchiveResultFilterSchema) UnsetCreatedAtLt()`

UnsetCreatedAtLt ensures that no value is present for CreatedAtLt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


