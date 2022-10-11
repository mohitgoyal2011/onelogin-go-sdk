# Rule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Name** | **string** | The name of the rule. | [optional] [default to null]
**Match** | **string** | Indicates how conditions should be matched. | [optional] [default to null]
**Enabled** | **bool** | Indicates if the rule is enabled or not. | [optional] [default to null]
**Position** | **int32** | Indicates the order of the rule. When &#x60;null&#x60; this will default to last position. | [optional] [default to null]
**Conditions** | [**[]Condition**](condition.md) | An array of conditions that the user must meet in order for the rule to be applied. | [optional] [default to null]
**Actions** | [**[]Action**](action.md) | An array of actions that will be applied to the users that are matched by the conditions. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

