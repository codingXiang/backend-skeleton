//demo 模組 module 的 Service Interface
package example

import "github.com/codingXiang/backend-skeleton/model"

//Service 用於封裝商業邏輯的方法
//go:generate mockgen -destination mock/mock_service.go -package mock -source service.go
type Service interface {
	CreateDepartment(data model.DepartmentInterface) (*model.Department, error)
	CreateUser(data model.UserInterface) (*model.User, error)
	GetDepartmentList() ([]*model.Department, error)
	GetDepartment(data model.DepartmentInterface) (*model.Department, error)
	GetUserList() ([]*model.User, error)
	GetUser(data model.UserInterface) (*model.User, error)
	ModifyDepartment(data model.DepartmentInterface) (*model.Department, error)
	ModifyUser(data model.UserInterface) (*model.User, error)
	DeleteDepartment(data model.DepartmentInterface) (error)
	DeleteUser(data model.UserInterface) (error)
}