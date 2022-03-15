# SaleResponseObjectSaleResponseProducts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to **int32** | Número de Producto o Item de la operación | [optional] 
**Name** | Pointer to **string** | Nombre del Producto | [optional] 
**Code** | Pointer to **string** | código de Producto | [optional] 
**Quantity** | Pointer to **float32** | Cantidad | [optional] 
**Unit** | Pointer to **string** | tipo de Unidad de Medida | [optional] 
**UnitAmount** | Pointer to **float32** | Valor unitario | [optional] 
**NetAmount** | Pointer to **float32** | Valor Total Neto sin impuestos | [optional] 
**TaxAmount** | Pointer to **float32** | Valor de los impuestos | [optional] 
**TotalAmount** | Pointer to **float32** | Valor Total | [optional] 

## Methods

### NewSaleResponseObjectSaleResponseProducts

`func NewSaleResponseObjectSaleResponseProducts() *SaleResponseObjectSaleResponseProducts`

NewSaleResponseObjectSaleResponseProducts instantiates a new SaleResponseObjectSaleResponseProducts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleResponseObjectSaleResponseProductsWithDefaults

`func NewSaleResponseObjectSaleResponseProductsWithDefaults() *SaleResponseObjectSaleResponseProducts`

NewSaleResponseObjectSaleResponseProductsWithDefaults instantiates a new SaleResponseObjectSaleResponseProducts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *SaleResponseObjectSaleResponseProducts) GetItem() int32`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *SaleResponseObjectSaleResponseProducts) GetItemOk() (*int32, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *SaleResponseObjectSaleResponseProducts) SetItem(v int32)`

SetItem sets Item field to given value.

### HasItem

`func (o *SaleResponseObjectSaleResponseProducts) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetName

`func (o *SaleResponseObjectSaleResponseProducts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SaleResponseObjectSaleResponseProducts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SaleResponseObjectSaleResponseProducts) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SaleResponseObjectSaleResponseProducts) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *SaleResponseObjectSaleResponseProducts) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SaleResponseObjectSaleResponseProducts) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SaleResponseObjectSaleResponseProducts) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SaleResponseObjectSaleResponseProducts) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetQuantity

`func (o *SaleResponseObjectSaleResponseProducts) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SaleResponseObjectSaleResponseProducts) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SaleResponseObjectSaleResponseProducts) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SaleResponseObjectSaleResponseProducts) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnit

`func (o *SaleResponseObjectSaleResponseProducts) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *SaleResponseObjectSaleResponseProducts) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *SaleResponseObjectSaleResponseProducts) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *SaleResponseObjectSaleResponseProducts) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnitAmount

`func (o *SaleResponseObjectSaleResponseProducts) GetUnitAmount() float32`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *SaleResponseObjectSaleResponseProducts) GetUnitAmountOk() (*float32, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *SaleResponseObjectSaleResponseProducts) SetUnitAmount(v float32)`

SetUnitAmount sets UnitAmount field to given value.

### HasUnitAmount

`func (o *SaleResponseObjectSaleResponseProducts) HasUnitAmount() bool`

HasUnitAmount returns a boolean if a field has been set.

### GetNetAmount

`func (o *SaleResponseObjectSaleResponseProducts) GetNetAmount() float32`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *SaleResponseObjectSaleResponseProducts) GetNetAmountOk() (*float32, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *SaleResponseObjectSaleResponseProducts) SetNetAmount(v float32)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *SaleResponseObjectSaleResponseProducts) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### GetTaxAmount

`func (o *SaleResponseObjectSaleResponseProducts) GetTaxAmount() float32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *SaleResponseObjectSaleResponseProducts) GetTaxAmountOk() (*float32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *SaleResponseObjectSaleResponseProducts) SetTaxAmount(v float32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *SaleResponseObjectSaleResponseProducts) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTotalAmount

`func (o *SaleResponseObjectSaleResponseProducts) GetTotalAmount() float32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *SaleResponseObjectSaleResponseProducts) GetTotalAmountOk() (*float32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *SaleResponseObjectSaleResponseProducts) SetTotalAmount(v float32)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *SaleResponseObjectSaleResponseProducts) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


