package environment

const (
	PROD EnvironmentValue = "production"
	DEV  EnvironmentValue = "development"
)

type EnvironmentValue string

func (e EnvironmentValue) String() string {
	return string(e)
}
