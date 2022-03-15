# PaymentMethodsResponseObjectPaymentMethodsResponseType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID interno para cada tipo de tarjeta | [optional] 
**Description** | Pointer to **string** | Descripción del tipo de tarejta | [optional] 
**ForeignIdentifier** | Pointer to **string** | ID externo de este metodo de pago, utilizado por el punto de venta para reconocer al medio de pago en su base de datos | [optional] 

## Methods

### NewPaymentMethodsResponseObjectPaymentMethodsResponseType

`func NewPaymentMethodsResponseObjectPaymentMethodsResponseType() *PaymentMethodsResponseObjectPaymentMethodsResponseType`

NewPaymentMethodsResponseObjectPaymentMethodsResponseType instantiates a new PaymentMethodsResponseObjectPaymentMethodsResponseType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodsResponseObjectPaymentMethodsResponseTypeWithDefaults

`func NewPaymentMethodsResponseObjectPaymentMethodsResponseTypeWithDefaults() *PaymentMethodsResponseObjectPaymentMethodsResponseType`

NewPaymentMethodsResponseObjectPaymentMethodsResponseTypeWithDefaults instantiates a new PaymentMethodsResponseObjectPaymentMethodsResponseType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponseType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponseType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponseType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponseType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponseType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponseType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponseType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponseType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForeignIdentifier

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponseType) GetForeignIdentifier() string`

GetForeignIdentifier returns the ForeignIdentifier field if non-nil, zero value otherwise.

### GetForeignIdentifierOk

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponseType) GetForeignIdentifierOk() (*string, bool)`

GetForeignIdentifierOk returns a tuple with the ForeignIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignIdentifier

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponseType) SetForeignIdentifier(v string)`

SetForeignIdentifier sets ForeignIdentifier field to given value.

### HasForeignIdentifier

`func (o *PaymentMethodsResponseObjectPaymentMethodsResponseType) HasForeignIdentifier() bool`

HasForeignIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


