# QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identificacion de Rubrro | [optional] 
**Name** | Pointer to **string** | Nombre del Rubro | [optional] 
**Description** | Pointer to **string** | Descripcion del Rubro | [optional] 
**Order** | Pointer to **int32** | Orden | [optional] 
**AdditionalInformation** | Pointer to [**[]QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation**](QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation.md) | En caso de que se requiera información adicional para poder completar la operación, como podrían ser ciertos datos ingresados por el vendedor para realizar verificaciones especificas (como los últimos 4 digitos), el código de seguridad de la tarjeta o la fecha de vencimiento, este elemento estará presente. | [optional] 

## Methods

### NewQueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness

`func NewQueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness() *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness`

NewQueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness instantiates a new QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusinessWithDefaults

`func NewQueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusinessWithDefaults() *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness`

NewQueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusinessWithDefaults instantiates a new QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrder

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) GetAdditionalInformation() []QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) GetAdditionalInformationOk() (*[]QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) SetAdditionalInformation(v []QueryCompaniesResponseObjectQueryCompaniesResponseAdditionalInformation)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *QueryLineOfBusinessResponseObjectQueryLineOfBusinessResponseLineOfBusiness) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


