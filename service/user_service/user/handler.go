package user

import (
	"net/http"
	"strconv"
	"time"

	"github.com/syedomair/backend-example/lib/request"
	"github.com/syedomair/backend-example/lib/response"

	"go.uber.org/zap"
)

type Controller struct {
	Logger *zap.Logger
	Repo   Repository
}

// GetAllUsers Public
func (c *Controller) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	methodName := "GetAllUsers"
	c.Logger.Debug("method start", zap.String("method_name", methodName))
	start := time.Now()

	limit, offset, orderby, sort, err := request.ValidateQueryString(r, "1000", "0", "name", "asc")
	if err != nil {
		response.ErrorResponseHelper(methodName, c.Logger, w, err.Error(), http.StatusBadRequest)
		return
	}

	actions, count, err := c.Repo.GetAllUserDB(limit, offset, orderby, sort)
	if err != nil {
		response.ErrorResponseHelper(methodName, c.Logger, w, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Logger.Debug("method end", zap.String("method_name", methodName), zap.Duration("since", time.Since(start)))
	response.SuccessResponseList(w, actions, strconv.Itoa(offset), strconv.Itoa(limit), count)
}
