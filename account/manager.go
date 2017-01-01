package account

import "bitbucket.com/sharingmachine/kwkcli/models"

type Manager interface {
	SignIn(username string, password string) (*models.User, error)
	SignUp(email string, username string, password string) (*models.User, error)
	Get() (*models.User, error)
	Signout() error
}