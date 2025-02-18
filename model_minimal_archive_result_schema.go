/*
ArchiveBox API

 <h3>Welcome to your ArchiveBox server's REST API <code>[v1 ALPHA]</code> homepage!</h3> <br/> <i><b>WARNING: This API is still in an early development stage and may change!</b></i> <br/> <ul> <li>⬅️ Manage your server: <a href=\"/admin/api/\"><b>Setup API Keys</b></a>, <a href=\"/admin/\">Go to your Server Admin UI</a>, <a href=\"/\">Go to your Snapshots list</a>  <li>💬 Ask questions and get help here: <a href=\"https://zulip.archivebox.io\">ArchiveBox Chat Forum</a></li> <li>🐞 Report API bugs here: <a href=\"https://github.com/ArchiveBox/ArchiveBox/issues\">Github Issues</a></li> <li>📚 ArchiveBox Documentation: <a href=\"https://github.com/ArchiveBox/ArchiveBox/wiki\">Github Wiki</a></li> <li>📜 See the API source code: <a href=\"https://github.com/ArchiveBox/ArchiveBox/blob/dev/archivebox/api\"><code>archivebox/api/</code></a></li> </ul> <small>Served by ArchiveBox v0.8.5rc50 (<a href=\"https://github.com/ArchiveBox/ArchiveBox/commit/1dff8bae4043efb678c1a6fb517dfa7e7a211150\"><code>1dff8bae</code></a>), API powered by <a href=\"https://django-ninja.dev/\"><code>django-ninja</code></a>.</small> 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the MinimalArchiveResultSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MinimalArchiveResultSchema{}

// MinimalArchiveResultSchema struct for MinimalArchiveResultSchema
type MinimalArchiveResultSchema struct {
	TYPE *string `json:"TYPE,omitempty"`
	Id string `json:"id"`
	Abid string `json:"abid"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatedAt time.Time `json:"created_at"`
	CreatedById string `json:"created_by_id"`
	CreatedByUsername string `json:"created_by_username"`
	Extractor string `json:"extractor"`
	CmdVersion NullableString `json:"cmd_version"`
	Cmd []string `json:"cmd"`
	Pwd string `json:"pwd"`
	Status string `json:"status"`
	Output string `json:"output"`
	StartTs NullableTime `json:"start_ts"`
	EndTs NullableTime `json:"end_ts"`
}

type _MinimalArchiveResultSchema MinimalArchiveResultSchema

// NewMinimalArchiveResultSchema instantiates a new MinimalArchiveResultSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinimalArchiveResultSchema(id string, abid string, modifiedAt time.Time, createdAt time.Time, createdById string, createdByUsername string, extractor string, cmdVersion NullableString, cmd []string, pwd string, status string, output string, startTs NullableTime, endTs NullableTime) *MinimalArchiveResultSchema {
	this := MinimalArchiveResultSchema{}
	var tYPE string = "core.models.ArchiveResult"
	this.TYPE = &tYPE
	this.Id = id
	this.Abid = abid
	this.ModifiedAt = modifiedAt
	this.CreatedAt = createdAt
	this.CreatedById = createdById
	this.CreatedByUsername = createdByUsername
	this.Extractor = extractor
	this.CmdVersion = cmdVersion
	this.Cmd = cmd
	this.Pwd = pwd
	this.Status = status
	this.Output = output
	this.StartTs = startTs
	this.EndTs = endTs
	return &this
}

// NewMinimalArchiveResultSchemaWithDefaults instantiates a new MinimalArchiveResultSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinimalArchiveResultSchemaWithDefaults() *MinimalArchiveResultSchema {
	this := MinimalArchiveResultSchema{}
	var tYPE string = "core.models.ArchiveResult"
	this.TYPE = &tYPE
	return &this
}

// GetTYPE returns the TYPE field value if set, zero value otherwise.
func (o *MinimalArchiveResultSchema) GetTYPE() string {
	if o == nil || IsNil(o.TYPE) {
		var ret string
		return ret
	}
	return *o.TYPE
}

// GetTYPEOk returns a tuple with the TYPE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalArchiveResultSchema) GetTYPEOk() (*string, bool) {
	if o == nil || IsNil(o.TYPE) {
		return nil, false
	}
	return o.TYPE, true
}

// HasTYPE returns a boolean if a field has been set.
func (o *MinimalArchiveResultSchema) HasTYPE() bool {
	if o != nil && !IsNil(o.TYPE) {
		return true
	}

	return false
}

// SetTYPE gets a reference to the given string and assigns it to the TYPE field.
func (o *MinimalArchiveResultSchema) SetTYPE(v string) {
	o.TYPE = &v
}

// GetId returns the Id field value
func (o *MinimalArchiveResultSchema) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MinimalArchiveResultSchema) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MinimalArchiveResultSchema) SetId(v string) {
	o.Id = v
}

// GetAbid returns the Abid field value
func (o *MinimalArchiveResultSchema) GetAbid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Abid
}

// GetAbidOk returns a tuple with the Abid field value
// and a boolean to check if the value has been set.
func (o *MinimalArchiveResultSchema) GetAbidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Abid, true
}

// SetAbid sets field value
func (o *MinimalArchiveResultSchema) SetAbid(v string) {
	o.Abid = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *MinimalArchiveResultSchema) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *MinimalArchiveResultSchema) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *MinimalArchiveResultSchema) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *MinimalArchiveResultSchema) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MinimalArchiveResultSchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *MinimalArchiveResultSchema) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedById returns the CreatedById field value
func (o *MinimalArchiveResultSchema) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *MinimalArchiveResultSchema) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *MinimalArchiveResultSchema) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetCreatedByUsername returns the CreatedByUsername field value
func (o *MinimalArchiveResultSchema) GetCreatedByUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedByUsername
}

// GetCreatedByUsernameOk returns a tuple with the CreatedByUsername field value
// and a boolean to check if the value has been set.
func (o *MinimalArchiveResultSchema) GetCreatedByUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByUsername, true
}

// SetCreatedByUsername sets field value
func (o *MinimalArchiveResultSchema) SetCreatedByUsername(v string) {
	o.CreatedByUsername = v
}

// GetExtractor returns the Extractor field value
func (o *MinimalArchiveResultSchema) GetExtractor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Extractor
}

// GetExtractorOk returns a tuple with the Extractor field value
// and a boolean to check if the value has been set.
func (o *MinimalArchiveResultSchema) GetExtractorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Extractor, true
}

// SetExtractor sets field value
func (o *MinimalArchiveResultSchema) SetExtractor(v string) {
	o.Extractor = v
}

// GetCmdVersion returns the CmdVersion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MinimalArchiveResultSchema) GetCmdVersion() string {
	if o == nil || o.CmdVersion.Get() == nil {
		var ret string
		return ret
	}

	return *o.CmdVersion.Get()
}

// GetCmdVersionOk returns a tuple with the CmdVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MinimalArchiveResultSchema) GetCmdVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CmdVersion.Get(), o.CmdVersion.IsSet()
}

// SetCmdVersion sets field value
func (o *MinimalArchiveResultSchema) SetCmdVersion(v string) {
	o.CmdVersion.Set(&v)
}

// GetCmd returns the Cmd field value
func (o *MinimalArchiveResultSchema) GetCmd() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Cmd
}

// GetCmdOk returns a tuple with the Cmd field value
// and a boolean to check if the value has been set.
func (o *MinimalArchiveResultSchema) GetCmdOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cmd, true
}

// SetCmd sets field value
func (o *MinimalArchiveResultSchema) SetCmd(v []string) {
	o.Cmd = v
}

// GetPwd returns the Pwd field value
func (o *MinimalArchiveResultSchema) GetPwd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pwd
}

// GetPwdOk returns a tuple with the Pwd field value
// and a boolean to check if the value has been set.
func (o *MinimalArchiveResultSchema) GetPwdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pwd, true
}

// SetPwd sets field value
func (o *MinimalArchiveResultSchema) SetPwd(v string) {
	o.Pwd = v
}

// GetStatus returns the Status field value
func (o *MinimalArchiveResultSchema) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MinimalArchiveResultSchema) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *MinimalArchiveResultSchema) SetStatus(v string) {
	o.Status = v
}

// GetOutput returns the Output field value
func (o *MinimalArchiveResultSchema) GetOutput() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Output
}

// GetOutputOk returns a tuple with the Output field value
// and a boolean to check if the value has been set.
func (o *MinimalArchiveResultSchema) GetOutputOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Output, true
}

// SetOutput sets field value
func (o *MinimalArchiveResultSchema) SetOutput(v string) {
	o.Output = v
}

// GetStartTs returns the StartTs field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *MinimalArchiveResultSchema) GetStartTs() time.Time {
	if o == nil || o.StartTs.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.StartTs.Get()
}

// GetStartTsOk returns a tuple with the StartTs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MinimalArchiveResultSchema) GetStartTsOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTs.Get(), o.StartTs.IsSet()
}

// SetStartTs sets field value
func (o *MinimalArchiveResultSchema) SetStartTs(v time.Time) {
	o.StartTs.Set(&v)
}

// GetEndTs returns the EndTs field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *MinimalArchiveResultSchema) GetEndTs() time.Time {
	if o == nil || o.EndTs.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.EndTs.Get()
}

// GetEndTsOk returns a tuple with the EndTs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MinimalArchiveResultSchema) GetEndTsOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndTs.Get(), o.EndTs.IsSet()
}

// SetEndTs sets field value
func (o *MinimalArchiveResultSchema) SetEndTs(v time.Time) {
	o.EndTs.Set(&v)
}

func (o MinimalArchiveResultSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MinimalArchiveResultSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TYPE) {
		toSerialize["TYPE"] = o.TYPE
	}
	toSerialize["id"] = o.Id
	toSerialize["abid"] = o.Abid
	toSerialize["modified_at"] = o.ModifiedAt
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["created_by_id"] = o.CreatedById
	toSerialize["created_by_username"] = o.CreatedByUsername
	toSerialize["extractor"] = o.Extractor
	toSerialize["cmd_version"] = o.CmdVersion.Get()
	toSerialize["cmd"] = o.Cmd
	toSerialize["pwd"] = o.Pwd
	toSerialize["status"] = o.Status
	toSerialize["output"] = o.Output
	toSerialize["start_ts"] = o.StartTs.Get()
	toSerialize["end_ts"] = o.EndTs.Get()
	return toSerialize, nil
}

func (o *MinimalArchiveResultSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"abid",
		"modified_at",
		"created_at",
		"created_by_id",
		"created_by_username",
		"extractor",
		"cmd_version",
		"cmd",
		"pwd",
		"status",
		"output",
		"start_ts",
		"end_ts",
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

	varMinimalArchiveResultSchema := _MinimalArchiveResultSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMinimalArchiveResultSchema)

	if err != nil {
		return err
	}

	*o = MinimalArchiveResultSchema(varMinimalArchiveResultSchema)

	return err
}

type NullableMinimalArchiveResultSchema struct {
	value *MinimalArchiveResultSchema
	isSet bool
}

func (v NullableMinimalArchiveResultSchema) Get() *MinimalArchiveResultSchema {
	return v.value
}

func (v *NullableMinimalArchiveResultSchema) Set(val *MinimalArchiveResultSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimalArchiveResultSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimalArchiveResultSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimalArchiveResultSchema(val *MinimalArchiveResultSchema) *NullableMinimalArchiveResultSchema {
	return &NullableMinimalArchiveResultSchema{value: val, isSet: true}
}

func (v NullableMinimalArchiveResultSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimalArchiveResultSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


