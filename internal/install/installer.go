package install

type Config struct {
	Version string
	User    UserConfig
	System  SystemConfig
	Packages PackageConfig
}

type UserConfig struct {
	Name     string
	FullName string
	Shell    string
}

type SystemConfig struct {
	Hostname string
	Timezone string
	Locale   string
}

type PackageConfig struct {
	Base    []string
	Dev     []string
	Desktop []string
}

type Installer interface {
	Install(config Config) error
	Validate(config Config) error
}
