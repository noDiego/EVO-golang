# SaleResponseObjectSaleResponsePaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID interno en la base de datos para este método de pago. | [optional] 
**ForeignIdentifier** | Pointer to **string** | ID externo de este método de pago, utilizado por el punto de venta para reconocer el medio de pago en su base de datos. | [optional] 
**Description** | Pointer to **string** | Describe al metodo de pago. | [optional] 
**Type** | Pointer to [**SaleResponseObjectSaleResponsePaymentMethodType**](SaleResponseObjectSaleResponsePaymentMethodType.md) |  | [optional] 
**DebitAccount** | Pointer to [**SaleResponseObjectSaleResponsePaymentMethodDebitAccount**](SaleResponseObjectSaleResponsePaymentMethodDebitAccount.md) |  | [optional] 
**Category** | Pointer to [**SaleResponseObjectSaleResponsePaymentMethodCategory**](SaleResponseObjectSaleResponsePaymentMethodCategory.md) |  | [optional] 
**Issuers** | Pointer to [**[]SaleResponseObjectSaleResponsePaymentMethodIssuers**](SaleResponseObjectSaleResponsePaymentMethodIssuers.md) | Cada elemento de este array representa a un emisor disponible para alguno de los planes informados bank. | [optional] 

## Methods

### NewSaleResponseObjectSaleResponsePaymentMethod

`func NewSaleResponseObjectSaleResponsePaymentMethod() *SaleResponseObjectSaleResponsePaymentMethod`

NewSaleResponseObjectSaleResponsePaymentMethod instantiates a new SaleResponseObjectSaleResponsePaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleResponseObjectSaleResponsePaymentMethodWithDefaults

`func NewSaleResponseObjectSaleResponsePaymentMethodWithDefaults() *SaleResponseObjectSaleResponsePaymentMethod`

NewSaleResponseObjectSaleResponsePaymentMethodWithDefaults instantiates a new SaleResponseObjectSaleResponsePaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SaleResponseObjectSaleResponsePaymentMethod) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SaleResponseObjectSaleResponsePaymentMethod) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SaleResponseObjectSaleResponsePaymentMethod) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SaleResponseObjectSaleResponsePaymentMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetForeignIdentifier

`func (o *SaleResponseObjectSaleResponsePaymentMethod) GetForeignIdentifier() string`

GetForeignIdentifier returns the ForeignIdentifier field if non-nil, zero value otherwise.

### GetForeignIdentifierOk

`func (o *SaleResponseObjectSaleResponsePaymentMethod) GetForeignIdentifierOk() (*string, bool)`

GetForeignIdentifierOk returns a tuple with the ForeignIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignIdentifier

`func (o *SaleResponseObjectSaleResponsePaymentMethod) SetForeignIdentifier(v string)`

SetForeignIdentifier sets ForeignIdentifier field to given value.

### HasForeignIdentifier

`func (o *SaleResponseObjectSaleResponsePaymentMethod) HasForeignIdentifier() bool`

HasForeignIdentifier returns a boolean if a field has been set.

### GetDescription

`func (o *SaleResponseObjectSaleResponsePaymentMethod) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SaleResponseObjectSaleResponsePaymentMethod) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SaleResponseObjectSaleResponsePaymentMethod) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SaleResponseObjectSaleResponsePaymentMethod) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *SaleResponseObjectSaleResponsePaymentMethod) GetType() SaleResponseObjectSaleResponsePaymentMethodType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SaleResponseObjectSaleResponsePaymentMethod) GetTypeOk() (*SaleResponseObjectSaleResponsePaymentMethodType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SaleResponseObjectSaleResponsePaymentMethod) SetType(v SaleResponseObjectSaleResponsePaymentMethodType)`

SetType sets Type field to given value.

### HasType

`func (o *SaleResponseObjectSaleResponsePaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDebitAccount

`func (o *SaleResponseObjectSaleResponsePaymentMethod) GetDebitAccount() SaleResponseObjectSaleResponsePaymentMethodDebitAccount`

GetDebitAccount returns the DebitAccount field if non-nil, zero value otherwise.

### GetDebitAccountOk

`func (o *SaleResponseObjectSaleResponsePaymentMethod) GetDebitAccountOk() (*SaleResponseObjectSaleResponsePaymentMethodDebitAccount, bool)`

GetDebitAccountOk returns a tuple with the DebitAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccount

`func (o *SaleResponseObjectSaleResponsePaymentMethod) SetDebitAccount(v SaleResponseObjectSaleResponsePaymentMethodDebitAccount)`

SetDebitAccount sets DebitAccount field to given value.

### HasDebitAccount

`func (o *SaleResponseObjectSaleResponsePaymentMethod) HasDebitAccount() bool`

HasDebitAccount returns a boolean if a field has been set.

### GetCategory

`func (o *SaleResponseObjectSaleResponsePaymentMethod) GetCategory() SaleResponseObjectSaleResponsePaymentMethodCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SaleResponseObjectSaleResponsePaymentMethod) GetCategoryOk() (*SaleResponseObjectSaleResponsePaymentMethodCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SaleResponseObjectSaleResponsePaymentMethod) SetCategory(v SaleResponseObjectSaleResponsePaymentMethodCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SaleResponseObjectSaleResponsePaymentMethod) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetIssuers

`func (o *SaleResponseObjectSaleResponsePaymentMethod) GetIssuers() []SaleResponseObjectSaleResponsePaymentMethodIssuers`

GetIssuers returns the Issuers field if non-nil, zero value otherwise.

### GetIssuersOk

`func (o *SaleResponseObjectSaleResponsePaymentMethod) GetIssuersOk() (*[]SaleResponseObjectSaleResponsePaymentMethodIssuers, bool)`

GetIssuersOk returns a tuple with the Issuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuers

`func (o *SaleResponseObjectSaleResponsePaymentMethod) SetIssuers(v []SaleResponseObjectSaleResponsePaymentMethodIssuers)`

SetIssuers sets Issuers field to given value.

### HasIssuers

`func (o *SaleResponseObjectSaleResponsePaymentMethod) HasIssuers() bool`

HasIssuers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


