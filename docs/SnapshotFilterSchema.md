# SnapshotFilterSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Abid** | Pointer to **NullableString** |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**CreatedByUsername** | Pointer to **string** |  | [optional] 
**CreatedAtGte** | Pointer to **time.Time** |  | [optional] 
**CreatedAtLt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedAtGte** | Pointer to **time.Time** |  | [optional] 
**ModifiedAtLt** | Pointer to **time.Time** |  | [optional] 
**Search** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Tag** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableString** |  | [optional] 
**BookmarkedAtGte** | Pointer to **NullableTime** |  | [optional] 
**BookmarkedAtLt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewSnapshotFilterSchema

`func NewSnapshotFilterSchema() *SnapshotFilterSchema`

NewSnapshotFilterSchema instantiates a new SnapshotFilterSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotFilterSchemaWithDefaults

`func NewSnapshotFilterSchemaWithDefaults() *SnapshotFilterSchema`

NewSnapshotFilterSchemaWithDefaults instantiates a new SnapshotFilterSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SnapshotFilterSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotFilterSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotFilterSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SnapshotFilterSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SnapshotFilterSchema) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SnapshotFilterSchema) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAbid

`func (o *SnapshotFilterSchema) GetAbid() string`

GetAbid returns the Abid field if non-nil, zero value otherwise.

### GetAbidOk

`func (o *SnapshotFilterSchema) GetAbidOk() (*string, bool)`

GetAbidOk returns a tuple with the Abid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbid

`func (o *SnapshotFilterSchema) SetAbid(v string)`

SetAbid sets Abid field to given value.

### HasAbid

`func (o *SnapshotFilterSchema) HasAbid() bool`

HasAbid returns a boolean if a field has been set.

### SetAbidNil

`func (o *SnapshotFilterSchema) SetAbidNil(b bool)`

 SetAbidNil sets the value for Abid to be an explicit nil

### UnsetAbid
`func (o *SnapshotFilterSchema) UnsetAbid()`

UnsetAbid ensures that no value is present for Abid, not even an explicit nil
### GetCreatedById

`func (o *SnapshotFilterSchema) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *SnapshotFilterSchema) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *SnapshotFilterSchema) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *SnapshotFilterSchema) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedByUsername

`func (o *SnapshotFilterSchema) GetCreatedByUsername() string`

GetCreatedByUsername returns the CreatedByUsername field if non-nil, zero value otherwise.

### GetCreatedByUsernameOk

`func (o *SnapshotFilterSchema) GetCreatedByUsernameOk() (*string, bool)`

GetCreatedByUsernameOk returns a tuple with the CreatedByUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUsername

`func (o *SnapshotFilterSchema) SetCreatedByUsername(v string)`

SetCreatedByUsername sets CreatedByUsername field to given value.

### HasCreatedByUsername

`func (o *SnapshotFilterSchema) HasCreatedByUsername() bool`

HasCreatedByUsername returns a boolean if a field has been set.

### GetCreatedAtGte

`func (o *SnapshotFilterSchema) GetCreatedAtGte() time.Time`

GetCreatedAtGte returns the CreatedAtGte field if non-nil, zero value otherwise.

### GetCreatedAtGteOk

`func (o *SnapshotFilterSchema) GetCreatedAtGteOk() (*time.Time, bool)`

GetCreatedAtGteOk returns a tuple with the CreatedAtGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtGte

`func (o *SnapshotFilterSchema) SetCreatedAtGte(v time.Time)`

SetCreatedAtGte sets CreatedAtGte field to given value.

### HasCreatedAtGte

`func (o *SnapshotFilterSchema) HasCreatedAtGte() bool`

HasCreatedAtGte returns a boolean if a field has been set.

### GetCreatedAtLt

`func (o *SnapshotFilterSchema) GetCreatedAtLt() time.Time`

GetCreatedAtLt returns the CreatedAtLt field if non-nil, zero value otherwise.

### GetCreatedAtLtOk

`func (o *SnapshotFilterSchema) GetCreatedAtLtOk() (*time.Time, bool)`

GetCreatedAtLtOk returns a tuple with the CreatedAtLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtLt

`func (o *SnapshotFilterSchema) SetCreatedAtLt(v time.Time)`

SetCreatedAtLt sets CreatedAtLt field to given value.

### HasCreatedAtLt

`func (o *SnapshotFilterSchema) HasCreatedAtLt() bool`

HasCreatedAtLt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SnapshotFilterSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SnapshotFilterSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SnapshotFilterSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SnapshotFilterSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *SnapshotFilterSchema) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SnapshotFilterSchema) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SnapshotFilterSchema) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SnapshotFilterSchema) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedAtGte

`func (o *SnapshotFilterSchema) GetModifiedAtGte() time.Time`

GetModifiedAtGte returns the ModifiedAtGte field if non-nil, zero value otherwise.

### GetModifiedAtGteOk

`func (o *SnapshotFilterSchema) GetModifiedAtGteOk() (*time.Time, bool)`

GetModifiedAtGteOk returns a tuple with the ModifiedAtGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAtGte

`func (o *SnapshotFilterSchema) SetModifiedAtGte(v time.Time)`

SetModifiedAtGte sets ModifiedAtGte field to given value.

### HasModifiedAtGte

`func (o *SnapshotFilterSchema) HasModifiedAtGte() bool`

HasModifiedAtGte returns a boolean if a field has been set.

### GetModifiedAtLt

`func (o *SnapshotFilterSchema) GetModifiedAtLt() time.Time`

GetModifiedAtLt returns the ModifiedAtLt field if non-nil, zero value otherwise.

### GetModifiedAtLtOk

`func (o *SnapshotFilterSchema) GetModifiedAtLtOk() (*time.Time, bool)`

GetModifiedAtLtOk returns a tuple with the ModifiedAtLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAtLt

`func (o *SnapshotFilterSchema) SetModifiedAtLt(v time.Time)`

SetModifiedAtLt sets ModifiedAtLt field to given value.

### HasModifiedAtLt

`func (o *SnapshotFilterSchema) HasModifiedAtLt() bool`

HasModifiedAtLt returns a boolean if a field has been set.

### GetSearch

`func (o *SnapshotFilterSchema) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *SnapshotFilterSchema) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *SnapshotFilterSchema) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *SnapshotFilterSchema) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *SnapshotFilterSchema) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *SnapshotFilterSchema) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetUrl

`func (o *SnapshotFilterSchema) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SnapshotFilterSchema) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SnapshotFilterSchema) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SnapshotFilterSchema) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SnapshotFilterSchema) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SnapshotFilterSchema) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetTag

`func (o *SnapshotFilterSchema) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *SnapshotFilterSchema) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *SnapshotFilterSchema) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *SnapshotFilterSchema) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *SnapshotFilterSchema) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *SnapshotFilterSchema) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetTitle

`func (o *SnapshotFilterSchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SnapshotFilterSchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SnapshotFilterSchema) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SnapshotFilterSchema) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SnapshotFilterSchema) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SnapshotFilterSchema) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTimestamp

`func (o *SnapshotFilterSchema) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SnapshotFilterSchema) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SnapshotFilterSchema) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SnapshotFilterSchema) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SnapshotFilterSchema) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SnapshotFilterSchema) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetBookmarkedAtGte

`func (o *SnapshotFilterSchema) GetBookmarkedAtGte() time.Time`

GetBookmarkedAtGte returns the BookmarkedAtGte field if non-nil, zero value otherwise.

### GetBookmarkedAtGteOk

`func (o *SnapshotFilterSchema) GetBookmarkedAtGteOk() (*time.Time, bool)`

GetBookmarkedAtGteOk returns a tuple with the BookmarkedAtGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarkedAtGte

`func (o *SnapshotFilterSchema) SetBookmarkedAtGte(v time.Time)`

SetBookmarkedAtGte sets BookmarkedAtGte field to given value.

### HasBookmarkedAtGte

`func (o *SnapshotFilterSchema) HasBookmarkedAtGte() bool`

HasBookmarkedAtGte returns a boolean if a field has been set.

### SetBookmarkedAtGteNil

`func (o *SnapshotFilterSchema) SetBookmarkedAtGteNil(b bool)`

 SetBookmarkedAtGteNil sets the value for BookmarkedAtGte to be an explicit nil

### UnsetBookmarkedAtGte
`func (o *SnapshotFilterSchema) UnsetBookmarkedAtGte()`

UnsetBookmarkedAtGte ensures that no value is present for BookmarkedAtGte, not even an explicit nil
### GetBookmarkedAtLt

`func (o *SnapshotFilterSchema) GetBookmarkedAtLt() time.Time`

GetBookmarkedAtLt returns the BookmarkedAtLt field if non-nil, zero value otherwise.

### GetBookmarkedAtLtOk

`func (o *SnapshotFilterSchema) GetBookmarkedAtLtOk() (*time.Time, bool)`

GetBookmarkedAtLtOk returns a tuple with the BookmarkedAtLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarkedAtLt

`func (o *SnapshotFilterSchema) SetBookmarkedAtLt(v time.Time)`

SetBookmarkedAtLt sets BookmarkedAtLt field to given value.

### HasBookmarkedAtLt

`func (o *SnapshotFilterSchema) HasBookmarkedAtLt() bool`

HasBookmarkedAtLt returns a boolean if a field has been set.

### SetBookmarkedAtLtNil

`func (o *SnapshotFilterSchema) SetBookmarkedAtLtNil(b bool)`

 SetBookmarkedAtLtNil sets the value for BookmarkedAtLt to be an explicit nil

### UnsetBookmarkedAtLt
`func (o *SnapshotFilterSchema) UnsetBookmarkedAtLt()`

UnsetBookmarkedAtLt ensures that no value is present for BookmarkedAtLt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


