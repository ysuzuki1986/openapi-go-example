package main

import (
	"context"
	"log"
	openapi "openapi-example/go-client-api"
)

func main() {
	log.Printf("Client start")
	cfg := openapi.NewConfiguration()
	api_client := openapi.NewAPIClient(cfg)
	ctx := context.Background()
	employeeId := "20"
	employee, _, _ := api_client.EmployeeApi.GetEmployee(ctx, employeeId)
	log.Printf("%s, %s, %s, %d\n", employee.FirstName, employee.LastName, employee.Id, employee.Salary)
	log.Printf("Client end")
}
