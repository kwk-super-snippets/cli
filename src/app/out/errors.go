package out

import (
	"fmt"
	"github.com/kwk-super-snippets/cli/src/style"
	"github.com/kwk-super-snippets/types/errs"
	"github.com/kwk-super-snippets/types/vwrite"
	"io"
	"os"
)

func NewErrHandler(w io.Writer) errs.Handler {
	return &handlerWrapper{Writer: vwrite.New(w)}
}

type handlerWrapper struct {
	vwrite.Writer
}

func (e *handlerWrapper) Handle(err error) {
	e.Write(errHandler(err))
}

func errHandler(e error) vwrite.Handler {
	return vwrite.HandlerFunc(func(w io.Writer) {
		if e == nil {
			return
		}
		LogErr(e)
		ce, ok := e.(*errs.Error)
		if ok {
			switch ce.Code {
			case errs.CodeInvalidArgument:
				invalidArgument(ce)
			case errs.CodeNotAuthenticated:
				NotAuthenticated.Write(w)
			case errs.CodeNotFound:
				notFound("").Write(w)
			case errs.CodeAlreadyExists:
				ItemExists.Write(w)
			case errs.CodePermissionDenied:
				notPermitted.Write(w)
				return
			case errs.CodeNotImplemented:
				notImplemented.Write(w)
				return
			case errs.CodeNotAvailable:
				notAvailable.Write(w)
			default:
				internalError(ce).Write(w)
			}
			return
		}
		internalError(e).Write(w)

	})
}

var overrides = map[errs.ErrCode]string{
	errs.CodeInvalidPassword: "Password must have one upper, one lower and one numeric",
	errs.CodeInvalidUsername: "Username must be bl",
	errs.CodeEmailTaken:      "That email has been taken",
}

func getMessageOverride(code errs.ErrCode) string {
	return overrides[code]
}

var NotAuthenticated = Warn(vwrite.HandlerFunc(func(w io.Writer) {
	fmt.Fprintln(w, "Please login to continue: kwk login")
}))

var ItemExists = Warn(vwrite.HandlerFunc(func(w io.Writer) {
	fmt.Fprintln(w, "An item with that identifier already exists.")
}))

func notFound(name string) vwrite.Handler {
	return Warn(vwrite.HandlerFunc(func(w io.Writer) {
		if name == "" {
			name = "Resource"
		}
		fmt.Fprintf(w, "%s couldn't be found.\n", name)
	}))
}

func internalError(err error) vwrite.Handler {
	return vwrite.HandlerFunc(func(w io.Writer) {
		ce, ok := err.(*errs.Error)
		if ok {
			o := getMessageOverride(ce.Code)
			if o != "" {
				ce.Message = o
			}
			fmt.Fprintln(w, ce.Message)
			return
		}
		LogErr(err)
		fmt.Fprintln(w, "🔥  We have a code RED error. \n- To report type: kwk upload-errors \n"+
			"- You can also try to upgrade: npm update kwkcli -g\n")
		os.Exit(1)
	})
}

var notImplemented = Warn(vwrite.HandlerFunc(func(w io.Writer) {
	fmt.Fprintln(w, "Your CLI may be out of date. Please run: kwk update")
}))

var notPermitted = Warn(vwrite.HandlerFunc(func(w io.Writer) {
	fmt.Fprintln(w, "Permission denied.")
}))

var notAvailable = Warn(vwrite.HandlerFunc(func(w io.Writer) {
	fmt.Fprintf(w, "%s  Kwk is DOWN! Please try again in a bit.\n", style.Fire)
}))

func invalidArgument(err *errs.Error) vwrite.Handler {
	return Warn(vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprintf(w, "%s.\n", err.Message)
	}))
}
