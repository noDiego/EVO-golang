# DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DebtAmount** | Pointer to **float32** |  | [optional] 
**AmountAvailable** | Pointer to **float32** |  | [optional] 
**PointsAvailable** | Pointer to **float32** |  | [optional] 
**MinimunPaymentAmount** | Pointer to **float32** |  | [optional] 
**CurrencyCode** | Pointer to **string** | código de Moneda - ISO 4217 &lt;https://en.wikipedia.org/wiki/ISO_4217 Se puede utilizar la Codificación Alfabética o Numérica &lt;br /&gt;   * Num   - Alpha - Description &lt;br /&gt;   * &#39;032&#39; - &#39;ARS&#39; - Pesos Argentinos &lt;br /&gt;   * &#39;152&#39; - &#39;CLP&#39; - Pesos Chilenos &lt;br/&gt;   * &#39;484&#39; - &#39;MXN&#39; - Pesos Mexicanos &lt;br/&gt;   * &#39;840&#39; - &#39;USD&#39; - dólares Americanos &lt;br/&gt;   * &#39;878&#39; - &#39;EUR&#39; - Euros &lt;br/&gt;   * &#39;858&#39; - &#39;UYU&#39; - Pesos Uruguayos &lt;br/&gt;   * &#39;878&#39; - &#39;EUR&#39; - Euros &lt;br/&gt;   * &#39;986&#39; - &#39;BRL&#39; - Real Brasileño | [optional] 

## Methods

### NewDebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus

`func NewDebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus() *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus`

NewDebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus instantiates a new DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatusWithDefaults

`func NewDebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatusWithDefaults() *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus`

NewDebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatusWithDefaults instantiates a new DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebtAmount

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) GetDebtAmount() float32`

GetDebtAmount returns the DebtAmount field if non-nil, zero value otherwise.

### GetDebtAmountOk

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) GetDebtAmountOk() (*float32, bool)`

GetDebtAmountOk returns a tuple with the DebtAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtAmount

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) SetDebtAmount(v float32)`

SetDebtAmount sets DebtAmount field to given value.

### HasDebtAmount

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) HasDebtAmount() bool`

HasDebtAmount returns a boolean if a field has been set.

### GetAmountAvailable

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) GetAmountAvailable() float32`

GetAmountAvailable returns the AmountAvailable field if non-nil, zero value otherwise.

### GetAmountAvailableOk

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) GetAmountAvailableOk() (*float32, bool)`

GetAmountAvailableOk returns a tuple with the AmountAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAvailable

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) SetAmountAvailable(v float32)`

SetAmountAvailable sets AmountAvailable field to given value.

### HasAmountAvailable

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) HasAmountAvailable() bool`

HasAmountAvailable returns a boolean if a field has been set.

### GetPointsAvailable

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) GetPointsAvailable() float32`

GetPointsAvailable returns the PointsAvailable field if non-nil, zero value otherwise.

### GetPointsAvailableOk

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) GetPointsAvailableOk() (*float32, bool)`

GetPointsAvailableOk returns a tuple with the PointsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsAvailable

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) SetPointsAvailable(v float32)`

SetPointsAvailable sets PointsAvailable field to given value.

### HasPointsAvailable

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) HasPointsAvailable() bool`

HasPointsAvailable returns a boolean if a field has been set.

### GetMinimunPaymentAmount

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) GetMinimunPaymentAmount() float32`

GetMinimunPaymentAmount returns the MinimunPaymentAmount field if non-nil, zero value otherwise.

### GetMinimunPaymentAmountOk

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) GetMinimunPaymentAmountOk() (*float32, bool)`

GetMinimunPaymentAmountOk returns a tuple with the MinimunPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimunPaymentAmount

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) SetMinimunPaymentAmount(v float32)`

SetMinimunPaymentAmount sets MinimunPaymentAmount field to given value.

### HasMinimunPaymentAmount

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) HasMinimunPaymentAmount() bool`

HasMinimunPaymentAmount returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


