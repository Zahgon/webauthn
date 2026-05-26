package metadata

import (
	"crypto/x509"
	"io"
	"net/http"
	"time"

	"github.com/go-viper/mapstructure/v2"
	"github.com/golang-jwt/jwt/v5"
)

// NewDecoder returns a new metadata decoder.
func NewDecoder(opts ...DecoderOption) (decoder *Decoder, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decoder handles decoding and specialized parsing of the metadata blob.
type Decoder struct {
	client                   *http.Client
	parser                   *jwt.Parser
	hook                     mapstructure.DecodeHookFunc
	root                     string
	ignoreEntryParsingErrors bool
}

// Parse handles parsing of the raw JSON values of the metadata blob. Should be used after using [Decoder.Decode] or
// [Decoder.DecodeBytes].
func (d *Decoder) Parse(payload *PayloadJSON) (metadata *Metadata, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decode the blob from an [io.Reader]. This function will close the [io.ReadCloser] after completing.
func (d *Decoder) Decode(r io.Reader) (payload *PayloadJSON, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DecodeBytes handles decoding raw bytes. If you have a read closer it's suggested to use [Decoder.Decode].
func (d *Decoder) DecodeBytes(bytes []byte) (payload *PayloadJSON, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 2. If the x5u attribute is present in the JWT Header.

// Never seen an x5u here, although it is in the spec.

// 3. If the x5u attribute is missing, the chain should be retrieved from the x5c attribute.

// If that attribute is missing as well, Metadata TOC signing trust anchor is considered the TOC signing certificate chain.

// The certificate chain MUST be verified to properly chain to the metadata TOC signing trust anchor.

// Chain validated, extract the TOC signing certificate from the chain. Create a buffer large enough to hold the
// certificate bytes.

// Decode the base64 certificate into the buffer.

// Parse the certificate from the buffer.

// 4. Verify the signature of the Metadata TOC object using the TOC signing certificate chain
// jwt.Parse() uses the TOC signing certificate public key internally to verify the signature.

// DecoderOption is a representation of a function that can set options within a decoder.
type DecoderOption func(decoder *Decoder) (err error)

// WithIgnoreEntryParsingErrors is a DecoderOption which ignores errors when parsing individual entries. The values for
// these entries will exist as an unparsed entry.
func WithIgnoreEntryParsingErrors() DecoderOption {
	_ = "STUB: not implemented"
	return *new(DecoderOption)
}

// WithRootCertificate overrides the root certificate used to validate the authenticity of the metadata payload.
func WithRootCertificate(value string) DecoderOption {
	_ = "STUB: not implemented"
	return *new(DecoderOption)
}

func validateChain(root string, chain []any) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func mdsParseX509Certificate(value string) (certificate *x509.Certificate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mdsParseTimePointer(format, value string) (parsed *time.Time, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
