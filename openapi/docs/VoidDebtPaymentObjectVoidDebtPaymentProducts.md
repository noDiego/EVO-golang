# VoidDebtPaymentObjectVoidDebtPaymentProducts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to **int32** | Número de Producto o Item de la operación | [optional] 
**Name** | Pointer to **string** | Nombre del Producto | [optional] 
**Code** | Pointer to **string** | c��digo de Producto | [optional] 
**Quantity** | Pointer to **float32** | Cantidad | [optional] 
**Unit** | Pointer to **string** | tipo de Unidad de Medida | [optional] 
**UnitAmount** | Pointer to **float32** | Valor unitario | [optional] 
**NetAmount** | Pointer to **float32** | Valor Total Neto sin impuestos | [optional] 
**TaxAmount** | Pointer to **float32** | Valor de los impuestos | [optional] 
**TotalAmount** | Pointer to **float32** | Valor Total | [optional] 

## Methods

### NewVoidDebtPaymentObjectVoidDebtPaymentProducts

`func NewVoidDebtPaymentObjectVoidDebtPaymentProducts() *VoidDebtPaymentObjectVoidDebtPaymentProducts`

NewVoidDebtPaymentObjectVoidDebtPaymentProducts instantiates a new VoidDebtPaymentObjectVoidDebtPaymentProducts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoidDebtPaymentObjectVoidDebtPaymentProductsWithDefaults

`func NewVoidDebtPaymentObjectVoidDebtPaymentProductsWithDefaults() *VoidDebtPaymentObjectVoidDebtPaymentProducts`

NewVoidDebtPaymentObjectVoidDebtPaymentProductsWithDefaults instantiates a new VoidDebtPaymentObjectVoidDebtPaymentProducts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) GetItem() int32`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) GetItemOk() (*int32, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) SetItem(v int32)`

SetItem sets Item field to given value.

### HasItem

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetName

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetQuantity

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnit

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnitAmount

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) GetUnitAmount() float32`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) GetUnitAmountOk() (*float32, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) SetUnitAmount(v float32)`

SetUnitAmount sets UnitAmount field to given value.

### HasUnitAmount

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) HasUnitAmount() bool`

HasUnitAmount returns a boolean if a field has been set.

### GetNetAmount

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) GetNetAmount() float32`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) GetNetAmountOk() (*float32, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) SetNetAmount(v float32)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### GetTaxAmount

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) GetTaxAmount() float32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) GetTaxAmountOk() (*float32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) SetTaxAmount(v float32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTotalAmount

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) GetTotalAmount() float32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) GetTotalAmountOk() (*float32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) SetTotalAmount(v float32)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *VoidDebtPaymentObjectVoidDebtPaymentProducts) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


