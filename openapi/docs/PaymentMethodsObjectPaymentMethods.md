# PaymentMethodsObjectPaymentMethods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyIdentification** | Pointer to **string** | ID que identifica la companía desde donde proviene la petición. | [optional] 
**SystemIdentification** | Pointer to **string** | ID que identifica el sistema desde donde proviene la petición. | [optional] 
**BranchIdentification** | Pointer to **string** | ID que identifica la sucursal desde donde proviene la petición. Esta sucursal pertenece a una determinada companía. | [optional] 
**POSIdentification** | Pointer to **string** | ID que identifica la caja o punto de venta desde donde proviene la petición. Este punto de venta pertenece a una determinada sucursal y companía. | [optional] 
**ServiceVersion** | Pointer to **string** | Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible. | [optional] 
**Sequence** | Pointer to **string** | Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma. | [optional] 
**Security** | Pointer to [**[]SaleObjectSaleSecurity**](SaleObjectSaleSecurity.md) | Datos asociados a la seguridad de la transacción o de elementos sensibles. | [optional] 

## Methods

### NewPaymentMethodsObjectPaymentMethods

`func NewPaymentMethodsObjectPaymentMethods() *PaymentMethodsObjectPaymentMethods`

NewPaymentMethodsObjectPaymentMethods instantiates a new PaymentMethodsObjectPaymentMethods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodsObjectPaymentMethodsWithDefaults

`func NewPaymentMethodsObjectPaymentMethodsWithDefaults() *PaymentMethodsObjectPaymentMethods`

NewPaymentMethodsObjectPaymentMethodsWithDefaults instantiates a new PaymentMethodsObjectPaymentMethods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyIdentification

`func (o *PaymentMethodsObjectPaymentMethods) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *PaymentMethodsObjectPaymentMethods) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *PaymentMethodsObjectPaymentMethods) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.

### HasCompanyIdentification

`func (o *PaymentMethodsObjectPaymentMethods) HasCompanyIdentification() bool`

HasCompanyIdentification returns a boolean if a field has been set.

### GetSystemIdentification

`func (o *PaymentMethodsObjectPaymentMethods) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *PaymentMethodsObjectPaymentMethods) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *PaymentMethodsObjectPaymentMethods) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.

### HasSystemIdentification

`func (o *PaymentMethodsObjectPaymentMethods) HasSystemIdentification() bool`

HasSystemIdentification returns a boolean if a field has been set.

### GetBranchIdentification

`func (o *PaymentMethodsObjectPaymentMethods) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *PaymentMethodsObjectPaymentMethods) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *PaymentMethodsObjectPaymentMethods) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *PaymentMethodsObjectPaymentMethods) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *PaymentMethodsObjectPaymentMethods) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *PaymentMethodsObjectPaymentMethods) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *PaymentMethodsObjectPaymentMethods) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *PaymentMethodsObjectPaymentMethods) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetServiceVersion

`func (o *PaymentMethodsObjectPaymentMethods) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *PaymentMethodsObjectPaymentMethods) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *PaymentMethodsObjectPaymentMethods) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *PaymentMethodsObjectPaymentMethods) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *PaymentMethodsObjectPaymentMethods) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *PaymentMethodsObjectPaymentMethods) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *PaymentMethodsObjectPaymentMethods) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *PaymentMethodsObjectPaymentMethods) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *PaymentMethodsObjectPaymentMethods) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *PaymentMethodsObjectPaymentMethods) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *PaymentMethodsObjectPaymentMethods) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *PaymentMethodsObjectPaymentMethods) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


