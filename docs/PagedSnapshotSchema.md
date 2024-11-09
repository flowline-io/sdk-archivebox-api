# PagedSnapshotSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalItems** | **int32** |  | 
**TotalPages** | **int32** |  | 
**Page** | **int32** |  | 
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 
**NumItems** | **int32** |  | 
**Items** | [**[]SnapshotSchema**](SnapshotSchema.md) |  | 

## Methods

### NewPagedSnapshotSchema

`func NewPagedSnapshotSchema(totalItems int32, totalPages int32, page int32, limit int32, offset int32, numItems int32, items []SnapshotSchema, ) *PagedSnapshotSchema`

NewPagedSnapshotSchema instantiates a new PagedSnapshotSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedSnapshotSchemaWithDefaults

`func NewPagedSnapshotSchemaWithDefaults() *PagedSnapshotSchema`

NewPagedSnapshotSchemaWithDefaults instantiates a new PagedSnapshotSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalItems

`func (o *PagedSnapshotSchema) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *PagedSnapshotSchema) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *PagedSnapshotSchema) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.


### GetTotalPages

`func (o *PagedSnapshotSchema) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PagedSnapshotSchema) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PagedSnapshotSchema) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.


### GetPage

`func (o *PagedSnapshotSchema) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PagedSnapshotSchema) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PagedSnapshotSchema) SetPage(v int32)`

SetPage sets Page field to given value.


### GetLimit

`func (o *PagedSnapshotSchema) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PagedSnapshotSchema) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PagedSnapshotSchema) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *PagedSnapshotSchema) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PagedSnapshotSchema) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PagedSnapshotSchema) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetNumItems

`func (o *PagedSnapshotSchema) GetNumItems() int32`

GetNumItems returns the NumItems field if non-nil, zero value otherwise.

### GetNumItemsOk

`func (o *PagedSnapshotSchema) GetNumItemsOk() (*int32, bool)`

GetNumItemsOk returns a tuple with the NumItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumItems

`func (o *PagedSnapshotSchema) SetNumItems(v int32)`

SetNumItems sets NumItems field to given value.


### GetItems

`func (o *PagedSnapshotSchema) GetItems() []SnapshotSchema`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PagedSnapshotSchema) GetItemsOk() (*[]SnapshotSchema, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PagedSnapshotSchema) SetItems(v []SnapshotSchema)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


