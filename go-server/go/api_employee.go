/*
 * oas3_go_flutter_exam Employee
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
  "encoding/json"
  "net/http"
)

// GetEmployee - get employee information
func GetEmployee(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  employeeID := r.URL.Query()["employeeId"]

  employee := Employee{
    Id:        employeeID[0],
    FirstName: "user-" + employeeID[0],
    LastName:  "test- " + employeeID[0],
    Salary:    12000000}
  json.NewEncoder(w).Encode(employee)
}
