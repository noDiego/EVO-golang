# SaleResponseObjectSaleResponsePlans

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identificador interno formado por el TEF por una combinación de datos del plan especifico. Si el plan no posee un código de emisor, el último campo no estará presente (ni el _). | [optional] 
**PlanCode** | Pointer to **string** | código de el plan. | [optional] 
**ForeignIdentifier** | Pointer to **string** | ID externo para que el punto de venta pueda reconocer el plan. | [optional] 
**Issuer** | Pointer to **string** | ID del emisor que se encuentra disponible para este plan. | [optional] 
**Description** | Pointer to **string** | Descripción del plan. | [optional] 
**FacilityPayments** | Pointer to **int32** | Cantidad de cuotas que permite este plan | [optional] 
**FacilityType** | Pointer to **string** | Tipo de Plan de Financiación | [optional] 
**CurrencyCode** | Pointer to **string** | código de Moneda - ISO 4217 &lt;https://en.wikipedia.org/wiki/ISO_4217 Se puede utilizar la Codificación Alfabética o Numérica &lt;br /&gt;   * Num   - Alpha - Description &lt;br /&gt;   * &#39;032&#39; - &#39;ARS&#39; - Pesos Argentinos &lt;br /&gt;   * &#39;152&#39; - &#39;CLP&#39; - Pesos Chilenos &lt;br/&gt;   * &#39;484&#39; - &#39;MXN&#39; - Pesos Mexicanos &lt;br/&gt;   * &#39;840&#39; - &#39;USD&#39; - dólares Americanos &lt;br/&gt;   * &#39;878&#39; - &#39;EUR&#39; - Euros &lt;br/&gt;   * &#39;858&#39; - &#39;UYU&#39; - Pesos Uruguayos &lt;br/&gt;   * &#39;878&#39; - &#39;EUR&#39; - Euros &lt;br/&gt;   * &#39;986&#39; - &#39;BRL&#39; - Real Brasileño | [optional] 
**Rate** | Pointer to **float32** | Porcentaje de recargo completo para el monto de la operación. | [optional] 
**Amount** | Pointer to **float32** | Monto debido al porcentaje de recargo. | [optional] 
**OfflineAmountLimit** | Pointer to **float32** | Monto máximo para operar de forma OFFLINE. | [optional] 
**Deferral** | Pointer to [**SaleResponseObjectSaleResponsePlansDeferral**](SaleResponseObjectSaleResponsePlansDeferral.md) |  | [optional] 
**Cashback** | Pointer to [**SaleResponseObjectSaleResponsePlansCashback**](SaleResponseObjectSaleResponsePlansCashback.md) |  | [optional] 
**TNA** | Pointer to **float32** | Se informará la tasa nominal anual, en casos en que el plan elegido para realizar la venta lo posea. Por ejemplo, el plan especial llamado Plan V de Prisma informará este valor, dado que se obtendrá dinámicamente, consultandolo instante a instante. | [optional] 
**TEM** | Pointer to **float32** | Se informará la tasa efectiva mensual, en casos en que el plan elegido para realizar la venta lo posea.  | [optional] 
**TEA** | Pointer to **float32** | Tasa Efectiva anual. Este campo estará presente solo si el tipo de plan es dinámico, o si fue ingresado un valor en la base de datos. | [optional] 
**CFT** | Pointer to **float32** | Costo Financiero Total. Este campo estará presente solo si el tipo de plan es dinámico, o si fue ingresado un valor en la base de datos.         | [optional] 
**ValueFacilityPayments** | Pointer to **float32** | Monto final a pagar en cada una de las cuotas en las que se divida la compra | [optional] 
**IssuerName** | Pointer to **string** | Nombre del emisor de este plan. Estará presente solo si existe un solo emisor para todos los planes, o si el plan es dinámico | [optional] 
**IsDynamic** | Pointer to **bool** | Flag que indica si el plan es del tipo dinamico o no | [optional] 
**Category** | Pointer to **string** | Campo que le permite al punto de venta elegir entre un plan u otro | [optional] 
**POSOrDeviceActions** | Pointer to **[]string** | Lista de Acciones que debe ejecutar el POS o el Dispositvo para el caso que este Plan sea seleccionado. Acciones para el Device &lt;b&gt;RequestPIN&lt;/b&gt; | [optional] 

## Methods

### NewSaleResponseObjectSaleResponsePlans

`func NewSaleResponseObjectSaleResponsePlans() *SaleResponseObjectSaleResponsePlans`

NewSaleResponseObjectSaleResponsePlans instantiates a new SaleResponseObjectSaleResponsePlans object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleResponseObjectSaleResponsePlansWithDefaults

`func NewSaleResponseObjectSaleResponsePlansWithDefaults() *SaleResponseObjectSaleResponsePlans`

NewSaleResponseObjectSaleResponsePlansWithDefaults instantiates a new SaleResponseObjectSaleResponsePlans object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SaleResponseObjectSaleResponsePlans) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SaleResponseObjectSaleResponsePlans) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SaleResponseObjectSaleResponsePlans) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SaleResponseObjectSaleResponsePlans) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlanCode

`func (o *SaleResponseObjectSaleResponsePlans) GetPlanCode() string`

GetPlanCode returns the PlanCode field if non-nil, zero value otherwise.

### GetPlanCodeOk

`func (o *SaleResponseObjectSaleResponsePlans) GetPlanCodeOk() (*string, bool)`

GetPlanCodeOk returns a tuple with the PlanCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCode

`func (o *SaleResponseObjectSaleResponsePlans) SetPlanCode(v string)`

SetPlanCode sets PlanCode field to given value.

### HasPlanCode

`func (o *SaleResponseObjectSaleResponsePlans) HasPlanCode() bool`

HasPlanCode returns a boolean if a field has been set.

### GetForeignIdentifier

`func (o *SaleResponseObjectSaleResponsePlans) GetForeignIdentifier() string`

GetForeignIdentifier returns the ForeignIdentifier field if non-nil, zero value otherwise.

### GetForeignIdentifierOk

`func (o *SaleResponseObjectSaleResponsePlans) GetForeignIdentifierOk() (*string, bool)`

GetForeignIdentifierOk returns a tuple with the ForeignIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignIdentifier

`func (o *SaleResponseObjectSaleResponsePlans) SetForeignIdentifier(v string)`

SetForeignIdentifier sets ForeignIdentifier field to given value.

### HasForeignIdentifier

`func (o *SaleResponseObjectSaleResponsePlans) HasForeignIdentifier() bool`

HasForeignIdentifier returns a boolean if a field has been set.

### GetIssuer

`func (o *SaleResponseObjectSaleResponsePlans) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SaleResponseObjectSaleResponsePlans) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SaleResponseObjectSaleResponsePlans) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *SaleResponseObjectSaleResponsePlans) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetDescription

`func (o *SaleResponseObjectSaleResponsePlans) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SaleResponseObjectSaleResponsePlans) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SaleResponseObjectSaleResponsePlans) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SaleResponseObjectSaleResponsePlans) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacilityPayments

`func (o *SaleResponseObjectSaleResponsePlans) GetFacilityPayments() int32`

GetFacilityPayments returns the FacilityPayments field if non-nil, zero value otherwise.

### GetFacilityPaymentsOk

`func (o *SaleResponseObjectSaleResponsePlans) GetFacilityPaymentsOk() (*int32, bool)`

GetFacilityPaymentsOk returns a tuple with the FacilityPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityPayments

`func (o *SaleResponseObjectSaleResponsePlans) SetFacilityPayments(v int32)`

SetFacilityPayments sets FacilityPayments field to given value.

### HasFacilityPayments

`func (o *SaleResponseObjectSaleResponsePlans) HasFacilityPayments() bool`

HasFacilityPayments returns a boolean if a field has been set.

### GetFacilityType

`func (o *SaleResponseObjectSaleResponsePlans) GetFacilityType() string`

GetFacilityType returns the FacilityType field if non-nil, zero value otherwise.

### GetFacilityTypeOk

`func (o *SaleResponseObjectSaleResponsePlans) GetFacilityTypeOk() (*string, bool)`

GetFacilityTypeOk returns a tuple with the FacilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityType

`func (o *SaleResponseObjectSaleResponsePlans) SetFacilityType(v string)`

SetFacilityType sets FacilityType field to given value.

### HasFacilityType

`func (o *SaleResponseObjectSaleResponsePlans) HasFacilityType() bool`

HasFacilityType returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *SaleResponseObjectSaleResponsePlans) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *SaleResponseObjectSaleResponsePlans) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *SaleResponseObjectSaleResponsePlans) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *SaleResponseObjectSaleResponsePlans) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetRate

`func (o *SaleResponseObjectSaleResponsePlans) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *SaleResponseObjectSaleResponsePlans) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *SaleResponseObjectSaleResponsePlans) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *SaleResponseObjectSaleResponsePlans) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetAmount

`func (o *SaleResponseObjectSaleResponsePlans) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SaleResponseObjectSaleResponsePlans) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SaleResponseObjectSaleResponsePlans) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SaleResponseObjectSaleResponsePlans) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetOfflineAmountLimit

`func (o *SaleResponseObjectSaleResponsePlans) GetOfflineAmountLimit() float32`

GetOfflineAmountLimit returns the OfflineAmountLimit field if non-nil, zero value otherwise.

### GetOfflineAmountLimitOk

`func (o *SaleResponseObjectSaleResponsePlans) GetOfflineAmountLimitOk() (*float32, bool)`

GetOfflineAmountLimitOk returns a tuple with the OfflineAmountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineAmountLimit

`func (o *SaleResponseObjectSaleResponsePlans) SetOfflineAmountLimit(v float32)`

SetOfflineAmountLimit sets OfflineAmountLimit field to given value.

### HasOfflineAmountLimit

`func (o *SaleResponseObjectSaleResponsePlans) HasOfflineAmountLimit() bool`

HasOfflineAmountLimit returns a boolean if a field has been set.

### GetDeferral

`func (o *SaleResponseObjectSaleResponsePlans) GetDeferral() SaleResponseObjectSaleResponsePlansDeferral`

GetDeferral returns the Deferral field if non-nil, zero value otherwise.

### GetDeferralOk

`func (o *SaleResponseObjectSaleResponsePlans) GetDeferralOk() (*SaleResponseObjectSaleResponsePlansDeferral, bool)`

GetDeferralOk returns a tuple with the Deferral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferral

`func (o *SaleResponseObjectSaleResponsePlans) SetDeferral(v SaleResponseObjectSaleResponsePlansDeferral)`

SetDeferral sets Deferral field to given value.

### HasDeferral

`func (o *SaleResponseObjectSaleResponsePlans) HasDeferral() bool`

HasDeferral returns a boolean if a field has been set.

### GetCashback

`func (o *SaleResponseObjectSaleResponsePlans) GetCashback() SaleResponseObjectSaleResponsePlansCashback`

GetCashback returns the Cashback field if non-nil, zero value otherwise.

### GetCashbackOk

`func (o *SaleResponseObjectSaleResponsePlans) GetCashbackOk() (*SaleResponseObjectSaleResponsePlansCashback, bool)`

GetCashbackOk returns a tuple with the Cashback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashback

`func (o *SaleResponseObjectSaleResponsePlans) SetCashback(v SaleResponseObjectSaleResponsePlansCashback)`

SetCashback sets Cashback field to given value.

### HasCashback

`func (o *SaleResponseObjectSaleResponsePlans) HasCashback() bool`

HasCashback returns a boolean if a field has been set.

### GetTNA

`func (o *SaleResponseObjectSaleResponsePlans) GetTNA() float32`

GetTNA returns the TNA field if non-nil, zero value otherwise.

### GetTNAOk

`func (o *SaleResponseObjectSaleResponsePlans) GetTNAOk() (*float32, bool)`

GetTNAOk returns a tuple with the TNA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTNA

`func (o *SaleResponseObjectSaleResponsePlans) SetTNA(v float32)`

SetTNA sets TNA field to given value.

### HasTNA

`func (o *SaleResponseObjectSaleResponsePlans) HasTNA() bool`

HasTNA returns a boolean if a field has been set.

### GetTEM

`func (o *SaleResponseObjectSaleResponsePlans) GetTEM() float32`

GetTEM returns the TEM field if non-nil, zero value otherwise.

### GetTEMOk

`func (o *SaleResponseObjectSaleResponsePlans) GetTEMOk() (*float32, bool)`

GetTEMOk returns a tuple with the TEM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEM

`func (o *SaleResponseObjectSaleResponsePlans) SetTEM(v float32)`

SetTEM sets TEM field to given value.

### HasTEM

`func (o *SaleResponseObjectSaleResponsePlans) HasTEM() bool`

HasTEM returns a boolean if a field has been set.

### GetTEA

`func (o *SaleResponseObjectSaleResponsePlans) GetTEA() float32`

GetTEA returns the TEA field if non-nil, zero value otherwise.

### GetTEAOk

`func (o *SaleResponseObjectSaleResponsePlans) GetTEAOk() (*float32, bool)`

GetTEAOk returns a tuple with the TEA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEA

`func (o *SaleResponseObjectSaleResponsePlans) SetTEA(v float32)`

SetTEA sets TEA field to given value.

### HasTEA

`func (o *SaleResponseObjectSaleResponsePlans) HasTEA() bool`

HasTEA returns a boolean if a field has been set.

### GetCFT

`func (o *SaleResponseObjectSaleResponsePlans) GetCFT() float32`

GetCFT returns the CFT field if non-nil, zero value otherwise.

### GetCFTOk

`func (o *SaleResponseObjectSaleResponsePlans) GetCFTOk() (*float32, bool)`

GetCFTOk returns a tuple with the CFT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCFT

`func (o *SaleResponseObjectSaleResponsePlans) SetCFT(v float32)`

SetCFT sets CFT field to given value.

### HasCFT

`func (o *SaleResponseObjectSaleResponsePlans) HasCFT() bool`

HasCFT returns a boolean if a field has been set.

### GetValueFacilityPayments

`func (o *SaleResponseObjectSaleResponsePlans) GetValueFacilityPayments() float32`

GetValueFacilityPayments returns the ValueFacilityPayments field if non-nil, zero value otherwise.

### GetValueFacilityPaymentsOk

`func (o *SaleResponseObjectSaleResponsePlans) GetValueFacilityPaymentsOk() (*float32, bool)`

GetValueFacilityPaymentsOk returns a tuple with the ValueFacilityPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFacilityPayments

`func (o *SaleResponseObjectSaleResponsePlans) SetValueFacilityPayments(v float32)`

SetValueFacilityPayments sets ValueFacilityPayments field to given value.

### HasValueFacilityPayments

`func (o *SaleResponseObjectSaleResponsePlans) HasValueFacilityPayments() bool`

HasValueFacilityPayments returns a boolean if a field has been set.

### GetIssuerName

`func (o *SaleResponseObjectSaleResponsePlans) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *SaleResponseObjectSaleResponsePlans) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *SaleResponseObjectSaleResponsePlans) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.

### HasIssuerName

`func (o *SaleResponseObjectSaleResponsePlans) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.

### GetIsDynamic

`func (o *SaleResponseObjectSaleResponsePlans) GetIsDynamic() bool`

GetIsDynamic returns the IsDynamic field if non-nil, zero value otherwise.

### GetIsDynamicOk

`func (o *SaleResponseObjectSaleResponsePlans) GetIsDynamicOk() (*bool, bool)`

GetIsDynamicOk returns a tuple with the IsDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDynamic

`func (o *SaleResponseObjectSaleResponsePlans) SetIsDynamic(v bool)`

SetIsDynamic sets IsDynamic field to given value.

### HasIsDynamic

`func (o *SaleResponseObjectSaleResponsePlans) HasIsDynamic() bool`

HasIsDynamic returns a boolean if a field has been set.

### GetCategory

`func (o *SaleResponseObjectSaleResponsePlans) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SaleResponseObjectSaleResponsePlans) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SaleResponseObjectSaleResponsePlans) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SaleResponseObjectSaleResponsePlans) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetPOSOrDeviceActions

`func (o *SaleResponseObjectSaleResponsePlans) GetPOSOrDeviceActions() []string`

GetPOSOrDeviceActions returns the POSOrDeviceActions field if non-nil, zero value otherwise.

### GetPOSOrDeviceActionsOk

`func (o *SaleResponseObjectSaleResponsePlans) GetPOSOrDeviceActionsOk() (*[]string, bool)`

GetPOSOrDeviceActionsOk returns a tuple with the POSOrDeviceActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPOSOrDeviceActions

`func (o *SaleResponseObjectSaleResponsePlans) SetPOSOrDeviceActions(v []string)`

SetPOSOrDeviceActions sets POSOrDeviceActions field to given value.

### HasPOSOrDeviceActions

`func (o *SaleResponseObjectSaleResponsePlans) HasPOSOrDeviceActions() bool`

HasPOSOrDeviceActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


