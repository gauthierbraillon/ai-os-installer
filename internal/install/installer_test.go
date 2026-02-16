package install

import "testing"

func TestUbuntuInstaller_Validate(t *testing.T) {
	installer := NewUbuntuInstaller()

	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: Config{
				User:   UserConfig{Name: "john", FullName: "John Doe", Shell: "bash"},
				System: SystemConfig{Hostname: "dev-machine", Timezone: "UTC", Locale: "en_US.UTF-8"},
			},
			wantErr: false,
		},
		{
			name: "missing user name",
			config: Config{
				System: SystemConfig{Hostname: "dev-machine"},
			},
			wantErr: true,
		},
		{
			name: "missing hostname",
			config: Config{
				User: UserConfig{Name: "john"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := installer.Validate(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
