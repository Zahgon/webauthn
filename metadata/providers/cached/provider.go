package cached

import (
	"io"
	"net/http"
	"os"

	"github.com/go-webauthn/webauthn/metadata"
)

// New returns a new cached Provider given a set of functional [Option]'s. This provider will download a new version and
// save it to the configured file path if it doesn't exist or if it's out of date by default.
func New(opts ...Option) (provider metadata.Provider, err error) {
	_ = "STUB: not implemented"
	return *new(metadata.Provider), nil
}

// Provider implements a [metadata.Provider] with a file-based cache.
type Provider struct {
	metadata.Provider

	name    string
	uri     string
	update  bool
	force   bool
	clock   metadata.Clock
	client  *http.Client
	decoder *metadata.Decoder
	newup   NewFunc
}

func (p *Provider) init() (err error) {
	var (
		f       *os.File
		rc      io.ReadCloser
		created bool
		mds     *metadata.Metadata
	)

	if f, created, err = doOpenOrCreate(p.name); err != nil {
		return err
	}

	defer f.Close()

	if created || p.force {
		if rc, err = p.get(); err != nil {
			return err
		}
	} else {
		if mds, err = p.parse(f); err != nil {
			return err
		}

		if p.outdated(mds) {
			if rc, err = p.get(); err != nil {
				return err
			}
		}
	}

	if rc != nil {
		if err = doTruncateCopyAndSeekStart(f, rc); err != nil {
			return err
		}

		if mds, err = p.parse(f); err != nil {
			return err
		}
	}

	var provider metadata.Provider

	if provider, err = p.newup(mds); err != nil {
		return err
	}

	p.Provider = provider

	return nil
}

func (p *Provider) parse(rc io.ReadCloser) (data *metadata.Metadata, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Provider) outdated(mds *metadata.Metadata) bool { _ = "STUB: not implemented"; return false }

func (p *Provider) get() (f io.ReadCloser, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}
