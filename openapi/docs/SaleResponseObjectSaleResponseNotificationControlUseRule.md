# SaleResponseObjectSaleResponseNotificationControlUseRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id del control de la regla | [optional] 
**Name** | Pointer to **string** | Nombre del control de la regla | [optional] 
**Template** | Pointer to **string** | Mensaje a enviar en la cadena de e-mails. En caso de no haber una notificación en ControlUse, se va a generar un objeto con los mensajes provenientes de Branch, Company, Channels y Platform en caso de que estos no se encuentren vacíos. | [optional] 

## Methods

### NewSaleResponseObjectSaleResponseNotificationControlUseRule

`func NewSaleResponseObjectSaleResponseNotificationControlUseRule() *SaleResponseObjectSaleResponseNotificationControlUseRule`

NewSaleResponseObjectSaleResponseNotificationControlUseRule instantiates a new SaleResponseObjectSaleResponseNotificationControlUseRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleResponseObjectSaleResponseNotificationControlUseRuleWithDefaults

`func NewSaleResponseObjectSaleResponseNotificationControlUseRuleWithDefaults() *SaleResponseObjectSaleResponseNotificationControlUseRule`

NewSaleResponseObjectSaleResponseNotificationControlUseRuleWithDefaults instantiates a new SaleResponseObjectSaleResponseNotificationControlUseRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SaleResponseObjectSaleResponseNotificationControlUseRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SaleResponseObjectSaleResponseNotificationControlUseRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SaleResponseObjectSaleResponseNotificationControlUseRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SaleResponseObjectSaleResponseNotificationControlUseRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SaleResponseObjectSaleResponseNotificationControlUseRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SaleResponseObjectSaleResponseNotificationControlUseRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SaleResponseObjectSaleResponseNotificationControlUseRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SaleResponseObjectSaleResponseNotificationControlUseRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplate

`func (o *SaleResponseObjectSaleResponseNotificationControlUseRule) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *SaleResponseObjectSaleResponseNotificationControlUseRule) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *SaleResponseObjectSaleResponseNotificationControlUseRule) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *SaleResponseObjectSaleResponseNotificationControlUseRule) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


