/*
ArchiveBox API

 <h3>Welcome to your ArchiveBox server's REST API <code>[v1 ALPHA]</code> homepage!</h3> <br/> <i><b>WARNING: This API is still in an early development stage and may change!</b></i> <br/> <ul> <li>⬅️ Manage your server: <a href=\"/admin/api/\"><b>Setup API Keys</b></a>, <a href=\"/admin/\">Go to your Server Admin UI</a>, <a href=\"/\">Go to your Snapshots list</a>  <li>💬 Ask questions and get help here: <a href=\"https://zulip.archivebox.io\">ArchiveBox Chat Forum</a></li> <li>🐞 Report API bugs here: <a href=\"https://github.com/ArchiveBox/ArchiveBox/issues\">Github Issues</a></li> <li>📚 ArchiveBox Documentation: <a href=\"https://github.com/ArchiveBox/ArchiveBox/wiki\">Github Wiki</a></li> <li>📜 See the API source code: <a href=\"https://github.com/ArchiveBox/ArchiveBox/blob/dev/archivebox/api\"><code>archivebox/api/</code></a></li> </ul> <small>Served by ArchiveBox v0.8.5rc50 (<a href=\"https://github.com/ArchiveBox/ArchiveBox/commit/1dff8bae4043efb678c1a6fb517dfa7e7a211150\"><code>1dff8bae</code></a>), API powered by <a href=\"https://django-ninja.dev/\"><code>django-ninja</code></a>.</small> 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)


// AsCsv struct for AsCsv
type AsCsv struct {
	Bool *bool
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AsCsv) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Bool
	err = json.Unmarshal(data, &dst.Bool);
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			return nil // data stored in dst.Bool, return on the first match
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AsCsv)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AsCsv) MarshalJSON() ([]byte, error) {
	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAsCsv struct {
	value *AsCsv
	isSet bool
}

func (v NullableAsCsv) Get() *AsCsv {
	return v.value
}

func (v *NullableAsCsv) Set(val *AsCsv) {
	v.value = val
	v.isSet = true
}

func (v NullableAsCsv) IsSet() bool {
	return v.isSet
}

func (v *NullableAsCsv) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsCsv(val *AsCsv) *NullableAsCsv {
	return &NullableAsCsv{value: val, isSet: true}
}

func (v NullableAsCsv) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsCsv) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


