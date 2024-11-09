# MinimalArchiveResultSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TYPE** | Pointer to **string** |  | [optional] [default to "core.models.ArchiveResult"]
**Id** | **string** |  | 
**Abid** | **string** |  | 
**ModifiedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**CreatedById** | **string** |  | 
**CreatedByUsername** | **string** |  | 
**Extractor** | **string** |  | 
**CmdVersion** | **NullableString** |  | 
**Cmd** | **[]string** |  | 
**Pwd** | **string** |  | 
**Status** | **string** |  | 
**Output** | **string** |  | 
**StartTs** | **NullableTime** |  | 
**EndTs** | **NullableTime** |  | 

## Methods

### NewMinimalArchiveResultSchema

`func NewMinimalArchiveResultSchema(id string, abid string, modifiedAt time.Time, createdAt time.Time, createdById string, createdByUsername string, extractor string, cmdVersion NullableString, cmd []string, pwd string, status string, output string, startTs NullableTime, endTs NullableTime, ) *MinimalArchiveResultSchema`

NewMinimalArchiveResultSchema instantiates a new MinimalArchiveResultSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalArchiveResultSchemaWithDefaults

`func NewMinimalArchiveResultSchemaWithDefaults() *MinimalArchiveResultSchema`

NewMinimalArchiveResultSchemaWithDefaults instantiates a new MinimalArchiveResultSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTYPE

`func (o *MinimalArchiveResultSchema) GetTYPE() string`

GetTYPE returns the TYPE field if non-nil, zero value otherwise.

### GetTYPEOk

`func (o *MinimalArchiveResultSchema) GetTYPEOk() (*string, bool)`

GetTYPEOk returns a tuple with the TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTYPE

`func (o *MinimalArchiveResultSchema) SetTYPE(v string)`

SetTYPE sets TYPE field to given value.

### HasTYPE

`func (o *MinimalArchiveResultSchema) HasTYPE() bool`

HasTYPE returns a boolean if a field has been set.

### GetId

`func (o *MinimalArchiveResultSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MinimalArchiveResultSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MinimalArchiveResultSchema) SetId(v string)`

SetId sets Id field to given value.


### GetAbid

`func (o *MinimalArchiveResultSchema) GetAbid() string`

GetAbid returns the Abid field if non-nil, zero value otherwise.

### GetAbidOk

`func (o *MinimalArchiveResultSchema) GetAbidOk() (*string, bool)`

GetAbidOk returns a tuple with the Abid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbid

`func (o *MinimalArchiveResultSchema) SetAbid(v string)`

SetAbid sets Abid field to given value.


### GetModifiedAt

`func (o *MinimalArchiveResultSchema) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *MinimalArchiveResultSchema) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *MinimalArchiveResultSchema) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetCreatedAt

`func (o *MinimalArchiveResultSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MinimalArchiveResultSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MinimalArchiveResultSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedById

`func (o *MinimalArchiveResultSchema) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *MinimalArchiveResultSchema) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *MinimalArchiveResultSchema) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetCreatedByUsername

`func (o *MinimalArchiveResultSchema) GetCreatedByUsername() string`

GetCreatedByUsername returns the CreatedByUsername field if non-nil, zero value otherwise.

### GetCreatedByUsernameOk

`func (o *MinimalArchiveResultSchema) GetCreatedByUsernameOk() (*string, bool)`

GetCreatedByUsernameOk returns a tuple with the CreatedByUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUsername

`func (o *MinimalArchiveResultSchema) SetCreatedByUsername(v string)`

SetCreatedByUsername sets CreatedByUsername field to given value.


### GetExtractor

`func (o *MinimalArchiveResultSchema) GetExtractor() string`

GetExtractor returns the Extractor field if non-nil, zero value otherwise.

### GetExtractorOk

`func (o *MinimalArchiveResultSchema) GetExtractorOk() (*string, bool)`

GetExtractorOk returns a tuple with the Extractor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractor

`func (o *MinimalArchiveResultSchema) SetExtractor(v string)`

SetExtractor sets Extractor field to given value.


### GetCmdVersion

`func (o *MinimalArchiveResultSchema) GetCmdVersion() string`

GetCmdVersion returns the CmdVersion field if non-nil, zero value otherwise.

### GetCmdVersionOk

`func (o *MinimalArchiveResultSchema) GetCmdVersionOk() (*string, bool)`

GetCmdVersionOk returns a tuple with the CmdVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdVersion

`func (o *MinimalArchiveResultSchema) SetCmdVersion(v string)`

SetCmdVersion sets CmdVersion field to given value.


### SetCmdVersionNil

`func (o *MinimalArchiveResultSchema) SetCmdVersionNil(b bool)`

 SetCmdVersionNil sets the value for CmdVersion to be an explicit nil

### UnsetCmdVersion
`func (o *MinimalArchiveResultSchema) UnsetCmdVersion()`

UnsetCmdVersion ensures that no value is present for CmdVersion, not even an explicit nil
### GetCmd

`func (o *MinimalArchiveResultSchema) GetCmd() []string`

GetCmd returns the Cmd field if non-nil, zero value otherwise.

### GetCmdOk

`func (o *MinimalArchiveResultSchema) GetCmdOk() (*[]string, bool)`

GetCmdOk returns a tuple with the Cmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmd

`func (o *MinimalArchiveResultSchema) SetCmd(v []string)`

SetCmd sets Cmd field to given value.


### GetPwd

`func (o *MinimalArchiveResultSchema) GetPwd() string`

GetPwd returns the Pwd field if non-nil, zero value otherwise.

### GetPwdOk

`func (o *MinimalArchiveResultSchema) GetPwdOk() (*string, bool)`

GetPwdOk returns a tuple with the Pwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwd

`func (o *MinimalArchiveResultSchema) SetPwd(v string)`

SetPwd sets Pwd field to given value.


### GetStatus

`func (o *MinimalArchiveResultSchema) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MinimalArchiveResultSchema) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MinimalArchiveResultSchema) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetOutput

`func (o *MinimalArchiveResultSchema) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *MinimalArchiveResultSchema) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *MinimalArchiveResultSchema) SetOutput(v string)`

SetOutput sets Output field to given value.


### GetStartTs

`func (o *MinimalArchiveResultSchema) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *MinimalArchiveResultSchema) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *MinimalArchiveResultSchema) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.


### SetStartTsNil

`func (o *MinimalArchiveResultSchema) SetStartTsNil(b bool)`

 SetStartTsNil sets the value for StartTs to be an explicit nil

### UnsetStartTs
`func (o *MinimalArchiveResultSchema) UnsetStartTs()`

UnsetStartTs ensures that no value is present for StartTs, not even an explicit nil
### GetEndTs

`func (o *MinimalArchiveResultSchema) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *MinimalArchiveResultSchema) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *MinimalArchiveResultSchema) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.


### SetEndTsNil

`func (o *MinimalArchiveResultSchema) SetEndTsNil(b bool)`

 SetEndTsNil sets the value for EndTs to be an explicit nil

### UnsetEndTs
`func (o *MinimalArchiveResultSchema) UnsetEndTs()`

UnsetEndTs ensures that no value is present for EndTs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


