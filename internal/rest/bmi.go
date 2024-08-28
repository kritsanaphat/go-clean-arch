package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/labstack/echo/v4"
)

type ResponseErrorBMI struct {
	Message string `json:"message"`
}

type ResponseResultBMI struct {
	Message string `json:"message"`
}
type BMIService interface {
	// Fetch(ctx context.Context, cursor string, num int64) ([]domain.Article, string, error)
	// GetByID(ctx context.Context, id int64) (domain.Article, error)
	// Update(ctx context.Context, ar *domain.Article) error
	// GetByTitle(ctx context.Context, title string) (domain.Article, error)
	StoreBMILog(context.Context, *domain.BMI) error
	// Delete(ctx context.Context, id int64) error
}

// ArticleHandler  represent the httphandler for article
type BMIHandler struct {
	Service BMIService
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewBMIHandler(e *echo.Echo, svc BMIService) {
	handler := &BMIHandler{
		Service: svc,
	}
	e.POST("/bmi_calculate", handler.CalculateBMI)

}

func (a *BMIHandler) CalculateBMI(c echo.Context) error {

	weight := c.QueryParam("weight")
	num_weight, err := strconv.Atoi(weight)
	if err != nil || num_weight == 0 {
		return c.JSON(http.StatusBadRequest, ResponseErrorBMI{Message: "Invalid weight"})
	}

	height := c.QueryParam("height")
	num_height, err := strconv.Atoi(height)
	if err != nil || num_height == 0 {
		return c.JSON(http.StatusBadRequest, ResponseErrorBMI{Message: "Invalid height"})
	}

	var response_message string
	bmi := float64(num_weight / (num_height * num_height))

	switch {
	case bmi < 18.5:
		response_message = "You are underweight."
	case bmi >= 18.5 && bmi < 24.9:
		response_message = "You have a normal weight"
	case bmi >= 25 && bmi < 29.9:
		response_message = "You are overweight."
	default:
		response_message = "You are obese."
	}
	bmi_info := domain.BMI{
		Weight:    weight,
		Height:    height,
		ResultBMI: response_message,
	}
	ctx := c.Request().Context()
	err = a.Service.StoreBMILog(ctx, &bmi_info)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseResultBMI{Message: response_message})
}
