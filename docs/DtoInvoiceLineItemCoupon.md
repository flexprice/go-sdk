# DtoInvoiceLineItemCoupon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CouponAssociationId** | Pointer to **string** |  | [optional] 
**CouponId** | **string** |  | 
**LineItemId** | **string** | price_id used to match the line item | 

## Methods

### NewDtoInvoiceLineItemCoupon

`func NewDtoInvoiceLineItemCoupon(couponId string, lineItemId string, ) *DtoInvoiceLineItemCoupon`

NewDtoInvoiceLineItemCoupon instantiates a new DtoInvoiceLineItemCoupon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoInvoiceLineItemCouponWithDefaults

`func NewDtoInvoiceLineItemCouponWithDefaults() *DtoInvoiceLineItemCoupon`

NewDtoInvoiceLineItemCouponWithDefaults instantiates a new DtoInvoiceLineItemCoupon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCouponAssociationId

`func (o *DtoInvoiceLineItemCoupon) GetCouponAssociationId() string`

GetCouponAssociationId returns the CouponAssociationId field if non-nil, zero value otherwise.

### GetCouponAssociationIdOk

`func (o *DtoInvoiceLineItemCoupon) GetCouponAssociationIdOk() (*string, bool)`

GetCouponAssociationIdOk returns a tuple with the CouponAssociationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponAssociationId

`func (o *DtoInvoiceLineItemCoupon) SetCouponAssociationId(v string)`

SetCouponAssociationId sets CouponAssociationId field to given value.

### HasCouponAssociationId

`func (o *DtoInvoiceLineItemCoupon) HasCouponAssociationId() bool`

HasCouponAssociationId returns a boolean if a field has been set.

### GetCouponId

`func (o *DtoInvoiceLineItemCoupon) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *DtoInvoiceLineItemCoupon) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *DtoInvoiceLineItemCoupon) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.


### GetLineItemId

`func (o *DtoInvoiceLineItemCoupon) GetLineItemId() string`

GetLineItemId returns the LineItemId field if non-nil, zero value otherwise.

### GetLineItemIdOk

`func (o *DtoInvoiceLineItemCoupon) GetLineItemIdOk() (*string, bool)`

GetLineItemIdOk returns a tuple with the LineItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemId

`func (o *DtoInvoiceLineItemCoupon) SetLineItemId(v string)`

SetLineItemId sets LineItemId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


