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

// checks if the PasswordAuthSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordAuthSchema{}

// PasswordAuthSchema Schema for a /get_api_token request
type PasswordAuthSchema struct {
	Username NullableString `json:"username,omitempty"`
	Password NullableString `json:"password,omitempty"`
}

// NewPasswordAuthSchema instantiates a new PasswordAuthSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordAuthSchema() *PasswordAuthSchema {
	this := PasswordAuthSchema{}
	return &this
}

// NewPasswordAuthSchemaWithDefaults instantiates a new PasswordAuthSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordAuthSchemaWithDefaults() *PasswordAuthSchema {
	this := PasswordAuthSchema{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordAuthSchema) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordAuthSchema) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *PasswordAuthSchema) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *PasswordAuthSchema) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *PasswordAuthSchema) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *PasswordAuthSchema) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordAuthSchema) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordAuthSchema) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *PasswordAuthSchema) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *PasswordAuthSchema) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *PasswordAuthSchema) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *PasswordAuthSchema) UnsetPassword() {
	o.Password.Unset()
}

func (o PasswordAuthSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordAuthSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	return toSerialize, nil
}

type NullablePasswordAuthSchema struct {
	value *PasswordAuthSchema
	isSet bool
}

func (v NullablePasswordAuthSchema) Get() *PasswordAuthSchema {
	return v.value
}

func (v *NullablePasswordAuthSchema) Set(val *PasswordAuthSchema) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordAuthSchema) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordAuthSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordAuthSchema(val *PasswordAuthSchema) *NullablePasswordAuthSchema {
	return &NullablePasswordAuthSchema{value: val, isSet: true}
}

func (v NullablePasswordAuthSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordAuthSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

