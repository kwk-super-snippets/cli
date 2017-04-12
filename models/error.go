package models

import (
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Code uint32



const (
	MaxProcessLevel = 3
	Code_Unspecified Code = 0
	Code_NotFound Code = 10
	Code_InvalidArgument Code = 20
	Code_UnexpectedEndOfTar Code = 30
	// Snippets
	Code_NoTags                 Code = 3000
	Code_NewFullKeyEmpty        Code = 3010
	Code_FullKeyExistsWhenClone Code = 3020
	Code_MultipleSnippetsFound  Code = 3030
	Code_TwoArgumentsRequiredForMove Code = 3040
	Code_SnippetNotVerified Code = 3050
	Code_SnippetVulnerable Code = 3060

	// Users
	Code_WrongCreds      Code = 4010
	Code_UsernameTaken   Code = 4020
	Code_EmailTaken      Code = 4030
	Code_EmptyToken      Code = 4040
	Code_InvalidEmail    Code = 4050
	Code_InvalidUsername Code = 4060
	Code_InvalidPassword Code = 4170

	Code_MultiplePouches  Code = 4210
	Code_IncompleteAlias  Code = 4220
	Code_AliasMaxSegments Code = 4230
	Code_NoSnippetName    Code = 4240
	Code_PouchMaxSegments    Code = 4250

	//Network
	Code_ApiDown Code = 5010

	Code_InvalidConfigSection Code = 6010
	Code_EnvironmentNotSupported Code = 6020

	//Runners
	Code_RunnerExitError Code = 700
	Code_ProcessTooDeep Code = 710

	//Writers
	Code_PrinterNotFound Code = 800
)

// ParseGrpcErr should be used at RPC service call level. i.e. the errors
// returned by the GRPC stubs need to be converted to local errors.
func ParseGrpcErr(e error) error {
	sts, _ := status.FromError(e)
	m := &ClientErr{}
	m.remoteCode = sts.Code()
	if m.remoteCode == codes.NotFound {
		return ErrOneLine(Code_NotFound, sts.Message())
	}
	if err := json.Unmarshal([]byte(sts.Message()), m); err != nil {
		m.Title = sts.Message()
		return m
	}
	return m
}

func ErrOneLine(c Code, desc string, args ...interface{}) error {
	return &ClientErr{Msgs: []Msg{{Code: c, Desc: fmt.Sprintf(desc, args...)}}}
}

type ClientErr struct {
	Msgs  []Msg
	Title string
	remoteCode codes.Code
}

func HasErrCode(err error, code Code) bool {
	e, ok := err.(*ClientErr)
	if !ok {
		return false
	}
	return e.Contains(code)
}

func (c *ClientErr) Contains(code Code) bool{
	for _, v := range c.Msgs {
		if v.Code == code {
			return true
		}
	}
	return false
}

func (c ClientErr) Error() string {
	return fmt.Sprintf("%s %+v %d", c.Title, c.Msgs, c.remoteCode)
}

type Msg struct {
	Code Code
	Desc string
}
