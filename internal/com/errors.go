package com

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/d-darac/lagra/internal/database/sqlc"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type AppError struct {
	Code    *ErrorCode
	Param   *string
	Message string
	Type    ErrorType
	Status  int
}

type ArrayNotGtErr struct {
	N, Param string
}

type ArrayNotGteErr struct {
	N, Param string
}

type ArrayNotLtErr struct {
	N, Param string
}

type ArrayNotLteErr struct {
	N, Param string
}

type CountryUnknownErr struct {
	Param *string
	Value string
}

type CurrencyUnknownErr struct {
	Param *string
	Value string
}

type ExclusiveParamsErr struct {
	A, B string
}

type FieldExpandDepthErr struct {
	Field string
}

type FieldUnexpandableErr struct {
	Field string
}

type InvalidIDErr struct {
	Param *string
	Value string
}

type InvalidStringErr struct {
	Param *string
}

type InvalidItemTypeErr struct {
	Param string
}

type JsonDecodeErr struct {
	Err error
}

type JsonTypeErr struct {
	Param, Type string
}

type MethodNotAllowedErr struct {
	Method, Path string
}

type ParameterInvalidErr struct {
	Param string
}

type ParameterMissingErr struct {
	Param string
}

type RequestTooLargeErr struct{}

type ResourceNotFoundErr struct {
	Param    *string
	ID       string
	Resource string
}

type RouteUnknownErr struct {
	Method, Path string
}

type StringNotGtErr struct {
	N, Param string
}

type StringNotGteErr struct {
	N, Param string
}

type StringNotLtErr struct {
	N, Param string
}

type StringNotLteErr struct {
	N, Param string
}

type ValueNotGtErr struct {
	N, Param string
}

type ValueNotGteErr struct {
	N, Param string
}

type ValueNotLtErr struct {
	N, Param string
}

type ValueNotLteErr struct {
	N, Param string
}

type ValueNotOneOfErr struct {
	Param, Set string
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *ArrayNotGtErr) Error() string {
	return fmt.Sprintf("The number of elements in the '%s' array must be greater than %s.", e.Param, e.N)
}

// func (e *ArrayNotGtErr) As(target any) bool {
// 	if target == nil {
// 		return false
// 	}
// 	appErr, ok := target.(*AppError)
// 	if ok {
// 		code := ParameterInvalid
// 		appErr.Code = &code
// 		appErr.Message = e.Error()
// 		appErr.Param = &e.Param
// 		appErr.Status = http.StatusBadRequest
// 		appErr.Type = InvalidRequestError
// 	}
// 	return ok
// }

func (e *ArrayNotGteErr) Error() string {
	return fmt.Sprintf("The number of elements in the '%s' array must be at least %s.", e.Param, e.N)
}

func (e *ArrayNotLtErr) Error() string {
	return fmt.Sprintf("The number of elements in the '%s' array must be fewer than %s.", e.Param, e.N)
}

func (e *ArrayNotLteErr) Error() string {
	return fmt.Sprintf("The number of elements in the '%s' array must be at most %s.", e.Param, e.N)
}

func (e *CountryUnknownErr) Error() string {
	return fmt.Sprintf("Country '%s' is unknown. "+
		"Try using a 2-character alphanumeric country code instead, such as 'DE', 'GB', or 'IE'. "+
		"A complete list of official country codes is available at https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements",
		e.Value)
}

func (e *CurrencyUnknownErr) Error() string {
	return fmt.Sprintf("Currency '%s' is unknown. "+
		"Try using a 3-character alphanumeric currency code instead, such as 'USD', 'EUR', or 'GBP'. "+
		"A complete list of official currency codes is available at https://en.wikipedia.org/wiki/ISO_4217#Active_codes_(list_one)",
		e.Value,
	)
}

func (e *ExclusiveParamsErr) Error() string {
	return fmt.Sprintf("Received both '%s' and '%s' parameters. Pass one at a time.", e.A, e.B)
}

func (e *FieldExpandDepthErr) Error() string {
	return fmt.Sprintf("Cannot expand more than 4 levels of field depth. Field: %s", e.Field)
}

func (e *FieldUnexpandableErr) Error() string {
	return fmt.Sprintf("Cannot expand field: %s", e.Field)
}

func (e *InvalidIDErr) Error() string {
	return fmt.Sprintf("Invalid id: '%s'", e.Value)
}

func (e *InvalidItemTypeErr) Error() string {
	return fmt.Sprintf("Value of '%s' must be one of: %s", e.Param, strings.Join([]string{string(sqlc.ItemTypeASSET), string(sqlc.ItemTypePRODUCT)}, ", "))
}

func (e *InvalidStringErr) Error() string {
	return "Invalid string."
}

func (e *JsonDecodeErr) Error() string {
	return fmt.Sprintf("Error parsing JSON: %s", e.Err.Error())
}

func (e *JsonTypeErr) Error() string {
	return fmt.Sprintf("Invalid %s.", e.Type)
}

func (e *MethodNotAllowedErr) Error() string {
	return fmt.Sprintf("Method '%s' not allowed on %s.", e.Method, e.Path)
}

func (e *ParameterInvalidErr) Error() string {
	return fmt.Sprintf("Parameter invalid: '%s'", e.Param)
}

func (e *ParameterMissingErr) Error() string {
	return fmt.Sprintf("Missing required param: '%s'", e.Param)
}

func (e *RequestTooLargeErr) Error() string {
	return "Request body too large."
}

func (e *ResourceNotFoundErr) Error() string {
	return fmt.Sprintf("%s with id '%s' not found.", cases.Title(language.English).String(e.Resource), e.ID)
}

func (e *RouteUnknownErr) Error() string {
	return fmt.Sprintf("Request to unknown route (%s: %s).", e.Method, e.Path)
}

func (e *StringNotGtErr) Error() string {
	return fmt.Sprintf("The length of the '%s' parameter must be greater than %s characters.", e.Param, e.N)
}

func (e *StringNotGteErr) Error() string {
	return fmt.Sprintf("The length of the '%s' parameter must be at least %s characters.", e.Param, e.N)
}

func (e *StringNotLtErr) Error() string {
	return fmt.Sprintf("The length of the '%s' parameter must be fewer than %s characters.", e.Param, e.N)
}

func (e *StringNotLteErr) Error() string {
	return fmt.Sprintf("The length of the '%s' parameter must be at most %s characters.", e.Param, e.N)
}

func (e *ValueNotGtErr) Error() string {
	return fmt.Sprintf("Value of '%s' param must be greater than '%s'.", e.Param, e.N)
}

func (e *ValueNotGteErr) Error() string {
	return fmt.Sprintf("Value of '%s' param must be greater than or equal to '%s'.", e.Param, e.N)
}

func (e *ValueNotLtErr) Error() string {
	return fmt.Sprintf("Value of '%s' param must be less than '%s'.", e.Param, e.N)
}

func (e *ValueNotLteErr) Error() string {
	return fmt.Sprintf("Value of '%s' param must be less than or equal to '%s'.", e.Param, e.N)
}

func (e *ValueNotOneOfErr) Error() string {
	return fmt.Sprintf("Value of '%s' must be one of: %s", e.Param, e.Set)
}

func ErrorToAppError(err error) *AppError {
	if arrayNotGtErr, ok := errors.AsType[*ArrayNotGtErr](err); ok {
		return &AppError{
			Code:    new(ParameterInvalid),
			Message: arrayNotGtErr.Error(),
			Param:   &arrayNotGtErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if arrayNotGteErr, ok := errors.AsType[*ArrayNotGteErr](err); ok {
		return &AppError{
			Code:    new(ParameterInvalid),
			Message: arrayNotGteErr.Error(),
			Param:   &arrayNotGteErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if arrayNotLtErr, ok := errors.AsType[*ArrayNotLtErr](err); ok {
		return &AppError{
			Code:    new(ParameterInvalid),
			Message: arrayNotLtErr.Error(),
			Param:   &arrayNotLtErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if arrayNotLteErr, ok := errors.AsType[*ArrayNotLteErr](err); ok {
		return &AppError{
			Code:    new(ParameterInvalid),
			Message: arrayNotLteErr.Error(),
			Param:   &arrayNotLteErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if countryUnknownErr, ok := errors.AsType[*CountryUnknownErr](err); ok {
		return &AppError{
			Code:    new(CountryCodeInvalid),
			Message: countryUnknownErr.Error(),
			Param:   countryUnknownErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if currencyUnknownErr, ok := errors.AsType[*CurrencyUnknownErr](err); ok {
		return &AppError{
			Code:    new(CurrencyCodeInvalid),
			Message: currencyUnknownErr.Value,
			Param:   currencyUnknownErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if exclusiveParamsErr, ok := errors.AsType[*ExclusiveParamsErr](err); ok {
		return &AppError{
			Message: exclusiveParamsErr.Error(),
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if fieldExpandDepthErr, ok := errors.AsType[*FieldExpandDepthErr](err); ok {
		return &AppError{
			Code:    new(FieldExpansionMaxDepth),
			Message: fieldExpandDepthErr.Error(),
			Param:   new("expand"),
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if fieldUnexpandableErr, ok := errors.AsType[*FieldUnexpandableErr](err); ok {
		return &AppError{
			Code:    new(FieldUnexpandable),
			Message: fieldUnexpandableErr.Error(),
			Param:   new("expand"),
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if invalidIdErr, ok := errors.AsType[*InvalidIDErr](err); ok {
		appErr := &AppError{
			Code:    nil,
			Message: invalidIdErr.Error(),
			Param:   invalidIdErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
		if invalidIdErr.Param != nil {
			code := ParameterInvalid
			appErr.Code = &code
		}
		return appErr
	}

	if invalidItemTypeErr, ok := errors.AsType[*InvalidItemTypeErr](err); ok {
		return &AppError{
			Code:    new(ParameterInvalid),
			Message: invalidItemTypeErr.Error(),
			Param:   &invalidItemTypeErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if invalidStringErr, ok := errors.AsType[*InvalidStringErr](err); ok {
		return &AppError{
			Code:    new(ParameterInvalid),
			Message: invalidStringErr.Error(),
			Param:   invalidStringErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if jsonDecodeErr, ok := errors.AsType[*JsonDecodeErr](err); ok {
		return &AppError{
			Message: jsonDecodeErr.Error(),
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if jsonTypeErr, ok := errors.AsType[*JsonTypeErr](err); ok {
		return &AppError{
			Code:    new(ParameterInvalid),
			Message: jsonTypeErr.Error(),
			Param:   &jsonTypeErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if methodNotAllowedErr, ok := errors.AsType[*MethodNotAllowedErr](err); ok {
		return &AppError{
			Message: methodNotAllowedErr.Error(),
			Status:  http.StatusMethodNotAllowed,
			Type:    InvalidRequestError,
		}
	}

	if parameterInvalidErr, ok := errors.AsType[*ParameterInvalidErr](err); ok {
		return &AppError{
			Code:    new(ParameterInvalid),
			Message: parameterInvalidErr.Error(),
			Param:   &parameterInvalidErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if parameterMissingErr, ok := errors.AsType[*ParameterMissingErr](err); ok {
		return &AppError{
			Code:    new(ParameterMissing),
			Message: parameterMissingErr.Error(),
			Status:  http.StatusBadRequest,
			Param:   &parameterMissingErr.Param,
			Type:    InvalidRequestError,
		}
	}

	if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
		if pgErr.Code == "23503" {
			column, value := fromPgErr23503(pgErr)
			return ErrorToAppError(new(ResourceNotFoundErr{
				ID:       value,
				Param:    &column,
				Resource: strings.Replace(column, "parent_", "", 1),
			}))
		}
	}

	if _, ok := errors.AsType[*RequestTooLargeErr](err); ok {
		return &AppError{
			Message: "Request body too large.",
			Status:  http.StatusRequestEntityTooLarge,
			Type:    InvalidRequestError,
		}
	}

	if resourceNotFoundErr, ok := errors.AsType[*ResourceNotFoundErr](err); ok {
		return &AppError{
			Code:    new(ResourceNotFound),
			Message: resourceNotFoundErr.Error(),
			Param:   resourceNotFoundErr.Param,
			Status:  http.StatusNotFound,
			Type:    InvalidRequestError,
		}
	}

	if routeUnknownErr, ok := errors.AsType[*RouteUnknownErr](err); ok {
		return &AppError{
			Message: routeUnknownErr.Error(),
			Status:  http.StatusNotFound,
			Type:    InvalidRequestError,
		}
	}

	if stringNotGtErr, ok := errors.AsType[*StringNotGtErr](err); ok {
		return &AppError{
			Code:    new(ParameterInvalid),
			Message: stringNotGtErr.Error(),
			Param:   &stringNotGtErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if stringNotGteErr, ok := errors.AsType[*StringNotGteErr](err); ok {
		return &AppError{
			Code:    new(ParameterInvalid),
			Message: stringNotGteErr.Error(),
			Param:   &stringNotGteErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if stringNotLtErr, ok := errors.AsType[*StringNotLtErr](err); ok {
		return &AppError{
			Code:    new(ParameterInvalid),
			Message: stringNotLtErr.Error(),
			Param:   &stringNotLtErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if stringNotLteErr, ok := errors.AsType[*StringNotLteErr](err); ok {
		return &AppError{
			Code:    new(ParameterInvalid),
			Message: stringNotLteErr.Error(),
			Param:   &stringNotLteErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if valueNotGtErr, ok := errors.AsType[*ValueNotGtErr](err); ok {
		return &AppError{
			Code:    new(ParameterInvalid),
			Message: valueNotGtErr.Error(),
			Param:   &valueNotGtErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if valueNotGteErr, ok := errors.AsType[*ValueNotGteErr](err); ok {
		return &AppError{
			Code:    new(ParameterInvalid),
			Message: valueNotGteErr.Error(),
			Param:   &valueNotGteErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if valueNotLtErr, ok := errors.AsType[*ValueNotLtErr](err); ok {
		return &AppError{
			Code:    new(ParameterInvalid),
			Message: valueNotLtErr.Error(),
			Param:   &valueNotLtErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if valueNotLteErr, ok := errors.AsType[*ValueNotLteErr](err); ok {
		return &AppError{
			Code:    new(ParameterInvalid),
			Message: valueNotLteErr.Error(),
			Param:   &valueNotLteErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	if valueNotOneOfErr, ok := errors.AsType[*ValueNotOneOfErr](err); ok {
		return &AppError{
			Code:    new(ParameterInvalid),
			Message: valueNotOneOfErr.Error(),
			Param:   &valueNotOneOfErr.Param,
			Status:  http.StatusBadRequest,
			Type:    InvalidRequestError,
		}
	}

	// if error is unexpected
	log.Print(err)

	return &AppError{
		Message: "Something went wrong.",
		Status:  http.StatusInternalServerError,
		Type:    ApiError,
	}
}

func fromPgErr23503(pgErr *pgconn.PgError) (column, value string) {
	re := regexp.MustCompile(`(?m)\((.*)\)=\((.*)\)`)
	if match := re.FindAllStringSubmatch(pgErr.Detail, 1); match != nil {
		if len(match[0]) == 3 {
			column = strings.Replace(match[0][1], "_id", "", 1)
			value = match[0][2]
		}
	}
	return
}
