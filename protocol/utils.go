package protocol

import (
	"crypto/x509"
	"time"

	"github.com/go-webauthn/webauthn/protocol/webauthncose"
)

func mustParseX509Certificate(der []byte) *x509.Certificate { _ = "STUB: not implemented"; return nil }

func mustParseX509CertificatePEM(raw []byte) *x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func attStatementParseX5CS(attStatement map[string]any, key string) (x5c []any, x5cs []*x509.Certificate, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func parseX5C(x5c []any) (x5cs []*x509.Certificate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// attStatementCertChainVerify allows verifying an attestation statement certificate chain and optionally allows
// mangling the not after value for purpose of just validating the attestation lineage. If you set mangleNotAfter to
// true this function should only be considered safe for determining lineage, and not hte validity of a chain in
// general.
//
// WARNING: Setting mangleNotAfter=true weakens security by accepting expired certificates.
func attStatementCertChainVerify(certs []*x509.Certificate, roots *x509.CertPool, mangleNotAfter bool, mangleNotAfterSafeTime time.Time) (chains [][]*x509.Certificate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isSelfSigned(c *x509.Certificate) bool { _ = "STUB: not implemented"; return false }

// This function is used to intentionally but conditionally mangle the certificate not after value to exclude it from
// the verification process. This should only be used in instances where all you care about is which certificates
// performed the signing.
//
// WARNING: Setting mangle=true weakens security by accepting expired certificates.
func certInsecureConditionalNotAfterMangle(cert *x509.Certificate, mangle bool, safe time.Time) (out *x509.Certificate) {
	_ = "STUB: not implemented"
	return nil
}

func verifyAttestationECDSAPublicKeyMatch(att AttestationObject, cert *x509.Certificate) (attPublicKeyData webauthncose.EC2PublicKeyData, err error) {
	_ = "STUB: not implemented"
	return *new(webauthncose.EC2PublicKeyData), nil
}

// ValidateRPID performs non-exhaustive checks to ensure the string is most likely a domain string as
// relying-party ID's are required to be. Effectively this can be an IP, localhost, or a string that contains a period.
// The relying-party ID must not contain scheme, port, path, query, or fragment components.
//
// See: https://www.w3.org/TR/webauthn/#rp-id
//
//nolint:gocyclo
func ValidateRPID(value string) (err error) { _ = "STUB: not implemented"; return nil }

// IsAttestationFormatString reports whether s is one of the WebAuthn-defined attestation statement format
// identifiers. Used to detect and migrate records from prior releases which stored
// the format string in the AttestationType field.
func IsAttestationFormatString(s string) bool { _ = "STUB: not implemented"; return false }
