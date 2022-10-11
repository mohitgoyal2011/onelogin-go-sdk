# FactorInner

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | MFA device identifier. | [optional] [default to null]
**Status** | **string** | accepted : factor has been verified. pending: registered but has not been verified. | [optional] [default to null]
**Default_** | **bool** | True &#x3D; is user&#x27;s default MFA device for OneLogin. | [optional] [default to null]
**AuthFactorName** | **string** | \&quot;Official\&quot; authentication factor name, as it appears to administrators in OneLogin. | [optional] [default to null]
**TypeDisplayName** | **string** | Authentication factor display name as it appears to users upon initial registration, as defined by admins at Settings &gt; Authentication Factors. | [optional] [default to null]
**UserDisplayName** | **string** | Authentication factor display name assigned by users when they enroll the device. | [optional] [default to null]
**ExpiresAt** | **string** | A short lived token that is required to Verify the Factor. This token expires based on the expires_in parameter passed in. | [optional] [default to null]
**FactorData** | [***FactorInnerFactorData**](factor_inner_factor_data.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

