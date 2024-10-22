package api

import (
	"net/http"
	"testTaskLamoda/internal/api/parameters"
	"testTaskLamoda/internal/api/responses"
	"testTaskLamoda/internal/lib/apihelper"
	"testTaskLamoda/internal/lib/jsonrpc"

	"github.com/go-playground/validator/v10"
)

func (serverAPI *ServerAPI) DeleteReserve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// обрабатываем запрос
		var reserveDeleteRequest parameters.ReserveDeleteRequest
		if err := apihelper.DecodeJSONBody(w, r, &reserveDeleteRequest); err != nil {
			apihelper.WriteErrorResponse(w, "", jsonrpc.ParseError, err)
			serverAPI.log.Error("Error while decoding json body", "error:", err)
			return
		}
		validate := validator.New()
		if err := validate.Struct(reserveDeleteRequest); err != nil {
			apihelper.WriteErrorResponse(w, reserveDeleteRequest.Id, jsonrpc.InvalidRequest, err)
			serverAPI.log.Error("Error while validate json body", "error:", err)
			return
		}
		serverAPI.log.Debug("deleteReserveHandler", "deleteReserveRequest: ", reserveDeleteRequest)

		// вызываем удаленный метод
		reserveDelete, err := serverAPI.Services.Reserve.Delete(r.Context(), reserveDeleteRequest.Params.Skus)
		if err != nil {
			apihelper.WriteErrorResponse(w, reserveDeleteRequest.Id, jsonrpc.InternalError, err)
			serverAPI.log.Error("Error while Reserve.Create", "error:", err)
			return
		}
		serverAPI.log.Debug("deleteReserveHandler", "responseDeleteResponse: ", reserveDelete)

		// отправляем ответ
		reserveDeleteResponse := responses.ReserveDeleteResponse{
			BaseJsonRpcResponse: responses.BaseJsonRpcResponse{
				Id:      reserveDeleteRequest.Id,
				Jsonrpc: jsonrpc.JsonRPCVersion,
			},
			Result: &reserveDelete,
		}
		apihelper.EncodeJsonBody(w, reserveDeleteResponse)
	}
}
