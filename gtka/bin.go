package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type bin struct {
	*container
	*gtk.Bin
}

func wrapBinSimple(v *gtk.Bin) *bin {
	if v == nil {
		return nil
	}
	return &bin{wrapContainerSimple(&v.Container), v}
}

func wrapBin(v *gtk.Bin, e error) (*bin, error) {
	return wrapBinSimple(v), e
}

func unwrapBin(v gtki.Bin) *gtk.Bin {
	if v == nil {
		return nil
	}
	return v.(*bin).Bin
}