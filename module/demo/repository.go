package demo

import (
	. "github.com/codingXiang/backend-skeleton/model"
)

type Repository interface {
	CreateDepartment(data *Department) (DepartmentModelInterface, error)
	CreateUser(data *User) (UserModelInterface, error)
	GetDepartmentList() ([]DepartmentModelInterface, error)
	GetDepartment(data *Department) (DepartmentModelInterface, error)
	GetUserList() ([]UserModelInterface, error)
	GetUser(data *User) (UserModelInterface, error)
	ModifyDepartment(data *Department) (DepartmentModelInterface, error)
	ModifyUser(data *User) (UserModelInterface, error)
	DeleteDepartment(data *Department) (error)
	DeleteUser(data *User) (error)
}
