package protocol

import (
	"github.com/go-webauthn/webauthn/metadata"
)

func init() {
	RegisterAttestationFormat(AttestationFormatCompound, attestationFormatValidationHandlerCompound)
}

// attestationFormatValidationHandlerCompound is the handler for the Compound Attestation Statement Format.
//
// The syntax of a Compound Attestation statement is defined by the following CDDL:
//
// $$attStmtType //= (
//
//	    fmt: "compound",
//	    attStmt: [2* nonCompoundAttStmt]
//	)
//
// nonCompoundAttStmt = { $$attStmtType } .within { fmt: text .ne "compound", * any => any }
//
// Specification: §8.9. Compound Attestation Statement Forma
//
// See: https://www.w3.org/TR/webauthn-3/#sctn-compound-attestation
//
//nolint:gocyclo
func attestationFormatValidationHandlerCompound(att AttestationObject, clientDataHash []byte, mds metadata.Provider) (attestationType string, x5cs []any, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}
