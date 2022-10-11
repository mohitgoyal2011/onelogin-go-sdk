/*
 * OneLogin API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0.0-alpha.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EnrollFactorRequest struct {
	// The identifier of the factor to enroll the user with.
	FactorId int32 `json:"factor_id"`
	// A name for the users device.
	DisplayName string `json:"display_name"`
	// Defaults to 120. Valid values are: 120-900 seconds.
	ExpiresIn string `json:"expires_in,omitempty"`
	// Defaults to false.
	Verified bool `json:"verified,omitempty"`
	// Redirects MagicLink success page to specified URL after 2 seconds.
	RedirectTo string `json:"redirect_to,omitempty"`
	// A message template that will be sent via SMS. Max length of the message after template items are inserted is 160 characters including the OTP code.
	CustomMessage string `json:"custom_message,omitempty"`
}
