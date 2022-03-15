# SaleResponseObjectSaleResponseNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Razón por la cuál se envía la notificación, puede ser BlackList, ExceptionList o ControlUse. | [optional] 
**Message** | Pointer to **string** | Mensaje a enviar el cuál puede ser proveniente de la sucursal, de la compañía, del canal o del sistema. Sólo se envía si la razón no es por Regla de control. | [optional] 
**ControlUseRule** | Pointer to [**SaleResponseObjectSaleResponseNotificationControlUseRule**](SaleResponseObjectSaleResponseNotificationControlUseRule.md) |  | [optional] 
**DistributionList** | Pointer to [**SaleResponseObjectSaleResponseNotificationDistributionList**](SaleResponseObjectSaleResponseNotificationDistributionList.md) |  | [optional] 

## Methods

### NewSaleResponseObjectSaleResponseNotification

`func NewSaleResponseObjectSaleResponseNotification() *SaleResponseObjectSaleResponseNotification`

NewSaleResponseObjectSaleResponseNotification instantiates a new SaleResponseObjectSaleResponseNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleResponseObjectSaleResponseNotificationWithDefaults

`func NewSaleResponseObjectSaleResponseNotificationWithDefaults() *SaleResponseObjectSaleResponseNotification`

NewSaleResponseObjectSaleResponseNotificationWithDefaults instantiates a new SaleResponseObjectSaleResponseNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *SaleResponseObjectSaleResponseNotification) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SaleResponseObjectSaleResponseNotification) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SaleResponseObjectSaleResponseNotification) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SaleResponseObjectSaleResponseNotification) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *SaleResponseObjectSaleResponseNotification) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SaleResponseObjectSaleResponseNotification) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SaleResponseObjectSaleResponseNotification) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SaleResponseObjectSaleResponseNotification) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetControlUseRule

`func (o *SaleResponseObjectSaleResponseNotification) GetControlUseRule() SaleResponseObjectSaleResponseNotificationControlUseRule`

GetControlUseRule returns the ControlUseRule field if non-nil, zero value otherwise.

### GetControlUseRuleOk

`func (o *SaleResponseObjectSaleResponseNotification) GetControlUseRuleOk() (*SaleResponseObjectSaleResponseNotificationControlUseRule, bool)`

GetControlUseRuleOk returns a tuple with the ControlUseRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlUseRule

`func (o *SaleResponseObjectSaleResponseNotification) SetControlUseRule(v SaleResponseObjectSaleResponseNotificationControlUseRule)`

SetControlUseRule sets ControlUseRule field to given value.

### HasControlUseRule

`func (o *SaleResponseObjectSaleResponseNotification) HasControlUseRule() bool`

HasControlUseRule returns a boolean if a field has been set.

### GetDistributionList

`func (o *SaleResponseObjectSaleResponseNotification) GetDistributionList() SaleResponseObjectSaleResponseNotificationDistributionList`

GetDistributionList returns the DistributionList field if non-nil, zero value otherwise.

### GetDistributionListOk

`func (o *SaleResponseObjectSaleResponseNotification) GetDistributionListOk() (*SaleResponseObjectSaleResponseNotificationDistributionList, bool)`

GetDistributionListOk returns a tuple with the DistributionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionList

`func (o *SaleResponseObjectSaleResponseNotification) SetDistributionList(v SaleResponseObjectSaleResponseNotificationDistributionList)`

SetDistributionList sets DistributionList field to given value.

### HasDistributionList

`func (o *SaleResponseObjectSaleResponseNotification) HasDistributionList() bool`

HasDistributionList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


