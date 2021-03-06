/* 
 * SMS Fusion API
 *
 * This is the SMS Fusion API
 *
 * OpenAPI spec version: 1.0.0
 * Contact: support@smsfusion.com.au
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type HlrCallback struct {

	// Unique ID for response
	Id string `json:"id,omitempty"`

	// The MSISDN of the number queried
	Msisdn int64 `json:"msisdn,omitempty"`

	// Short status code of the response
	Status string `json:"status,omitempty"`

	// ISO 3166-2 country code
	Country string `json:"country,omitempty"`

	// The operator attached to the MSISDN
	Operator string `json:"operator,omitempty"`

	// MCC and MNC of MSDISDN
	Mccmnc int32 `json:"mccmnc,omitempty"`

	// If the response code is temporary or permenant
	Duration string `json:"duration,omitempty"`

	// Full status code of the response
	Message string `json:"message,omitempty"`

	// Extended explanation of the status code
	Extended string `json:"extended,omitempty"`

	// Cost of the response
	Cost float32 `json:"cost,omitempty"`
}
