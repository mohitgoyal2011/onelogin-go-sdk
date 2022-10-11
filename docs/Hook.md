# Hook

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Hook unique ID in OneLogin. | [optional] [default to null]
**Type_** | **string** | A string describing the type of hook. e.g. &#x60;pre-authentication&#x60; | [default to null]
**Disabled** | **bool** | Boolean to enable or disable the hook. Disabled hooks will not run. | [default to true]
**Timeout** | **int32** | The number of seconds to allow the hook function to run before before timing out. Maximum timeout varies based on the type of hook. | [default to 1]
**EnvVars** | **[]string** | Environment Variable objects that will be available via process.env.ENV_VAR_NAME in the hook code. | [default to null]
**Runtime** | **string** | The Smart Hooks supported Node.js version to execute this hook with. | [default to null]
**Retries** | **int32** | Number of retries if execution fails. | [default to 0]
**Packages** | [***interface{}**](interface{}.md) | An object containing NPM packages that are bundled with the hook function. | [default to null]
**Function** | **string** | A base64 encoded string containing the javascript function code. | [default to null]
**ContextVersion** | **string** | The semantic version of the content that will be injected into this hook. | [optional] [default to null]
**Status** | **string** | String describing the state of the hook function. When a hook is ready and disabled is false it will be executed. | [optional] [default to null]
**Options** | [***HookOptions**](hook_options.md) |  | [optional] [default to null]
**Conditions** | [**[]HookConditionsInner**](hook_conditions_inner.md) | An array of objects that let you limit the execution of a hook to users in specific roles. | [optional] [default to null]
**CreatedAt** | **string** | ISO8601 format date that they hook function was created. | [optional] [default to null]
**UpdatedAt** | **string** | ISO8601 format date that they hook function was last updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

