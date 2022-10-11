# TrackEventRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verb** | **string** | Verbs are used to distinguish between different types of events. | [default to null]
**Ip** | **string** | The IP address of the User&#x27;s request. | [default to null]
**UserAgent** | **string** | The user agent of the User&#x27;s request. | [default to null]
**User** | [***RiskUser**](risk_user.md) |  | [default to null]
**Source** | [***Source**](source.md) |  | [optional] [default to null]
**Session** | [***Session**](session.md) |  | [optional] [default to null]
**Device** | [***RiskDevice**](risk_device.md) |  | [optional] [default to null]
**Fp** | **string** | Set to the value of the __tdli_fp cookie. | [optional] [default to null]
**Published** | **string** | Date and time of the event in IS08601 format. Useful for preloading old events. Defaults to date time this API request is received. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

