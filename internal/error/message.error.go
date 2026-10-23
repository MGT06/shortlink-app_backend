package errormsg

import "fmt"

var ErrEmailPasswordIncorrect = fmt.Errorf("email or password are incorrect")
var ErrEmailExist = fmt.Errorf("email is exist")
var ErrNoRowAffected = fmt.Errorf("no row affected")
var ErrSlugTaken = fmt.Errorf("slug already taken")
var ErrLinkNotFound = fmt.Errorf("link not found")
var ErrSlugInvalid = fmt.Errorf("The slug must be 3–50 characters long and may only contain letters, numbers, and hyphens.")
var ErrSlugReserved = fmt.Errorf("slug includes reserved words")