/*
 * Package main ORY Keto
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.sh
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type UpsertOryAccessControlPolicy struct {
	Body OryAccessControlPolicy `json:"Body,omitempty"`

	// The ORY Access Control Policy flavor. Can be \"regex\", \"glob\", and \"exact\".  in: path
	Flavor string `json:"flavor"`
}
