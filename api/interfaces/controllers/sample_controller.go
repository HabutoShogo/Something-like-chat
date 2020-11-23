package controllers

import (
	"strconv"
	"fmt"

	// "sample/domain"
	"sample/interfaces/database"
	"sample/usecase"

	"github.com/labstack/echo"
)

type SampleController struct {
	Interactor usecase.SampleInteractor
}


func NewSampleController(sqlHandler database.SqlHandler) *SampleController {
	return &SampleController{
		Interactor: usecase.SampleInteractor{
			SampleRepository: &database.SampleRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *SampleController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	fmt.Println(id)
	fmt.Println(user)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}

func (controller *SampleController) Index(c echo.Context) (err error) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
	return
}