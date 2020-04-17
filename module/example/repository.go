//demo 模組 module 的 Repository Interface
package example

import (
	"github.com/codingXiang/backend-skeleton/model"
)

//Repository 用於與資料庫進行存取的封裝方法
//go:generate mockgen -destination mock/mock_repository.go -package mock -source repository.go
type Repository interface {
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
