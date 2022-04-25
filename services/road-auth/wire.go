package main

import (
	"github.com/eastack-projects/road/services/road-auth/utils"
	"github.com/google/wire"
)

func InitializePasswordEncoder() utils.PasswordEncoder {
	wire.Build(utils.NewPasswordEncoder())
	return utils.PasswordEncoder{}
}
