/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
)

// checks if the APIServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIServer{}

// APIServer struct for APIServer
type APIServer struct {
	AuthorizationServer APIServerAuthorizationServer `json:"authorizationServer"`
	// An array of string that specifies the possible base URLs that an end-user will use to access the APIs hosted on the customer's API server. Multiple base URLs may be specified to support cases where the same API may be available from multiple URLs (for example, from a user-friendly domain URL and an internal domain URL). Base URLs must be valid absolute URLs with the https or http scheme. If the path component is non-empty, it must not end in a trailing slash. The path must not contain empty backslash, dot, or double-dot segments. It must not have a query or fragment present, and the host portion of the authority must be a DNS hostname or valid IP (IPv4 or IPv6). The length must be less than or equal to 256 characters.
	BaseURLs []string `json:"baseURLs"`
	// A string that specifies the resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// A string that specifies the API server resource name. The name value must be unique among all API servers, and it must be a valid resource name.
	Name string `json:"name"`
	// A map from the operation name to the operation object. Each key must be valid ObjectName, and each value must be a valid operation. Each key must be unique within the operations object, which means the operation key is unique within an API server. No duplicate operation values are allowed; operations with the same paths and methods members are not allowed. The operations object is limited to 25 keys (25 individual operations).
	Operations map[string]interface{} `json:"operations,omitempty"`
}

// NewAPIServer instantiates a new APIServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServer(authorizationServer APIServerAuthorizationServer, baseURLs []string, name string) *APIServer {
	this := APIServer{}
	this.AuthorizationServer = authorizationServer
	this.BaseURLs = baseURLs
	this.Name = name
	return &this
}

// NewAPIServerWithDefaults instantiates a new APIServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIServerWithDefaults() *APIServer {
	this := APIServer{}
	return &this
}

// GetAuthorizationServer returns the AuthorizationServer field value
func (o *APIServer) GetAuthorizationServer() APIServerAuthorizationServer {
	if o == nil {
		var ret APIServerAuthorizationServer
		return ret
	}

	return o.AuthorizationServer
}

// GetAuthorizationServerOk returns a tuple with the AuthorizationServer field value
// and a boolean to check if the value has been set.
func (o *APIServer) GetAuthorizationServerOk() (*APIServerAuthorizationServer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationServer, true
}

// SetAuthorizationServer sets field value
func (o *APIServer) SetAuthorizationServer(v APIServerAuthorizationServer) {
	o.AuthorizationServer = v
}

// GetBaseURLs returns the BaseURLs field value
func (o *APIServer) GetBaseURLs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BaseURLs
}

// GetBaseURLsOk returns a tuple with the BaseURLs field value
// and a boolean to check if the value has been set.
func (o *APIServer) GetBaseURLsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaseURLs, true
}

// SetBaseURLs sets field value
func (o *APIServer) SetBaseURLs(v []string) {
	o.BaseURLs = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *APIServer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *APIServer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *APIServer) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *APIServer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *APIServer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *APIServer) SetName(v string) {
	o.Name = v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *APIServer) GetOperations() map[string]interface{} {
	if o == nil || IsNil(o.Operations) {
		var ret map[string]interface{}
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServer) GetOperationsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Operations) {
		return map[string]interface{}{}, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *APIServer) HasOperations() bool {
	if o != nil && !IsNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given map[string]interface{} and assigns it to the Operations field.
func (o *APIServer) SetOperations(v map[string]interface{}) {
	o.Operations = v
}

func (o APIServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authorizationServer"] = o.AuthorizationServer
	toSerialize["baseURLs"] = o.BaseURLs
	// skip: id is readOnly
	toSerialize["name"] = o.Name
	if !IsNil(o.Operations) {
		toSerialize["operations"] = o.Operations
	}
	return toSerialize, nil
}

type NullableAPIServer struct {
	value *APIServer
	isSet bool
}

func (v NullableAPIServer) Get() *APIServer {
	return v.value
}

func (v *NullableAPIServer) Set(val *APIServer) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIServer) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIServer(val *APIServer) *NullableAPIServer {
	return &NullableAPIServer{value: val, isSet: true}
}

func (v NullableAPIServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


