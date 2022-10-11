# GenerateSamlAssertionRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsernameOrEmail** | **string** | Set this to the username or email of the OneLogin user accessing the app for which you want to generate a SAML token. | [default to null]
**Password** | **string** | Password of the OneLogin user accessing the app for which you want to generate a SAML token. | [default to null]
**AppId** | **string** | App ID of the app for which you want to generate a SAML token. This is the app ID in OneLogin. | [default to null]
**Subdomain** | **string** | Set to the subdomain of the OneLogin user accessing the app for which you want to generate a SAML token. | [default to null]
**IpAddress** | **string** | Whitelisted IP address, if MFA is required and you need to honor IP address whitelisting defined in MFA policies. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

