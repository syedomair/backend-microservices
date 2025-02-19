package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestErrorResponseHelper(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	recorder := httptest.NewRecorder()
	methodName := "TestMethod"
	errorMessage := "Test Error Message"
	httpStatus := http.StatusBadRequest

	ErrorResponseHelper(methodName, logger, recorder, errorMessage, httpStatus)

	assert.Equal(t, httpStatus, recorder.Code)
	assert.Equal(t, "application/json;charset=utf-8", recorder.Header().Get("Content-Type"))

	var response map[string]interface{}
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, Failure, response["result"])
	assert.Equal(t, 2, len(response))
}

func TestSuccessResponseHelper(t *testing.T) {
	recorder := httptest.NewRecorder()
	class := map[string]string{"key": "value"}
	httpStatus := http.StatusOK

	SuccessResponseHelper(recorder, class, httpStatus)

	assert.Equal(t, httpStatus, recorder.Code)
	assert.Equal(t, "application/json;charset=utf-8", recorder.Header().Get("Content-Type"))
	assert.Equal(t, "*", recorder.Header().Get("Access-Control-Allow-Origin"))

	var response map[string]interface{}
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, Success, response["result"])
	assert.Equal(t, 2, len(response))
}

func TestSuccessResponseList(t *testing.T) {
	recorder := httptest.NewRecorder()
	class := []string{"item1", "item2"}
	offset := "0"
	limit := "10"
	count := "2"

	SuccessResponseList(recorder, class, offset, limit, count)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	var response map[string]interface{}
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, Success, response["result"])
	assert.Equal(t, 2, len(response))
}

func TestErrorResponse(t *testing.T) {
	message := "Test Error Message"
	response := errorResponse(message)

	var result map[string]interface{}
	err := json.Unmarshal(response, &result)
	assert.NoError(t, err)
	assert.Equal(t, Failure, result["result"])
	assert.Equal(t, 2, len(result))
}

func TestSuccessResponse(t *testing.T) {
	class := map[string]string{"key": "value"}
	response := successResponse(class)
	var result map[string]interface{}
	err := json.Unmarshal(response, &result)
	assert.NoError(t, err)
	assert.Equal(t, Success, result["result"])
	assert.Equal(t, 2, len(result))
}

func TestCommonResponse(t *testing.T) {
	class := map[string]string{"key": "value"}
	result := "test_result"
	response := commonResponse(class, result)

	var resultMap map[string]interface{}
	err := json.Unmarshal(response, &resultMap)
	assert.NoError(t, err)
	assert.Equal(t, result, resultMap["result"])
	assert.Equal(t, 2, len(resultMap))
}
