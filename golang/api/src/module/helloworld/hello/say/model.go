// Package hellosay models
package hellosay

/*
// RequestBody specifies RequestBody
type RequestBody struct {

}
*/

// Sample specifies repo/select_sample.sql model respose
type Sample struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Grettings string `db:"grettings" json:"grettings"`
}

// ResponseBody specifies the response body
type ResponseBody struct {
	Message string `json:"message"`
}
