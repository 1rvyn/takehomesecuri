package graph

import (
	"context"

	"github.com/1rvyn/takehomesecuri/database"
)

type Resolver struct{}

func (r *mutationResolver) CreateEmployee(ctx context.Context, input model.EmployeeInput) (*model.EmployeeInput, error) {
	employee := &model.Employee{
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Username:     input.Username,
		Password:     input.Password,
		Email:        input.Email,
		DOB:          input.DOB,
		DepartmentID: input.DepartmentID,
		Position:     input.Position,
	}

	if err := database.Database.Db.Create(&employee).Error; err != nil {
		return nil, err
	}

	return employee, nil
}
