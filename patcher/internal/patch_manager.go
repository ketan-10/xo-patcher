package internal

import (
	"context"
	"errors"

	"github.com/google/wire"
)

type IPatchRunner interface {
	Run(ctx context.Context) error
}

type IPatchManager interface {
	Run(ctx context.Context, name string) error
	RegisterPatch(name string, runner IPatchRunner) error
}

type PatchManager struct {
	patchList map[string]IPatchRunner
}

func InitPatchManager() *PatchManager {
	return &PatchManager{
		patchList: map[string]IPatchRunner{},
	}
}

var NerPatchManager = wire.NewSet(InitPatchManager, wire.Bind(new(IPatchManager), new(PatchManager)))

func (pm *PatchManager) RegisterPatch(name string, runner IPatchRunner) error {
	if _, ok := pm.patchList[name]; ok {
		return errors.New("patch Already exists")
	}
	pm.patchList[name] = runner
	return nil
}

func (pm *PatchManager) Run(ctx context.Context, name string) error {
	if runner, ok := pm.patchList[name]; ok {
		runner.Run(ctx)
		return nil
	}
	return errors.New("patch by name: " + name + " not found")
}
