package http

import (
	"github.com/astaxie/beego/validation"
	"github.com/codingXiang/backend-skeleton/model"
	"github.com/codingXiang/backend-skeleton/module/demo"
	"github.com/codingXiang/backend-skeleton/module/demo/delivery"
	cx "github.com/codingXiang/cxgateway/delivery"
	"github.com/codingXiang/cxgateway/pkg/e"
	"github.com/codingXiang/go-logger"
	"github.com/gin-gonic/gin"
)

type DemoHttpHandler struct {
	gateway     cx.HttpHandler
	demoService demo.Service
}

func NewDemoHandler(gateway cx.HttpHandler, demoService demo.Service) delivery.HttpHandler {
	var handler = &DemoHttpHandler{
		gateway:     gateway,
		demoService: demoService,
	}
	logger.Log.Info("Setup Demo Handler")
	/*
		v1 版本的 user API
	  */
	logger.Log.Debug("use routing `/v1`")
	v1 := gateway.GetApiRoute().Group("/v1")
	{
		// demo 相關的 routing
		logger.Log.Debug("use routing `/v1`")

		// app 相關的 routing
		logger.Log.Debug("use routing `/v1/demo`")
		demo := v1.Group("/demo")
		{
			department := demo.Group("/department")
			{
				department.GET("", e.Wrapper(handler.GetDepartment))
				department.GET("/:id", e.Wrapper(handler.GetDepartment))
				department.POST("", e.Wrapper(handler.CreateDepartment))
				department.PATCH("/:id", e.Wrapper(handler.ModifyDepartment))
				department.DELETE("/:id", e.Wrapper(handler.DeleteDepartment))
			}

			user := demo.Group("/user")
			{
				user.GET("", e.Wrapper(handler.GetUser))
				user.GET("/:id", e.Wrapper(handler.GetUser))
				user.POST("", e.Wrapper(handler.CreateUser))
				user.PATCH("/:id", e.Wrapper(handler.ModifyUser))
				user.DELETE("/:id", e.Wrapper(handler.DeleteUser))
			}
		}
	}
	return handler
}

func (d *DemoHttpHandler) CreateDepartment(c *gin.Context) error {
	var (
		valid = new(validation.Validation)
		data  *model.Department
	)

	//綁定參數
	if err := d.gateway.GetHandler().BindBody(c, &data); err != nil {
		return err
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

func (d *DemoHttpHandler) CreateUser(c *gin.Context) error {
	var (
		valid = new(validation.Validation)
		data *model.User
	)

	//綁定參數
	if err := d.gateway.GetHandler().BindBody(c, &data); err != nil {
		return err
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

func (d *DemoHttpHandler) DeleteDepartment(c *gin.Context) error {
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

func (d *DemoHttpHandler) DeleteUser(c *gin.Context) error {
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

func (d *DemoHttpHandler) GetDepartment(c *gin.Context) error {
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
			c.JSON(e.StatusSuccess("get department success", result))
		}
		return nil
	}

}

func (d *DemoHttpHandler) GetUser(c *gin.Context) error {
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

func (d *DemoHttpHandler) ModifyDepartment(c *gin.Context) error {
	var (
		data *model.Department
	)

	//綁定參數
	if err := d.gateway.GetHandler().BindBody(c, &data); err != nil {
		return err
	}
	data.ID = c.Params.ByName("id")

	if result, err := d.demoService.ModifyDepartment(data); err != nil {
		return e.UnknownError(err.Error())
	} else {
		c.JSON(e.StatusCreated("modify department success", result))
		return nil
	}
}

func (d *DemoHttpHandler) ModifyUser(c *gin.Context) error {
	var (
		data *model.User
	)

	//綁定參數
	if err := d.gateway.GetHandler().BindBody(c, &data); err != nil {
		return err
	}
	data.ID = c.Params.ByName("id")

	if result, err := d.demoService.ModifyUser(data); err != nil {
		return e.UnknownError(err.Error())
	} else {
		c.JSON(e.StatusCreated("modify user success", result))
		return nil
	}
}
