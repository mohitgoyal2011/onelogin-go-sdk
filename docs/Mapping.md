# Mapping

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Name** | **string** | The name of the mapping. | [default to null]
**Enabled** | **bool** | Indicates if the mapping is enabled or not. | [default to null]
**Match** | **string** | Indicates how conditions should be matched. | [default to null]
**Position** | **int32** | Indicates the order of the mapping. When &#x60;null&#x60; this will default to last position. | [default to null]
**Conditions** | [**[]Condition**](condition.md) | An array of conditions that the user must meet in order for the mapping to be applied. | [optional] [default to null]
**Actions** | [**[]Action**](action.md) | An array of actions that will be applied to the users that are matched by the conditions. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

