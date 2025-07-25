package store

import (
	"context"
	"fmt"
	"strings"

	"oras.land/oras-go/pkg/content"

	"hauler.dev/go/hauler/internal/flags"
	"hauler.dev/go/hauler/pkg/cosign"
	"hauler.dev/go/hauler/pkg/log"
	"hauler.dev/go/hauler/pkg/store"
)

func CopyCmd(ctx context.Context, o *flags.CopyOpts, s *store.Layout, targetRef string, ro *flags.CliRootOpts) error {
	l := log.FromContext(ctx)

	if o.Username != "" || o.Password != "" {
		return fmt.Errorf("--username/--password have been deprecated, please use 'hauler login'")
	}

	components := strings.SplitN(targetRef, "://", 2)
	switch components[0] {
	case "dir":
		l.Debugf("identified directory target reference of [%s]", components[1])
		fs := content.NewFile(components[1])
		defer fs.Close()

		_, err := s.CopyAll(ctx, fs, nil)
		if err != nil {
			return err
		}

	case "registry":
		l.Debugf("identified registry target reference of [%s]", components[1])
		ropts := content.RegistryOptions{
			Insecure:  o.Insecure,
			PlainHTTP: o.PlainHTTP,
		}

		err := cosign.LoadImages(ctx, s, components[1], o.Only, ropts, ro)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("detecting protocol from [%s]", targetRef)
	}

	l.Infof("copied artifacts to [%s]", components[1])
	return nil
}
