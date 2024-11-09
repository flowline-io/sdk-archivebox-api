/*
ArchiveBox API

 <h3>Welcome to your ArchiveBox server's REST API <code>[v1 ALPHA]</code> homepage!</h3> <br/> <i><b>WARNING: This API is still in an early development stage and may change!</b></i> <br/> <ul> <li>⬅️ Manage your server: <a href=\"/admin/api/\"><b>Setup API Keys</b></a>, <a href=\"/admin/\">Go to your Server Admin UI</a>, <a href=\"/\">Go to your Snapshots list</a>  <li>💬 Ask questions and get help here: <a href=\"https://zulip.archivebox.io\">ArchiveBox Chat Forum</a></li> <li>🐞 Report API bugs here: <a href=\"https://github.com/ArchiveBox/ArchiveBox/issues\">Github Issues</a></li> <li>📚 ArchiveBox Documentation: <a href=\"https://github.com/ArchiveBox/ArchiveBox/wiki\">Github Wiki</a></li> <li>📜 See the API source code: <a href=\"https://github.com/ArchiveBox/ArchiveBox/blob/dev/archivebox/api\"><code>archivebox/api/</code></a></li> </ul> <small>Served by ArchiveBox v0.8.5rc50 (<a href=\"https://github.com/ArchiveBox/ArchiveBox/commit/1dff8bae4043efb678c1a6fb517dfa7e7a211150\"><code>1dff8bae</code></a>), API powered by <a href=\"https://django-ninja.dev/\"><code>django-ninja</code></a>.</small> 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ListCommandSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCommandSchema{}

// ListCommandSchema struct for ListCommandSchema
type ListCommandSchema struct {
	FilterPatterns []string `json:"filter_patterns,omitempty"`
	FilterType *string `json:"filter_type,omitempty"`
	Status NullableStatusChoices `json:"status,omitempty"`
	After NullableFloat32 `json:"after,omitempty"`
	Before NullableFloat32 `json:"before,omitempty"`
	Sort *string `json:"sort,omitempty"`
	AsJson *bool `json:"as_json,omitempty"`
	AsHtml *bool `json:"as_html,omitempty"`
	AsCsv *AsCsv `json:"as_csv,omitempty"`
	WithHeaders *bool `json:"with_headers,omitempty"`
}

// NewListCommandSchema instantiates a new ListCommandSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCommandSchema() *ListCommandSchema {
	this := ListCommandSchema{}
	var filterType string = "substring"
	this.FilterType = &filterType
	var sort string = "bookmarked_at"
	this.Sort = &sort
	var asJson bool = true
	this.AsJson = &asJson
	var asHtml bool = false
	this.AsHtml = &asHtml
	// var asCsv AsCsv = timestamp,url FIXME
	// this.AsCsv = &asCsv FIXME
	var withHeaders bool = false
	this.WithHeaders = &withHeaders
	return &this
}

// NewListCommandSchemaWithDefaults instantiates a new ListCommandSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCommandSchemaWithDefaults() *ListCommandSchema {
	this := ListCommandSchema{}
	var filterType string = "substring"
	this.FilterType = &filterType
	var sort string = "bookmarked_at"
	this.Sort = &sort
	var asJson bool = true
	this.AsJson = &asJson
	var asHtml bool = false
	this.AsHtml = &asHtml
	// var asCsv AsCsv = timestamp,url FIXME
	// this.AsCsv = &asCsv FIXME
	var withHeaders bool = false
	this.WithHeaders = &withHeaders
	return &this
}

// GetFilterPatterns returns the FilterPatterns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListCommandSchema) GetFilterPatterns() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.FilterPatterns
}

// GetFilterPatternsOk returns a tuple with the FilterPatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListCommandSchema) GetFilterPatternsOk() ([]string, bool) {
	if o == nil || IsNil(o.FilterPatterns) {
		return nil, false
	}
	return o.FilterPatterns, true
}

// HasFilterPatterns returns a boolean if a field has been set.
func (o *ListCommandSchema) HasFilterPatterns() bool {
	if o != nil && !IsNil(o.FilterPatterns) {
		return true
	}

	return false
}

// SetFilterPatterns gets a reference to the given []string and assigns it to the FilterPatterns field.
func (o *ListCommandSchema) SetFilterPatterns(v []string) {
	o.FilterPatterns = v
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *ListCommandSchema) GetFilterType() string {
	if o == nil || IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCommandSchema) GetFilterTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *ListCommandSchema) HasFilterType() bool {
	if o != nil && !IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *ListCommandSchema) SetFilterType(v string) {
	o.FilterType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListCommandSchema) GetStatus() StatusChoices {
	if o == nil || IsNil(o.Status.Get()) {
		var ret StatusChoices
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListCommandSchema) GetStatusOk() (*StatusChoices, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ListCommandSchema) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableStatusChoices and assigns it to the Status field.
func (o *ListCommandSchema) SetStatus(v StatusChoices) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *ListCommandSchema) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ListCommandSchema) UnsetStatus() {
	o.Status.Unset()
}

// GetAfter returns the After field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListCommandSchema) GetAfter() float32 {
	if o == nil || IsNil(o.After.Get()) {
		var ret float32
		return ret
	}
	return *o.After.Get()
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListCommandSchema) GetAfterOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.After.Get(), o.After.IsSet()
}

// HasAfter returns a boolean if a field has been set.
func (o *ListCommandSchema) HasAfter() bool {
	if o != nil && o.After.IsSet() {
		return true
	}

	return false
}

// SetAfter gets a reference to the given NullableFloat32 and assigns it to the After field.
func (o *ListCommandSchema) SetAfter(v float32) {
	o.After.Set(&v)
}
// SetAfterNil sets the value for After to be an explicit nil
func (o *ListCommandSchema) SetAfterNil() {
	o.After.Set(nil)
}

// UnsetAfter ensures that no value is present for After, not even an explicit nil
func (o *ListCommandSchema) UnsetAfter() {
	o.After.Unset()
}

// GetBefore returns the Before field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListCommandSchema) GetBefore() float32 {
	if o == nil || IsNil(o.Before.Get()) {
		var ret float32
		return ret
	}
	return *o.Before.Get()
}

// GetBeforeOk returns a tuple with the Before field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListCommandSchema) GetBeforeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Before.Get(), o.Before.IsSet()
}

// HasBefore returns a boolean if a field has been set.
func (o *ListCommandSchema) HasBefore() bool {
	if o != nil && o.Before.IsSet() {
		return true
	}

	return false
}

// SetBefore gets a reference to the given NullableFloat32 and assigns it to the Before field.
func (o *ListCommandSchema) SetBefore(v float32) {
	o.Before.Set(&v)
}
// SetBeforeNil sets the value for Before to be an explicit nil
func (o *ListCommandSchema) SetBeforeNil() {
	o.Before.Set(nil)
}

// UnsetBefore ensures that no value is present for Before, not even an explicit nil
func (o *ListCommandSchema) UnsetBefore() {
	o.Before.Unset()
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ListCommandSchema) GetSort() string {
	if o == nil || IsNil(o.Sort) {
		var ret string
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCommandSchema) GetSortOk() (*string, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ListCommandSchema) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given string and assigns it to the Sort field.
func (o *ListCommandSchema) SetSort(v string) {
	o.Sort = &v
}

// GetAsJson returns the AsJson field value if set, zero value otherwise.
func (o *ListCommandSchema) GetAsJson() bool {
	if o == nil || IsNil(o.AsJson) {
		var ret bool
		return ret
	}
	return *o.AsJson
}

// GetAsJsonOk returns a tuple with the AsJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCommandSchema) GetAsJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.AsJson) {
		return nil, false
	}
	return o.AsJson, true
}

// HasAsJson returns a boolean if a field has been set.
func (o *ListCommandSchema) HasAsJson() bool {
	if o != nil && !IsNil(o.AsJson) {
		return true
	}

	return false
}

// SetAsJson gets a reference to the given bool and assigns it to the AsJson field.
func (o *ListCommandSchema) SetAsJson(v bool) {
	o.AsJson = &v
}

// GetAsHtml returns the AsHtml field value if set, zero value otherwise.
func (o *ListCommandSchema) GetAsHtml() bool {
	if o == nil || IsNil(o.AsHtml) {
		var ret bool
		return ret
	}
	return *o.AsHtml
}

// GetAsHtmlOk returns a tuple with the AsHtml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCommandSchema) GetAsHtmlOk() (*bool, bool) {
	if o == nil || IsNil(o.AsHtml) {
		return nil, false
	}
	return o.AsHtml, true
}

// HasAsHtml returns a boolean if a field has been set.
func (o *ListCommandSchema) HasAsHtml() bool {
	if o != nil && !IsNil(o.AsHtml) {
		return true
	}

	return false
}

// SetAsHtml gets a reference to the given bool and assigns it to the AsHtml field.
func (o *ListCommandSchema) SetAsHtml(v bool) {
	o.AsHtml = &v
}

// GetAsCsv returns the AsCsv field value if set, zero value otherwise.
func (o *ListCommandSchema) GetAsCsv() AsCsv {
	if o == nil || IsNil(o.AsCsv) {
		var ret AsCsv
		return ret
	}
	return *o.AsCsv
}

// GetAsCsvOk returns a tuple with the AsCsv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCommandSchema) GetAsCsvOk() (*AsCsv, bool) {
	if o == nil || IsNil(o.AsCsv) {
		return nil, false
	}
	return o.AsCsv, true
}

// HasAsCsv returns a boolean if a field has been set.
func (o *ListCommandSchema) HasAsCsv() bool {
	if o != nil && !IsNil(o.AsCsv) {
		return true
	}

	return false
}

// SetAsCsv gets a reference to the given AsCsv and assigns it to the AsCsv field.
func (o *ListCommandSchema) SetAsCsv(v AsCsv) {
	o.AsCsv = &v
}

// GetWithHeaders returns the WithHeaders field value if set, zero value otherwise.
func (o *ListCommandSchema) GetWithHeaders() bool {
	if o == nil || IsNil(o.WithHeaders) {
		var ret bool
		return ret
	}
	return *o.WithHeaders
}

// GetWithHeadersOk returns a tuple with the WithHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCommandSchema) GetWithHeadersOk() (*bool, bool) {
	if o == nil || IsNil(o.WithHeaders) {
		return nil, false
	}
	return o.WithHeaders, true
}

// HasWithHeaders returns a boolean if a field has been set.
func (o *ListCommandSchema) HasWithHeaders() bool {
	if o != nil && !IsNil(o.WithHeaders) {
		return true
	}

	return false
}

// SetWithHeaders gets a reference to the given bool and assigns it to the WithHeaders field.
func (o *ListCommandSchema) SetWithHeaders(v bool) {
	o.WithHeaders = &v
}

func (o ListCommandSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCommandSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FilterPatterns != nil {
		toSerialize["filter_patterns"] = o.FilterPatterns
	}
	if !IsNil(o.FilterType) {
		toSerialize["filter_type"] = o.FilterType
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.After.IsSet() {
		toSerialize["after"] = o.After.Get()
	}
	if o.Before.IsSet() {
		toSerialize["before"] = o.Before.Get()
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !IsNil(o.AsJson) {
		toSerialize["as_json"] = o.AsJson
	}
	if !IsNil(o.AsHtml) {
		toSerialize["as_html"] = o.AsHtml
	}
	if !IsNil(o.AsCsv) {
		toSerialize["as_csv"] = o.AsCsv
	}
	if !IsNil(o.WithHeaders) {
		toSerialize["with_headers"] = o.WithHeaders
	}
	return toSerialize, nil
}

type NullableListCommandSchema struct {
	value *ListCommandSchema
	isSet bool
}

func (v NullableListCommandSchema) Get() *ListCommandSchema {
	return v.value
}

func (v *NullableListCommandSchema) Set(val *ListCommandSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableListCommandSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableListCommandSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCommandSchema(val *ListCommandSchema) *NullableListCommandSchema {
	return &NullableListCommandSchema{value: val, isSet: true}
}

func (v NullableListCommandSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCommandSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


