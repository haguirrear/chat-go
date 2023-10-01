package service

import "errors"

var ErrorUserAlreadyCreated = errors.New("The user exists already in db")
var ErrorEmailOrPasswordIncorrect = errors.New("Email or Password incorrects")
