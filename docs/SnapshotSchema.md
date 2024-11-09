# SnapshotSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TYPE** | Pointer to **string** |  | [optional] [default to "core.models.Snapshot"]
**Id** | **string** |  | 
**Abid** | **string** |  | 
**CreatedById** | **string** |  | 
**CreatedByUsername** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**ModifiedAt** | **time.Time** |  | 
**BookmarkedAt** | **time.Time** |  | 
**DownloadedAt** | **NullableTime** |  | 
**Url** | **string** |  | 
**Tags** | **[]string** |  | 
**Title** | **NullableString** |  | 
**Timestamp** | **string** |  | 
**ArchivePath** | **string** |  | 
**NumArchiveresults** | **int32** |  | 
**Archiveresults** | [**[]MinimalArchiveResultSchema**](MinimalArchiveResultSchema.md) |  | 

## Methods

### NewSnapshotSchema

`func NewSnapshotSchema(id string, abid string, createdById string, createdByUsername string, createdAt time.Time, modifiedAt time.Time, bookmarkedAt time.Time, downloadedAt NullableTime, url string, tags []string, title NullableString, timestamp string, archivePath string, numArchiveresults int32, archiveresults []MinimalArchiveResultSchema, ) *SnapshotSchema`

NewSnapshotSchema instantiates a new SnapshotSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotSchemaWithDefaults

`func NewSnapshotSchemaWithDefaults() *SnapshotSchema`

NewSnapshotSchemaWithDefaults instantiates a new SnapshotSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTYPE

`func (o *SnapshotSchema) GetTYPE() string`

GetTYPE returns the TYPE field if non-nil, zero value otherwise.

### GetTYPEOk

`func (o *SnapshotSchema) GetTYPEOk() (*string, bool)`

GetTYPEOk returns a tuple with the TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTYPE

`func (o *SnapshotSchema) SetTYPE(v string)`

SetTYPE sets TYPE field to given value.

### HasTYPE

`func (o *SnapshotSchema) HasTYPE() bool`

HasTYPE returns a boolean if a field has been set.

### GetId

`func (o *SnapshotSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotSchema) SetId(v string)`

SetId sets Id field to given value.


### GetAbid

`func (o *SnapshotSchema) GetAbid() string`

GetAbid returns the Abid field if non-nil, zero value otherwise.

### GetAbidOk

`func (o *SnapshotSchema) GetAbidOk() (*string, bool)`

GetAbidOk returns a tuple with the Abid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbid

`func (o *SnapshotSchema) SetAbid(v string)`

SetAbid sets Abid field to given value.


### GetCreatedById

`func (o *SnapshotSchema) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *SnapshotSchema) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *SnapshotSchema) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetCreatedByUsername

`func (o *SnapshotSchema) GetCreatedByUsername() string`

GetCreatedByUsername returns the CreatedByUsername field if non-nil, zero value otherwise.

### GetCreatedByUsernameOk

`func (o *SnapshotSchema) GetCreatedByUsernameOk() (*string, bool)`

GetCreatedByUsernameOk returns a tuple with the CreatedByUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUsername

`func (o *SnapshotSchema) SetCreatedByUsername(v string)`

SetCreatedByUsername sets CreatedByUsername field to given value.


### GetCreatedAt

`func (o *SnapshotSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SnapshotSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SnapshotSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *SnapshotSchema) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SnapshotSchema) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SnapshotSchema) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetBookmarkedAt

`func (o *SnapshotSchema) GetBookmarkedAt() time.Time`

GetBookmarkedAt returns the BookmarkedAt field if non-nil, zero value otherwise.

### GetBookmarkedAtOk

`func (o *SnapshotSchema) GetBookmarkedAtOk() (*time.Time, bool)`

GetBookmarkedAtOk returns a tuple with the BookmarkedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarkedAt

`func (o *SnapshotSchema) SetBookmarkedAt(v time.Time)`

SetBookmarkedAt sets BookmarkedAt field to given value.


### GetDownloadedAt

`func (o *SnapshotSchema) GetDownloadedAt() time.Time`

GetDownloadedAt returns the DownloadedAt field if non-nil, zero value otherwise.

### GetDownloadedAtOk

`func (o *SnapshotSchema) GetDownloadedAtOk() (*time.Time, bool)`

GetDownloadedAtOk returns a tuple with the DownloadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadedAt

`func (o *SnapshotSchema) SetDownloadedAt(v time.Time)`

SetDownloadedAt sets DownloadedAt field to given value.


### SetDownloadedAtNil

`func (o *SnapshotSchema) SetDownloadedAtNil(b bool)`

 SetDownloadedAtNil sets the value for DownloadedAt to be an explicit nil

### UnsetDownloadedAt
`func (o *SnapshotSchema) UnsetDownloadedAt()`

UnsetDownloadedAt ensures that no value is present for DownloadedAt, not even an explicit nil
### GetUrl

`func (o *SnapshotSchema) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SnapshotSchema) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SnapshotSchema) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTags

`func (o *SnapshotSchema) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SnapshotSchema) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SnapshotSchema) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetTitle

`func (o *SnapshotSchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SnapshotSchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SnapshotSchema) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *SnapshotSchema) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SnapshotSchema) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTimestamp

`func (o *SnapshotSchema) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SnapshotSchema) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SnapshotSchema) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetArchivePath

`func (o *SnapshotSchema) GetArchivePath() string`

GetArchivePath returns the ArchivePath field if non-nil, zero value otherwise.

### GetArchivePathOk

`func (o *SnapshotSchema) GetArchivePathOk() (*string, bool)`

GetArchivePathOk returns a tuple with the ArchivePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivePath

`func (o *SnapshotSchema) SetArchivePath(v string)`

SetArchivePath sets ArchivePath field to given value.


### GetNumArchiveresults

`func (o *SnapshotSchema) GetNumArchiveresults() int32`

GetNumArchiveresults returns the NumArchiveresults field if non-nil, zero value otherwise.

### GetNumArchiveresultsOk

`func (o *SnapshotSchema) GetNumArchiveresultsOk() (*int32, bool)`

GetNumArchiveresultsOk returns a tuple with the NumArchiveresults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumArchiveresults

`func (o *SnapshotSchema) SetNumArchiveresults(v int32)`

SetNumArchiveresults sets NumArchiveresults field to given value.


### GetArchiveresults

`func (o *SnapshotSchema) GetArchiveresults() []MinimalArchiveResultSchema`

GetArchiveresults returns the Archiveresults field if non-nil, zero value otherwise.

### GetArchiveresultsOk

`func (o *SnapshotSchema) GetArchiveresultsOk() (*[]MinimalArchiveResultSchema, bool)`

GetArchiveresultsOk returns a tuple with the Archiveresults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveresults

`func (o *SnapshotSchema) SetArchiveresults(v []MinimalArchiveResultSchema)`

SetArchiveresults sets Archiveresults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


