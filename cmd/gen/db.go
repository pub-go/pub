package main

import (
	"code.gopub.tech/pub/dal/model"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithDefaultQuery,
	})

	g.ApplyBasic(model.User{}, model.Option{})

	g.Execute()
}
