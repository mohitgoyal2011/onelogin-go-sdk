/*
 * OneLogin API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0.0-alpha.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Hook struct {
	// The Hook unique ID in OneLogin.
	Id string `json:"id,omitempty"`
	// A string describing the type of hook. e.g. `pre-authentication`
	Type_ string `json:"type"`
	// Boolean to enable or disable the hook. Disabled hooks will not run.
	Disabled bool `json:"disabled"`
	// The number of seconds to allow the hook function to run before before timing out. Maximum timeout varies based on the type of hook.
	Timeout int32 `json:"timeout"`
	// Environment Variable objects that will be available via process.env.ENV_VAR_NAME in the hook code.
	EnvVars []string `json:"env_vars"`
	// The Smart Hooks supported Node.js version to execute this hook with.
	Runtime string `json:"runtime"`
	// Number of retries if execution fails.
	Retries int32 `json:"retries"`
	// An object containing NPM packages that are bundled with the hook function.
	Packages *interface{} `json:"packages"`
	// A base64 encoded string containing the javascript function code.
	Function string `json:"function"`
	// The semantic version of the content that will be injected into this hook.
	ContextVersion string `json:"context_version,omitempty"`
	// String describing the state of the hook function. When a hook is ready and disabled is false it will be executed.
	Status string `json:"status,omitempty"`
	Options *HookOptions `json:"options,omitempty"`
	// An array of objects that let you limit the execution of a hook to users in specific roles.
	Conditions []HookConditionsInner `json:"conditions,omitempty"`
	// ISO8601 format date that they hook function was created.
	CreatedAt string `json:"created_at,omitempty"`
	// ISO8601 format date that they hook function was last updated.
	UpdatedAt string `json:"updated_at,omitempty"`
}
