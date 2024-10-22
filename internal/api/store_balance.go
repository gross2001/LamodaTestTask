package api

import (
	"net/http"
	"testTaskLamoda/internal/api/parameters"
	"testTaskLamoda/internal/api/responses"
	"testTaskLamoda/internal/lib/apihelper"
	"testTaskLamoda/internal/lib/jsonrpc"

	"github.com/go-playground/validator/v10"
)

func (serverAPI *ServerAPI) StoreBalace() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// обрабатываем запрос
		var storeBalanceRequest parameters.StoreBalanceRequest
		if err := apihelper.DecodeJSONBody(w, r, &storeBalanceRequest); err != nil {
			apihelper.WriteErrorResponse(w, "", jsonrpc.ParseError, err)
			serverAPI.log.Error("Error while decoding json body", "error:", err)
			return
		}
		validate := validator.New()
		if err := validate.Struct(storeBalanceRequest); err != nil {
			apihelper.WriteErrorResponse(w, storeBalanceRequest.Id, jsonrpc.InvalidRequest, err)
			serverAPI.log.Error("Error while validate json body", "error:", err)
			return
		}
		serverAPI.log.Debug("storeBalanceRequestHandler", "storeBalanceRequest: ", storeBalanceRequest)

		// вызываем удаленный метод
		storeBalance, err := serverAPI.Services.Store.Balance(r.Context(), storeBalanceRequest.Params.Store_Id)
		if err != nil {
			apihelper.WriteErrorResponse(w, storeBalanceRequest.Id, jsonrpc.InternalError, err)
			serverAPI.log.Error("Error while Reserve.Create", "error:", err)
			return
		}
		serverAPI.log.Debug("storeBalanceHandler", "storeBalanceResponse: ", storeBalance)
		// отправляем ответ
		storeBalanceResponse := responses.StoreBalanceResponse{
			BaseJsonRpcResponse: responses.BaseJsonRpcResponse{
				Id:      storeBalanceRequest.Id,
				Jsonrpc: jsonrpc.JsonRPCVersion,
			},
			Result: &storeBalance,
		}
		apihelper.EncodeJsonBody(w, storeBalanceResponse)
	}
}
