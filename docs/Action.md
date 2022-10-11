# Action

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action to apply | [optional] [default to null]
**Value** | **[]string** | Only applicable to provisioned and set_* actions. Items in the array will be a plain text string or valid value for the selected action. | [optional] [default to null]
**Expression** | **string** | A regular expression to extract a value. Applies to provisionable, multi-selects, and string actions. | [optional] [default to null]
**Scriplet** | **string** | A hash containing scriptlet code that returns a value. | [optional] [default to null]
**Macro** | **string** | A template to construct a value. Applies to default, string, and list actions. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

