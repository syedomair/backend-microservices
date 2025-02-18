package response

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

const (
	Success = "success"
	Failure = "failure"
)

// ErrorResponseHelper func
func ErrorResponseHelper(methodName string, logger *zap.Logger,
	w http.ResponseWriter, errorMessage string, httpStatus int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(httpStatus)
	w.Write(errorResponse(errorMessage))
	logger.Debug("error response", zap.String("method", methodName), zap.String("error_message", errorMessage))
}

// SuccessResponseHelper func
func SuccessResponseHelper(w http.ResponseWriter, class interface{}, httpStatus int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(httpStatus)
	w.Write(successResponse(class))
	//c.Logger.Log("METHOD", "SuccessResponseHelper", "END", c.successResponse(class))
}

// SuccessResponseList func
func SuccessResponseList(w http.ResponseWriter, class interface{}, offset string, limit string, count string) {
	tempResponse := make(map[string]interface{})
	tempResponse["count"] = count
	tempResponse["list"] = class

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(successResponse(tempResponse))
	//c.Logger.Log("METHOD", "SuccessResponseList", "END", c.successResponse(tempResponse))
}

func errorResponse(message string) []byte {
	class := map[string]string{"message": message}
	return commonResponse(class, Failure)
}

func successResponse(class interface{}) []byte {
	return commonResponse(class, Success)
}

func commonResponse(class interface{}, result string) []byte {
	response := make(map[string]interface{})
	response["result"] = result
	response["data"] = class
	jsonByte, _ := json.Marshal(response)
	return jsonByte
}
