/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the NotificationsSettingsPhoneDeliverySettingsCustomRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationsSettingsPhoneDeliverySettingsCustomRequest{}

// NotificationsSettingsPhoneDeliverySettingsCustomRequest struct for NotificationsSettingsPhoneDeliverySettingsCustomRequest
type NotificationsSettingsPhoneDeliverySettingsCustomRequest struct {
	DeliveryMethod EnumNotificationsSettingsPhoneDeliverySettingsCustomDeliveryMethod `json:"deliveryMethod"`
	// The provider's remote gateway or customer gateway URL. For requests using the POST method, use the provider's remote gateway URL. For requests using the GET method, use the provider's remote gateway URL, including the `${to}` and `${message}` mandatory variables, and the optional `${from}` variable, for example: `https://api.transmitsms.com/send-sms.json?to=${to}&from=${from}&message=${message}` 
	Url string `json:"url"`
	// The notification's request body. The body should include the ${to} and ${message} mandatory variables. For some vendors, the optional ${from} variable may also be required. For example: `messageType=ARN&message=${message}&phoneNumber=${to}&sender=${from}` In addition, you can use the following optional variables: `${voice}` - the type of voice configured for notifications `${locale}` - locale `${otp}` - OTP `${user.user.name}` - user's username `${user.name.given}` - user's given name `${user.name.family}` - user's family name You can also use dynamic variables in the body. For more information, see [Dynamic variables](https://apidocs.pingidentity.com/pingone/platform/v1/api/#notifications-templates-dynamic-variables). 
	Body string `json:"body"`
	// A map of the notification's request headers 
	Headers map[string]string `json:"headers"`
	Method EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod `json:"method"`
	PhoneNumberFormat EnumNotificationsSettingsPhoneDeliverySettingsCustomNumberFormat `json:"phoneNumberFormat"`
	// For voice OTP notifications only. An opening tag which is commonly used by custom providers for defining a pause between each number in the OTP number string. Possible value: `<Say>` 
	BeforeTag *string `json:"beforeTag,omitempty"`
	// For voice OTP notifications only. A closing tag which is commonly used by custom providers for defining a pause between each number in the OTP number string. Possible value: `</Say> <Pause length=\"1\"/>` 
	AfterTag *string `json:"afterTag,omitempty"`
}

// NewNotificationsSettingsPhoneDeliverySettingsCustomRequest instantiates a new NotificationsSettingsPhoneDeliverySettingsCustomRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsSettingsPhoneDeliverySettingsCustomRequest(deliveryMethod EnumNotificationsSettingsPhoneDeliverySettingsCustomDeliveryMethod, url string, body string, headers map[string]string, method EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod, phoneNumberFormat EnumNotificationsSettingsPhoneDeliverySettingsCustomNumberFormat) *NotificationsSettingsPhoneDeliverySettingsCustomRequest {
	this := NotificationsSettingsPhoneDeliverySettingsCustomRequest{}
	this.DeliveryMethod = deliveryMethod
	this.Url = url
	this.Body = body
	this.Headers = headers
	this.Method = method
	this.PhoneNumberFormat = phoneNumberFormat
	return &this
}

// NewNotificationsSettingsPhoneDeliverySettingsCustomRequestWithDefaults instantiates a new NotificationsSettingsPhoneDeliverySettingsCustomRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsSettingsPhoneDeliverySettingsCustomRequestWithDefaults() *NotificationsSettingsPhoneDeliverySettingsCustomRequest {
	this := NotificationsSettingsPhoneDeliverySettingsCustomRequest{}
	var phoneNumberFormat EnumNotificationsSettingsPhoneDeliverySettingsCustomNumberFormat = ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSCUSTOMNUMBERFORMAT_FULL
	this.PhoneNumberFormat = phoneNumberFormat
	return &this
}

// GetDeliveryMethod returns the DeliveryMethod field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) GetDeliveryMethod() EnumNotificationsSettingsPhoneDeliverySettingsCustomDeliveryMethod {
	if o == nil {
		var ret EnumNotificationsSettingsPhoneDeliverySettingsCustomDeliveryMethod
		return ret
	}

	return o.DeliveryMethod
}

// GetDeliveryMethodOk returns a tuple with the DeliveryMethod field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) GetDeliveryMethodOk() (*EnumNotificationsSettingsPhoneDeliverySettingsCustomDeliveryMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryMethod, true
}

// SetDeliveryMethod sets field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) SetDeliveryMethod(v EnumNotificationsSettingsPhoneDeliverySettingsCustomDeliveryMethod) {
	o.DeliveryMethod = v
}

// GetUrl returns the Url field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) SetUrl(v string) {
	o.Url = v
}

// GetBody returns the Body field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) SetBody(v string) {
	o.Body = v
}

// GetHeaders returns the Headers field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) GetHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) GetHeadersOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headers, true
}

// SetHeaders sets field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetMethod returns the Method field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) GetMethod() EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod {
	if o == nil {
		var ret EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) GetMethodOk() (*EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) SetMethod(v EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod) {
	o.Method = v
}

// GetPhoneNumberFormat returns the PhoneNumberFormat field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) GetPhoneNumberFormat() EnumNotificationsSettingsPhoneDeliverySettingsCustomNumberFormat {
	if o == nil {
		var ret EnumNotificationsSettingsPhoneDeliverySettingsCustomNumberFormat
		return ret
	}

	return o.PhoneNumberFormat
}

// GetPhoneNumberFormatOk returns a tuple with the PhoneNumberFormat field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) GetPhoneNumberFormatOk() (*EnumNotificationsSettingsPhoneDeliverySettingsCustomNumberFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumberFormat, true
}

// SetPhoneNumberFormat sets field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) SetPhoneNumberFormat(v EnumNotificationsSettingsPhoneDeliverySettingsCustomNumberFormat) {
	o.PhoneNumberFormat = v
}

// GetBeforeTag returns the BeforeTag field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) GetBeforeTag() string {
	if o == nil || IsNil(o.BeforeTag) {
		var ret string
		return ret
	}
	return *o.BeforeTag
}

// GetBeforeTagOk returns a tuple with the BeforeTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) GetBeforeTagOk() (*string, bool) {
	if o == nil || IsNil(o.BeforeTag) {
		return nil, false
	}
	return o.BeforeTag, true
}

// HasBeforeTag returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) HasBeforeTag() bool {
	if o != nil && !IsNil(o.BeforeTag) {
		return true
	}

	return false
}

// SetBeforeTag gets a reference to the given string and assigns it to the BeforeTag field.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) SetBeforeTag(v string) {
	o.BeforeTag = &v
}

// GetAfterTag returns the AfterTag field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) GetAfterTag() string {
	if o == nil || IsNil(o.AfterTag) {
		var ret string
		return ret
	}
	return *o.AfterTag
}

// GetAfterTagOk returns a tuple with the AfterTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) GetAfterTagOk() (*string, bool) {
	if o == nil || IsNil(o.AfterTag) {
		return nil, false
	}
	return o.AfterTag, true
}

// HasAfterTag returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) HasAfterTag() bool {
	if o != nil && !IsNil(o.AfterTag) {
		return true
	}

	return false
}

// SetAfterTag gets a reference to the given string and assigns it to the AfterTag field.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomRequest) SetAfterTag(v string) {
	o.AfterTag = &v
}

func (o NotificationsSettingsPhoneDeliverySettingsCustomRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationsSettingsPhoneDeliverySettingsCustomRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deliveryMethod"] = o.DeliveryMethod
	toSerialize["url"] = o.Url
	toSerialize["body"] = o.Body
	toSerialize["headers"] = o.Headers
	toSerialize["method"] = o.Method
	toSerialize["phoneNumberFormat"] = o.PhoneNumberFormat
	if !IsNil(o.BeforeTag) {
		toSerialize["beforeTag"] = o.BeforeTag
	}
	if !IsNil(o.AfterTag) {
		toSerialize["afterTag"] = o.AfterTag
	}
	return toSerialize, nil
}

type NullableNotificationsSettingsPhoneDeliverySettingsCustomRequest struct {
	value *NotificationsSettingsPhoneDeliverySettingsCustomRequest
	isSet bool
}

func (v NullableNotificationsSettingsPhoneDeliverySettingsCustomRequest) Get() *NotificationsSettingsPhoneDeliverySettingsCustomRequest {
	return v.value
}

func (v *NullableNotificationsSettingsPhoneDeliverySettingsCustomRequest) Set(val *NotificationsSettingsPhoneDeliverySettingsCustomRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsSettingsPhoneDeliverySettingsCustomRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsSettingsPhoneDeliverySettingsCustomRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsSettingsPhoneDeliverySettingsCustomRequest(val *NotificationsSettingsPhoneDeliverySettingsCustomRequest) *NullableNotificationsSettingsPhoneDeliverySettingsCustomRequest {
	return &NullableNotificationsSettingsPhoneDeliverySettingsCustomRequest{value: val, isSet: true}
}

func (v NullableNotificationsSettingsPhoneDeliverySettingsCustomRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsSettingsPhoneDeliverySettingsCustomRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


