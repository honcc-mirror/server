package accounts

type AccountFieldTypeEnum string

const (
	TextField     AccountFieldTypeEnum = "text"
	EmailField    AccountFieldTypeEnum = "email"
	PasswordField AccountFieldTypeEnum = "password"
	NumberField   AccountFieldTypeEnum = "number"
)

type AccountFieldName string
type AccountFields map[AccountFieldName]AccountFieldTypeEnum
type AccountData map[AccountFieldName]string

type Account struct {
	ID           string
	Displayname  string
	CustomFields AccountFields
}
