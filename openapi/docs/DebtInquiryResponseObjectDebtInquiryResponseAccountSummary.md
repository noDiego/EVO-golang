# DebtInquiryResponseObjectDebtInquiryResponseAccountSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DueDate** | Pointer to **time.Time** | .. | [optional] 
**NextDueDate** | Pointer to **time.Time** | .. | [optional] 
**NextClosingDate** | Pointer to **time.Time** | .. | [optional] 
**AccountStatus** | Pointer to [**[]DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus**](DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus.md) | .. | [optional] 

## Methods

### NewDebtInquiryResponseObjectDebtInquiryResponseAccountSummary

`func NewDebtInquiryResponseObjectDebtInquiryResponseAccountSummary() *DebtInquiryResponseObjectDebtInquiryResponseAccountSummary`

NewDebtInquiryResponseObjectDebtInquiryResponseAccountSummary instantiates a new DebtInquiryResponseObjectDebtInquiryResponseAccountSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebtInquiryResponseObjectDebtInquiryResponseAccountSummaryWithDefaults

`func NewDebtInquiryResponseObjectDebtInquiryResponseAccountSummaryWithDefaults() *DebtInquiryResponseObjectDebtInquiryResponseAccountSummary`

NewDebtInquiryResponseObjectDebtInquiryResponseAccountSummaryWithDefaults instantiates a new DebtInquiryResponseObjectDebtInquiryResponseAccountSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDueDate

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummary) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummary) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummary) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummary) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetNextDueDate

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummary) GetNextDueDate() time.Time`

GetNextDueDate returns the NextDueDate field if non-nil, zero value otherwise.

### GetNextDueDateOk

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummary) GetNextDueDateOk() (*time.Time, bool)`

GetNextDueDateOk returns a tuple with the NextDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDueDate

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummary) SetNextDueDate(v time.Time)`

SetNextDueDate sets NextDueDate field to given value.

### HasNextDueDate

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummary) HasNextDueDate() bool`

HasNextDueDate returns a boolean if a field has been set.

### GetNextClosingDate

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummary) GetNextClosingDate() time.Time`

GetNextClosingDate returns the NextClosingDate field if non-nil, zero value otherwise.

### GetNextClosingDateOk

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummary) GetNextClosingDateOk() (*time.Time, bool)`

GetNextClosingDateOk returns a tuple with the NextClosingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextClosingDate

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummary) SetNextClosingDate(v time.Time)`

SetNextClosingDate sets NextClosingDate field to given value.

### HasNextClosingDate

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummary) HasNextClosingDate() bool`

HasNextClosingDate returns a boolean if a field has been set.

### GetAccountStatus

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummary) GetAccountStatus() []DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummary) GetAccountStatusOk() (*[]DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummary) SetAccountStatus(v []DebtInquiryResponseObjectDebtInquiryResponseAccountSummaryAccountStatus)`

SetAccountStatus sets AccountStatus field to given value.

### HasAccountStatus

`func (o *DebtInquiryResponseObjectDebtInquiryResponseAccountSummary) HasAccountStatus() bool`

HasAccountStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


