package model

type (
	//DepartmentModelInterface  部門模型的實作介面
	DepartmentModelInterface interface {
		GetID() string
		GetName() string
		GetPreviousDepartment() DepartmentModelInterface
	}
	//UserModelInterface User 模型的實作介面
	UserModelInterface interface {
		GetID() string
		GetUserName() string
		GetEmail() string
		GetQQ() string
		GetLocation() string
		GetEXT() int
		GetDepartment() DepartmentModelInterface
	}
)

type (
	//Department 部門模型
	Department struct {
		ID                   string      `json:"id,omitempty" gorm:"primary_key;comment:'部門編號'"`                                                //部門編號
		Name                 string      `json:"name,omitempty" gorm:"comment:'部門名稱'"`                                                          //部門名稱
		PreviousDepartment   *Department `json:"previousDepartment,omitempty" gorm:"foreignkey:PreviousDepartmentID;association_foreignkey:ID"` //上級部門
		PreviousDepartmentID string      `json:"previousDepartmentID,omitempty" gorm:"comment:'上級部門'"`                                          //上級部門 ID
	}

	//User
	User struct {
		ID           string      `json:"id,omitempty" gorm:"primary_key;comment:'使用者 ID'"`                              //使用者 ID
		Username     string      `json:"username,omitempty" gorm:"comment:'使用者名稱'"`                                     //使用者名稱
		Email        string      `json:"email,omitempty" gorm:"comment:'信箱'"`                                           //使用者信箱
		QQ           string      `json:"qq,omitempty" gorm:"comment:'QQ'"`                                              //使用者QQ
		Location     string      `location:"email,omitempty" gorm:"comment:'所在地'"`                                      //所在地
		EXT          int         `json:"ext,omitempty" gorm:"comment:'分機'"`                                             //分機
		Department   *Department `json:"department,omitempty" gorm:"foreignkey:departmentID;association_foreignkey:ID"` //部門
		DepartmentID string      `json:"departmentID,omitempty" gorm:"comment:'部門'"`                                    //部門 ID
	}
)

//GetID 用於取得部門 ID
func (d *Department) GetID() string {
	return d.ID
}

//GetName 用於取得部門名稱
func (d *Department) GetName() string {
	return d.Name
}

//GetPreiousDepartment 用於取得上級部門
func (d *Department) GetPreviousDepartment() DepartmentModelInterface {
	return d.PreviousDepartment
}

//GetID 用於取得部門 ID
func (u *User) GetID() string {
	return u.ID
}

//GetUsername 用於取得使用者名稱
func (u *User) GetUserName() string {
	return u.Username
}

//GetEmail 用於取得使用者信箱
func (u *User) GetEmail() string {
	return u.Email
}

//GetQQ 用於取得使用者 QQ
func (u *User) GetQQ() string {
	return u.QQ
}

//GetLocation 用於取得使用者所在地
func (u *User) GetLocation() string {
	return u.Location
}

//GetEXT 用於取得使用者分機
func (u *User) GetEXT() int {
	return u.EXT
}

//GetDepartment 用於取得使用者部門
func (u *User) GetDepartment() DepartmentModelInterface {
	return u.Department
}
