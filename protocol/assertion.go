package protocol

import (
	"io"
	"net/http"
)

// The CredentialAssertionResponse is the raw response returned to the Relying Party from an authenticator when we request a
// credential for login/assertion.
type CredentialAssertionResponse struct {
	PublicKeyCredential

	AssertionResponse AuthenticatorAssertionResponse `json:"response"`
}

// The ParsedCredentialAssertionData is the parsed [CredentialAssertionResponse] that has been marshalled into a format
// that allows us to verify the client and authenticator data inside the response.
type ParsedCredentialAssertionData struct {
	ParsedPublicKeyCredential

	Response ParsedAssertionResponse
	Raw      CredentialAssertionResponse
}

// The AuthenticatorAssertionResponse contains the raw authenticator assertion data and is parsed into
// [ParsedAssertionResponse].
type AuthenticatorAssertionResponse struct {
	AuthenticatorResponse

	AuthenticatorData URLEncodedBase64 `json:"authenticatorData"`
	Signature         URLEncodedBase64 `json:"signature"`
	UserHandle        URLEncodedBase64 `json:"userHandle,omitempty"`
}

// ParsedAssertionResponse is the parsed form of [AuthenticatorAssertionResponse].
type ParsedAssertionResponse struct {
	CollectedClientData CollectedClientData
	AuthenticatorData   AuthenticatorData
	Signature           []byte
	UserHandle          []byte
}

// ParseCredentialRequestResponse parses a login/assertion response from a [*http.Request]. The request body is
// automatically drained and closed after parsing.
//
// This is the standard entry point when using [net/http]. For implementations that don't use [net/http], see
// [ParseCredentialRequestResponseBody] (accepts an [io.Reader]) or [ParseCredentialRequestResponseBytes] (accepts a
// []byte).
func ParseCredentialRequestResponse(response *http.Request) (*ParsedCredentialAssertionData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseCredentialRequestResponseBody parses a login/assertion response from an [io.Reader]. The caller is responsible
// for closing the reader if applicable.
//
// This is the framework-agnostic variant of [ParseCredentialRequestResponse]. For a [*http.Request] use
// [ParseCredentialRequestResponse] instead. For raw bytes use [ParseCredentialRequestResponseBytes].
func ParseCredentialRequestResponseBody(body io.Reader) (par *ParsedCredentialAssertionData, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseCredentialRequestResponseBytes parses a login/assertion response from raw bytes.
//
// See also [ParseCredentialRequestResponse] (for [*http.Request]) and [ParseCredentialRequestResponseBody] (for
// [io.Reader]).
func ParseCredentialRequestResponseBytes(data []byte) (par *ParsedCredentialAssertionData, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse validates and parses the [CredentialAssertionResponse] into a [ParsedCredentialAssertionData]. Most
// implementations should use [ParseCredentialRequestResponse], [ParseCredentialRequestResponseBody], or
// [ParseCredentialRequestResponseBytes] instead of calling this method directly.
func (car CredentialAssertionResponse) Parse() (par *ParsedCredentialAssertionData, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Step 5. Let JSONtext be the result of running UTF-8 decode on the value of cData.
// We don't call it cData but this is Step 5 in the spec.

// Verify the remaining elements of the assertion data by following the steps outlined in the referenced specification
// documentation. It's important to note that the credentialBytes field is the CBOR representation of the credential.
//
// Specification: §7.2 Verifying an Authentication Assertion (https://www.w3.org/TR/webauthn/#sctn-verifying-assertion)
func (p *ParsedCredentialAssertionData) Verify(storedChallenge string, relyingPartyID, appID string, rpOrigins, rpTopOrigins []string, rpTopOriginsVerify TopOriginVerificationMode, allowCrossOrigin, verifyUser, verifyUserPresence bool, credentialBytes []byte) error {
	_ = "STUB: not implemented"
	// Steps 4 through 6 in verifying the assertion data (https://www.w3.org/TR/webauthn/#verifying-assertion) are
	// "assertive" steps, i.e. "Let JSONtext be the result of running UTF-8 decode on the value of cData."
	// We handle these steps in part as we verify but also beforehand
	//
	// Handle steps 7 through 10 of assertion by verifying stored data against the Collected Client Data
	// returned by the authenticator.
	return nil
}

// Begin Step 11. Verify that the rpIdHash in authData is the SHA-256 hash of the RP ID expected by the RP.

// Handle steps 11 through 14, verifying the authenticator data.

// Step 15. Let hash be the result of computing a hash over the cData using SHA-256.

// Step 16. Using the credential public key looked up in step 3, verify that sig is
// a valid signature over the binary concatenation of authData and hash.

//nolint:gocritic // This is intentional.

// If the Session Data does not contain the appID extension or it wasn't reported as used by the Client/RP then we
// use the standard CTAP2 public key parser.
