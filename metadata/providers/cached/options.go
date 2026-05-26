package cached

import (
	"net/http"

	"github.com/go-webauthn/webauthn/metadata"
)

// Option describes an optional pattern for this provider.
type Option func(provider *Provider) (err error)

// NewFunc describes the type used to create the underlying provider.
type NewFunc func(mds *metadata.Metadata) (provider metadata.Provider, err error)

// WithPath sets the path name for the cached file. This option is REQUIRED.
func WithPath(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUpdate is used to enable or disable the update. By default it's set to true.
func WithUpdate(update bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithForceUpdate is used to force an update on creation. This will forcibly overwrite the file if possible.
func WithForceUpdate(force bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNew customizes the NewFunc. By default we just create a fairly standard [memory.Provider] with strict defaults.
func WithNew(newup NewFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDecoder sets the decoder to be used for this provider. By default this is a decoder with the entry parsing errors
// configured to skip that entry.
func WithDecoder(decoder *metadata.Decoder) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetadataURL configures the URL to get the metadata from. This shouldn't be modified unless you know what you're
// doing as we use the [metadata.ProductionMDSURL] which is safe in most instances.
func WithMetadataURL(uri string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithClient configures the [*http.Client] used to get the MDS3 blob.
func WithClient(client *http.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithClock allows injection of a [metadata.Clock] to check the up-to-date status of a blob.
func WithClock(clock metadata.Clock) Option { _ = "STUB: not implemented"; return *new(Option) }
