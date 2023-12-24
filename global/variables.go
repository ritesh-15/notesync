package global

import (
	"github.com/go-playground/validator/v10"
	"github.com/ritesh-15/notesync-backend/utils"
)

var ValidatorInstance = validator.New()
var MyValidator = utils.NewValidator(ValidatorInstance)
