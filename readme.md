# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0.0-alpha.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/onelogin/onelogin-go-sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://onelogininc.onelogin.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**ActivateFactor**](docs/DefaultApi.md#activatefactor) | **Post** /api/2/mfa/users/{user_id}/verifications | 
*DefaultApi* | [**AddAccessTokenClaim**](docs/DefaultApi.md#addaccesstokenclaim) | **Post** /api/2/api_authorizations/{id}/claims | 
*DefaultApi* | [**AddClientApp**](docs/DefaultApi.md#addclientapp) | **Post** /api/2/api_authorizations/{id}/clients | 
*DefaultApi* | [**AddRoleAdmins**](docs/DefaultApi.md#addroleadmins) | **Post** /api/2/roles/{role_id}/admins | 
*DefaultApi* | [**AddRoleUsers**](docs/DefaultApi.md#addroleusers) | **Post** /api/2/roles/{role_id}/users | 
*DefaultApi* | [**AddScope**](docs/DefaultApi.md#addscope) | **Post** /api/2/api_authorizations/{id}/scopes | 
*DefaultApi* | [**BulkMappingSort**](docs/DefaultApi.md#bulkmappingsort) | **Put** /api/2/apps/mappings/sort | 
*DefaultApi* | [**BulkSort**](docs/DefaultApi.md#bulksort) | **Put** /api/2/apps/{app_id}/rules/sort | 
*DefaultApi* | [**CreateApp**](docs/DefaultApi.md#createapp) | **Post** /api/2/apps | 
*DefaultApi* | [**CreateAuthorizationServer**](docs/DefaultApi.md#createauthorizationserver) | **Post** /api/2/api_authorizations | 
*DefaultApi* | [**CreateEnvironmentVariable**](docs/DefaultApi.md#createenvironmentvariable) | **Post** /api/2/hooks/envs | 
*DefaultApi* | [**CreateHook**](docs/DefaultApi.md#createhook) | **Post** /api/2/hooks | 
*DefaultApi* | [**CreateMapping**](docs/DefaultApi.md#createmapping) | **Post** /api/2/mappings | 
*DefaultApi* | [**CreateRiskRule**](docs/DefaultApi.md#createriskrule) | **Post** /api/2/risk/rules | 
*DefaultApi* | [**CreateRoles**](docs/DefaultApi.md#createroles) | **Post** /api/2/roles | 
*DefaultApi* | [**CreateRule**](docs/DefaultApi.md#createrule) | **Post** /api/2/apps/{app_id}/rules | 
*DefaultApi* | [**CreateUser**](docs/DefaultApi.md#createuser) | **Post** /api/2/users | 
*DefaultApi* | [**DeleteAccessTokenClaim**](docs/DefaultApi.md#deleteaccesstokenclaim) | **Delete** /api/2/api_authorizations/{id}/claims/{claim_id} | 
*DefaultApi* | [**DeleteApp**](docs/DefaultApi.md#deleteapp) | **Delete** /api/2/apps/{app_id} | 
*DefaultApi* | [**DeleteAppParameter**](docs/DefaultApi.md#deleteappparameter) | **Delete** /api/2/apps/{app_id}/parameters/{parameter_id} | 
*DefaultApi* | [**DeleteAuthorizationServer**](docs/DefaultApi.md#deleteauthorizationserver) | **Delete** /api/2/api_authorizations/{id} | 
*DefaultApi* | [**DeleteEnvironmentVariable**](docs/DefaultApi.md#deleteenvironmentvariable) | **Delete** /api/2/hooks/envs/{envvar_id} | 
*DefaultApi* | [**DeleteFactor**](docs/DefaultApi.md#deletefactor) | **Delete** /api/2/mfa/users/{user_id}/devices/{device_id} | 
*DefaultApi* | [**DeleteHook**](docs/DefaultApi.md#deletehook) | **Delete** /api/2/hooks/{hook_id} | 
*DefaultApi* | [**DeleteMapping**](docs/DefaultApi.md#deletemapping) | **Delete** /api/2/mappings/{mapping_id} | 
*DefaultApi* | [**DeleteRiskRule**](docs/DefaultApi.md#deleteriskrule) | **Delete** /api/2/risk/rules/{risk_rule_id} | 
*DefaultApi* | [**DeleteRole**](docs/DefaultApi.md#deleterole) | **Delete** /api/2/roles/{role_id} | 
*DefaultApi* | [**DeleteRule**](docs/DefaultApi.md#deleterule) | **Delete** /api/2/apps/{app_id}/rules/{rule_id} | 
*DefaultApi* | [**DeleteScope**](docs/DefaultApi.md#deletescope) | **Delete** /api/2/api_authorizations/{id}/scopes/{scope_id} | 
*DefaultApi* | [**DeleteUser**](docs/DefaultApi.md#deleteuser) | **Delete** /api/2/users/{user_id} | 
*DefaultApi* | [**DryRunMapping**](docs/DefaultApi.md#dryrunmapping) | **Post** /api/2/mappings/{mapping_id}/dryrun | 
*DefaultApi* | [**EnrollFactor**](docs/DefaultApi.md#enrollfactor) | **Post** /api/2/mfa/users/{user_id}/registrations | 
*DefaultApi* | [**GenerateMfaToken**](docs/DefaultApi.md#generatemfatoken) | **Post** /api/2/mfs/users/{user_id}/mfa_token | 
*DefaultApi* | [**GenerateSamlAssertion**](docs/DefaultApi.md#generatesamlassertion) | **Post** /api/2/saml_assertion | 
*DefaultApi* | [**GenerateToken**](docs/DefaultApi.md#generatetoken) | **Post** /auth/oauth2/v2/token | 
*DefaultApi* | [**GetApp**](docs/DefaultApi.md#getapp) | **Get** /api/2/apps/{app_id} | 
*DefaultApi* | [**GetAuthorizationServer**](docs/DefaultApi.md#getauthorizationserver) | **Get** /api/2/api_authorizations/{id} | 
*DefaultApi* | [**GetAvailableFactors**](docs/DefaultApi.md#getavailablefactors) | **Get** /api/2/mfa/users/{user_id}/factors | 
*DefaultApi* | [**GetClientApps**](docs/DefaultApi.md#getclientapps) | **Get** /api/2/api_authorizations/{id}/clients | 
*DefaultApi* | [**GetEnrolledFactors**](docs/DefaultApi.md#getenrolledfactors) | **Get** /api/2/mfa/users/{user_id}/devices | 
*DefaultApi* | [**GetEnvironmentVariable**](docs/DefaultApi.md#getenvironmentvariable) | **Get** /api/2/hooks/envs/{envvar_id} | 
*DefaultApi* | [**GetHook**](docs/DefaultApi.md#gethook) | **Get** /api/2/hooks/{hook_id} | 
*DefaultApi* | [**GetLogs**](docs/DefaultApi.md#getlogs) | **Get** /api/2/hooks/{hook_id}/logs | 
*DefaultApi* | [**GetMapping**](docs/DefaultApi.md#getmapping) | **Get** /api/2/mappings/{mapping_id} | 
*DefaultApi* | [**GetRateLimit**](docs/DefaultApi.md#getratelimit) | **Get** /auth/rate_limit | 
*DefaultApi* | [**GetRiskRule**](docs/DefaultApi.md#getriskrule) | **Get** /api/2/risk/rules/{risk_rule_id} | 
*DefaultApi* | [**GetRiskScore**](docs/DefaultApi.md#getriskscore) | **Post** /api/2/risk/verify | 
*DefaultApi* | [**GetRole**](docs/DefaultApi.md#getrole) | **Get** /api/2/roles/{role_id} | 
*DefaultApi* | [**GetRoleAdmins**](docs/DefaultApi.md#getroleadmins) | **Get** /api/2/roles/{role_id}/admins | 
*DefaultApi* | [**GetRoleApps**](docs/DefaultApi.md#getroleapps) | **Get** /api/2/roles/{role_id}/apps | 
*DefaultApi* | [**GetRoleUsers**](docs/DefaultApi.md#getroleusers) | **Get** /api/2/roles/{role_id}/users | 
*DefaultApi* | [**GetRule**](docs/DefaultApi.md#getrule) | **Get** /api/2/apps/{app_id}/rules/{rule_id} | 
*DefaultApi* | [**GetScoreInsights**](docs/DefaultApi.md#getscoreinsights) | **Get** /api/2/risk/scores | 
*DefaultApi* | [**GetUser**](docs/DefaultApi.md#getuser) | **Get** /api/2/users/{user_id} | 
*DefaultApi* | [**GetUserApps**](docs/DefaultApi.md#getuserapps) | **Get** /api/2/users/{user_id}/apps | 
*DefaultApi* | [**ListAccessTokenClaims**](docs/DefaultApi.md#listaccesstokenclaims) | **Get** /api/2/api_authorizations/{id}/claims | 
*DefaultApi* | [**ListActionValues**](docs/DefaultApi.md#listactionvalues) | **Get** /api/2/apps/{app_id}/rules/actions/{actuion_value}/values | 
*DefaultApi* | [**ListActions**](docs/DefaultApi.md#listactions) | **Get** /api/2/apps/{app_id}/rules/actions | 
*DefaultApi* | [**ListAppUsers**](docs/DefaultApi.md#listappusers) | **Get** /api/2/apps/{app_id}/users | 
*DefaultApi* | [**ListApps**](docs/DefaultApi.md#listapps) | **Get** /api/2/apps | 
*DefaultApi* | [**ListAuthorizationServers**](docs/DefaultApi.md#listauthorizationservers) | **Get** /api/2/api_authorizations | 
*DefaultApi* | [**ListConditionOperators**](docs/DefaultApi.md#listconditionoperators) | **Get** /api/2/apps/{app_id}/rules/conditions/{condition_value}/operators | 
*DefaultApi* | [**ListConditionValues**](docs/DefaultApi.md#listconditionvalues) | **Get** /api/2/apps/{app_id}/rules/conditions/{condition_value}/values | 
*DefaultApi* | [**ListConditions**](docs/DefaultApi.md#listconditions) | **Get** /api/2/apps/{app_id}/rules/conditions | 
*DefaultApi* | [**ListConnectors**](docs/DefaultApi.md#listconnectors) | **Get** /api/2/connectors | 
*DefaultApi* | [**ListEnvironmentVariables**](docs/DefaultApi.md#listenvironmentvariables) | **Get** /api/2/hooks/envs | 
*DefaultApi* | [**ListHooks**](docs/DefaultApi.md#listhooks) | **Get** /api/2/hooks | 
*DefaultApi* | [**ListMappingActionValues**](docs/DefaultApi.md#listmappingactionvalues) | **Get** /api/2/apps/mappings/actions/{actuion_value}/values | 
*DefaultApi* | [**ListMappingActions**](docs/DefaultApi.md#listmappingactions) | **Get** /api/2/apps/mappings/actions | 
*DefaultApi* | [**ListMappingConditionOperators**](docs/DefaultApi.md#listmappingconditionoperators) | **Get** /api/2/apps/mappings/conditions/{condition_value}/operators | 
*DefaultApi* | [**ListMappingConditionValues**](docs/DefaultApi.md#listmappingconditionvalues) | **Get** /api/2/apps/mappings/conditions/{condition_value}/values | 
*DefaultApi* | [**ListMappingConditions**](docs/DefaultApi.md#listmappingconditions) | **Get** /api/2/apps/mappings/conditions | 
*DefaultApi* | [**ListMappings**](docs/DefaultApi.md#listmappings) | **Get** /api/2/mappings | 
*DefaultApi* | [**ListRiskRules**](docs/DefaultApi.md#listriskrules) | **Get** /api/2/risk/rules | 
*DefaultApi* | [**ListRoles**](docs/DefaultApi.md#listroles) | **Get** /api/2/roles | 
*DefaultApi* | [**ListRules**](docs/DefaultApi.md#listrules) | **Get** /api/2/apps/{app_id}/rules | 
*DefaultApi* | [**ListScopes**](docs/DefaultApi.md#listscopes) | **Get** /api/2/api_authorizations/{id}/scopes | 
*DefaultApi* | [**ListUsers**](docs/DefaultApi.md#listusers) | **Get** /api/2/users | 
*DefaultApi* | [**RemoveClientApp**](docs/DefaultApi.md#removeclientapp) | **Delete** /api/2/api_authorizations/{id}/clients/{client_app_id} | 
*DefaultApi* | [**RemoveRoleAdmins**](docs/DefaultApi.md#removeroleadmins) | **Delete** /api/2/roles/{role_id}/admins | 
*DefaultApi* | [**RemoveRoleUsers**](docs/DefaultApi.md#removeroleusers) | **Delete** /api/2/roles/{role_id}/users | 
*DefaultApi* | [**RevokeToken**](docs/DefaultApi.md#revoketoken) | **Post** /auth/oauth2/revoke | 
*DefaultApi* | [**SetRoleApps**](docs/DefaultApi.md#setroleapps) | **Put** /api/2/roles/{role_id}/apps | 
*DefaultApi* | [**TrackEvent**](docs/DefaultApi.md#trackevent) | **Post** /api/2/risk/events | 
*DefaultApi* | [**UpdateAccessTokenClaim**](docs/DefaultApi.md#updateaccesstokenclaim) | **Put** /api/2/api_authorizations/{id}/claims/{claim_id} | 
*DefaultApi* | [**UpdateApp**](docs/DefaultApi.md#updateapp) | **Put** /api/2/apps/{app_id} | 
*DefaultApi* | [**UpdateAuthorizationServer**](docs/DefaultApi.md#updateauthorizationserver) | **Put** /api/2/api_authorizations/{id} | 
*DefaultApi* | [**UpdateClientApp**](docs/DefaultApi.md#updateclientapp) | **Put** /api/2/api_authorizations/{id}/clients/{client_app_id} | 
*DefaultApi* | [**UpdateEnvironmentVariable**](docs/DefaultApi.md#updateenvironmentvariable) | **Put** /api/2/hooks/envs/{envvar_id} | 
*DefaultApi* | [**UpdateHook**](docs/DefaultApi.md#updatehook) | **Put** /api/2/hooks/{hook_id} | 
*DefaultApi* | [**UpdateMapping**](docs/DefaultApi.md#updatemapping) | **Put** /api/2/mappings/{mapping_id} | 
*DefaultApi* | [**UpdateRiskRule**](docs/DefaultApi.md#updateriskrule) | **Put** /api/2/risk/rules/{risk_rule_id} | 
*DefaultApi* | [**UpdateRole**](docs/DefaultApi.md#updaterole) | **Put** /api/2/roles/{role_id} | 
*DefaultApi* | [**UpdateRule**](docs/DefaultApi.md#updaterule) | **Put** /api/2/apps/{app_id}/rules/{rule_id} | 
*DefaultApi* | [**UpdateScope**](docs/DefaultApi.md#updatescope) | **Put** /api/2/api_authorizations/{id}/scopes/{scope_id} | 
*DefaultApi* | [**UpdateUser**](docs/DefaultApi.md#updateuser) | **Put** /api/2/users/{user_id} | 
*DefaultApi* | [**VerifyEnrollment**](docs/DefaultApi.md#verifyenrollment) | **Put** /api/2/mfa/users/{user_id}/registrations/{registration_id} | 
*DefaultApi* | [**VerifyEnrollmentVoiceProtect**](docs/DefaultApi.md#verifyenrollmentvoiceprotect) | **Get** /api/2/mfa/users/{user_id}/registrations/{registration_id} | 
*DefaultApi* | [**VerifyFactor**](docs/DefaultApi.md#verifyfactor) | **Put** /api/2/mfa/users/{user_id}/verifications/{verification_id} | 
*DefaultApi* | [**VerifyFactorSaml**](docs/DefaultApi.md#verifyfactorsaml) | **Post** /api/2/saml_assertion/verify_factor | 
*DefaultApi* | [**VerifyFactorVoice**](docs/DefaultApi.md#verifyfactorvoice) | **Get** /api/2/mfa/users/{user_id}/verifications/{verification_id} | 


## Documentation For Models

 - [Action](docs/Action.md)
 - [ActivateFactorRequest](docs/ActivateFactorRequest.md)
 - [AddAccessTokenClaimRequest](docs/AddAccessTokenClaimRequest.md)
 - [AddClientAppRequest](docs/AddClientAppRequest.md)
 - [AddRoleUsers200ResponseInner](docs/AddRoleUsers200ResponseInner.md)
 - [AddScopeRequest](docs/AddScopeRequest.md)
 - [AuthMethod](docs/AuthMethod.md)
 - [AuthServerConfiguration](docs/AuthServerConfiguration.md)
 - [ClientApp](docs/ClientApp.md)
 - [Condition](docs/Condition.md)
 - [Connector](docs/Connector.md)
 - [CreateAuthorizationServerRequest](docs/CreateAuthorizationServerRequest.md)
 - [CreateEnvironmentVariableRequest](docs/CreateEnvironmentVariableRequest.md)
 - [CreateRoles201ResponseInner](docs/CreateRoles201ResponseInner.md)
 - [Device](docs/Device.md)
 - [EnrollFactorRequest](docs/EnrollFactorRequest.md)
 - [Envvar](docs/Envvar.md)
 - [ErrorStatus](docs/ErrorStatus.md)
 - [ErrorStatusErrorsInner](docs/ErrorStatusErrorsInner.md)
 - [FactorInner](docs/FactorInner.md)
 - [FactorInnerFactorData](docs/FactorInnerFactorData.md)
 - [GenerateMfaToken200Response](docs/GenerateMfaToken200Response.md)
 - [GenerateMfaToken422Response](docs/GenerateMfaToken422Response.md)
 - [GenerateMfaToken422ResponseDetails](docs/GenerateMfaToken422ResponseDetails.md)
 - [GenerateMfaTokenRequest](docs/GenerateMfaTokenRequest.md)
 - [GenerateSamlAssertionRequest](docs/GenerateSamlAssertionRequest.md)
 - [GenerateToken200Response](docs/GenerateToken200Response.md)
 - [GenerateToken400Response](docs/GenerateToken400Response.md)
 - [GenerateTokenRequest](docs/GenerateTokenRequest.md)
 - [GetAuthorizationServer200Response](docs/GetAuthorizationServer200Response.md)
 - [GetAvailableFactors200ResponseInner](docs/GetAvailableFactors200ResponseInner.md)
 - [GetClientApps200ResponseInner](docs/GetClientApps200ResponseInner.md)
 - [GetClientApps200ResponseInnerScopesInner](docs/GetClientApps200ResponseInnerScopesInner.md)
 - [GetRateLimit200Response](docs/GetRateLimit200Response.md)
 - [GetRateLimit200ResponseData](docs/GetRateLimit200ResponseData.md)
 - [GetRiskScore200Response](docs/GetRiskScore200Response.md)
 - [GetRiskScore400Response](docs/GetRiskScore400Response.md)
 - [GetRiskScoreRequest](docs/GetRiskScoreRequest.md)
 - [GetScoreInsights200Response](docs/GetScoreInsights200Response.md)
 - [GetScoreInsights200ResponseScores](docs/GetScoreInsights200ResponseScores.md)
 - [GetUserApps200ResponseInner](docs/GetUserApps200ResponseInner.md)
 - [Hook](docs/Hook.md)
 - [HookConditionsInner](docs/HookConditionsInner.md)
 - [HookOptions](docs/HookOptions.md)
 - [HookStatus](docs/HookStatus.md)
 - [Id](docs/Id.md)
 - [ListAccessTokenClaims200ResponseInner](docs/ListAccessTokenClaims200ResponseInner.md)
 - [ListActions200ResponseInner](docs/ListActions200ResponseInner.md)
 - [ListAppUsers200ResponseInner](docs/ListAppUsers200ResponseInner.md)
 - [ListAuthorizationServers200ResponseInner](docs/ListAuthorizationServers200ResponseInner.md)
 - [ListAuthorizationServers200ResponseInnerConfiguration](docs/ListAuthorizationServers200ResponseInnerConfiguration.md)
 - [ListConditionOperators200ResponseInner](docs/ListConditionOperators200ResponseInner.md)
 - [ListConditionValues200ResponseInner](docs/ListConditionValues200ResponseInner.md)
 - [ListConditions200ResponseInner](docs/ListConditions200ResponseInner.md)
 - [ListMappingConditionOperators200ResponseInner](docs/ListMappingConditionOperators200ResponseInner.md)
 - [ListMappingConditions200ResponseInner](docs/ListMappingConditions200ResponseInner.md)
 - [ListScopes200ResponseInner](docs/ListScopes200ResponseInner.md)
 - [Log](docs/Log.md)
 - [Mapping](docs/Mapping.md)
 - [Registration](docs/Registration.md)
 - [RemoveRoleUsersRequest](docs/RemoveRoleUsersRequest.md)
 - [RevokeTokenRequest](docs/RevokeTokenRequest.md)
 - [RiskDevice](docs/RiskDevice.md)
 - [RiskRule](docs/RiskRule.md)
 - [RiskUser](docs/RiskUser.md)
 - [Role](docs/Role.md)
 - [Rule](docs/Rule.md)
 - [RuleId](docs/RuleId.md)
 - [Schema](docs/Schema.md)
 - [Schema1](docs/Schema1.md)
 - [Schema1AddedBy](docs/Schema1AddedBy.md)
 - [SchemaProvisioning](docs/SchemaProvisioning.md)
 - [Session](docs/Session.md)
 - [SetRoleApps200ResponseInner](docs/SetRoleApps200ResponseInner.md)
 - [Source](docs/Source.md)
 - [Status](docs/Status.md)
 - [Status1](docs/Status1.md)
 - [Status2](docs/Status2.md)
 - [Status2Status](docs/Status2Status.md)
 - [TrackEventRequest](docs/TrackEventRequest.md)
 - [UpdateAuthorizationServer400Response](docs/UpdateAuthorizationServer400Response.md)
 - [UpdateClientAppRequest](docs/UpdateClientAppRequest.md)
 - [UpdateEnvironmentVariableRequest](docs/UpdateEnvironmentVariableRequest.md)
 - [UpdateRole200Response](docs/UpdateRole200Response.md)
 - [User](docs/User.md)
 - [VerifyEnrollmentRequest](docs/VerifyEnrollmentRequest.md)
 - [VerifyFactorRequest](docs/VerifyFactorRequest.md)
 - [VerifyFactorSaml200Response](docs/VerifyFactorSaml200Response.md)
 - [VerifyFactorSamlRequest](docs/VerifyFactorSamlRequest.md)
 - [VerifyFactorVoice200ResponseInner](docs/VerifyFactorVoice200ResponseInner.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



