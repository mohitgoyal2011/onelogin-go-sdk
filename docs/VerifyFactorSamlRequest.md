# VerifyFactorSamlRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **string** | App ID of the app for which you want to generate a SAML token. This is the app ID in OneLogin. | [default to null]
**DeviceId** | **string** | Provide the MFA device_id you are submitting for verification. The device_id is supplied by the Generate SAML Assertion API. | [default to null]
**StateToken** | **string** | state_token associated with the MFA device_id you are submitting. The state_token is supplied by the Generate SAML Assertion API. | [default to null]
**OtpToken** | **string** | Provide the OTP value for the MFA factor you are submitting for verification. | [optional] [default to null]
**DoNotNotify** | **bool** | When verifying MFA via Protect Push, set this to true to stop additional push notifications being sent to the OneLogin Protect device. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

