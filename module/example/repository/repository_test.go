//Package repository_test 用於測試 demo 模組的 Repository
package repository_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/codingXiang/backend-skeleton/model"
	"github.com/codingXiang/backend-skeleton/module/example"
	"github.com/codingXiang/backend-skeleton/module/example/repository"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
)

//Suite 集成 sql mock, repository
type Suite struct {
	suite.Suite
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository example.Repository
}

//初始化 Suite
func (s *Suite) SetupSuite() {
	/*
		宣告參數
	 */
	var (
		db  *sql.DB //SQL連線
		err error   //錯誤
	)
	// 初始化 sql mock，建立 db 的 instance
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)
	// 透過 gorm 建立 mysql 的 instance
	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)
	// 設定 log 模式
	s.DB.LogMode(false)
	// 設定要測試的 repository
	s.repository = repository.NewExampleRepository(s.DB)
}

//AfterTest 用於測試完畢之後的檢查
func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

//TestStart 為測試程式進入點
func TestStart(t *testing.T) {
	suite.Run(t, new(Suite))
}

//TestRepositoryCreateDepartment 用於測試 Repository 中的 CreateDepartment
func (s *Suite) TestRepositoryCreateDepartment() {
	//設定測試資料
	var testData = &model.Department{
		ID:                   "CX001",
		Name:                 "CX",
		PreviousDepartmentID: "",
	}
	//sql-mock 開始
	s.mock.ExpectBegin()
	//sql-mock 執行 insert sql
	s.mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `departments` (`id`,`name`,`previous_department_id`) " +
			"VALUES (?,?,?)")).
		WithArgs(testData.ID, testData.Name, testData.PreviousDepartmentID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	//sql-mock 送出上述指令
	s.mock.ExpectCommit()
	//實際執行 CreateDepartment 方法
	data, err := s.repository.CreateDepartment(testData)
	//判斷 CreateDepartment 中的 sql 邏輯與 sql-mock 中定義的是否有差異
	require.NoError(s.T(), err)
	//判斷 CreateDepartment 中的執行結果是否有誤
	require.NotNil(s.T(), data)
}

//TestCreateUser 用於測試 Repository 中的 CreateUser
func (s *Suite) TestCreateUser() {
	//設定測試資料
	var testData = &model.User{
		ID:           "11634",
		Username:     "xiang",
		DepartmentID: "CX001",
		Email:        "xianglai6658@digiwin.com",
		Location:     "台中",
		EXT:          4665,
	}
	//sql-mock 開始
	s.mock.ExpectBegin()
	//sql-mock 執行 insert sql
	s.mock.ExpectExec("^INSERT INTO (.*)").WillReturnResult(sqlmock.NewResult(1, 1))
	//sql-mock 送出上述指令
	s.mock.ExpectCommit()
	//實際執行 CreateUser 方法
	data, err := s.repository.CreateUser(testData)
	//判斷 CreateUser 中的 sql 邏輯與 sql-mock 中定義的是否有差異
	require.NoError(s.T(), err)
	//判斷 CreateUser 中的執行結果是否有誤
	require.NotNil(s.T(), data)
}

//TestGetDepartment 用於測試 Repository 中的 GetDepartment
func (s *Suite) TestGetDepartment() {
	//設定測試資料
	var testData = &model.Department{
		ID:                   "CX001",
		Name:                 "CX",
		PreviousDepartmentID: "",
	}
	//sql-mock 查詢
	s.mock.ExpectQuery("^SELECT (.*)").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "previous_department_id"}).
			AddRow(testData.ID, testData.Name, testData.PreviousDepartmentID))
	s.mock.ExpectQuery("^SELECT (.*)").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "previous_department_id"}).
			AddRow(testData.ID, testData.Name, testData.PreviousDepartmentID))
	//實際執行 GetDepartment 方法
	data, err := s.repository.GetDepartment(testData)
	//判斷 GetDepartment 中的 sql 邏輯與 sql-mock 中定義的是否有差異
	require.NoError(s.T(), err)
	//判斷 GetDepartment 中的執行結果是否有誤
	require.NotNil(s.T(), data)
}

//TestGetUser 用於測試 Repository 中的 GetUser
func (s *Suite) TestGetUser() {
	//設定測試資料
	var (
		testDepartmentData = &model.Department{
			ID:                   "CX001",
			Name:                 "CX",
			PreviousDepartmentID: "",
		}
		testUserData = &model.User{
			ID:           "11634",
			Username:     "xiang",
			DepartmentID: "CX001",
			Email:        "xianglai6658@digiwin.com",
			Location:     "台中",
			EXT:          4665,
		}
	)
	//sql-mock 查詢
	s.mock.ExpectQuery("^SELECT (.*)").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "previous_department_id"}).
			AddRow(testDepartmentData.ID, testDepartmentData.Name, testDepartmentData.PreviousDepartmentID))
	s.mock.ExpectQuery("^SELECT (.*)").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "qq", "location", "ext", "department_id"}).
			AddRow(testUserData.ID, testUserData.Username, testUserData.Email, testUserData.QQ, testUserData.Location, testUserData.EXT, testUserData.DepartmentID))
	//實際執行 GetUser 方法
	data, err := s.repository.GetUser(testUserData)
	//判斷 GetUser 中的 sql 邏輯與 sql-mock 中定義的是否有差異
	require.NoError(s.T(), err)
	//判斷 GetUser 中的執行結果是否有誤
	require.NotNil(s.T(), data)
}

//TestGetUserList 用於測試 Repository 中的 GetUserList
func (s *Suite) TestGetUserList() {
	//設定測試資料
	var testData = &model.User{
		ID:           "11634",
		Username:     "xiang",
		DepartmentID: "CX001",
		Email:        "xianglai6658@digiwin.com",
		Location:     "台中",
		EXT:          4665,
	}
	//sql-mock 查詢
	s.mock.ExpectQuery("^SELECT (.*)").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "qq", "location", "ext", "department_id"}).
			AddRow(testData.ID, testData.Username, testData.Email, testData.QQ, testData.Location, testData.EXT, testData.DepartmentID))
	//實際執行 GetUserList 方法
	data, err := s.repository.GetUserList()
	//判斷 GetUserList 中的 sql 邏輯與 sql-mock 中定義的是否有差異
	require.NoError(s.T(), err)
	//判斷 GetUserList 中的執行結果是否有誤
	require.NotNil(s.T(), data)
}

//TestGetDepartmentList 用於測試 Repository 中的 GetDepartmentList
func (s *Suite) TestGetDepartmentList() {
	//設定測試資料
	var testData = &model.Department{
		ID:                   "CX001",
		Name:                 "CX",
		PreviousDepartmentID: "",
	}
	//sql-mock 查詢
	s.mock.ExpectQuery("^SELECT (.*)").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "previous_department_id"}).
			AddRow(testData.ID, testData.Name, testData.PreviousDepartmentID))
	//實際執行 GetDepartmentList 方法
	data, err := s.repository.GetDepartmentList()
	//判斷 GetDepartmentList 中的 sql 邏輯與 sql-mock 中定義的是否有差異
	require.NoError(s.T(), err)
	//判斷 GetDepartmentList 中的執行結果是否有誤
	require.NotNil(s.T(), data)
}

//TestRModifyDepartment 用於測試 Repository 中的 ModifyDepartment
func (s *Suite) TestRModifyDepartment() {
	//設定測試資料
	var testData = &model.Department{
		ID:                   "CX001",
		Name:                 "CX",
		PreviousDepartmentID: "",
	}
	//sql-mock 開始
	s.mock.ExpectBegin()
	//sql-mock 執行 update sql
	s.mock.ExpectExec("^UPDATE (.*)").WillReturnResult(sqlmock.NewResult(0, 1))
	//sql-mock 送出上述指令
	s.mock.ExpectCommit()
	//實際執行 ModifyDepartment 方法
	data, err := s.repository.ModifyDepartment(testData)
	//判斷 ModifyDepartment 中的 sql 邏輯與 sql-mock 中定義的是否有差異
	require.NoError(s.T(), err)
	//判斷 ModifyDepartment 中的執行結果是否有誤
	require.NotNil(s.T(), data)
}

//TestModifyUser 用於測試 Repository 中的 ModifyUser
func (s *Suite) TestModifyUser() {
	//設定測試資料
	var testData = &model.User{
		ID:           "11634",
		Username:     "xiang",
		DepartmentID: "CX001",
		Email:        "xianglai6658@digiwin.com",
		Location:     "台中",
		EXT:          4665,
	}
	//sql-mock 開始
	s.mock.ExpectBegin()
	//sql-mock 執行 update sql
	s.mock.ExpectExec("^UPDATE (.*)").WillReturnResult(sqlmock.NewResult(0, 1))
	//sql-mock 送出上述指令
	s.mock.ExpectCommit()
	//實際執行 ModifyUser 方法
	data, err := s.repository.ModifyUser(testData)
	//判斷 ModifyUser 中的 sql 邏輯與 sql-mock 中定義的是否有差異
	require.NoError(s.T(), err)
	//判斷 ModifyUser 中的執行結果是否有誤
	require.NotNil(s.T(), data)
}

//TestDeleteDepartment 用於測試 Repository 中的 DeleteDepartment
func (s *Suite) TestDeleteDepartment() {
	//設定測試資料
	var testData = &model.Department{
		ID:                   "CX001",
		Name:                 "CX",
		PreviousDepartmentID: "",
	}
	//sql-mock 開始
	s.mock.ExpectBegin()
	//sql-mock 執行 delete sql
	s.mock.ExpectExec("^DELETE (.*)").WillReturnResult(sqlmock.NewResult(1, 1))
	//sql-mock 送出上述指令
	s.mock.ExpectCommit()
	//實際執行 DeleteDepartment 方法
	err := s.repository.DeleteDepartment(testData)
	//判斷 DeleteDepartment 中的 sql 邏輯與 sql-mock 中定義的是否有差異
	require.NoError(s.T(), err)
}

//TestDeleteUser 用於測試 Repository 中的 DeleteUser
func (s *Suite) TestDeleteUser() {
	//設定測試資料
	var testData = &model.User{
		ID:           "11634",
		Username:     "xiang",
		DepartmentID: "CX001",
		Email:        "xianglai6658@digiwin.com",
		Location:     "台中",
		EXT:          4665,
	}
	//sql-mock 開始
	s.mock.ExpectBegin()
	//sql-mock 執行 delete sql
	s.mock.ExpectExec("^DELETE (.*)").WillReturnResult(sqlmock.NewResult(1, 1))
	//sql-mock 送出上述指令
	s.mock.ExpectCommit()
	//實際執行 DeleteUser 方法
	err := s.repository.DeleteUser(testData)
	//判斷 DeleteUser 中的 sql 邏輯與 sql-mock 中定義的是否有差異
	require.NoError(s.T(), err)
}
