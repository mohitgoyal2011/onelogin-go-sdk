# ActivateFactorRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **int32** | Required. Specifies the factor to be verified. | [optional] [default to null]
**ExpiresIn** | **int32** | Optional. Sets the window of time in seconds that the factor must be verified within.  | [optional] [default to null]
**RedirectTo** | **string** | Optional. Only applies to Email MagicLink factor. | [optional] [default to null]
**CustomMessage** | **string** | Optional. Only applies to SMS factor. A message template that will be sent via SMS. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

