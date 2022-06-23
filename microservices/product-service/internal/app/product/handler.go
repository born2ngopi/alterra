package product

import (
	"fmt"
	"strconv"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) Get(c echo.Context) error {

	payload := new(dto.SearchGetRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Find(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result.Datas, "Get users success", &result.PaginationInfo).Send(c)
}

func (h *handler) GetByID(c echo.Context) error {

	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	result, err := h.service.FindByID(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *handler) Create(c echo.Context) error {

	payload := new(dto.CreateProductRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	result, err := h.service.Create(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)

}

func (h *handler) Update(c echo.Context) error {
	payload := new(dto.UpdateProductRequest)
	if err := c.Bind(payload); err != nil {
		fmt.Println("bind", err.Error())
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		fmt.Println("validate", err.Error())
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	strID := c.Param("id")
	ID, err := strconv.Atoi(strID)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	// param := new(dto.ByIDRequest)
	// if err := c.Bind(param); err != nil {
	// 	fmt.Println("bind id", err.Error())
	// 	return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	// }
	// if err := c.Validate(param); err != nil {
	// 	response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
	// 	return response.Send(c)
	// }

	result, err := h.service.Update(c.Request().Context(), uint(ID), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *handler) Delete(c echo.Context) error {

	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	result, err := h.service.Delete(c.Request().Context(), payload.ID)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}
