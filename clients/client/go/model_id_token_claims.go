/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.60
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// IDTokenClaims IDTokenClaims represent the claims used in open id connect requests
type IDTokenClaims struct {
	Acr *string `json:"acr,omitempty"`
	Amr []string `json:"amr,omitempty"`
	AtHash *string `json:"at_hash,omitempty"`
	Aud []string `json:"aud,omitempty"`
	AuthTime *time.Time `json:"auth_time,omitempty"`
	CHash *string `json:"c_hash,omitempty"`
	Exp *time.Time `json:"exp,omitempty"`
	Ext map[string]interface{} `json:"ext,omitempty"`
	Iat *time.Time `json:"iat,omitempty"`
	Iss *string `json:"iss,omitempty"`
	Jti *string `json:"jti,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	Rat *time.Time `json:"rat,omitempty"`
	Sub *string `json:"sub,omitempty"`
}

// NewIDTokenClaims instantiates a new IDTokenClaims object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIDTokenClaims() *IDTokenClaims {
	this := IDTokenClaims{}
	return &this
}

// NewIDTokenClaimsWithDefaults instantiates a new IDTokenClaims object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIDTokenClaimsWithDefaults() *IDTokenClaims {
	this := IDTokenClaims{}
	return &this
}

// GetAcr returns the Acr field value if set, zero value otherwise.
func (o *IDTokenClaims) GetAcr() string {
	if o == nil || o.Acr == nil {
		var ret string
		return ret
	}
	return *o.Acr
}

// GetAcrOk returns a tuple with the Acr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDTokenClaims) GetAcrOk() (*string, bool) {
	if o == nil || o.Acr == nil {
		return nil, false
	}
	return o.Acr, true
}

// HasAcr returns a boolean if a field has been set.
func (o *IDTokenClaims) HasAcr() bool {
	if o != nil && o.Acr != nil {
		return true
	}

	return false
}

// SetAcr gets a reference to the given string and assigns it to the Acr field.
func (o *IDTokenClaims) SetAcr(v string) {
	o.Acr = &v
}

// GetAmr returns the Amr field value if set, zero value otherwise.
func (o *IDTokenClaims) GetAmr() []string {
	if o == nil || o.Amr == nil {
		var ret []string
		return ret
	}
	return o.Amr
}

// GetAmrOk returns a tuple with the Amr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDTokenClaims) GetAmrOk() ([]string, bool) {
	if o == nil || o.Amr == nil {
		return nil, false
	}
	return o.Amr, true
}

// HasAmr returns a boolean if a field has been set.
func (o *IDTokenClaims) HasAmr() bool {
	if o != nil && o.Amr != nil {
		return true
	}

	return false
}

// SetAmr gets a reference to the given []string and assigns it to the Amr field.
func (o *IDTokenClaims) SetAmr(v []string) {
	o.Amr = v
}

// GetAtHash returns the AtHash field value if set, zero value otherwise.
func (o *IDTokenClaims) GetAtHash() string {
	if o == nil || o.AtHash == nil {
		var ret string
		return ret
	}
	return *o.AtHash
}

// GetAtHashOk returns a tuple with the AtHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDTokenClaims) GetAtHashOk() (*string, bool) {
	if o == nil || o.AtHash == nil {
		return nil, false
	}
	return o.AtHash, true
}

// HasAtHash returns a boolean if a field has been set.
func (o *IDTokenClaims) HasAtHash() bool {
	if o != nil && o.AtHash != nil {
		return true
	}

	return false
}

// SetAtHash gets a reference to the given string and assigns it to the AtHash field.
func (o *IDTokenClaims) SetAtHash(v string) {
	o.AtHash = &v
}

// GetAud returns the Aud field value if set, zero value otherwise.
func (o *IDTokenClaims) GetAud() []string {
	if o == nil || o.Aud == nil {
		var ret []string
		return ret
	}
	return o.Aud
}

// GetAudOk returns a tuple with the Aud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDTokenClaims) GetAudOk() ([]string, bool) {
	if o == nil || o.Aud == nil {
		return nil, false
	}
	return o.Aud, true
}

// HasAud returns a boolean if a field has been set.
func (o *IDTokenClaims) HasAud() bool {
	if o != nil && o.Aud != nil {
		return true
	}

	return false
}

// SetAud gets a reference to the given []string and assigns it to the Aud field.
func (o *IDTokenClaims) SetAud(v []string) {
	o.Aud = v
}

// GetAuthTime returns the AuthTime field value if set, zero value otherwise.
func (o *IDTokenClaims) GetAuthTime() time.Time {
	if o == nil || o.AuthTime == nil {
		var ret time.Time
		return ret
	}
	return *o.AuthTime
}

// GetAuthTimeOk returns a tuple with the AuthTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDTokenClaims) GetAuthTimeOk() (*time.Time, bool) {
	if o == nil || o.AuthTime == nil {
		return nil, false
	}
	return o.AuthTime, true
}

// HasAuthTime returns a boolean if a field has been set.
func (o *IDTokenClaims) HasAuthTime() bool {
	if o != nil && o.AuthTime != nil {
		return true
	}

	return false
}

// SetAuthTime gets a reference to the given time.Time and assigns it to the AuthTime field.
func (o *IDTokenClaims) SetAuthTime(v time.Time) {
	o.AuthTime = &v
}

// GetCHash returns the CHash field value if set, zero value otherwise.
func (o *IDTokenClaims) GetCHash() string {
	if o == nil || o.CHash == nil {
		var ret string
		return ret
	}
	return *o.CHash
}

// GetCHashOk returns a tuple with the CHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDTokenClaims) GetCHashOk() (*string, bool) {
	if o == nil || o.CHash == nil {
		return nil, false
	}
	return o.CHash, true
}

// HasCHash returns a boolean if a field has been set.
func (o *IDTokenClaims) HasCHash() bool {
	if o != nil && o.CHash != nil {
		return true
	}

	return false
}

// SetCHash gets a reference to the given string and assigns it to the CHash field.
func (o *IDTokenClaims) SetCHash(v string) {
	o.CHash = &v
}

// GetExp returns the Exp field value if set, zero value otherwise.
func (o *IDTokenClaims) GetExp() time.Time {
	if o == nil || o.Exp == nil {
		var ret time.Time
		return ret
	}
	return *o.Exp
}

// GetExpOk returns a tuple with the Exp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDTokenClaims) GetExpOk() (*time.Time, bool) {
	if o == nil || o.Exp == nil {
		return nil, false
	}
	return o.Exp, true
}

// HasExp returns a boolean if a field has been set.
func (o *IDTokenClaims) HasExp() bool {
	if o != nil && o.Exp != nil {
		return true
	}

	return false
}

// SetExp gets a reference to the given time.Time and assigns it to the Exp field.
func (o *IDTokenClaims) SetExp(v time.Time) {
	o.Exp = &v
}

// GetExt returns the Ext field value if set, zero value otherwise.
func (o *IDTokenClaims) GetExt() map[string]interface{} {
	if o == nil || o.Ext == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Ext
}

// GetExtOk returns a tuple with the Ext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDTokenClaims) GetExtOk() (map[string]interface{}, bool) {
	if o == nil || o.Ext == nil {
		return nil, false
	}
	return o.Ext, true
}

// HasExt returns a boolean if a field has been set.
func (o *IDTokenClaims) HasExt() bool {
	if o != nil && o.Ext != nil {
		return true
	}

	return false
}

// SetExt gets a reference to the given map[string]interface{} and assigns it to the Ext field.
func (o *IDTokenClaims) SetExt(v map[string]interface{}) {
	o.Ext = v
}

// GetIat returns the Iat field value if set, zero value otherwise.
func (o *IDTokenClaims) GetIat() time.Time {
	if o == nil || o.Iat == nil {
		var ret time.Time
		return ret
	}
	return *o.Iat
}

// GetIatOk returns a tuple with the Iat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDTokenClaims) GetIatOk() (*time.Time, bool) {
	if o == nil || o.Iat == nil {
		return nil, false
	}
	return o.Iat, true
}

// HasIat returns a boolean if a field has been set.
func (o *IDTokenClaims) HasIat() bool {
	if o != nil && o.Iat != nil {
		return true
	}

	return false
}

// SetIat gets a reference to the given time.Time and assigns it to the Iat field.
func (o *IDTokenClaims) SetIat(v time.Time) {
	o.Iat = &v
}

// GetIss returns the Iss field value if set, zero value otherwise.
func (o *IDTokenClaims) GetIss() string {
	if o == nil || o.Iss == nil {
		var ret string
		return ret
	}
	return *o.Iss
}

// GetIssOk returns a tuple with the Iss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDTokenClaims) GetIssOk() (*string, bool) {
	if o == nil || o.Iss == nil {
		return nil, false
	}
	return o.Iss, true
}

// HasIss returns a boolean if a field has been set.
func (o *IDTokenClaims) HasIss() bool {
	if o != nil && o.Iss != nil {
		return true
	}

	return false
}

// SetIss gets a reference to the given string and assigns it to the Iss field.
func (o *IDTokenClaims) SetIss(v string) {
	o.Iss = &v
}

// GetJti returns the Jti field value if set, zero value otherwise.
func (o *IDTokenClaims) GetJti() string {
	if o == nil || o.Jti == nil {
		var ret string
		return ret
	}
	return *o.Jti
}

// GetJtiOk returns a tuple with the Jti field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDTokenClaims) GetJtiOk() (*string, bool) {
	if o == nil || o.Jti == nil {
		return nil, false
	}
	return o.Jti, true
}

// HasJti returns a boolean if a field has been set.
func (o *IDTokenClaims) HasJti() bool {
	if o != nil && o.Jti != nil {
		return true
	}

	return false
}

// SetJti gets a reference to the given string and assigns it to the Jti field.
func (o *IDTokenClaims) SetJti(v string) {
	o.Jti = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *IDTokenClaims) GetNonce() string {
	if o == nil || o.Nonce == nil {
		var ret string
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDTokenClaims) GetNonceOk() (*string, bool) {
	if o == nil || o.Nonce == nil {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *IDTokenClaims) HasNonce() bool {
	if o != nil && o.Nonce != nil {
		return true
	}

	return false
}

// SetNonce gets a reference to the given string and assigns it to the Nonce field.
func (o *IDTokenClaims) SetNonce(v string) {
	o.Nonce = &v
}

// GetRat returns the Rat field value if set, zero value otherwise.
func (o *IDTokenClaims) GetRat() time.Time {
	if o == nil || o.Rat == nil {
		var ret time.Time
		return ret
	}
	return *o.Rat
}

// GetRatOk returns a tuple with the Rat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDTokenClaims) GetRatOk() (*time.Time, bool) {
	if o == nil || o.Rat == nil {
		return nil, false
	}
	return o.Rat, true
}

// HasRat returns a boolean if a field has been set.
func (o *IDTokenClaims) HasRat() bool {
	if o != nil && o.Rat != nil {
		return true
	}

	return false
}

// SetRat gets a reference to the given time.Time and assigns it to the Rat field.
func (o *IDTokenClaims) SetRat(v time.Time) {
	o.Rat = &v
}

// GetSub returns the Sub field value if set, zero value otherwise.
func (o *IDTokenClaims) GetSub() string {
	if o == nil || o.Sub == nil {
		var ret string
		return ret
	}
	return *o.Sub
}

// GetSubOk returns a tuple with the Sub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDTokenClaims) GetSubOk() (*string, bool) {
	if o == nil || o.Sub == nil {
		return nil, false
	}
	return o.Sub, true
}

// HasSub returns a boolean if a field has been set.
func (o *IDTokenClaims) HasSub() bool {
	if o != nil && o.Sub != nil {
		return true
	}

	return false
}

// SetSub gets a reference to the given string and assigns it to the Sub field.
func (o *IDTokenClaims) SetSub(v string) {
	o.Sub = &v
}

func (o IDTokenClaims) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Acr != nil {
		toSerialize["acr"] = o.Acr
	}
	if o.Amr != nil {
		toSerialize["amr"] = o.Amr
	}
	if o.AtHash != nil {
		toSerialize["at_hash"] = o.AtHash
	}
	if o.Aud != nil {
		toSerialize["aud"] = o.Aud
	}
	if o.AuthTime != nil {
		toSerialize["auth_time"] = o.AuthTime
	}
	if o.CHash != nil {
		toSerialize["c_hash"] = o.CHash
	}
	if o.Exp != nil {
		toSerialize["exp"] = o.Exp
	}
	if o.Ext != nil {
		toSerialize["ext"] = o.Ext
	}
	if o.Iat != nil {
		toSerialize["iat"] = o.Iat
	}
	if o.Iss != nil {
		toSerialize["iss"] = o.Iss
	}
	if o.Jti != nil {
		toSerialize["jti"] = o.Jti
	}
	if o.Nonce != nil {
		toSerialize["nonce"] = o.Nonce
	}
	if o.Rat != nil {
		toSerialize["rat"] = o.Rat
	}
	if o.Sub != nil {
		toSerialize["sub"] = o.Sub
	}
	return json.Marshal(toSerialize)
}

type NullableIDTokenClaims struct {
	value *IDTokenClaims
	isSet bool
}

func (v NullableIDTokenClaims) Get() *IDTokenClaims {
	return v.value
}

func (v *NullableIDTokenClaims) Set(val *IDTokenClaims) {
	v.value = val
	v.isSet = true
}

func (v NullableIDTokenClaims) IsSet() bool {
	return v.isSet
}

func (v *NullableIDTokenClaims) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIDTokenClaims(val *IDTokenClaims) *NullableIDTokenClaims {
	return &NullableIDTokenClaims{value: val, isSet: true}
}

func (v NullableIDTokenClaims) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIDTokenClaims) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


