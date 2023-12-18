package main

import (
	"known-anchors/dal/db/model"

	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "dal/db/dao",
		ModelPkgPath: "entity",

		Mode:              gen.WithDefaultQuery | gen.WithoutContext | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  true,
	})

	g.ApplyBasic(model.Post{}, model.User{}, model.Comment{}, model.Deck{}, model.Card{})

	g.ApplyInterface(func(model.Method) {}, model.User{}, model.Post{}, model.Comment{}, model.Deck{}, model.Card{})

	g.ApplyInterface(func(model.UserMethod) {}, model.User{})

	g.ApplyInterface(func(model.DeckMethod) {}, model.Deck{})
	g.Execute()
}
