package protocol

// Error is a struct that describes specific error conditions in a structured format.
type Error struct {
	// Short name for the type of error that has occurred.
	Type string `json:"type"`

	// Additional details about the error.
	Details string `json:"error"`

	// Information to help debug the error.
	DevInfo string `json:"debug"`

	// Inner error.
	Err error `json:"-"`
}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Error) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *Error) WithDetails(details string) *Error { _ = "STUB: not implemented"; return nil }

func (e *Error) WithInfo(info string) *Error { _ = "STUB: not implemented"; return nil }

func (e *Error) WithError(err error) *Error { _ = "STUB: not implemented"; return nil }

// ErrorUnknownCredential is a special Error which signals the fact the provided credential is unknown. The reason this
// specific error type is useful is so that the relying-party can send a signal to the Authenticator that the
// credential has been removed.
type ErrorUnknownCredential struct {
	Err *Error
}

func (e *ErrorUnknownCredential) Error() string { _ = "STUB: not implemented"; return "" }

func (e *ErrorUnknownCredential) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *ErrorUnknownCredential) copy() ErrorUnknownCredential {
	_ = "STUB: not implemented"
	return *new(ErrorUnknownCredential)
}

func (e *ErrorUnknownCredential) WithDetails(details string) *ErrorUnknownCredential {
	_ = "STUB: not implemented"
	return nil
}

func (e *ErrorUnknownCredential) WithInfo(info string) *ErrorUnknownCredential {
	_ = "STUB: not implemented"
	return nil
}

func (e *ErrorUnknownCredential) WithError(err error) *ErrorUnknownCredential {
	_ = "STUB: not implemented"
	return nil
}

var (
	ErrBadRequest = &Error{
		Type:    "invalid_request",
		Details: "Error reading the request data",
	}
	ErrPolicyRestriction = &Error{
		Type:    "policy_restriction",
		Details: "Policy restriction prevented the operation from completing",
	}
	ErrChallengeMismatch = &Error{
		Type:    "challenge_mismatch",
		Details: "Stored challenge and received challenge do not match",
	}
	ErrParsingData = &Error{
		Type:    "parse_error",
		Details: "Error parsing the authenticator response",
	}
	ErrAuthData = &Error{
		Type:    "auth_data",
		Details: "Error verifying the authenticator data",
	}
	ErrVerification = &Error{
		Type:    "verification_error",
		Details: "Error validating the authenticator response",
	}
	ErrAttestation = &Error{
		Type:    "attestation_error",
		Details: "Error validating the attestation data provided",
	}
	ErrInvalidAttestation = &Error{
		Type:    "invalid_attestation",
		Details: "Invalid attestation data",
	}
	ErrMetadata = &Error{
		Type:    "invalid_metadata",
		Details: "",
	}
	ErrAttestationFormat = &Error{
		Type:    "invalid_attestation",
		Details: "Invalid attestation format",
	}
	ErrAttestationCertificate = &Error{
		Type:    "invalid_certificate",
		Details: "Invalid attestation certificate",
	}
	ErrAssertionSignature = &Error{
		Type:    "invalid_signature",
		Details: "Assertion Signature against auth data and client hash is not valid",
	}
	ErrUnsupportedKey = &Error{
		Type:    "invalid_key_type",
		Details: "Unsupported Public Key Type",
	}
	ErrUnsupportedAlgorithm = &Error{
		Type:    "unsupported_key_algorithm",
		Details: "Unsupported public key algorithm",
	}
	ErrNotSpecImplemented = &Error{
		Type:    "spec_unimplemented",
		Details: "This field is not yet supported by the WebAuthn spec",
	}
	ErrNotImplemented = &Error{
		Type:    "not_implemented",
		Details: "This field is not yet supported by this library",
	}
)
