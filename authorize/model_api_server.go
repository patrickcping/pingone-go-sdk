/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the APIServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIServer{}

// APIServer struct for APIServer
type APIServer struct {
	Links map[string]LinksHATEOASValue `json:"_links,omitempty"`
	AccessControl *APIServerAccessControl `json:"accessControl,omitempty"`
	AuthorizationServer APIServerAuthorizationServer `json:"authorizationServer"`
	// An array of string that specifies the possible base URLs that an end-user will use to access the APIs hosted on the customer's API server. Multiple base URLs may be specified to support cases where the same API may be available from multiple URLs (for example, from a user-friendly domain URL and an internal domain URL). Base URLs must be valid absolute URLs with the https or http scheme. If the path component is non-empty, it must not end in a trailing slash. The path must not contain empty backslash, dot, or double-dot segments. It must not have a query or fragment present, and the host portion of the authority must be a DNS hostname or valid IP (IPv4 or IPv6). The length must be less than or equal to 256 characters.
	BaseUrls []string `json:"baseUrls"`
	Directory *APIServerDirectory `json:"directory,omitempty"`
	// A string that specifies the resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// A string that specifies the API server resource name. The name value must be unique among all API servers, and it must be a valid resource name.
	Name string `json:"name"`
	Policy *APIServerPolicy `json:"policy,omitempty"`
}

type _APIServer APIServer

// NewAPIServer instantiates a new APIServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServer(authorizationServer APIServerAuthorizationServer, baseUrls []string, name string) *APIServer {
	this := APIServer{}
	this.AuthorizationServer = authorizationServer
	this.BaseUrls = baseUrls
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

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *APIServer) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServer) GetLinksOk() (map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return map[string]LinksHATEOASValue{}, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *APIServer) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *APIServer) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = v
}

// GetAccessControl returns the AccessControl field value if set, zero value otherwise.
func (o *APIServer) GetAccessControl() APIServerAccessControl {
	if o == nil || IsNil(o.AccessControl) {
		var ret APIServerAccessControl
		return ret
	}
	return *o.AccessControl
}

// GetAccessControlOk returns a tuple with the AccessControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServer) GetAccessControlOk() (*APIServerAccessControl, bool) {
	if o == nil || IsNil(o.AccessControl) {
		return nil, false
	}
	return o.AccessControl, true
}

// HasAccessControl returns a boolean if a field has been set.
func (o *APIServer) HasAccessControl() bool {
	if o != nil && !IsNil(o.AccessControl) {
		return true
	}

	return false
}

// SetAccessControl gets a reference to the given APIServerAccessControl and assigns it to the AccessControl field.
func (o *APIServer) SetAccessControl(v APIServerAccessControl) {
	o.AccessControl = &v
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

// GetBaseUrls returns the BaseUrls field value
func (o *APIServer) GetBaseUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BaseUrls
}

// GetBaseUrlsOk returns a tuple with the BaseUrls field value
// and a boolean to check if the value has been set.
func (o *APIServer) GetBaseUrlsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaseUrls, true
}

// SetBaseUrls sets field value
func (o *APIServer) SetBaseUrls(v []string) {
	o.BaseUrls = v
}

// GetDirectory returns the Directory field value if set, zero value otherwise.
func (o *APIServer) GetDirectory() APIServerDirectory {
	if o == nil || IsNil(o.Directory) {
		var ret APIServerDirectory
		return ret
	}
	return *o.Directory
}

// GetDirectoryOk returns a tuple with the Directory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServer) GetDirectoryOk() (*APIServerDirectory, bool) {
	if o == nil || IsNil(o.Directory) {
		return nil, false
	}
	return o.Directory, true
}

// HasDirectory returns a boolean if a field has been set.
func (o *APIServer) HasDirectory() bool {
	if o != nil && !IsNil(o.Directory) {
		return true
	}

	return false
}

// SetDirectory gets a reference to the given APIServerDirectory and assigns it to the Directory field.
func (o *APIServer) SetDirectory(v APIServerDirectory) {
	o.Directory = &v
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

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *APIServer) GetPolicy() APIServerPolicy {
	if o == nil || IsNil(o.Policy) {
		var ret APIServerPolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServer) GetPolicyOk() (*APIServerPolicy, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *APIServer) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given APIServerPolicy and assigns it to the Policy field.
func (o *APIServer) SetPolicy(v APIServerPolicy) {
	o.Policy = &v
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
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.AccessControl) {
		toSerialize["accessControl"] = o.AccessControl
	}
	toSerialize["authorizationServer"] = o.AuthorizationServer
	toSerialize["baseUrls"] = o.BaseUrls
	if !IsNil(o.Directory) {
		toSerialize["directory"] = o.Directory
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	return toSerialize, nil
}

func (o *APIServer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authorizationServer",
		"baseUrls",
		"name",
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

	varAPIServer := _APIServer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAPIServer)

	if err != nil {
		return err
	}

	*o = APIServer(varAPIServer)

	return err
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


