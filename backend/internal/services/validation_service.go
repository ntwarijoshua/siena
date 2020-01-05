package services

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/ntwarijoshua/siena/internal/models"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
)

type ValidationService struct{
	validator *validator.Validate
	dataLayer *models.DataStore
	translator locales.Translator
	uni *ut.UniversalTranslator
	lang ut.Translator
	logger *logrus.Logger
}

func (vs *ValidationService)GetValidator() *validator.Validate {
	return vs.validator
}

func (vs *ValidationService)InitializeValidator()  {
	vs.validator = validator.New()
	vs.translator = en.New()
	vs.uni = ut.New(vs.translator, vs.translator)
	_ = vs.validator.RegisterValidation("is_unique", func(fl validator.FieldLevel) bool {
		email := fl.Field().String()
		user, err := vs.dataLayer.GetUserByEmail(email)
		if err == nil && user.Email == email {
			return false
		}
		return true
	})

	//register custom message
	var found = false
	vs.lang, found = vs.uni.FindTranslator("en")
	if !found {
		vs.logger.Fatal("Could not find translator")
	}
	if err := enTranslations.RegisterDefaultTranslations(vs.validator, vs.lang); err != nil {
		vs.logger.Fatal(err)
	}

	_ = vs.validator.RegisterTranslation("required", vs.lang, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = vs.validator.RegisterTranslation("email", vs.lang, func(ut ut.Translator) error {
		return ut.Add("email", "{0} should be a valid email", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	_ = vs.validator.RegisterTranslation("is_unique", vs.lang, func(ut ut.Translator) error {
		return ut.Add("is_unique", "{0} An account with the same email already exists", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("is_unique", fe.Field())
		return t
	})
}

func (vs *ValidationService)GenerateValidationResponse(err error) []string {
	var errMessages []string
	for _,e := range err.(validator.ValidationErrors) {
		errMessages = append(errMessages, e.Translate(vs.lang))
	}
	return errMessages
}


