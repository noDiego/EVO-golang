# QueryCompaniesResponseObjectQueryCompaniesResponseCompanies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identificacion de la Compania | [optional] 
**Name** | Pointer to **string** | Nombre de la compania | [optional] 
**Description** | Pointer to **string** | Descripcion de la compania | [optional] 
**LineOfBusinessIdentification** | Pointer to **string** | Rubro de Pertenencias de la Compania | [optional] 
**Order** | Pointer to **int32** | Orden | [optional] 
**AmountType** | Pointer to **int32** | Tipo de Importe que se puede usar para pagar esta empresa | [optional] 
**EnableRecurringPayment** | Pointer to **bool** | Indicador de Habilitacion de Pagos Recurrentes | [optional] 
**Provider** | Pointer to **string** | Identificador del Proveedor del Servicio | [optional] 
**AdditionalInformation** | Pointer to [**[]QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation**](QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 

## Methods

### NewQueryCompaniesResponseObjectQueryCompaniesResponseCompanies

`func NewQueryCompaniesResponseObjectQueryCompaniesResponseCompanies() *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies`

NewQueryCompaniesResponseObjectQueryCompaniesResponseCompanies instantiates a new QueryCompaniesResponseObjectQueryCompaniesResponseCompanies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCompaniesResponseObjectQueryCompaniesResponseCompaniesWithDefaults

`func NewQueryCompaniesResponseObjectQueryCompaniesResponseCompaniesWithDefaults() *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies`

NewQueryCompaniesResponseObjectQueryCompaniesResponseCompaniesWithDefaults instantiates a new QueryCompaniesResponseObjectQueryCompaniesResponseCompanies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLineOfBusinessIdentification

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) GetLineOfBusinessIdentification() string`

GetLineOfBusinessIdentification returns the LineOfBusinessIdentification field if non-nil, zero value otherwise.

### GetLineOfBusinessIdentificationOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) GetLineOfBusinessIdentificationOk() (*string, bool)`

GetLineOfBusinessIdentificationOk returns a tuple with the LineOfBusinessIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineOfBusinessIdentification

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) SetLineOfBusinessIdentification(v string)`

SetLineOfBusinessIdentification sets LineOfBusinessIdentification field to given value.

### HasLineOfBusinessIdentification

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) HasLineOfBusinessIdentification() bool`

HasLineOfBusinessIdentification returns a boolean if a field has been set.

### GetOrder

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetAmountType

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) GetAmountType() int32`

GetAmountType returns the AmountType field if non-nil, zero value otherwise.

### GetAmountTypeOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) GetAmountTypeOk() (*int32, bool)`

GetAmountTypeOk returns a tuple with the AmountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountType

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) SetAmountType(v int32)`

SetAmountType sets AmountType field to given value.

### HasAmountType

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) HasAmountType() bool`

HasAmountType returns a boolean if a field has been set.

### GetEnableRecurringPayment

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) GetEnableRecurringPayment() bool`

GetEnableRecurringPayment returns the EnableRecurringPayment field if non-nil, zero value otherwise.

### GetEnableRecurringPaymentOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) GetEnableRecurringPaymentOk() (*bool, bool)`

GetEnableRecurringPaymentOk returns a tuple with the EnableRecurringPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecurringPayment

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) SetEnableRecurringPayment(v bool)`

SetEnableRecurringPayment sets EnableRecurringPayment field to given value.

### HasEnableRecurringPayment

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) HasEnableRecurringPayment() bool`

HasEnableRecurringPayment returns a boolean if a field has been set.

### GetProvider

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) GetAdditionalInformation() []QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) GetAdditionalInformationOk() (*[]QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) SetAdditionalInformation(v []QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *QueryCompaniesResponseObjectQueryCompaniesResponseCompanies) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


