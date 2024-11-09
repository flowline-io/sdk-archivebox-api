# TagSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TYPE** | Pointer to **string** |  | [optional] [default to "core.models.Tag"]
**Id** | **string** |  | 
**Abid** | **string** |  | 
**ModifiedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**CreatedById** | **string** |  | 
**CreatedByUsername** | **string** |  | 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**NumSnapshots** | **int32** |  | 
**Snapshots** | [**[]SnapshotSchema**](SnapshotSchema.md) |  | 

## Methods

### NewTagSchema

`func NewTagSchema(id string, abid string, modifiedAt time.Time, createdAt time.Time, createdById string, createdByUsername string, name string, slug string, numSnapshots int32, snapshots []SnapshotSchema, ) *TagSchema`

NewTagSchema instantiates a new TagSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagSchemaWithDefaults

`func NewTagSchemaWithDefaults() *TagSchema`

NewTagSchemaWithDefaults instantiates a new TagSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTYPE

`func (o *TagSchema) GetTYPE() string`

GetTYPE returns the TYPE field if non-nil, zero value otherwise.

### GetTYPEOk

`func (o *TagSchema) GetTYPEOk() (*string, bool)`

GetTYPEOk returns a tuple with the TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTYPE

`func (o *TagSchema) SetTYPE(v string)`

SetTYPE sets TYPE field to given value.

### HasTYPE

`func (o *TagSchema) HasTYPE() bool`

HasTYPE returns a boolean if a field has been set.

### GetId

`func (o *TagSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagSchema) SetId(v string)`

SetId sets Id field to given value.


### GetAbid

`func (o *TagSchema) GetAbid() string`

GetAbid returns the Abid field if non-nil, zero value otherwise.

### GetAbidOk

`func (o *TagSchema) GetAbidOk() (*string, bool)`

GetAbidOk returns a tuple with the Abid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbid

`func (o *TagSchema) SetAbid(v string)`

SetAbid sets Abid field to given value.


### GetModifiedAt

`func (o *TagSchema) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *TagSchema) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *TagSchema) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetCreatedAt

`func (o *TagSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TagSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TagSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedById

`func (o *TagSchema) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TagSchema) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TagSchema) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetCreatedByUsername

`func (o *TagSchema) GetCreatedByUsername() string`

GetCreatedByUsername returns the CreatedByUsername field if non-nil, zero value otherwise.

### GetCreatedByUsernameOk

`func (o *TagSchema) GetCreatedByUsernameOk() (*string, bool)`

GetCreatedByUsernameOk returns a tuple with the CreatedByUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUsername

`func (o *TagSchema) SetCreatedByUsername(v string)`

SetCreatedByUsername sets CreatedByUsername field to given value.


### GetName

`func (o *TagSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagSchema) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *TagSchema) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *TagSchema) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *TagSchema) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetNumSnapshots

`func (o *TagSchema) GetNumSnapshots() int32`

GetNumSnapshots returns the NumSnapshots field if non-nil, zero value otherwise.

### GetNumSnapshotsOk

`func (o *TagSchema) GetNumSnapshotsOk() (*int32, bool)`

GetNumSnapshotsOk returns a tuple with the NumSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSnapshots

`func (o *TagSchema) SetNumSnapshots(v int32)`

SetNumSnapshots sets NumSnapshots field to given value.


### GetSnapshots

`func (o *TagSchema) GetSnapshots() []SnapshotSchema`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *TagSchema) GetSnapshotsOk() (*[]SnapshotSchema, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *TagSchema) SetSnapshots(v []SnapshotSchema)`

SetSnapshots sets Snapshots field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


