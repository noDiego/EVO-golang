# OrderFinalObjectOrderFinal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemIdentification** | Pointer to **string** | ID que identifica el sistema desde donde proviene la petición. | [optional] 
**CompanyIdentification** | Pointer to **string** | ID que identifica la companía desde donde proviene la petición. | [optional] 
**BranchIdentification** | Pointer to **string** | ID que identifica la sucursal desde donde proviene la petición. Esta sucursal pertenece a una determinada companía. | [optional] 
**POSIdentification** | Pointer to **string** | ID que identifica la caja o punto de venta desde donde proviene la petición. Este punto de venta pertenece a una determinada sucursal y companía. | [optional] 
**ServiceVersion** | Pointer to **string** | Versión del Servicio de la Plataforma con la cual se quiere transaccionar, en caso de no ser especificado será atendido por la última versión del servicio disponible. | [optional] 
**Sequence** | Pointer to **string** | Retornado en todas las respuesta que el POS/PINPAD debe enviar en el próximo requerimiento. En caso de que el POS no lo envíe, envíe vacío o con un valor que no corresponde se produce “La Ruptura de Secuencia” y la plataforma si la última transacción que realizó el POS no esta confirmada y esta Aprobada genera entonces una reversa de la misma. | [optional] 
**Security** | Pointer to [**[]SaleObjectSaleSecurity**](SaleObjectSaleSecurity.md) | Datos asociados a la seguridad de la transacción o de elementos sensibles. | [optional] 
**MerchantIdentification** | Pointer to **string** | . | [optional] 
**TransactionIdentification** | Pointer to **string** | . | [optional] 
**TransactionDescription** | Pointer to **string** | Descripción del tipo de operación que se realizará | [optional] 
**PaymentToken** | Pointer to **string** | . | [optional] 
**InitialIdentification** | Pointer to **string** | . | [optional] 

## Methods

### NewOrderFinalObjectOrderFinal

`func NewOrderFinalObjectOrderFinal() *OrderFinalObjectOrderFinal`

NewOrderFinalObjectOrderFinal instantiates a new OrderFinalObjectOrderFinal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderFinalObjectOrderFinalWithDefaults

`func NewOrderFinalObjectOrderFinalWithDefaults() *OrderFinalObjectOrderFinal`

NewOrderFinalObjectOrderFinalWithDefaults instantiates a new OrderFinalObjectOrderFinal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemIdentification

`func (o *OrderFinalObjectOrderFinal) GetSystemIdentification() string`

GetSystemIdentification returns the SystemIdentification field if non-nil, zero value otherwise.

### GetSystemIdentificationOk

`func (o *OrderFinalObjectOrderFinal) GetSystemIdentificationOk() (*string, bool)`

GetSystemIdentificationOk returns a tuple with the SystemIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentification

`func (o *OrderFinalObjectOrderFinal) SetSystemIdentification(v string)`

SetSystemIdentification sets SystemIdentification field to given value.

### HasSystemIdentification

`func (o *OrderFinalObjectOrderFinal) HasSystemIdentification() bool`

HasSystemIdentification returns a boolean if a field has been set.

### GetCompanyIdentification

`func (o *OrderFinalObjectOrderFinal) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *OrderFinalObjectOrderFinal) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *OrderFinalObjectOrderFinal) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.

### HasCompanyIdentification

`func (o *OrderFinalObjectOrderFinal) HasCompanyIdentification() bool`

HasCompanyIdentification returns a boolean if a field has been set.

### GetBranchIdentification

`func (o *OrderFinalObjectOrderFinal) GetBranchIdentification() string`

GetBranchIdentification returns the BranchIdentification field if non-nil, zero value otherwise.

### GetBranchIdentificationOk

`func (o *OrderFinalObjectOrderFinal) GetBranchIdentificationOk() (*string, bool)`

GetBranchIdentificationOk returns a tuple with the BranchIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIdentification

`func (o *OrderFinalObjectOrderFinal) SetBranchIdentification(v string)`

SetBranchIdentification sets BranchIdentification field to given value.

### HasBranchIdentification

`func (o *OrderFinalObjectOrderFinal) HasBranchIdentification() bool`

HasBranchIdentification returns a boolean if a field has been set.

### GetPOSIdentification

`func (o *OrderFinalObjectOrderFinal) GetPOSIdentification() string`

GetPOSIdentification returns the POSIdentification field if non-nil, zero value otherwise.

### GetPOSIdentificationOk

`func (o *OrderFinalObjectOrderFinal) GetPOSIdentificationOk() (*string, bool)`

GetPOSIdentificationOk returns a tuple with the POSIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSIdentification

`func (o *OrderFinalObjectOrderFinal) SetPOSIdentification(v string)`

SetPOSIdentification sets POSIdentification field to given value.

### HasPOSIdentification

`func (o *OrderFinalObjectOrderFinal) HasPOSIdentification() bool`

HasPOSIdentification returns a boolean if a field has been set.

### GetServiceVersion

`func (o *OrderFinalObjectOrderFinal) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *OrderFinalObjectOrderFinal) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *OrderFinalObjectOrderFinal) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *OrderFinalObjectOrderFinal) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSequence

`func (o *OrderFinalObjectOrderFinal) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *OrderFinalObjectOrderFinal) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *OrderFinalObjectOrderFinal) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *OrderFinalObjectOrderFinal) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSecurity

`func (o *OrderFinalObjectOrderFinal) GetSecurity() []SaleObjectSaleSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *OrderFinalObjectOrderFinal) GetSecurityOk() (*[]SaleObjectSaleSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *OrderFinalObjectOrderFinal) SetSecurity(v []SaleObjectSaleSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *OrderFinalObjectOrderFinal) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetMerchantIdentification

`func (o *OrderFinalObjectOrderFinal) GetMerchantIdentification() string`

GetMerchantIdentification returns the MerchantIdentification field if non-nil, zero value otherwise.

### GetMerchantIdentificationOk

`func (o *OrderFinalObjectOrderFinal) GetMerchantIdentificationOk() (*string, bool)`

GetMerchantIdentificationOk returns a tuple with the MerchantIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantIdentification

`func (o *OrderFinalObjectOrderFinal) SetMerchantIdentification(v string)`

SetMerchantIdentification sets MerchantIdentification field to given value.

### HasMerchantIdentification

`func (o *OrderFinalObjectOrderFinal) HasMerchantIdentification() bool`

HasMerchantIdentification returns a boolean if a field has been set.

### GetTransactionIdentification

`func (o *OrderFinalObjectOrderFinal) GetTransactionIdentification() string`

GetTransactionIdentification returns the TransactionIdentification field if non-nil, zero value otherwise.

### GetTransactionIdentificationOk

`func (o *OrderFinalObjectOrderFinal) GetTransactionIdentificationOk() (*string, bool)`

GetTransactionIdentificationOk returns a tuple with the TransactionIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIdentification

`func (o *OrderFinalObjectOrderFinal) SetTransactionIdentification(v string)`

SetTransactionIdentification sets TransactionIdentification field to given value.

### HasTransactionIdentification

`func (o *OrderFinalObjectOrderFinal) HasTransactionIdentification() bool`

HasTransactionIdentification returns a boolean if a field has been set.

### GetTransactionDescription

`func (o *OrderFinalObjectOrderFinal) GetTransactionDescription() string`

GetTransactionDescription returns the TransactionDescription field if non-nil, zero value otherwise.

### GetTransactionDescriptionOk

`func (o *OrderFinalObjectOrderFinal) GetTransactionDescriptionOk() (*string, bool)`

GetTransactionDescriptionOk returns a tuple with the TransactionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDescription

`func (o *OrderFinalObjectOrderFinal) SetTransactionDescription(v string)`

SetTransactionDescription sets TransactionDescription field to given value.

### HasTransactionDescription

`func (o *OrderFinalObjectOrderFinal) HasTransactionDescription() bool`

HasTransactionDescription returns a boolean if a field has been set.

### GetPaymentToken

`func (o *OrderFinalObjectOrderFinal) GetPaymentToken() string`

GetPaymentToken returns the PaymentToken field if non-nil, zero value otherwise.

### GetPaymentTokenOk

`func (o *OrderFinalObjectOrderFinal) GetPaymentTokenOk() (*string, bool)`

GetPaymentTokenOk returns a tuple with the PaymentToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentToken

`func (o *OrderFinalObjectOrderFinal) SetPaymentToken(v string)`

SetPaymentToken sets PaymentToken field to given value.

### HasPaymentToken

`func (o *OrderFinalObjectOrderFinal) HasPaymentToken() bool`

HasPaymentToken returns a boolean if a field has been set.

### GetInitialIdentification

`func (o *OrderFinalObjectOrderFinal) GetInitialIdentification() string`

GetInitialIdentification returns the InitialIdentification field if non-nil, zero value otherwise.

### GetInitialIdentificationOk

`func (o *OrderFinalObjectOrderFinal) GetInitialIdentificationOk() (*string, bool)`

GetInitialIdentificationOk returns a tuple with the InitialIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialIdentification

`func (o *OrderFinalObjectOrderFinal) SetInitialIdentification(v string)`

SetInitialIdentification sets InitialIdentification field to given value.

### HasInitialIdentification

`func (o *OrderFinalObjectOrderFinal) HasInitialIdentification() bool`

HasInitialIdentification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


