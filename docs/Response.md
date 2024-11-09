# Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TYPE** | Pointer to **string** |  | [optional] [default to "core.models.Tag"]
**Id** | **string** |  | 
**Abid** | **string** |  | 
**CreatedById** | **string** |  | 
**CreatedByUsername** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**ModifiedAt** | **time.Time** |  | 
**BookmarkedAt** | **time.Time** |  | 
**DownloadedAt** | **time.Time** |  | 
**Url** | **string** |  | 
**Tags** | **[]string** |  | 
**Title** | **string** |  | 
**Timestamp** | **string** |  | 
**ArchivePath** | **string** |  | 
**NumArchiveresults** | **int32** |  | 
**Archiveresults** | [**[]MinimalArchiveResultSchema**](MinimalArchiveResultSchema.md) |  | 
**Extractor** | **string** |  | 
**CmdVersion** | **string** |  | 
**Cmd** | **[]string** |  | 
**Pwd** | **string** |  | 
**Status** | **string** |  | 
**Output** | **string** |  | 
**StartTs** | **time.Time** |  | 
**EndTs** | **time.Time** |  | 
**SnapshotId** | **string** |  | 
**SnapshotAbid** | **string** |  | 
**SnapshotTimestamp** | **string** |  | 
**SnapshotUrl** | **string** |  | 
**SnapshotTags** | **[]string** |  | 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**NumSnapshots** | **int32** |  | 
**Snapshots** | [**[]SnapshotSchema**](SnapshotSchema.md) |  | 

## Methods

### NewResponse

`func NewResponse(id string, abid string, createdById string, createdByUsername string, createdAt time.Time, modifiedAt time.Time, bookmarkedAt time.Time, downloadedAt time.Time, url string, tags []string, title string, timestamp string, archivePath string, numArchiveresults int32, archiveresults []MinimalArchiveResultSchema, extractor string, cmdVersion string, cmd []string, pwd string, status string, output string, startTs time.Time, endTs time.Time, snapshotId string, snapshotAbid string, snapshotTimestamp string, snapshotUrl string, snapshotTags []string, name string, slug string, numSnapshots int32, snapshots []SnapshotSchema, ) *Response`

NewResponse instantiates a new Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWithDefaults

`func NewResponseWithDefaults() *Response`

NewResponseWithDefaults instantiates a new Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTYPE

`func (o *Response) GetTYPE() string`

GetTYPE returns the TYPE field if non-nil, zero value otherwise.

### GetTYPEOk

`func (o *Response) GetTYPEOk() (*string, bool)`

GetTYPEOk returns a tuple with the TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTYPE

`func (o *Response) SetTYPE(v string)`

SetTYPE sets TYPE field to given value.

### HasTYPE

`func (o *Response) HasTYPE() bool`

HasTYPE returns a boolean if a field has been set.

### GetId

`func (o *Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Response) SetId(v string)`

SetId sets Id field to given value.


### GetAbid

`func (o *Response) GetAbid() string`

GetAbid returns the Abid field if non-nil, zero value otherwise.

### GetAbidOk

`func (o *Response) GetAbidOk() (*string, bool)`

GetAbidOk returns a tuple with the Abid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbid

`func (o *Response) SetAbid(v string)`

SetAbid sets Abid field to given value.


### GetCreatedById

`func (o *Response) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Response) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Response) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetCreatedByUsername

`func (o *Response) GetCreatedByUsername() string`

GetCreatedByUsername returns the CreatedByUsername field if non-nil, zero value otherwise.

### GetCreatedByUsernameOk

`func (o *Response) GetCreatedByUsernameOk() (*string, bool)`

GetCreatedByUsernameOk returns a tuple with the CreatedByUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUsername

`func (o *Response) SetCreatedByUsername(v string)`

SetCreatedByUsername sets CreatedByUsername field to given value.


### GetCreatedAt

`func (o *Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *Response) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Response) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Response) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetBookmarkedAt

`func (o *Response) GetBookmarkedAt() time.Time`

GetBookmarkedAt returns the BookmarkedAt field if non-nil, zero value otherwise.

### GetBookmarkedAtOk

`func (o *Response) GetBookmarkedAtOk() (*time.Time, bool)`

GetBookmarkedAtOk returns a tuple with the BookmarkedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarkedAt

`func (o *Response) SetBookmarkedAt(v time.Time)`

SetBookmarkedAt sets BookmarkedAt field to given value.


### GetDownloadedAt

`func (o *Response) GetDownloadedAt() time.Time`

GetDownloadedAt returns the DownloadedAt field if non-nil, zero value otherwise.

### GetDownloadedAtOk

`func (o *Response) GetDownloadedAtOk() (*time.Time, bool)`

GetDownloadedAtOk returns a tuple with the DownloadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadedAt

`func (o *Response) SetDownloadedAt(v time.Time)`

SetDownloadedAt sets DownloadedAt field to given value.


### GetUrl

`func (o *Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Response) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTags

`func (o *Response) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Response) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Response) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetTitle

`func (o *Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Response) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTimestamp

`func (o *Response) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Response) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Response) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetArchivePath

`func (o *Response) GetArchivePath() string`

GetArchivePath returns the ArchivePath field if non-nil, zero value otherwise.

### GetArchivePathOk

`func (o *Response) GetArchivePathOk() (*string, bool)`

GetArchivePathOk returns a tuple with the ArchivePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivePath

`func (o *Response) SetArchivePath(v string)`

SetArchivePath sets ArchivePath field to given value.


### GetNumArchiveresults

`func (o *Response) GetNumArchiveresults() int32`

GetNumArchiveresults returns the NumArchiveresults field if non-nil, zero value otherwise.

### GetNumArchiveresultsOk

`func (o *Response) GetNumArchiveresultsOk() (*int32, bool)`

GetNumArchiveresultsOk returns a tuple with the NumArchiveresults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumArchiveresults

`func (o *Response) SetNumArchiveresults(v int32)`

SetNumArchiveresults sets NumArchiveresults field to given value.


### GetArchiveresults

`func (o *Response) GetArchiveresults() []MinimalArchiveResultSchema`

GetArchiveresults returns the Archiveresults field if non-nil, zero value otherwise.

### GetArchiveresultsOk

`func (o *Response) GetArchiveresultsOk() (*[]MinimalArchiveResultSchema, bool)`

GetArchiveresultsOk returns a tuple with the Archiveresults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveresults

`func (o *Response) SetArchiveresults(v []MinimalArchiveResultSchema)`

SetArchiveresults sets Archiveresults field to given value.


### GetExtractor

`func (o *Response) GetExtractor() string`

GetExtractor returns the Extractor field if non-nil, zero value otherwise.

### GetExtractorOk

`func (o *Response) GetExtractorOk() (*string, bool)`

GetExtractorOk returns a tuple with the Extractor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractor

`func (o *Response) SetExtractor(v string)`

SetExtractor sets Extractor field to given value.


### GetCmdVersion

`func (o *Response) GetCmdVersion() string`

GetCmdVersion returns the CmdVersion field if non-nil, zero value otherwise.

### GetCmdVersionOk

`func (o *Response) GetCmdVersionOk() (*string, bool)`

GetCmdVersionOk returns a tuple with the CmdVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdVersion

`func (o *Response) SetCmdVersion(v string)`

SetCmdVersion sets CmdVersion field to given value.


### GetCmd

`func (o *Response) GetCmd() []string`

GetCmd returns the Cmd field if non-nil, zero value otherwise.

### GetCmdOk

`func (o *Response) GetCmdOk() (*[]string, bool)`

GetCmdOk returns a tuple with the Cmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmd

`func (o *Response) SetCmd(v []string)`

SetCmd sets Cmd field to given value.


### GetPwd

`func (o *Response) GetPwd() string`

GetPwd returns the Pwd field if non-nil, zero value otherwise.

### GetPwdOk

`func (o *Response) GetPwdOk() (*string, bool)`

GetPwdOk returns a tuple with the Pwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwd

`func (o *Response) SetPwd(v string)`

SetPwd sets Pwd field to given value.


### GetStatus

`func (o *Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetOutput

`func (o *Response) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *Response) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *Response) SetOutput(v string)`

SetOutput sets Output field to given value.


### GetStartTs

`func (o *Response) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *Response) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *Response) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.


### GetEndTs

`func (o *Response) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *Response) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *Response) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.


### GetSnapshotId

`func (o *Response) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *Response) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *Response) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### GetSnapshotAbid

`func (o *Response) GetSnapshotAbid() string`

GetSnapshotAbid returns the SnapshotAbid field if non-nil, zero value otherwise.

### GetSnapshotAbidOk

`func (o *Response) GetSnapshotAbidOk() (*string, bool)`

GetSnapshotAbidOk returns a tuple with the SnapshotAbid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotAbid

`func (o *Response) SetSnapshotAbid(v string)`

SetSnapshotAbid sets SnapshotAbid field to given value.


### GetSnapshotTimestamp

`func (o *Response) GetSnapshotTimestamp() string`

GetSnapshotTimestamp returns the SnapshotTimestamp field if non-nil, zero value otherwise.

### GetSnapshotTimestampOk

`func (o *Response) GetSnapshotTimestampOk() (*string, bool)`

GetSnapshotTimestampOk returns a tuple with the SnapshotTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTimestamp

`func (o *Response) SetSnapshotTimestamp(v string)`

SetSnapshotTimestamp sets SnapshotTimestamp field to given value.


### GetSnapshotUrl

`func (o *Response) GetSnapshotUrl() string`

GetSnapshotUrl returns the SnapshotUrl field if non-nil, zero value otherwise.

### GetSnapshotUrlOk

`func (o *Response) GetSnapshotUrlOk() (*string, bool)`

GetSnapshotUrlOk returns a tuple with the SnapshotUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotUrl

`func (o *Response) SetSnapshotUrl(v string)`

SetSnapshotUrl sets SnapshotUrl field to given value.


### GetSnapshotTags

`func (o *Response) GetSnapshotTags() []string`

GetSnapshotTags returns the SnapshotTags field if non-nil, zero value otherwise.

### GetSnapshotTagsOk

`func (o *Response) GetSnapshotTagsOk() (*[]string, bool)`

GetSnapshotTagsOk returns a tuple with the SnapshotTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTags

`func (o *Response) SetSnapshotTags(v []string)`

SetSnapshotTags sets SnapshotTags field to given value.


### GetName

`func (o *Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Response) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *Response) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Response) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Response) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetNumSnapshots

`func (o *Response) GetNumSnapshots() int32`

GetNumSnapshots returns the NumSnapshots field if non-nil, zero value otherwise.

### GetNumSnapshotsOk

`func (o *Response) GetNumSnapshotsOk() (*int32, bool)`

GetNumSnapshotsOk returns a tuple with the NumSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSnapshots

`func (o *Response) SetNumSnapshots(v int32)`

SetNumSnapshots sets NumSnapshots field to given value.


### GetSnapshots

`func (o *Response) GetSnapshots() []SnapshotSchema`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *Response) GetSnapshotsOk() (*[]SnapshotSchema, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *Response) SetSnapshots(v []SnapshotSchema)`

SetSnapshots sets Snapshots field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


