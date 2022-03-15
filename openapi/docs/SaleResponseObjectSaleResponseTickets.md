# SaleResponseObjectSaleResponseTickets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Tipo de Ticket | [optional] 
**Action** | Pointer to **[]string** | Acción a ejecutar para cada Ticket por el punto de venta | [optional] 
**DeviceAction** | Pointer to **[]string** |  | [optional] 
**ExecutedAction** | Pointer to **[]string** | Acción que ejecutó la plataforma para los tickets | [optional] 
**Layout** | Pointer to [**[]SaleResponseObjectSaleResponseLayout**](SaleResponseObjectSaleResponseLayout.md) | Formato del Ticket a Imprimir | [optional] 

## Methods

### NewSaleResponseObjectSaleResponseTickets

`func NewSaleResponseObjectSaleResponseTickets() *SaleResponseObjectSaleResponseTickets`

NewSaleResponseObjectSaleResponseTickets instantiates a new SaleResponseObjectSaleResponseTickets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleResponseObjectSaleResponseTicketsWithDefaults

`func NewSaleResponseObjectSaleResponseTicketsWithDefaults() *SaleResponseObjectSaleResponseTickets`

NewSaleResponseObjectSaleResponseTicketsWithDefaults instantiates a new SaleResponseObjectSaleResponseTickets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SaleResponseObjectSaleResponseTickets) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SaleResponseObjectSaleResponseTickets) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SaleResponseObjectSaleResponseTickets) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SaleResponseObjectSaleResponseTickets) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAction

`func (o *SaleResponseObjectSaleResponseTickets) GetAction() []string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SaleResponseObjectSaleResponseTickets) GetActionOk() (*[]string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SaleResponseObjectSaleResponseTickets) SetAction(v []string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SaleResponseObjectSaleResponseTickets) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDeviceAction

`func (o *SaleResponseObjectSaleResponseTickets) GetDeviceAction() []string`

GetDeviceAction returns the DeviceAction field if non-nil, zero value otherwise.

### GetDeviceActionOk

`func (o *SaleResponseObjectSaleResponseTickets) GetDeviceActionOk() (*[]string, bool)`

GetDeviceActionOk returns a tuple with the DeviceAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAction

`func (o *SaleResponseObjectSaleResponseTickets) SetDeviceAction(v []string)`

SetDeviceAction sets DeviceAction field to given value.

### HasDeviceAction

`func (o *SaleResponseObjectSaleResponseTickets) HasDeviceAction() bool`

HasDeviceAction returns a boolean if a field has been set.

### GetExecutedAction

`func (o *SaleResponseObjectSaleResponseTickets) GetExecutedAction() []string`

GetExecutedAction returns the ExecutedAction field if non-nil, zero value otherwise.

### GetExecutedActionOk

`func (o *SaleResponseObjectSaleResponseTickets) GetExecutedActionOk() (*[]string, bool)`

GetExecutedActionOk returns a tuple with the ExecutedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAction

`func (o *SaleResponseObjectSaleResponseTickets) SetExecutedAction(v []string)`

SetExecutedAction sets ExecutedAction field to given value.

### HasExecutedAction

`func (o *SaleResponseObjectSaleResponseTickets) HasExecutedAction() bool`

HasExecutedAction returns a boolean if a field has been set.

### GetLayout

`func (o *SaleResponseObjectSaleResponseTickets) GetLayout() []SaleResponseObjectSaleResponseLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *SaleResponseObjectSaleResponseTickets) GetLayoutOk() (*[]SaleResponseObjectSaleResponseLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *SaleResponseObjectSaleResponseTickets) SetLayout(v []SaleResponseObjectSaleResponseLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *SaleResponseObjectSaleResponseTickets) HasLayout() bool`

HasLayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


