# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateFactor**](DefaultApi.md#ActivateFactor) | **Post** /api/2/mfa/users/{user_id}/verifications | 
[**AddAccessTokenClaim**](DefaultApi.md#AddAccessTokenClaim) | **Post** /api/2/api_authorizations/{id}/claims | 
[**AddClientApp**](DefaultApi.md#AddClientApp) | **Post** /api/2/api_authorizations/{id}/clients | 
[**AddRoleAdmins**](DefaultApi.md#AddRoleAdmins) | **Post** /api/2/roles/{role_id}/admins | 
[**AddRoleUsers**](DefaultApi.md#AddRoleUsers) | **Post** /api/2/roles/{role_id}/users | 
[**AddScope**](DefaultApi.md#AddScope) | **Post** /api/2/api_authorizations/{id}/scopes | 
[**BulkMappingSort**](DefaultApi.md#BulkMappingSort) | **Put** /api/2/apps/mappings/sort | 
[**BulkSort**](DefaultApi.md#BulkSort) | **Put** /api/2/apps/{app_id}/rules/sort | 
[**CreateApp**](DefaultApi.md#CreateApp) | **Post** /api/2/apps | 
[**CreateAuthorizationServer**](DefaultApi.md#CreateAuthorizationServer) | **Post** /api/2/api_authorizations | 
[**CreateEnvironmentVariable**](DefaultApi.md#CreateEnvironmentVariable) | **Post** /api/2/hooks/envs | 
[**CreateHook**](DefaultApi.md#CreateHook) | **Post** /api/2/hooks | 
[**CreateMapping**](DefaultApi.md#CreateMapping) | **Post** /api/2/mappings | 
[**CreateRiskRule**](DefaultApi.md#CreateRiskRule) | **Post** /api/2/risk/rules | 
[**CreateRoles**](DefaultApi.md#CreateRoles) | **Post** /api/2/roles | 
[**CreateRule**](DefaultApi.md#CreateRule) | **Post** /api/2/apps/{app_id}/rules | 
[**CreateUser**](DefaultApi.md#CreateUser) | **Post** /api/2/users | 
[**DeleteAccessTokenClaim**](DefaultApi.md#DeleteAccessTokenClaim) | **Delete** /api/2/api_authorizations/{id}/claims/{claim_id} | 
[**DeleteApp**](DefaultApi.md#DeleteApp) | **Delete** /api/2/apps/{app_id} | 
[**DeleteAppParameter**](DefaultApi.md#DeleteAppParameter) | **Delete** /api/2/apps/{app_id}/parameters/{parameter_id} | 
[**DeleteAuthorizationServer**](DefaultApi.md#DeleteAuthorizationServer) | **Delete** /api/2/api_authorizations/{id} | 
[**DeleteEnvironmentVariable**](DefaultApi.md#DeleteEnvironmentVariable) | **Delete** /api/2/hooks/envs/{envvar_id} | 
[**DeleteFactor**](DefaultApi.md#DeleteFactor) | **Delete** /api/2/mfa/users/{user_id}/devices/{device_id} | 
[**DeleteHook**](DefaultApi.md#DeleteHook) | **Delete** /api/2/hooks/{hook_id} | 
[**DeleteMapping**](DefaultApi.md#DeleteMapping) | **Delete** /api/2/mappings/{mapping_id} | 
[**DeleteRiskRule**](DefaultApi.md#DeleteRiskRule) | **Delete** /api/2/risk/rules/{risk_rule_id} | 
[**DeleteRole**](DefaultApi.md#DeleteRole) | **Delete** /api/2/roles/{role_id} | 
[**DeleteRule**](DefaultApi.md#DeleteRule) | **Delete** /api/2/apps/{app_id}/rules/{rule_id} | 
[**DeleteScope**](DefaultApi.md#DeleteScope) | **Delete** /api/2/api_authorizations/{id}/scopes/{scope_id} | 
[**DeleteUser**](DefaultApi.md#DeleteUser) | **Delete** /api/2/users/{user_id} | 
[**DryRunMapping**](DefaultApi.md#DryRunMapping) | **Post** /api/2/mappings/{mapping_id}/dryrun | 
[**EnrollFactor**](DefaultApi.md#EnrollFactor) | **Post** /api/2/mfa/users/{user_id}/registrations | 
[**GenerateMfaToken**](DefaultApi.md#GenerateMfaToken) | **Post** /api/2/mfs/users/{user_id}/mfa_token | 
[**GenerateSamlAssertion**](DefaultApi.md#GenerateSamlAssertion) | **Post** /api/2/saml_assertion | 
[**GenerateToken**](DefaultApi.md#GenerateToken) | **Post** /auth/oauth2/v2/token | 
[**GetApp**](DefaultApi.md#GetApp) | **Get** /api/2/apps/{app_id} | 
[**GetAuthorizationServer**](DefaultApi.md#GetAuthorizationServer) | **Get** /api/2/api_authorizations/{id} | 
[**GetAvailableFactors**](DefaultApi.md#GetAvailableFactors) | **Get** /api/2/mfa/users/{user_id}/factors | 
[**GetClientApps**](DefaultApi.md#GetClientApps) | **Get** /api/2/api_authorizations/{id}/clients | 
[**GetEnrolledFactors**](DefaultApi.md#GetEnrolledFactors) | **Get** /api/2/mfa/users/{user_id}/devices | 
[**GetEnvironmentVariable**](DefaultApi.md#GetEnvironmentVariable) | **Get** /api/2/hooks/envs/{envvar_id} | 
[**GetHook**](DefaultApi.md#GetHook) | **Get** /api/2/hooks/{hook_id} | 
[**GetLogs**](DefaultApi.md#GetLogs) | **Get** /api/2/hooks/{hook_id}/logs | 
[**GetMapping**](DefaultApi.md#GetMapping) | **Get** /api/2/mappings/{mapping_id} | 
[**GetRateLimit**](DefaultApi.md#GetRateLimit) | **Get** /auth/rate_limit | 
[**GetRiskRule**](DefaultApi.md#GetRiskRule) | **Get** /api/2/risk/rules/{risk_rule_id} | 
[**GetRiskScore**](DefaultApi.md#GetRiskScore) | **Post** /api/2/risk/verify | 
[**GetRole**](DefaultApi.md#GetRole) | **Get** /api/2/roles/{role_id} | 
[**GetRoleAdmins**](DefaultApi.md#GetRoleAdmins) | **Get** /api/2/roles/{role_id}/admins | 
[**GetRoleApps**](DefaultApi.md#GetRoleApps) | **Get** /api/2/roles/{role_id}/apps | 
[**GetRoleUsers**](DefaultApi.md#GetRoleUsers) | **Get** /api/2/roles/{role_id}/users | 
[**GetRule**](DefaultApi.md#GetRule) | **Get** /api/2/apps/{app_id}/rules/{rule_id} | 
[**GetScoreInsights**](DefaultApi.md#GetScoreInsights) | **Get** /api/2/risk/scores | 
[**GetUser**](DefaultApi.md#GetUser) | **Get** /api/2/users/{user_id} | 
[**GetUserApps**](DefaultApi.md#GetUserApps) | **Get** /api/2/users/{user_id}/apps | 
[**ListAccessTokenClaims**](DefaultApi.md#ListAccessTokenClaims) | **Get** /api/2/api_authorizations/{id}/claims | 
[**ListActionValues**](DefaultApi.md#ListActionValues) | **Get** /api/2/apps/{app_id}/rules/actions/{actuion_value}/values | 
[**ListActions**](DefaultApi.md#ListActions) | **Get** /api/2/apps/{app_id}/rules/actions | 
[**ListAppUsers**](DefaultApi.md#ListAppUsers) | **Get** /api/2/apps/{app_id}/users | 
[**ListApps**](DefaultApi.md#ListApps) | **Get** /api/2/apps | 
[**ListAuthorizationServers**](DefaultApi.md#ListAuthorizationServers) | **Get** /api/2/api_authorizations | 
[**ListConditionOperators**](DefaultApi.md#ListConditionOperators) | **Get** /api/2/apps/{app_id}/rules/conditions/{condition_value}/operators | 
[**ListConditionValues**](DefaultApi.md#ListConditionValues) | **Get** /api/2/apps/{app_id}/rules/conditions/{condition_value}/values | 
[**ListConditions**](DefaultApi.md#ListConditions) | **Get** /api/2/apps/{app_id}/rules/conditions | 
[**ListConnectors**](DefaultApi.md#ListConnectors) | **Get** /api/2/connectors | 
[**ListEnvironmentVariables**](DefaultApi.md#ListEnvironmentVariables) | **Get** /api/2/hooks/envs | 
[**ListHooks**](DefaultApi.md#ListHooks) | **Get** /api/2/hooks | 
[**ListMappingActionValues**](DefaultApi.md#ListMappingActionValues) | **Get** /api/2/apps/mappings/actions/{actuion_value}/values | 
[**ListMappingActions**](DefaultApi.md#ListMappingActions) | **Get** /api/2/apps/mappings/actions | 
[**ListMappingConditionOperators**](DefaultApi.md#ListMappingConditionOperators) | **Get** /api/2/apps/mappings/conditions/{condition_value}/operators | 
[**ListMappingConditionValues**](DefaultApi.md#ListMappingConditionValues) | **Get** /api/2/apps/mappings/conditions/{condition_value}/values | 
[**ListMappingConditions**](DefaultApi.md#ListMappingConditions) | **Get** /api/2/apps/mappings/conditions | 
[**ListMappings**](DefaultApi.md#ListMappings) | **Get** /api/2/mappings | 
[**ListRiskRules**](DefaultApi.md#ListRiskRules) | **Get** /api/2/risk/rules | 
[**ListRoles**](DefaultApi.md#ListRoles) | **Get** /api/2/roles | 
[**ListRules**](DefaultApi.md#ListRules) | **Get** /api/2/apps/{app_id}/rules | 
[**ListScopes**](DefaultApi.md#ListScopes) | **Get** /api/2/api_authorizations/{id}/scopes | 
[**ListUsers**](DefaultApi.md#ListUsers) | **Get** /api/2/users | 
[**RemoveClientApp**](DefaultApi.md#RemoveClientApp) | **Delete** /api/2/api_authorizations/{id}/clients/{client_app_id} | 
[**RemoveRoleAdmins**](DefaultApi.md#RemoveRoleAdmins) | **Delete** /api/2/roles/{role_id}/admins | 
[**RemoveRoleUsers**](DefaultApi.md#RemoveRoleUsers) | **Delete** /api/2/roles/{role_id}/users | 
[**RevokeToken**](DefaultApi.md#RevokeToken) | **Post** /auth/oauth2/revoke | 
[**SetRoleApps**](DefaultApi.md#SetRoleApps) | **Put** /api/2/roles/{role_id}/apps | 
[**TrackEvent**](DefaultApi.md#TrackEvent) | **Post** /api/2/risk/events | 
[**UpdateAccessTokenClaim**](DefaultApi.md#UpdateAccessTokenClaim) | **Put** /api/2/api_authorizations/{id}/claims/{claim_id} | 
[**UpdateApp**](DefaultApi.md#UpdateApp) | **Put** /api/2/apps/{app_id} | 
[**UpdateAuthorizationServer**](DefaultApi.md#UpdateAuthorizationServer) | **Put** /api/2/api_authorizations/{id} | 
[**UpdateClientApp**](DefaultApi.md#UpdateClientApp) | **Put** /api/2/api_authorizations/{id}/clients/{client_app_id} | 
[**UpdateEnvironmentVariable**](DefaultApi.md#UpdateEnvironmentVariable) | **Put** /api/2/hooks/envs/{envvar_id} | 
[**UpdateHook**](DefaultApi.md#UpdateHook) | **Put** /api/2/hooks/{hook_id} | 
[**UpdateMapping**](DefaultApi.md#UpdateMapping) | **Put** /api/2/mappings/{mapping_id} | 
[**UpdateRiskRule**](DefaultApi.md#UpdateRiskRule) | **Put** /api/2/risk/rules/{risk_rule_id} | 
[**UpdateRole**](DefaultApi.md#UpdateRole) | **Put** /api/2/roles/{role_id} | 
[**UpdateRule**](DefaultApi.md#UpdateRule) | **Put** /api/2/apps/{app_id}/rules/{rule_id} | 
[**UpdateScope**](DefaultApi.md#UpdateScope) | **Put** /api/2/api_authorizations/{id}/scopes/{scope_id} | 
[**UpdateUser**](DefaultApi.md#UpdateUser) | **Put** /api/2/users/{user_id} | 
[**VerifyEnrollment**](DefaultApi.md#VerifyEnrollment) | **Put** /api/2/mfa/users/{user_id}/registrations/{registration_id} | 
[**VerifyEnrollmentVoiceProtect**](DefaultApi.md#VerifyEnrollmentVoiceProtect) | **Get** /api/2/mfa/users/{user_id}/registrations/{registration_id} | 
[**VerifyFactor**](DefaultApi.md#VerifyFactor) | **Put** /api/2/mfa/users/{user_id}/verifications/{verification_id} | 
[**VerifyFactorSaml**](DefaultApi.md#VerifyFactorSaml) | **Post** /api/2/saml_assertion/verify_factor | 
[**VerifyFactorVoice**](DefaultApi.md#VerifyFactorVoice) | **Get** /api/2/mfa/users/{user_id}/verifications/{verification_id} | 

# **ActivateFactor**
> ActivateFactor(ctx, body, authorization, userId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ActivateFactorRequest**](ActivateFactorRequest.md)|  | 
  **authorization** | **string**|  | 
  **userId** | **int32**| Set to the id of the user. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddAccessTokenClaim**
> Id AddAccessTokenClaim(ctx, body, authorization, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddAccessTokenClaimRequest**](AddAccessTokenClaimRequest.md)|  | 
  **authorization** | **string**|  | 
  **id** | **int32**|  | 

### Return type

[**Id**](id.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddClientApp**
> ClientApp AddClientApp(ctx, body, authorization, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddClientAppRequest**](AddClientAppRequest.md)|  | 
  **authorization** | **string**|  | 
  **id** | **int32**|  | 

### Return type

[**ClientApp**](client_app.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRoleAdmins**
> []AddRoleUsers200ResponseInner AddRoleAdmins(ctx, body, authorization, roleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]int32**](int32.md)|  | 
  **authorization** | **string**|  | 
  **roleId** | [**int32**](.md)| Set to the id of the role you want to return. | 

### Return type

[**[]AddRoleUsers200ResponseInner**](addRoleUsers_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRoleUsers**
> []AddRoleUsers200ResponseInner AddRoleUsers(ctx, body, authorization, roleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]int32**](int32.md)|  | 
  **authorization** | **string**|  | 
  **roleId** | [**int32**](.md)| Set to the id of the role you want to return. | 

### Return type

[**[]AddRoleUsers200ResponseInner**](addRoleUsers_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddScope**
> Id AddScope(ctx, body, authorization, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddScopeRequest**](AddScopeRequest.md)|  | 
  **authorization** | **string**|  | 
  **id** | **int32**|  | 

### Return type

[**Id**](id.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkMappingSort**
> []int32 BulkMappingSort(ctx, body, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]int32**](int32.md)| The request body must contain an array of User Mapping IDs in the desired order. | 
  **authorization** | **string**|  | 

### Return type

[**[]int32**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkSort**
> []int32 BulkSort(ctx, body, authorization, appId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]int32**](int32.md)| The request body must contain an array of App Rule IDs in the desired order. | 
  **authorization** | **string**|  | 
  **appId** | [**int32**](.md)|  | 

### Return type

[**[]int32**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateApp**
> Schema CreateApp(ctx, body, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Schema**](Schema.md)|  | 
  **authorization** | **string**|  | 

### Return type

[**Schema**](schema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAuthorizationServer**
> Id CreateAuthorizationServer(ctx, body, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateAuthorizationServerRequest**](CreateAuthorizationServerRequest.md)|  | 
  **authorization** | **string**|  | 

### Return type

[**Id**](id.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEnvironmentVariable**
> Envvar CreateEnvironmentVariable(ctx, body, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateEnvironmentVariableRequest**](CreateEnvironmentVariableRequest.md)|  | 
  **authorization** | **string**|  | 

### Return type

[**Envvar**](envvar.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHook**
> CreateHook(ctx, body, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Hook**](Hook.md)|  | 
  **authorization** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMapping**
> int32 CreateMapping(ctx, body, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Mapping**](Mapping.md)|  | 
  **authorization** | **string**|  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRiskRule**
> CreateRiskRule(ctx, body, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RiskRule**](RiskRule.md)|  | 
  **authorization** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRoles**
> []CreateRoles201ResponseInner CreateRoles(ctx, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 

### Return type

[**[]CreateRoles201ResponseInner**](createRoles_201_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRule**
> RuleId CreateRule(ctx, body, authorization, appId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rule**](Rule.md)|  | 
  **authorization** | **string**|  | 
  **appId** | [**int32**](.md)|  | 

### Return type

[**RuleId**](rule_id.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> User CreateUser(ctx, body, authorization, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**User**](User.md)|  | 
  **authorization** | **string**|  | 
 **optional** | ***DefaultApiCreateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiCreateUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mappings** | **optional.**| Controls how mappings will be applied to the user on creation. Defaults to async. | 
 **validatePolicy** | **optional.**| Will passwords validate against the User Policy? Defaults to true. | 

### Return type

[**User**](user.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccessTokenClaim**
> DeleteAccessTokenClaim(ctx, authorization, id, claimId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **id** | **int32**|  | 
  **claimId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApp**
> DeleteApp(ctx, authorization, appId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **appId** | [**int32**](.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAppParameter**
> DeleteAppParameter(ctx, authorization, appId, parameterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **appId** | [**int32**](.md)|  | 
  **parameterId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAuthorizationServer**
> DeleteAuthorizationServer(ctx, authorization, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnvironmentVariable**
> DeleteEnvironmentVariable(ctx, authorization, envvarId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **envvarId** | [**string**](.md)| Set to the id of the Hook Environment Variable that you want to fetch. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFactor**
> DeleteFactor(ctx, authorization, userId, deviceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **userId** | **int32**| Set to the id of the user. | 
  **deviceId** | **int32**| Set to the device_id of the MFA device. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHook**
> DeleteHook(ctx, authorization, hookId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **hookId** | [**string**](.md)| Set to the id of the Hook that you want to return. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMapping**
> DeleteMapping(ctx, authorization, mappingId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **mappingId** | **int32**| The id of the user mapping to locate. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRiskRule**
> RiskRule DeleteRiskRule(ctx, authorization, riskRuleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **riskRuleId** | [**string**](.md)|  | 

### Return type

[**RiskRule**](risk_rule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRole**
> DeleteRole(ctx, authorization, roleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **roleId** | [**int32**](.md)| Set to the id of the role you want to return. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRule**
> DeleteRule(ctx, authorization, appId, ruleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **appId** | [**int32**](.md)|  | 
  **ruleId** | [**int32**](.md)| The id of the app rule to locate. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteScope**
> DeleteScope(ctx, authorization, id, scopeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **id** | **int32**|  | 
  **scopeId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> DeleteUser(ctx, authorization, userId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **userId** | [**int32**](.md)| Set to the id of the user that you want to return. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DryRunMapping**
> []interface{} DryRunMapping(ctx, body, authorization, mappingId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]int32**](int32.md)| Request body is a list of user IDs tested against the mapping conditions to verify that the mapping would be applied | 
  **authorization** | **string**|  | 
  **mappingId** | **int32**| The id of the user mapping to locate. | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnrollFactor**
> [][]FactorInner EnrollFactor(ctx, body, authorization, userId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnrollFactorRequest**](EnrollFactorRequest.md)|  | 
  **authorization** | **string**|  | 
  **userId** | **int32**| Set to the id of the user. | 

### Return type

[**[][]FactorInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateMfaToken**
> GenerateMfaToken200Response GenerateMfaToken(ctx, body, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GenerateMfaTokenRequest**](GenerateMfaTokenRequest.md)|  | 
  **authorization** | **string**|  | 

### Return type

[**GenerateMfaToken200Response**](generateMfaToken_200_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateSamlAssertion**
> GenerateSamlAssertion(ctx, body, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GenerateSamlAssertionRequest**](GenerateSamlAssertionRequest.md)|  | 
  **authorization** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateToken**
> GenerateToken200Response GenerateToken(ctx, body, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GenerateTokenRequest**](GenerateTokenRequest.md)|  | 
  **authorization** | **string**|  | 

### Return type

[**GenerateToken200Response**](generateToken_200_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApp**
> Schema GetApp(ctx, authorization, appId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **appId** | [**int32**](.md)|  | 

### Return type

[**Schema**](schema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthorizationServer**
> GetAuthorizationServer200Response GetAuthorizationServer(ctx, authorization, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **id** | **int32**|  | 

### Return type

[**GetAuthorizationServer200Response**](getAuthorizationServer_200_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvailableFactors**
> []GetAvailableFactors200ResponseInner GetAvailableFactors(ctx, authorization, userId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **userId** | **int32**| Set to the id of the user. | 

### Return type

[**[]GetAvailableFactors200ResponseInner**](getAvailableFactors_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientApps**
> []GetClientApps200ResponseInner GetClientApps(ctx, authorization, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **id** | **int32**|  | 

### Return type

[**[]GetClientApps200ResponseInner**](getClientApps_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnrolledFactors**
> []Device GetEnrolledFactors(ctx, authorization, userId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **userId** | **int32**| Set to the id of the user. | 

### Return type

[**[]Device**](device.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironmentVariable**
> Envvar GetEnvironmentVariable(ctx, authorization, envvarId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **envvarId** | [**string**](.md)| Set to the id of the Hook Environment Variable that you want to fetch. | 

### Return type

[**Envvar**](envvar.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHook**
> Hook GetHook(ctx, authorization, hookId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **hookId** | [**string**](.md)| Set to the id of the Hook that you want to return. | 

### Return type

[**Hook**](hook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogs**
> []Log GetLogs(ctx, authorization, hookId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **hookId** | [**string**](.md)| Set to the id of the Hook that you want to return. | 
 **optional** | ***DefaultApiGetLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **optional.Int32**| The page number of results to return. | 
 **cursor** | **optional.String**| Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **requestId** | **optional.String**| Returns logs that contain this request_id. | 
 **correlationId** | **optional.String**| Returns logs that contain this correlation_id. | 

### Return type

[**[]Log**](log.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMapping**
> Mapping GetMapping(ctx, authorization, mappingId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **mappingId** | **int32**| The id of the user mapping to locate. | 

### Return type

[**Mapping**](mapping.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRateLimit**
> GetRateLimit200Response GetRateLimit(ctx, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 

### Return type

[**GetRateLimit200Response**](getRateLimit_200_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRiskRule**
> GetRiskRule(ctx, authorization, riskRuleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **riskRuleId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRiskScore**
> GetRiskScore200Response GetRiskScore(ctx, body, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GetRiskScoreRequest**](GetRiskScoreRequest.md)|  | 
  **authorization** | **string**|  | 

### Return type

[**GetRiskScore200Response**](getRiskScore_200_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRole**
> Role GetRole(ctx, authorization, roleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **roleId** | [**int32**](.md)| Set to the id of the role you want to return. | 

### Return type

[**Role**](role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleAdmins**
> []Schema1 GetRoleAdmins(ctx, authorization, roleId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **roleId** | [**int32**](.md)| Set to the id of the role you want to return. | 
 **optional** | ***DefaultApiGetRoleAdminsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetRoleAdminsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **optional.Int32**| The page number of results to return. | 
 **cursor** | **optional.String**| Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **name** | **optional.String**| Allows you to filter on first name, last name, username, and email address. | 
 **includeUnassigned** | **optional.Bool**| Optional. Defaults to false. Include users that aren’t assigned to the role. | 

### Return type

[**[]Schema1**](schema_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: applcation/json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleApps**
> []Schema GetRoleApps(ctx, authorization, roleId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **roleId** | [**int32**](.md)| Set to the id of the role you want to return. | 
 **optional** | ***DefaultApiGetRoleAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetRoleAppsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **optional.Int32**| The page number of results to return. | 
 **cursor** | **optional.String**| Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **assigned** | **optional.Bool**| Optional. Defaults to true. Returns all apps not yet assigned to the role. | 

### Return type

[**[]Schema**](schema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: applcation/json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleUsers**
> []Schema1 GetRoleUsers(ctx, authorization, roleId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **roleId** | [**int32**](.md)| Set to the id of the role you want to return. | 
 **optional** | ***DefaultApiGetRoleUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetRoleUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **optional.Int32**| The page number of results to return. | 
 **cursor** | **optional.String**| Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **name** | **optional.String**| Allows you to filter on first name, last name, username, and email address. | 
 **includeUnassigned** | **optional.Bool**| Optional. Defaults to false. Include users that aren’t assigned to the role. | 

### Return type

[**[]Schema1**](schema_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: applcation/json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRule**
> Rule GetRule(ctx, authorization, appId, ruleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **appId** | [**int32**](.md)|  | 
  **ruleId** | [**int32**](.md)| The id of the app rule to locate. | 

### Return type

[**Rule**](rule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScoreInsights**
> GetScoreInsights200Response GetScoreInsights(ctx, authorization, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
 **optional** | ***DefaultApiGetScoreInsightsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetScoreInsightsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **before** | **optional.String**| Optional ISO8601 formatted date string. Defaults to current date. Maximum date is 90 days ago. | 
 **after** | **optional.String**| Optional ISO8601 formatted date string. Defaults to 30 days ago. Maximum date is 90 days ago. | 

### Return type

[**GetScoreInsights200Response**](getScoreInsights_200_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> User GetUser(ctx, authorization, userId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **userId** | [**int32**](.md)| Set to the id of the user that you want to return. | 

### Return type

[**User**](user.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserApps**
> []GetUserApps200ResponseInner GetUserApps(ctx, authorization, userId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **userId** | [**int32**](.md)| Set to the id of the user that you want to return. | 
 **optional** | ***DefaultApiGetUserAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetUserAppsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ignoreVisibility** | **optional.Bool**| Defaults to &#x60;false&#x60;. When &#x60;true&#x60; will show all apps that are assigned to a user regardless of their portal visibility setting. | 

### Return type

[**[]GetUserApps200ResponseInner**](getUserApps_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccessTokenClaims**
> []ListAccessTokenClaims200ResponseInner ListAccessTokenClaims(ctx, authorization, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **id** | **int32**|  | 

### Return type

[**[]ListAccessTokenClaims200ResponseInner**](listAccessTokenClaims_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListActionValues**
> []ListConditionValues200ResponseInner ListActionValues(ctx, authorization, appId, actionValue)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **appId** | [**int32**](.md)|  | 
  **actionValue** | **string**| The value for the selected action. | 

### Return type

[**[]ListConditionValues200ResponseInner**](listConditionValues_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListActions**
> []ListActions200ResponseInner ListActions(ctx, authorization, appId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **appId** | [**int32**](.md)|  | 

### Return type

[**[]ListActions200ResponseInner**](listActions_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAppUsers**
> []ListAppUsers200ResponseInner ListAppUsers(ctx, authorization, appId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **appId** | [**int32**](.md)|  | 
 **optional** | ***DefaultApiListAppUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListAppUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **optional.Int32**| The page number of results to return. | 
 **cursor** | **optional.String**| Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 

### Return type

[**[]ListAppUsers200ResponseInner**](listAppUsers_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApps**
> []Schema ListApps(ctx, authorization, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
 **optional** | ***DefaultApiListAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListAppsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **optional.Int32**| The page number of results to return. | 
 **cursor** | **optional.String**| Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **name** | [**optional.Interface of string**](.md)| The name or partial name of the app to search for. When using a partial name you must append a wildcard &#x60;*&#x60; | 
 **connectorId** | [**optional.Interface of int32**](.md)| Returns all apps based off a specific connector. See List Connectors for a complete list of Connector IDs. | 
 **authMethod** | [**optional.Interface of AuthMethod**](.md)| Returns all apps based of a given type. | 

### Return type

[**[]Schema**](schema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAuthorizationServers**
> []ListAuthorizationServers200ResponseInner ListAuthorizationServers(ctx, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 

### Return type

[**[]ListAuthorizationServers200ResponseInner**](listAuthorizationServers_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConditionOperators**
> []ListConditionOperators200ResponseInner ListConditionOperators(ctx, authorization, appId, conditionValue)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **appId** | [**int32**](.md)|  | 
  **conditionValue** | **string**| The value for the selected condition. | 

### Return type

[**[]ListConditionOperators200ResponseInner**](listConditionOperators_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConditionValues**
> []ListConditionValues200ResponseInner ListConditionValues(ctx, authorization, appId, conditionValue)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **appId** | [**int32**](.md)|  | 
  **conditionValue** | **string**| The value for the selected condition. | 

### Return type

[**[]ListConditionValues200ResponseInner**](listConditionValues_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConditions**
> []ListConditions200ResponseInner ListConditions(ctx, authorization, appId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **appId** | [**int32**](.md)|  | 

### Return type

[**[]ListConditions200ResponseInner**](listConditions_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConnectors**
> []Connector ListConnectors(ctx, authorization, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
 **optional** | ***DefaultApiListConnectorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListConnectorsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **optional.Int32**| The page number of results to return. | 
 **cursor** | **optional.String**| Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **name** | [**optional.Interface of string**](.md)| The name or partial name of the connector to search for. When using a partial name you must append a wildcard &#x60;*&#x60; | 
 **authMethod** | [**optional.Interface of AuthMethod**](.md)| Returns all connectors of a given type. | 

### Return type

[**[]Connector**](connector.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEnvironmentVariables**
> []Envvar ListEnvironmentVariables(ctx, authorization, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
 **optional** | ***DefaultApiListEnvironmentVariablesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListEnvironmentVariablesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **optional.Int32**| The page number of results to return. | 
 **cursor** | **optional.String**| Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 

### Return type

[**[]Envvar**](envvar.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHooks**
> []Hook ListHooks(ctx, authorization, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
 **optional** | ***DefaultApiListHooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListHooksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **optional.Int32**| The page number of results to return. | 
 **cursor** | **optional.String**| Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 

### Return type

[**[]Hook**](hook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMappingActionValues**
> []ListConditionValues200ResponseInner ListMappingActionValues(ctx, authorization, actionValue)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **actionValue** | **string**| The value for the selected action. | 

### Return type

[**[]ListConditionValues200ResponseInner**](listConditionValues_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMappingActions**
> []ListActions200ResponseInner ListMappingActions(ctx, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 

### Return type

[**[]ListActions200ResponseInner**](listActions_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMappingConditionOperators**
> []ListMappingConditionOperators200ResponseInner ListMappingConditionOperators(ctx, authorization, conditionValue)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **conditionValue** | **string**| The value for the selected condition. | 

### Return type

[**[]ListMappingConditionOperators200ResponseInner**](listMappingConditionOperators_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMappingConditionValues**
> []ListConditionValues200ResponseInner ListMappingConditionValues(ctx, authorization, conditionValue)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **conditionValue** | **string**| The value for the selected condition. | 

### Return type

[**[]ListConditionValues200ResponseInner**](listConditionValues_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMappingConditions**
> []ListMappingConditions200ResponseInner ListMappingConditions(ctx, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 

### Return type

[**[]ListMappingConditions200ResponseInner**](listMappingConditions_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMappings**
> []Mapping ListMappings(ctx, authorization, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
 **optional** | ***DefaultApiListMappingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListMappingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enabled** | **optional.Bool**| Defaults to true. When set to &#x60;false&#x60; will return all disabled mappings. | [default to true]
 **hasCondition** | **optional.String**| Filters Mappings based on their Conditions. | 
 **hasConditionType** | **optional.String**| Filters Mappings based on their condition types. | 
 **hasAction** | **optional.String**| Filters Mappings based on their Actions. | 
 **hasActionType** | **optional.String**| Filters Mappings based on their action types. | 

### Return type

[**[]Mapping**](mapping.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRiskRules**
> ListRiskRules(ctx, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRoles**
> []Role ListRoles(ctx, authorization, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
 **optional** | ***DefaultApiListRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListRolesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **optional.Int32**| The page number of results to return. | 
 **cursor** | **optional.String**| Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **name** | [**optional.Interface of string**](.md)| Optional. Filters by role name. | 
 **appId** | [**optional.Interface of string**](.md)| Optional. Returns roles that contain this app name. | 
 **fields** | **optional.String**| Optional. Comma delimited list of fields to return. | 

### Return type

[**[]Role**](role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRules**
> []Rule ListRules(ctx, authorization, appId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **appId** | [**int32**](.md)|  | 
 **optional** | ***DefaultApiListRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enabled** | **optional.Bool**| Defaults to true. When set to &#x60;false&#x60; will return all disabled rules. | 
 **hasCondition** | **optional.String**| Filters Rules based on their Conditions. | 
 **hasConditionType** | **optional.String**| Filters Rules based on their condition types. | 
 **hasAction** | **optional.String**| Filters Rules based on their Actions. | 
 **hasActionType** | **optional.String**| Filters Rules based on their action types. | 

### Return type

[**[]Rule**](rule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListScopes**
> []ListScopes200ResponseInner ListScopes(ctx, authorization, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **id** | **int32**|  | 

### Return type

[**[]ListScopes200ResponseInner**](listScopes_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsers**
> []User ListUsers(ctx, authorization, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
 **optional** | ***DefaultApiListUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **optional.Int32**| The page number of results to return. | 
 **cursor** | **optional.String**| Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **createdSince** | **optional.String**| An ISO8601 timestamp value that returns all users created after a given date &amp; time. | 
 **createdUntil** | **optional.String**| An ISO8601 timestamp value that returns all users created before a given date &amp; time. | 
 **updatedSince** | **optional.String**| An ISO8601 timestamp value that returns all users updated after a given date &amp; time. | 
 **updatedUntil** | **optional.String**| An ISO8601 timestamp value that returns all users updated before a given date &amp; time. | 
 **lastLoginSince** | **optional.String**| An ISO8601 timestamp value that returns all users that logged in after a given date &amp; time. | 
 **lastLoginUntil** | **optional.String**|  | 
 **firstname** | **optional.String**| The first name of the user | 
 **lastname** | **optional.String**| The last name of the user | 
 **email** | **optional.String**| The email address of the user | 
 **memberOf** | **optional.String**| The member_of attribute in AD | 
 **username** | **optional.String**| The username for the user | 
 **samaccountname** | **optional.String**| The AD login name for the user | 
 **directoryId** | **optional.String**| The ID in OneLogin of the Directory that the user belongs to | 
 **externalId** | **optional.String**| An external identifier that has been set on the user | 
 **appId** | **optional.String**| The ID of a OneLogin Application | 
 **userIds** | **optional.String**| A comma separated list of OneLogin User IDs | 
 **customAttributesAttributeName** | **optional.String**| The short name of a custom attribute. Note that the attribute name is prefixed with custom_attributes. | 
 **fields** | **optional.String**| A comma separated list user attributes to return. | 

### Return type

[**[]User**](user.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveClientApp**
> RemoveClientApp(ctx, authorization, id, clientAppId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **id** | **int32**|  | 
  **clientAppId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveRoleAdmins**
> RemoveRoleAdmins(ctx, body, authorization, roleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveRoleUsersRequest**](RemoveRoleUsersRequest.md)|  | 
  **authorization** | **string**|  | 
  **roleId** | [**int32**](.md)| Set to the id of the role you want to return. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveRoleUsers**
> RemoveRoleUsers(ctx, body, authorization, roleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveRoleUsersRequest**](RemoveRoleUsersRequest.md)|  | 
  **authorization** | **string**|  | 
  **roleId** | [**int32**](.md)| Set to the id of the role you want to return. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeToken**
> GenerateToken400Response RevokeToken(ctx, authorization, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
 **optional** | ***DefaultApiRevokeTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiRevokeTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RevokeTokenRequest**](RevokeTokenRequest.md)|  | 

### Return type

[**GenerateToken400Response**](generateToken_400_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetRoleApps**
> []SetRoleApps200ResponseInner SetRoleApps(ctx, body, authorization, roleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]int32**](int32.md)|  | 
  **authorization** | **string**|  | 
  **roleId** | [**int32**](.md)| Set to the id of the role you want to return. | 

### Return type

[**[]SetRoleApps200ResponseInner**](setRoleApps_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrackEvent**
> TrackEvent(ctx, body, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TrackEventRequest**](TrackEventRequest.md)|  | 
  **authorization** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccessTokenClaim**
> Id UpdateAccessTokenClaim(ctx, body, authorization, id, claimId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddAccessTokenClaimRequest**](AddAccessTokenClaimRequest.md)|  | 
  **authorization** | **string**|  | 
  **id** | **int32**|  | 
  **claimId** | **int32**|  | 

### Return type

[**Id**](id.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApp**
> Schema UpdateApp(ctx, body, authorization, appId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Schema**](Schema.md)|  | 
  **authorization** | **string**|  | 
  **appId** | [**int32**](.md)|  | 

### Return type

[**Schema**](schema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuthorizationServer**
> Id UpdateAuthorizationServer(ctx, body, authorization, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateAuthorizationServerRequest**](CreateAuthorizationServerRequest.md)|  | 
  **authorization** | **string**|  | 
  **id** | **int32**|  | 

### Return type

[**Id**](id.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClientApp**
> ClientApp UpdateClientApp(ctx, body, authorization, id, clientAppId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateClientAppRequest**](UpdateClientAppRequest.md)|  | 
  **authorization** | **string**|  | 
  **id** | **int32**|  | 
  **clientAppId** | **int32**|  | 

### Return type

[**ClientApp**](client_app.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnvironmentVariable**
> Envvar UpdateEnvironmentVariable(ctx, body, authorization, envvarId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateEnvironmentVariableRequest**](UpdateEnvironmentVariableRequest.md)|  | 
  **authorization** | **string**|  | 
  **envvarId** | [**string**](.md)| Set to the id of the Hook Environment Variable that you want to fetch. | 

### Return type

[**Envvar**](envvar.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHook**
> Hook UpdateHook(ctx, body, authorization, hookId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Hook**](Hook.md)|  | 
  **authorization** | **string**|  | 
  **hookId** | [**string**](.md)| Set to the id of the Hook that you want to return. | 

### Return type

[**Hook**](hook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMapping**
> int32 UpdateMapping(ctx, body, authorization, mappingId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Mapping**](Mapping.md)|  | 
  **authorization** | **string**|  | 
  **mappingId** | **int32**| The id of the user mapping to locate. | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRiskRule**
> RiskRule UpdateRiskRule(ctx, body, authorization, riskRuleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RiskRule**](RiskRule.md)|  | 
  **authorization** | **string**|  | 
  **riskRuleId** | [**string**](.md)|  | 

### Return type

[**RiskRule**](risk_rule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRole**
> UpdateRole200Response UpdateRole(ctx, body, authorization, roleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Role**](Role.md)|  | 
  **authorization** | **string**|  | 
  **roleId** | [**int32**](.md)| Set to the id of the role you want to return. | 

### Return type

[**UpdateRole200Response**](updateRole_200_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRule**
> RuleId UpdateRule(ctx, body, authorization, appId, ruleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Rule**](Rule.md)|  | 
  **authorization** | **string**|  | 
  **appId** | [**int32**](.md)|  | 
  **ruleId** | [**int32**](.md)| The id of the app rule to locate. | 

### Return type

[**RuleId**](rule_id.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateScope**
> Id UpdateScope(ctx, body, authorization, id, scopeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddScopeRequest**](AddScopeRequest.md)|  | 
  **authorization** | **string**|  | 
  **id** | **int32**|  | 
  **scopeId** | **int32**|  | 

### Return type

[**Id**](id.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> User UpdateUser(ctx, body, authorization, userId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**User**](User.md)|  | 
  **authorization** | **string**|  | 
  **userId** | [**int32**](.md)| Set to the id of the user that you want to return. | 
 **optional** | ***DefaultApiUpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUpdateUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mappings** | **optional.**| Controls how mappings will be applied to the user on creation. Defaults to async. | 
 **validatePolicy** | **optional.**| Will passwords validate against the User Policy? Defaults to true. | 

### Return type

[**User**](user.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyEnrollment**
> []Registration VerifyEnrollment(ctx, body, authorization, userId, registrationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VerifyEnrollmentRequest**](VerifyEnrollmentRequest.md)|  | 
  **authorization** | **string**|  | 
  **userId** | **int32**| Set to the id of the user. | 
  **registrationId** | **int32**| Set to the uuid of the registration. This was included in the response as part of the initial request in Enroll Factor. | 

### Return type

[**[]Registration**](registration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyEnrollmentVoiceProtect**
> []Registration VerifyEnrollmentVoiceProtect(ctx, authorization, userId, registrationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **userId** | **int32**| Set to the id of the user. | 
  **registrationId** | **int32**| Set to the uuid of the registration. This was included in the response as part of the initial request in Enroll Factor. | 

### Return type

[**[]Registration**](registration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyFactor**
> Status2 VerifyFactor(ctx, body, authorization, userId, verificationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VerifyFactorRequest**](VerifyFactorRequest.md)|  | 
  **authorization** | **string**|  | 
  **userId** | **int32**| Set to the id of the user. | 
  **verificationId** | **int32**| The verification_id is returned on activation of the factor or you can get the device_id using the Activate Factor API call. | 

### Return type

[**Status2**](status_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyFactorSaml**
> VerifyFactorSaml200Response VerifyFactorSaml(ctx, body, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VerifyFactorSamlRequest**](VerifyFactorSamlRequest.md)|  | 
  **authorization** | **string**|  | 

### Return type

[**VerifyFactorSaml200Response**](verifyFactorSaml_200_response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyFactorVoice**
> []VerifyFactorVoice200ResponseInner VerifyFactorVoice(ctx, authorization, userId, verificationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **userId** | **int32**| Set to the id of the user. | 
  **verificationId** | **int32**| The verification_id is returned on activation of the factor or you can get the device_id using the Activate Factor API call. | 

### Return type

[**[]VerifyFactorVoice200ResponseInner**](verifyFactorVoice_200_response_inner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

