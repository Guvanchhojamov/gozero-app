package response

import (
	"fmt"
	"github.com/go-errors/errors"
	"google.golang.org/grpc/codes"
)

func NewRpcErrResponse(code codes.Code, errMsg *errors.Error) error {
	return errors.New(fmt.Errorf("\"%s\":\"%s\"", code, errMsg.ErrorStack()))
}
