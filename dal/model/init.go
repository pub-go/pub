package model

var autoGen = []any{}
var autoMigrates = []any{}

func AutoGen() []any      { return autoGen }
func AutoMigrates() []any { return autoMigrates }

func register(a ...any) {
	autoGen = append(autoGen, a...)
	for _, m := range a {
		m := m
		autoMigrates = append(autoMigrates, &m)
	}
}
