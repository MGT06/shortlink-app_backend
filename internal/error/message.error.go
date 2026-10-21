package errormsg

import "fmt"

var ErrEmailPasswordIncorrect = fmt.Errorf("email or password are incorrect")
var ErrEmailExist = fmt.Errorf("email is exist")
var ErrNoRowAffected = fmt.Errorf("no row affected")