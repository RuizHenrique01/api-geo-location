package middlewares

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type validationError struct {
	ActualTag string `json:"tag"`
	Namespace string `json:"namespace"`
	Kind      string `json:"kind"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	Param     string `json:"param"`
}

func HandleValidatorError(model interface{}) iris.Handler {
	return func(ctx iris.Context) {
		err := ctx.ReadJSON(*&model)
		if err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				validationErrors := make([]validationError, 0, len(errs))
				for _, validationErr := range errs {
					validationErrors = append(validationErrors, validationError{
						ActualTag: validationErr.ActualTag(),
						Namespace: validationErr.Namespace(),
						Kind:      validationErr.Kind().String(),
						Type:      validationErr.Type().String(),
						Value:     fmt.Sprintf("%v", validationErr.Value()),
						Param:     validationErr.Param(),
					})
				}

				ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
					Title("Validation error").
					Detail("One or more fields failed to be validated").
					Type("/user/validation-errors").
					Key("errors", validationErrors))
			}

			ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
				Title(err.Error()))
		}

		ctx.Values().Set("model", model)
		ctx.Next()
	}

}
