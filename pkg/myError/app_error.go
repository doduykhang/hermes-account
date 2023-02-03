package myError

import "errors"

var WrongCredential = errors.New("wrong email or password")
var EmailExists = errors.New("email exists")
