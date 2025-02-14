package none

import (
	"context"
	"fmt"

	"github.com/reproducible-containers/repro-get/pkg/cache"
	"github.com/reproducible-containers/repro-get/pkg/distro"
	"github.com/reproducible-containers/repro-get/pkg/filespec"
)

const Name = "none"

var ErrNotImplemented = fmt.Errorf("distro driver %q does not implement the requested feature", Name)

func New() distro.Distro {
	d := &none{
		info: distro.Info{
			Name:             Name,
			DefaultProviders: nil,
		},
	}
	return d
}

type none struct {
	info distro.Info
}

func (d *none) Info() distro.Info {
	return d.info
}

func (d *none) GenerateHash(ctx context.Context, hw distro.HashWriter, opts distro.HashOpts) error {
	return ErrNotImplemented
}

func (d *none) PackageName(sp filespec.FileSpec) (string, error) {
	return "", ErrNotImplemented
}

func (d *none) IsPackageVersionInstalled(ctx context.Context, sp filespec.FileSpec) (bool, error) {
	// No need to return ErrNotImplemented
	return false, nil
}

func (d *none) InstallPackages(ctx context.Context, c *cache.Cache, pkgs []filespec.FileSpec, opts distro.InstallOpts) error {
	if len(pkgs) == 0 {
		return nil
	}
	return ErrNotImplemented
}

func (d *none) GenerateDockerfile(ctx context.Context, dir string, args distro.DockerfileTemplateArgs, opts distro.DockerfileOpts) error {
	return ErrNotImplemented
}
