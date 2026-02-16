package install

import "fmt"

type UbuntuInstaller struct{}

func NewUbuntuInstaller() *UbuntuInstaller {
	return &UbuntuInstaller{}
}

func (u *UbuntuInstaller) Validate(config Config) error {
	if config.User.Name == "" {
		return fmt.Errorf("user name is required")
	}
	if config.System.Hostname == "" {
		return fmt.Errorf("hostname is required")
	}
	return nil
}

func (u *UbuntuInstaller) Install(config Config) error {
	return fmt.Errorf("not implemented")
}
