# Schema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Apps unique ID in OneLogin. | [optional] [default to null]
**ConnectorId** | **int32** | ID of the apps underlying connector. | [optional] [default to null]
**Name** | **string** | App name. | [optional] [default to null]
**Description** | **string** | Freeform description of the app. | [optional] [default to null]
**Notes** | **string** | Freeform notes about the app. | [optional] [default to null]
**PolicyId** | **int32** | The security policy assigned to the app. | [optional] [default to null]
**BrandId** | **int32** | The custom login page branding to use for this app. Applies to app initiated logins via OIDC or SAML. | [optional] [default to null]
**IconUrl** | **string** | A link to the apps icon url. | [optional] [default to null]
**Visible** | **bool** | Indicates if the app is visible in the OneLogin portal. | [optional] [default to null]
**AuthMethod** | **int32** | An ID indicating the type of app. | [optional] [default to null]
**TabId** | **int32** | ID of the OneLogin portal tab that the app is assigned to. | [optional] [default to null]
**CreatedAt** | **string** | The date the app was created. | [optional] [default to null]
**UpdatedAt** | **string** | The date the app was last updated. | [optional] [default to null]
**RoleIds** | **[]int32** | List of Role IDs that are assigned to the app. On App Create or Update the entire array is replaced with the values provided. | [optional] [default to null]
**AllowAssumedSignin** | **bool** | Indicates whether or not administrators can access the app as a user that they have assumed control over. | [optional] [default to null]
**Provisioning** | [***SchemaProvisioning**](schema_provisioning.md) |  | [optional] [default to null]
**Sso** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**Configuration** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**Parameters** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**EnforcementPoint** | [***interface{}**](interface{}.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

