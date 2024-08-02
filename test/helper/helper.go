package helper

import (
	"github.com/seppedelanghe/mizu/pkg/engine"
)

type SingleEntityScene struct {
	setup func(w engine.World)
}

func RunSingleSceneGame(setup func(w engine.World)) {
	scene := &SingleEntityScene{setup}
	_ = engine.NewGame(scene).Update()
}

func (s *SingleEntityScene) Setup(w engine.World) {
	s.setup(w)
}
