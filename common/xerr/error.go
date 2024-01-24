package xerr

import (
	"fmt"
	pkg_err "github.com/pkg/errors"
	"google.golang.org/grpc/status"
	"lightleap-c/common/xerr/code"
	"lightleap-c/common/xerr/msg"
)

type CodeError struct {
	ErrCode uint32
	ErrMsg  string
	err     error
}

func (e *CodeError) GetCode() uint32 {
	return e.ErrCode
}

func (e *CodeError) GetMsg() string {
	return e.ErrMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode: %d, ErrMsg: %s, Err: %+v", e.ErrCode, e.ErrMsg, e.err)
}

func NewErrCode(errCode uint32) *CodeError {
	ec := errCode
	return &CodeError{ErrCode: errCode, ErrMsg: msg.GetMsg(ec), err: pkg_err.New(msg.GetMsg(ec))}
}

func New(errCode uint32, errMsg string) *CodeError {
	return &CodeError{ErrCode: errCode, ErrMsg: errMsg, err: pkg_err.New(errMsg)}
}

func NewErrorWithMsg(errMsg string) *CodeError {
	return &CodeError{ErrCode: code.CODE_ERROR, ErrMsg: errMsg, err: pkg_err.New(errMsg)}
}

func Cause(err error) error {
	if err != nil {
		var (
			codeErr *CodeError
			//rpcErr  *status.Status
		)

		if st, ok := status.FromError(err); ok {
			err = st.Err()
		}
		if pkg_err.As(err, &codeErr) {
			err = pkg_err.Cause(codeErr.err)
		}
	}
	return err
}

func Code(err error) uint32 {
	if err != nil {
		var codeErr *CodeError
		if pkg_err.As(err, &codeErr) {
			return codeErr.GetCode()
		}
		return code.CODE_ERROR
	}
	return code.CODE_OK
}

func Msg(err error) string {
	if err != nil {
		var codeErr *CodeError
		if pkg_err.As(err, &codeErr) {
			return codeErr.GetMsg()
		}
		return msg.GetMsg(code.CODE_ERROR)
	}
	return ""
}

func Wrap(err error, code uint32, msg string) error {
	if err == nil {
		return nil
	}

	if msg == "" {
		return &CodeError{
			err:     err,
			ErrMsg:  msg,
			ErrCode: code,
		}
	}

	return &CodeError{
		err:     pkg_err.Wrap(err, msg),
		ErrCode: code,
		ErrMsg:  msg,
	}
}
