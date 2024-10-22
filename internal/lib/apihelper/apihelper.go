package apihelper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testTaskLamoda/internal/api/responses"
	"testTaskLamoda/internal/lib/jsonrpc"
)

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	ct := r.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			return fmt.Errorf("Content-Type header is not application/json")
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	dec := json.NewDecoder(r.Body)
	//dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var result error

		switch {
		case errors.As(err, &syntaxError):
			result = fmt.Errorf("request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			result = fmt.Errorf("request body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			result = fmt.Errorf("request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			result = fmt.Errorf("request body must not be empty")
		case err.Error() == "http: request body too large":
			result = fmt.Errorf("request body must not be larger than 1MB")
		default:
			result = err
		}
		return result
	}

	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return fmt.Errorf("request body must only contain a single JSON object")
	}
	return nil
}

func WriteErrorResponse(w http.ResponseWriter, id string, code string, err error) {
	baseJsonRpcResponse := &responses.BaseJsonRpcResponse{
		Jsonrpc: jsonrpc.JsonRPCVersion,
		Id:      id,
		Error: &responses.Error{
			Code:    code,
			Message: err.Error(),
		},
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	error := json.NewEncoder(w).Encode(baseJsonRpcResponse)
	if error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func EncodeJsonBody(w http.ResponseWriter, src interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(src)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
