# EnrollFactorRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactorId** | **int32** | The identifier of the factor to enroll the user with. | [default to null]
**DisplayName** | **string** | A name for the users device. | [default to null]
**ExpiresIn** | **string** | Defaults to 120. Valid values are: 120-900 seconds. | [optional] [default to null]
**Verified** | **bool** | Defaults to false. | [optional] [default to null]
**RedirectTo** | **string** | Redirects MagicLink success page to specified URL after 2 seconds. | [optional] [default to null]
**CustomMessage** | **string** | A message template that will be sent via SMS. Max length of the message after template items are inserted is 160 characters including the OTP code. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

