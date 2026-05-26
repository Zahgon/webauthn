package memory

import (
	"github.com/google/uuid"

	"github.com/go-webauthn/webauthn/metadata"
)

// Option describes an optional pattern for this provider.
type Option func(provider *Provider) (err error)

// WithMetadata provides the required metadata for the memory provider.
func WithMetadata(mds map[uuid.UUID]*metadata.Entry) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithValidateEntry requires that the provided metadata has an entry for the given authenticator to be considered
// valid. By default an AAGUID which has a zero value should fail validation if [WithValidateEntryPermitZeroAAGUID] is not
// provided with the value of true. Default is true.
func WithValidateEntry(require bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithValidateEntryPermitZeroAAGUID is an option that permits a zero'd AAGUID from an attestation statement to
// automatically pass metadata validations. Generally helpful to use with [WithValidateEntry]. Default is false.
func WithValidateEntryPermitZeroAAGUID(permit bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithValidateTrustAnchor when set to true enables the validation of the attestation statement against the trust anchor
// from the metadata. Default is true.
func WithValidateTrustAnchor(validate bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithValidateStatus when set to true enables the validation of the attestation statements AAGUID against the desired
// and undesired [metadata.AuthenticatorStatus] lists. Default is true.
func WithValidateStatus(validate bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithValidateAttestationTypes when set to true enables the validation of the attestation statements type against the
// known types the authenticator can produce. Default is true.
func WithValidateAttestationTypes(validate bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithStatusUndesired provides the list of statuses which are considered undesirable for status report validation
// purposes. Should be used with [WithValidateStatus] set to true.
func WithStatusUndesired(statuses []metadata.AuthenticatorStatus) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithStatusDesired provides the list of statuses which are considered desired and will be required for status report
// validation purposes. Should be used with [WithValidateStatus] set to true.
func WithStatusDesired(statuses []metadata.AuthenticatorStatus) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
