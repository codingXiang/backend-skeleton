package http

import (
	"github.com/astaxie/beego/validation"
	"github.com/codingXiang/backend-skeleton/model"
	"github.com/codingXiang/backend-skeleton/module/example"
	"github.com/codingXiang/backend-skeleton/module/example/delivery"
	cx "github.com/codingXiang/cxgateway/delivery"
	"github.com/codingXiang/cxgateway/pkg/e"
	"github.com/codingXiang/cxgateway/pkg/util"
	"github.com/codingXiang/go-logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ExampleHttpHandler struct {
	gateway     cx.HttpHandler
	demoService example.Service
}

func NewExampleHandler(gateway cx.HttpHandler, demoService example.Service) delivery.HttpHandler {
	var handler = &ExampleHttpHandler{
		gateway:     gateway,
		demoService: demoService,
	}
	logger.Log.Info("Setup Example Handler")
	/*
		v1 版本的 user API
	  */
	logger.Log.Debug("use routing `/v1`")
	v1 := gateway.GetApiRoute().Group("/v1")
	{
		logger.Log.Debug("use routing `/v1/department`")
		department := v1.Group("/department")
		{
			department.GET("", e.Wrapper(handler.GetDepartment))
			department.GET("/:id", e.Wrapper(handler.GetDepartment))
			department.POST("", e.Wrapper(handler.CreateDepartment))
			department.PATCH("/:id", e.Wrapper(handler.ModifyDepartment))
			department.DELETE("/:id", e.Wrapper(handler.DeleteDepartment))
		}
		logger.Log.Debug("use routing `/v1/user`")
		user := v1.Group("/user")
		{
			user.GET("", e.Wrapper(handler.GetUser))
			user.GET("/:id", e.Wrapper(handler.GetUser))
			user.POST("", e.Wrapper(handler.CreateUser))
			user.PATCH("/:id", e.Wrapper(handler.ModifyUser))
			user.DELETE("/:id", e.Wrapper(handler.DeleteUser))
		}

	}
	return handler
}

func (d *ExampleHttpHandler) CreateDepartment(c *gin.Context) error {
	var (
		valid = new(validation.Validation)
		data  = new(model.Department)
	)

	//綁定參數
	var err = c.ShouldBindWith(data, binding.JSON)
	if err != nil {
		return e.ParameterError("error parameter, please check your parameter again.")
	}

	//驗證表單資訊是否填寫充足
	valid.Required(&data.ID, "id")
	valid.Required(&data.Name, "name")
	if err := d.gateway.GetHandler().ValidValidation(valid); err != nil {
		return err
	}

	if result, err := d.demoService.CreateDepartment(data); err != nil {
		return e.UnknownError(err.Error())
	} else {
		c.JSON(e.StatusCreated("create department success", result))
		return nil
	}

}

func (d *ExampleHttpHandler) CreateUser(c *gin.Context) error {
	var (
		valid = new(validation.Validation)
		data  = new(model.User)
	)

	//綁定參數
	var err = c.ShouldBindWith(data, binding.JSON)
	if err != nil {
		return e.ParameterError("error parameter, please check your parameter again.")
	}

	//驗證表單資訊是否填寫充足
	valid.Required(&data.ID, "id")
	valid.Required(&data.Username, "name")
	valid.Required(&data.DepartmentID, "departmentID")
	valid.Email(&data.Email, "email")
	if err := d.gateway.GetHandler().ValidValidation(valid); err != nil {
		return err
	}

	if result, err := d.demoService.CreateUser(data); err != nil {
		return e.UnknownError(err.Error())
	} else {
		c.JSON(e.StatusCreated("create user success", result))
		return nil
	}
}

func (d *ExampleHttpHandler) DeleteDepartment(c *gin.Context) error {
	var (
		data = new(model.Department)
	)
	data.ID = c.Params.ByName("id")
	if err := d.demoService.DeleteDepartment(data); err != nil {
		return e.UnknownError(err.Error())
	} else {
		c.JSON(e.StatusNoContent("delete department success"))
	}
	return nil
}

func (d *ExampleHttpHandler) DeleteUser(c *gin.Context) error {
	var (
		data = new(model.User)
	)
	data.ID = c.Params.ByName("id")
	if err := d.demoService.DeleteUser(data); err != nil {
		return e.UnknownError(err.Error())
	} else {
		c.JSON(e.StatusNoContent("delete user success"))
	}
	return nil
}

func (d *ExampleHttpHandler) GetDepartment(c *gin.Context) error {
	var (
		data = new(model.Department)
	)

	if id := c.Params.ByName("id"); id == "" {
		logger.Log.Debug("id is empty")
		if result, err := d.demoService.GetDepartmentList(); err != nil {
			logger.Log.Error("get department list failed", err)
			return e.UnknownError(err.Error())
		} else {
			c.JSON(e.StatusSuccess("get department list success", result))
		}
		return nil
	} else {
		data.ID = id
		if result, err := d.demoService.GetDepartment(data); err != nil {
			return e.UnknownError(err.Error())
		} else {
			c.JSON(e.StatusSuccess(util.GetI18nData(c).GetMessage("orm.create.success", nil), result))
		}
		return nil
	}

}

func (d *ExampleHttpHandler) GetUser(c *gin.Context) error {
	var (
		data = new(model.User)
	)

	if id := c.Params.ByName("id"); id == "" {
		logger.Log.Debug("id is empty")
		if result, err := d.demoService.GetUserList(); err != nil {
			logger.Log.Error("get user list failed", err)
			return e.UnknownError(err.Error())
		} else {
			c.JSON(e.StatusSuccess("get user list success", result))
		}
		return nil
	} else {
		data.ID = id
		if result, err := d.demoService.GetUser(data); err != nil {
			return e.UnknownError(err.Error())
		} else {
			c.JSON(e.StatusSuccess("get user success", result))
		}
		return nil
	}
}

func (d *ExampleHttpHandler) ModifyDepartment(c *gin.Context) error {
	var (
		data = new(model.Department)
	)

	//綁定參數
	var err = c.ShouldBindWith(data, binding.JSON)
	if err != nil {
		return e.ParameterError("error parameter, please check your parameter again.")
	}
	data.ID = c.Params.ByName("id")

	if result, err := d.demoService.ModifyDepartment(data); err != nil {
		return e.UnknownError(err.Error())
	} else {
		c.JSON(e.StatusCreated("modify department success", result))
		return nil
	}
}

func (d *ExampleHttpHandler) ModifyUser(c *gin.Context) error {
	var (
		data = new(model.User)
	)

	//綁定參數
	var err = c.ShouldBindWith(data, binding.JSON)
	if err != nil {
		return e.ParameterError("error parameter, please check your parameter again.")
	}

	data.ID = c.Params.ByName("id")

	if result, err := d.demoService.ModifyUser(data); err != nil {
		return e.UnknownError(err.Error())
	} else {
		c.JSON(e.StatusCreated("modify user success", result))
		return nil
	}
}
