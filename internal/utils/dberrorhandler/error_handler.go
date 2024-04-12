package dberrorhandler

import (
    "github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"

    "github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/msg/logmsg"

	"github.com/kebin6/simple-douyin-rpc/ent"
)

// DefaultEntError returns errors dealing with default functions.
func DefaultEntError(logger logx.Logger, err error, detail any) error {
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logger.Errorw(err.Error(), logx.Field("detail", detail))
			return errorx.NewInvalidArgumentError(errormsg.TargetNotFound)
		case ent.IsConstraintError(err):
			logger.Errorw(err.Error(), logx.Field("detail", detail))
			return errorx.NewInvalidArgumentError(errormsg.ConstraintError)
		case ent.IsValidationError(err):
			logger.Errorw(err.Error(), logx.Field("detail", detail))
			return errorx.NewInvalidArgumentError(errormsg.ValidationError)
		case ent.IsNotSingular(err):
			logger.Errorw(err.Error(), logx.Field("detail", detail))
			return errorx.NewInvalidArgumentError(errormsg.NotSingularError)
		default:
			logger.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return errorx.NewInternalError(errormsg.DatabaseError)
		}
	}
	return err
}
