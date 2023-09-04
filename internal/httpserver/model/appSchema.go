package model

type Schema map[string]TypeSchema

type AppSchema struct {
	Data Schema
	Common
}
