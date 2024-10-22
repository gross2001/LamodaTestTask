package jsonrpc

const JsonRPCVersion = "2.0"

const (
	ParseError     = "-32700" // Invalid JSON was received by the server. An error occurred on the server while parsing the JSON text.
	InvalidRequest = "-32600" // The JSON sent is not a valid Request object.
	MethodNotFound = "-32601" // The method does not exist / is not available.
	InvalidParams  = "-32602" // Invalid method parameter(s).
	InternalError  = "-32603" // Internal JSON-RPC error.
	ServerError    = "-32000" //  -32000 to -32099 Reserved for implementation-defined server-errors.

)
