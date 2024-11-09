# PagedArchiveResultSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalItems** | **int32** |  | 
**TotalPages** | **int32** |  | 
**Page** | **int32** |  | 
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 
**NumItems** | **int32** |  | 
**Items** | [**[]ArchiveResultSchema**](ArchiveResultSchema.md) |  | 

## Methods

### NewPagedArchiveResultSchema

`func NewPagedArchiveResultSchema(totalItems int32, totalPages int32, page int32, limit int32, offset int32, numItems int32, items []ArchiveResultSchema, ) *PagedArchiveResultSchema`

NewPagedArchiveResultSchema instantiates a new PagedArchiveResultSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedArchiveResultSchemaWithDefaults

`func NewPagedArchiveResultSchemaWithDefaults() *PagedArchiveResultSchema`

NewPagedArchiveResultSchemaWithDefaults instantiates a new PagedArchiveResultSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalItems

`func (o *PagedArchiveResultSchema) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *PagedArchiveResultSchema) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *PagedArchiveResultSchema) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.


### GetTotalPages

`func (o *PagedArchiveResultSchema) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PagedArchiveResultSchema) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PagedArchiveResultSchema) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.


### GetPage

`func (o *PagedArchiveResultSchema) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PagedArchiveResultSchema) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PagedArchiveResultSchema) SetPage(v int32)`

SetPage sets Page field to given value.


### GetLimit

`func (o *PagedArchiveResultSchema) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PagedArchiveResultSchema) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PagedArchiveResultSchema) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *PagedArchiveResultSchema) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PagedArchiveResultSchema) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PagedArchiveResultSchema) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetNumItems

`func (o *PagedArchiveResultSchema) GetNumItems() int32`

GetNumItems returns the NumItems field if non-nil, zero value otherwise.

### GetNumItemsOk

`func (o *PagedArchiveResultSchema) GetNumItemsOk() (*int32, bool)`

GetNumItemsOk returns a tuple with the NumItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumItems

`func (o *PagedArchiveResultSchema) SetNumItems(v int32)`

SetNumItems sets NumItems field to given value.


### GetItems

`func (o *PagedArchiveResultSchema) GetItems() []ArchiveResultSchema`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PagedArchiveResultSchema) GetItemsOk() (*[]ArchiveResultSchema, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PagedArchiveResultSchema) SetItems(v []ArchiveResultSchema)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


