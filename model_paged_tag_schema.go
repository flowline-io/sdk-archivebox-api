/*
ArchiveBox API

 <h3>Welcome to your ArchiveBox server's REST API <code>[v1 ALPHA]</code> homepage!</h3> <br/> <i><b>WARNING: This API is still in an early development stage and may change!</b></i> <br/> <ul> <li>⬅️ Manage your server: <a href=\"/admin/api/\"><b>Setup API Keys</b></a>, <a href=\"/admin/\">Go to your Server Admin UI</a>, <a href=\"/\">Go to your Snapshots list</a>  <li>💬 Ask questions and get help here: <a href=\"https://zulip.archivebox.io\">ArchiveBox Chat Forum</a></li> <li>🐞 Report API bugs here: <a href=\"https://github.com/ArchiveBox/ArchiveBox/issues\">Github Issues</a></li> <li>📚 ArchiveBox Documentation: <a href=\"https://github.com/ArchiveBox/ArchiveBox/wiki\">Github Wiki</a></li> <li>📜 See the API source code: <a href=\"https://github.com/ArchiveBox/ArchiveBox/blob/dev/archivebox/api\"><code>archivebox/api/</code></a></li> </ul> <small>Served by ArchiveBox v0.8.5rc50 (<a href=\"https://github.com/ArchiveBox/ArchiveBox/commit/1dff8bae4043efb678c1a6fb517dfa7e7a211150\"><code>1dff8bae</code></a>), API powered by <a href=\"https://django-ninja.dev/\"><code>django-ninja</code></a>.</small> 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PagedTagSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagedTagSchema{}

// PagedTagSchema struct for PagedTagSchema
type PagedTagSchema struct {
	TotalItems int32 `json:"total_items"`
	TotalPages int32 `json:"total_pages"`
	Page int32 `json:"page"`
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
	NumItems int32 `json:"num_items"`
	Items []TagSchema `json:"items"`
}

type _PagedTagSchema PagedTagSchema

// NewPagedTagSchema instantiates a new PagedTagSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedTagSchema(totalItems int32, totalPages int32, page int32, limit int32, offset int32, numItems int32, items []TagSchema) *PagedTagSchema {
	this := PagedTagSchema{}
	this.TotalItems = totalItems
	this.TotalPages = totalPages
	this.Page = page
	this.Limit = limit
	this.Offset = offset
	this.NumItems = numItems
	this.Items = items
	return &this
}

// NewPagedTagSchemaWithDefaults instantiates a new PagedTagSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedTagSchemaWithDefaults() *PagedTagSchema {
	this := PagedTagSchema{}
	return &this
}

// GetTotalItems returns the TotalItems field value
func (o *PagedTagSchema) GetTotalItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value
// and a boolean to check if the value has been set.
func (o *PagedTagSchema) GetTotalItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalItems, true
}

// SetTotalItems sets field value
func (o *PagedTagSchema) SetTotalItems(v int32) {
	o.TotalItems = v
}

// GetTotalPages returns the TotalPages field value
func (o *PagedTagSchema) GetTotalPages() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value
// and a boolean to check if the value has been set.
func (o *PagedTagSchema) GetTotalPagesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPages, true
}

// SetTotalPages sets field value
func (o *PagedTagSchema) SetTotalPages(v int32) {
	o.TotalPages = v
}

// GetPage returns the Page field value
func (o *PagedTagSchema) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *PagedTagSchema) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *PagedTagSchema) SetPage(v int32) {
	o.Page = v
}

// GetLimit returns the Limit field value
func (o *PagedTagSchema) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PagedTagSchema) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *PagedTagSchema) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *PagedTagSchema) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *PagedTagSchema) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *PagedTagSchema) SetOffset(v int32) {
	o.Offset = v
}

// GetNumItems returns the NumItems field value
func (o *PagedTagSchema) GetNumItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumItems
}

// GetNumItemsOk returns a tuple with the NumItems field value
// and a boolean to check if the value has been set.
func (o *PagedTagSchema) GetNumItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumItems, true
}

// SetNumItems sets field value
func (o *PagedTagSchema) SetNumItems(v int32) {
	o.NumItems = v
}

// GetItems returns the Items field value
func (o *PagedTagSchema) GetItems() []TagSchema {
	if o == nil {
		var ret []TagSchema
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PagedTagSchema) GetItemsOk() ([]TagSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *PagedTagSchema) SetItems(v []TagSchema) {
	o.Items = v
}

func (o PagedTagSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagedTagSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total_items"] = o.TotalItems
	toSerialize["total_pages"] = o.TotalPages
	toSerialize["page"] = o.Page
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["num_items"] = o.NumItems
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *PagedTagSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total_items",
		"total_pages",
		"page",
		"limit",
		"offset",
		"num_items",
		"items",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPagedTagSchema := _PagedTagSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPagedTagSchema)

	if err != nil {
		return err
	}

	*o = PagedTagSchema(varPagedTagSchema)

	return err
}

type NullablePagedTagSchema struct {
	value *PagedTagSchema
	isSet bool
}

func (v NullablePagedTagSchema) Get() *PagedTagSchema {
	return v.value
}

func (v *NullablePagedTagSchema) Set(val *PagedTagSchema) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedTagSchema) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedTagSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedTagSchema(val *PagedTagSchema) *NullablePagedTagSchema {
	return &NullablePagedTagSchema{value: val, isSet: true}
}

func (v NullablePagedTagSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedTagSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

