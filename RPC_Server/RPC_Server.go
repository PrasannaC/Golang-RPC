package RPC_Server

import (
	"Go_RPC/Common"
	"errors"
)

type StringOps struct {
}

func (t *StringOps) Concat(args *Common.Args, resultString *string) error {
	if args == nil || args.StringOne == "" || args.StringTwo == "" {
		return errors.New("Invalid Parameters.")
	}
	*resultString = string("")
	*resultString = args.StringOne + args.StringTwo
	return nil
}
