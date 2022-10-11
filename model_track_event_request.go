/*
 * OneLogin API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0.0-alpha.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TrackEventRequest struct {
	// Verbs are used to distinguish between different types of events.
	Verb string `json:"verb"`
	// The IP address of the User's request.
	Ip string `json:"ip"`
	// The user agent of the User's request.
	UserAgent string `json:"user_agent"`
	User *RiskUser `json:"user"`
	Source *Source `json:"source,omitempty"`
	Session *Session `json:"session,omitempty"`
	Device *RiskDevice `json:"device,omitempty"`
	// Set to the value of the __tdli_fp cookie.
	Fp string `json:"fp,omitempty"`
	// Date and time of the event in IS08601 format. Useful for preloading old events. Defaults to date time this API request is received.
	Published string `json:"published,omitempty"`
}
