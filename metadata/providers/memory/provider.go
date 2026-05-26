package memory

import (
	"context"

	"github.com/google/uuid"

	"github.com/go-webauthn/webauthn/metadata"
)

// New returns a new memory Provider given a set of functional Option's.
func New(opts ...Option) (provider metadata.Provider, err error) {
	_ = "STUB: not implemented"
	return *new(metadata.Provider), nil
}

// Provider is a concrete implementation of the [metadata.Provider] that utilizes memory for validation. This provider is
// a simple one-shot that doesn't perform any locking, provide dynamic functionality, or download the metadata at any
// stage (it expects it's provided via one of the Option's).
type Provider struct {
	mds             map[uuid.UUID]*metadata.Entry
	desired         []metadata.AuthenticatorStatus
	undesired       []metadata.AuthenticatorStatus
	entry           bool
	entryPermitZero bool
	anchors         bool
	status          bool
	attestation     bool
}

func (p *Provider) GetEntry(ctx context.Context, aaguid uuid.UUID) (entry *metadata.Entry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Provider) GetValidateEntry(ctx context.Context) (require bool) {
	_ = "STUB: not implemented"
	return false
}

func (p *Provider) GetValidateEntryPermitZeroAAGUID(ctx context.Context) (skip bool) {
	_ = "STUB: not implemented"
	return false
}

func (p *Provider) GetValidateTrustAnchor(ctx context.Context) (validate bool) {
	_ = "STUB: not implemented"
	return false
}

func (p *Provider) GetValidateStatus(ctx context.Context) (validate bool) {
	_ = "STUB: not implemented"
	return false
}

func (p *Provider) GetValidateAttestationTypes(ctx context.Context) (validate bool) {
	_ = "STUB: not implemented"
	return false
}

func (p *Provider) ValidateStatusReports(ctx context.Context, reports []metadata.StatusReport) (err error) {
	_ = "STUB: not implemented"
	return nil
}

var (
	_ metadata.Provider = (*Provider)(nil)
)
