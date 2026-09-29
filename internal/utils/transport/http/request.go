package transport_http_utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func GetJson[T any](r *http.Request) (T, error) {
	var req T
	var reqNil T

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return reqNil, fmt.Errorf("Failed to parse: %w", err)
	}

	return req, nil
}

func GetAndFilterJson[T any](validator *validator.Validate, r *http.Request) (T, error) {
	var req T
	var reqNil T

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return reqNil, fmt.Errorf("Failed to parse: %w", err)
	}

	if err := validator.Struct(req); err != nil {
		return reqNil, fmt.Errorf("The structure failed validation: %w", err)
	}

	return req, nil
}
