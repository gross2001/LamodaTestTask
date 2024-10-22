package api

import (
	"net/http"
	"testTaskLamoda/internal/api/parameters"
	"testTaskLamoda/internal/api/responses"
	"testTaskLamoda/internal/lib/apihelper"
	"testTaskLamoda/internal/lib/jsonrpc"

	"github.com/go-playground/validator/v10"
)

func (serverAPI *ServerAPI) CreateReserve() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// обрабатываем запрос
		var reserveCreateRequest parameters.ReserveCreateRequest
		if err := apihelper.DecodeJSONBody(w, r, &reserveCreateRequest); err != nil {
			apihelper.WriteErrorResponse(w, "", jsonrpc.ParseError, err)
			serverAPI.log.Error("Error while decoding json body", "error:", err)
			return
		}

		validate := validator.New()
		if err := validate.Struct(reserveCreateRequest); err != nil {
			apihelper.WriteErrorResponse(w, reserveCreateRequest.Id, jsonrpc.InvalidRequest, err)
			serverAPI.log.Error("Error while validate json body", "error:", err)
			return
		}
		serverAPI.log.Debug("createReserveHandler", "reserveCreateRequest: ", reserveCreateRequest)

		// вызываем удаленный метод
		reserveCreate, err := serverAPI.Services.Reserve.Create(r.Context(), reserveCreateRequest.Params.Skus)
		if err != nil {
			apihelper.WriteErrorResponse(w, reserveCreateRequest.Id, jsonrpc.InternalError, err)
			serverAPI.log.Error("Error while Reserve.Create", "error:", err)
			return
		}
		serverAPI.log.Debug("createReserveHandler", "responseCreateResponse: ", reserveCreate)

		// отправляем ответ
		reserveCreateResponse := responses.ReserveCreateResponse{
			BaseJsonRpcResponse: responses.BaseJsonRpcResponse{
				Id:      reserveCreateRequest.Id,
				Jsonrpc: jsonrpc.JsonRPCVersion,
			},
			Result: &reserveCreate,
		}
		apihelper.EncodeJsonBody(w, reserveCreateResponse)
	}
}
