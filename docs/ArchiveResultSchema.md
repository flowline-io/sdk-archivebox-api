# ArchiveResultSchema

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
**SnapshotId** | **string** |  | 
**SnapshotAbid** | **string** |  | 
**SnapshotTimestamp** | **string** |  | 
**SnapshotUrl** | **string** |  | 
**SnapshotTags** | **[]string** |  | 

## Methods

### NewArchiveResultSchema

`func NewArchiveResultSchema(id string, abid string, modifiedAt time.Time, createdAt time.Time, createdById string, createdByUsername string, extractor string, cmdVersion NullableString, cmd []string, pwd string, status string, output string, startTs NullableTime, endTs NullableTime, snapshotId string, snapshotAbid string, snapshotTimestamp string, snapshotUrl string, snapshotTags []string, ) *ArchiveResultSchema`

NewArchiveResultSchema instantiates a new ArchiveResultSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveResultSchemaWithDefaults

`func NewArchiveResultSchemaWithDefaults() *ArchiveResultSchema`

NewArchiveResultSchemaWithDefaults instantiates a new ArchiveResultSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTYPE

`func (o *ArchiveResultSchema) GetTYPE() string`

GetTYPE returns the TYPE field if non-nil, zero value otherwise.

### GetTYPEOk

`func (o *ArchiveResultSchema) GetTYPEOk() (*string, bool)`

GetTYPEOk returns a tuple with the TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTYPE

`func (o *ArchiveResultSchema) SetTYPE(v string)`

SetTYPE sets TYPE field to given value.

### HasTYPE

`func (o *ArchiveResultSchema) HasTYPE() bool`

HasTYPE returns a boolean if a field has been set.

### GetId

`func (o *ArchiveResultSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArchiveResultSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArchiveResultSchema) SetId(v string)`

SetId sets Id field to given value.


### GetAbid

`func (o *ArchiveResultSchema) GetAbid() string`

GetAbid returns the Abid field if non-nil, zero value otherwise.

### GetAbidOk

`func (o *ArchiveResultSchema) GetAbidOk() (*string, bool)`

GetAbidOk returns a tuple with the Abid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbid

`func (o *ArchiveResultSchema) SetAbid(v string)`

SetAbid sets Abid field to given value.


### GetModifiedAt

`func (o *ArchiveResultSchema) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ArchiveResultSchema) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ArchiveResultSchema) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetCreatedAt

`func (o *ArchiveResultSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArchiveResultSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArchiveResultSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedById

`func (o *ArchiveResultSchema) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *ArchiveResultSchema) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *ArchiveResultSchema) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetCreatedByUsername

`func (o *ArchiveResultSchema) GetCreatedByUsername() string`

GetCreatedByUsername returns the CreatedByUsername field if non-nil, zero value otherwise.

### GetCreatedByUsernameOk

`func (o *ArchiveResultSchema) GetCreatedByUsernameOk() (*string, bool)`

GetCreatedByUsernameOk returns a tuple with the CreatedByUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUsername

`func (o *ArchiveResultSchema) SetCreatedByUsername(v string)`

SetCreatedByUsername sets CreatedByUsername field to given value.


### GetExtractor

`func (o *ArchiveResultSchema) GetExtractor() string`

GetExtractor returns the Extractor field if non-nil, zero value otherwise.

### GetExtractorOk

`func (o *ArchiveResultSchema) GetExtractorOk() (*string, bool)`

GetExtractorOk returns a tuple with the Extractor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractor

`func (o *ArchiveResultSchema) SetExtractor(v string)`

SetExtractor sets Extractor field to given value.


### GetCmdVersion

`func (o *ArchiveResultSchema) GetCmdVersion() string`

GetCmdVersion returns the CmdVersion field if non-nil, zero value otherwise.

### GetCmdVersionOk

`func (o *ArchiveResultSchema) GetCmdVersionOk() (*string, bool)`

GetCmdVersionOk returns a tuple with the CmdVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdVersion

`func (o *ArchiveResultSchema) SetCmdVersion(v string)`

SetCmdVersion sets CmdVersion field to given value.


### SetCmdVersionNil

`func (o *ArchiveResultSchema) SetCmdVersionNil(b bool)`

 SetCmdVersionNil sets the value for CmdVersion to be an explicit nil

### UnsetCmdVersion
`func (o *ArchiveResultSchema) UnsetCmdVersion()`

UnsetCmdVersion ensures that no value is present for CmdVersion, not even an explicit nil
### GetCmd

`func (o *ArchiveResultSchema) GetCmd() []string`

GetCmd returns the Cmd field if non-nil, zero value otherwise.

### GetCmdOk

`func (o *ArchiveResultSchema) GetCmdOk() (*[]string, bool)`

GetCmdOk returns a tuple with the Cmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmd

`func (o *ArchiveResultSchema) SetCmd(v []string)`

SetCmd sets Cmd field to given value.


### GetPwd

`func (o *ArchiveResultSchema) GetPwd() string`

GetPwd returns the Pwd field if non-nil, zero value otherwise.

### GetPwdOk

`func (o *ArchiveResultSchema) GetPwdOk() (*string, bool)`

GetPwdOk returns a tuple with the Pwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwd

`func (o *ArchiveResultSchema) SetPwd(v string)`

SetPwd sets Pwd field to given value.


### GetStatus

`func (o *ArchiveResultSchema) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArchiveResultSchema) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArchiveResultSchema) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetOutput

`func (o *ArchiveResultSchema) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ArchiveResultSchema) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ArchiveResultSchema) SetOutput(v string)`

SetOutput sets Output field to given value.


### GetStartTs

`func (o *ArchiveResultSchema) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *ArchiveResultSchema) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *ArchiveResultSchema) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.


### SetStartTsNil

`func (o *ArchiveResultSchema) SetStartTsNil(b bool)`

 SetStartTsNil sets the value for StartTs to be an explicit nil

### UnsetStartTs
`func (o *ArchiveResultSchema) UnsetStartTs()`

UnsetStartTs ensures that no value is present for StartTs, not even an explicit nil
### GetEndTs

`func (o *ArchiveResultSchema) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *ArchiveResultSchema) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *ArchiveResultSchema) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.


### SetEndTsNil

`func (o *ArchiveResultSchema) SetEndTsNil(b bool)`

 SetEndTsNil sets the value for EndTs to be an explicit nil

### UnsetEndTs
`func (o *ArchiveResultSchema) UnsetEndTs()`

UnsetEndTs ensures that no value is present for EndTs, not even an explicit nil
### GetSnapshotId

`func (o *ArchiveResultSchema) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ArchiveResultSchema) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ArchiveResultSchema) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### GetSnapshotAbid

`func (o *ArchiveResultSchema) GetSnapshotAbid() string`

GetSnapshotAbid returns the SnapshotAbid field if non-nil, zero value otherwise.

### GetSnapshotAbidOk

`func (o *ArchiveResultSchema) GetSnapshotAbidOk() (*string, bool)`

GetSnapshotAbidOk returns a tuple with the SnapshotAbid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotAbid

`func (o *ArchiveResultSchema) SetSnapshotAbid(v string)`

SetSnapshotAbid sets SnapshotAbid field to given value.


### GetSnapshotTimestamp

`func (o *ArchiveResultSchema) GetSnapshotTimestamp() string`

GetSnapshotTimestamp returns the SnapshotTimestamp field if non-nil, zero value otherwise.

### GetSnapshotTimestampOk

`func (o *ArchiveResultSchema) GetSnapshotTimestampOk() (*string, bool)`

GetSnapshotTimestampOk returns a tuple with the SnapshotTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTimestamp

`func (o *ArchiveResultSchema) SetSnapshotTimestamp(v string)`

SetSnapshotTimestamp sets SnapshotTimestamp field to given value.


### GetSnapshotUrl

`func (o *ArchiveResultSchema) GetSnapshotUrl() string`

GetSnapshotUrl returns the SnapshotUrl field if non-nil, zero value otherwise.

### GetSnapshotUrlOk

`func (o *ArchiveResultSchema) GetSnapshotUrlOk() (*string, bool)`

GetSnapshotUrlOk returns a tuple with the SnapshotUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotUrl

`func (o *ArchiveResultSchema) SetSnapshotUrl(v string)`

SetSnapshotUrl sets SnapshotUrl field to given value.


### GetSnapshotTags

`func (o *ArchiveResultSchema) GetSnapshotTags() []string`

GetSnapshotTags returns the SnapshotTags field if non-nil, zero value otherwise.

### GetSnapshotTagsOk

`func (o *ArchiveResultSchema) GetSnapshotTagsOk() (*[]string, bool)`

GetSnapshotTagsOk returns a tuple with the SnapshotTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTags

`func (o *ArchiveResultSchema) SetSnapshotTags(v []string)`

SetSnapshotTags sets SnapshotTags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


