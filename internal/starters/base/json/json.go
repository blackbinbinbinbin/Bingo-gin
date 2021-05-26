package json

import (
	"github.com/json-iterator/go"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters"
)

var json jsoniter.API

func Json() jsoniter.API {
	starters.Check(json)
	return json
}

type JsonBaseStarter struct {
	starters.BaseStarter
}

func (d JsonBaseStarter) Setup(ctx starters.StarterContext) {
	json = jsoniter.ConfigCompatibleWithStandardLibrary
}