package plugins

import (
	"honcc/server/internal/users"
	"honcc/server/pkg/accounts"
	"honcc/server/pkg/posts"
)

type BasePlugin interface {
	Metadata() PluginMetadata
}

type PlatformPlugin interface {
	BasePlugin

	AccountCreationFields() accounts.AccountFields
	Accounts() []accounts.Account
	AccountsForUser(users.UserId) []accounts.Account
	CreateAccount(accounts.AccountData) error
	UpdateAccount(accounts.Account) error
	DeleteAccount(accounts.Account) error

	PublishPost(posts.Post) error
}
