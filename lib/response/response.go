package response

/*
import (
	"encoding/json"
	"net/http"
)

const (
	Success = "success"
	Failure = "failure"
)

// ErrorResponseHelper func
func ErrorResponseHelper(methodName string, w http.ResponseWriter, errorMessage string, httpStatus int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(httpStatus)
	w.Write(errorResponse(errorMessage))
}

// SuccessResponseHelper func
func SuccessResponseHelper(w http.ResponseWriter, class interface{}, httpStatus int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(httpStatus)
	w.Write(successResponse(class))
}

// SuccessResponseList func
func SuccessResponseList(w http.ResponseWriter, class interface{}, offset string, limit string, count string) {
	tempResponse := make(map[string]interface{})
	tempResponse["count"] = count
	tempResponse["offset"] = offset
	tempResponse["limit"] = limit
	tempResponse["list"] = class

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(successResponse(tempResponse))
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

*/
import (
	"encoding/json"
	"net/http"
)

const (
	Success = "success"
	Failure = "failure"
)

// ResponseHandler is an interface for handling responses.
type ResponseHandler interface {
	Handle(w http.ResponseWriter, class interface{}, httpStatus int)
}

// BaseResponseHandler is a basic implementation of ResponseHandler.
type BaseResponseHandler struct{}

func (b *BaseResponseHandler) Handle(w http.ResponseWriter, class interface{}, httpStatus int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(httpStatus)
	w.Write(successResponse(class))
}

// SuccessResponseDecorator adds additional headers to the response.
type SuccessResponseDecorator struct {
	ResponseHandler
}

func (s *SuccessResponseDecorator) Handle(w http.ResponseWriter, class interface{}, httpStatus int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	s.ResponseHandler.Handle(w, class, httpStatus)
}

// ErrorResponseDecorator handles error responses.
type ErrorResponseDecorator struct {
	ResponseHandler
}

func (e *ErrorResponseDecorator) Handle(w http.ResponseWriter, class interface{}, httpStatus int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(httpStatus)
	w.Write(errorResponse(class.(string)))
}

// SuccessResponseListDecorator handles list responses with additional metadata.
type SuccessResponseListDecorator struct {
	ResponseHandler
}

func (s *SuccessResponseListDecorator) Handle(w http.ResponseWriter, class interface{}, offset string, limit string, count string) {
	tempResponse := make(map[string]interface{})
	tempResponse["count"] = count
	tempResponse["offset"] = offset
	tempResponse["limit"] = limit
	tempResponse["list"] = class

	s.ResponseHandler.Handle(w, tempResponse, http.StatusOK)
}

// ErrorResponseHelper function
func ErrorResponseHelper(methodName string, w http.ResponseWriter, errorMessage string, httpStatus int) {
	decorator := &ErrorResponseDecorator{&BaseResponseHandler{}}
	decorator.Handle(w, errorMessage, httpStatus)
}

// SuccessResponseHelper function
func SuccessResponseHelper(w http.ResponseWriter, class interface{}, httpStatus int) {
	decorator := &SuccessResponseDecorator{&BaseResponseHandler{}}
	decorator.Handle(w, class, httpStatus)
}

// SuccessResponseList function
func SuccessResponseList(w http.ResponseWriter, class interface{}, offset string, limit string, count string) {
	decorator := &SuccessResponseListDecorator{&BaseResponseHandler{}}
	decorator.Handle(w, class, offset, limit, count)
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
