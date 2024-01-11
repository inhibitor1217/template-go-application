package env

type Stage = string

const (
	StageDevelopment Stage = "development"
	StageTest        Stage = "test"
)

type Env struct {
	Stage Stage
}
