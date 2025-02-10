package user

import (
	"backend/lib/request"
	"backend/lib/response"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"
)

const (
	errorCodePrefix = "01"
	actionID        = "action_id"
	taskID          = "task_id"
	commNameID      = "comm_name_id"
)

type Controller struct {
	Logger *zap.Logger
	Repo   Repository
}

// GetAllActions Public
func (c *Controller) GetAllActions(w http.ResponseWriter, r *http.Request) {
	methodName := "GetAllActions"
	c.Logger.Debug(request.GetRequestID(r), zap.String("m", methodName))
	start := time.Now()

	c.Repo.SetRequestID(request.GetRequestID(r))

	limit, offset, orderby, sort, err := request.ValidateQueryString(r, "1000", "0", "id", "asc")
	if err != nil {
		//response.ErrorResponseHelper(request.GetRequestID(r), methodName, c.Logger, w, errorCodePrefix+"116", err.Error(), http.StatusBadRequest)
		return
	}

	actions, count, err := c.Repo.GetAllActionDB(limit, offset, orderby, sort)
	if err != nil {
		//response.ErrorResponseHelper(request.GetRequestID(r), methodName, c.Logger, w, errorCodePrefix+"117", err.Error(), http.StatusInternalServerError)
		return
	}
	c.Logger.Debug(request.GetRequestID(r), zap.String("m", methodName), zap.Duration("since", time.Since(start)))
	response.SuccessResponseList(w, actions, strconv.Itoa(offset), strconv.Itoa(limit), count)
}
