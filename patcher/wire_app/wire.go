//go:build wireinject
// +build wireinject

package wire_app

import (
	"github.com/google/wire"
	"github.com/ketan-10/xo-patcher/patcher/internal"
	"github.com/ketan-10/xo-patcher/patcher/patches"
	"github.com/ketan-10/xo-patcher/patcher/xo_gen/xo_wire"
)

// To inject all patch to App
// This will allow calls to InitPatch method
// and we can also we can directly call run method of individual patch

var patchSet = wire.NewSet(
	patches.NewOne,
)

type PatchGroup struct {
	One patches.IOne
}

type App struct {
	PatchGroup PatchGroup
}

var globalSet = wire.NewSet(
	xo_wire.RepositorySet,
	wire.Struct(new(App), "*"),
	wire.Struct(new(PatchGroup), "*"),
	internal.NerPatchManager,
	patchSet,
)

func GetApp() (*App, func(), error) {
	wire.Build(globalSet)
	return &App{}, nil, nil
}
