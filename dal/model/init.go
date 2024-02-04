package model

var autoGen = []any{}
var autoMigrates = []any{}

func AutoGen() []any      { return autoGen }
func AutoMigrates() []any { return autoMigrates }

func register(a any) {
	autoGen = append(autoGen, a)
	autoMigrates = append(autoMigrates, &a)
}
