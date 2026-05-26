package protocol

import (
	"net/url"
)

// CollectedClientData represents the contextual bindings of both the WebAuthn Relying Party
// and the client. It is a key-value mapping whose keys are strings. Values can be any type
// that has a valid encoding in JSON. Its structure is defined by the following Web IDL.
//
// Specification: §5.8.1. Client Data Used in WebAuthn Signatures (https://www.w3.org/TR/webauthn/#dictdef-collectedclientdata)
type CollectedClientData struct {
	// Type contains the string "webauthn.create" when creating new credentials, and "webauthn.get" when getting an
	// assertion from an existing credential. The purpose of this member is to prevent certain types of signature
	// confusion attacks (where an attacker substitutes one legitimate signature for another).
	Type CeremonyType `json:"type"`

	// Challenge contains the base64url encoding of the challenge provided by the Relying Party.
	Challenge string `json:"challenge"`

	// Origin contains the fully qualified origin of the requester, as provided to the authenticator by the client.
	Origin string `json:"origin"`

	// TopOrigin contains the fully qualified top-level origin of the requester when the client is cross-origin.
	// This is only present when CrossOrigin is true.
	//
	// WebAuthn Level 3.
	TopOrigin string `json:"topOrigin,omitempty"`

	// CrossOrigin indicates whether the calling context is an iframe that is not same-origin with its ancestor.
	//
	// WebAuthn Level 3.
	CrossOrigin bool `json:"crossOrigin,omitempty"`

	// TokenBinding contains information about the state of the Token Binding protocol.
	TokenBinding *TokenBinding `json:"tokenBinding,omitempty"`

	// Hint is an opaque field that may be added by the client. Chromium-based browsers include this field to remind
	// implementers not to perform string comparison on the clientDataJSON.
	Hint string `json:"new_keys_may_be_added_here,omitempty"`
}

// CeremonyType represents the type of WebAuthn ceremony being performed.
//
// Specification: §5.8.1. Client Data Used in WebAuthn Signatures (https://www.w3.org/TR/webauthn/#dom-collectedclientdata-type)
type CeremonyType string

const (
	// CreateCeremony is the ceremony type for credential registration ("webauthn.create").
	CreateCeremony CeremonyType = "webauthn.create"

	// AssertCeremony is the ceremony type for authentication assertion ("webauthn.get").
	AssertCeremony CeremonyType = "webauthn.get"
)

// TokenBinding contains information about the state of the Token Binding protocol used when communicating with the
// Relying Party. Its absence indicates that the client doesn't support token binding.
//
// Specification: §5.8.1. Client Data Used in WebAuthn Signatures (https://www.w3.org/TR/webauthn/#dom-collectedclientdata-tokenbinding)
type TokenBinding struct {
	Status TokenBindingStatus `json:"status"`
	ID     string             `json:"id,omitempty"`
}

// TokenBindingStatus represents the state of Token Binding between the client and the Relying Party.
type TokenBindingStatus string

const (
	// Present indicates token binding was used when communicating with the
	// Relying Party. In this case, the id member MUST be present.
	Present TokenBindingStatus = "present"

	// Supported indicates the client supports token binding, but it was not
	// negotiated when communicating with the Relying Party.
	Supported TokenBindingStatus = "supported"

	// NotSupported indicates token binding not supported
	// when communicating with the Relying Party.
	NotSupported TokenBindingStatus = "not-supported"
)

// FullyQualifiedOrigin returns the origin per the HTML spec: (scheme)://(host)[:(port)].
func FullyQualifiedOrigin(rawOrigin string) (fqOrigin string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Verify handles steps 3 through 6 of verifying the registering client data of a
// new credential and steps 7 through 10 of verifying an authentication assertion
// See https://www.w3.org/TR/webauthn/#registering-a-new-credential
// and https://www.w3.org/TR/webauthn/#verifying-assertion
//
// Note: the rpTopOriginsVerify parameter does not accept the TopOriginVerificationMode value of
// TopOriginDefaultVerificationMode as it's expected this value is updated by the config validation process.
//
//nolint:gocyclo
func (c *CollectedClientData) Verify(storedChallenge string, ceremony CeremonyType, rpOrigins, rpTopOrigins []string, rpTopOriginsVerify TopOriginVerificationMode, allowCrossOrigin bool) (err error) {
	_ = "STUB: not implemented"
	// Registration Step 3. Verify that the value of C.type is webauthn.create.
	return nil
}

// Assertion Step 7. Verify that the value of C.type is the string webauthn.get.

// Registration Step 4. Verify that the value of C.challenge matches the challenge
// that was sent to the authenticator in the create() call.

// Assertion Step 8. Verify that the value of C.challenge matches the challenge
// that was sent to the authenticator in the PublicKeyCredentialRequestOptions
// passed to the get() call.

// Registration Step 5 & Assertion Step 9. Verify that the value of C.origin matches
// the Relying Party's origin.

// Registration Step 6 and Assertion Step 10. Verify that the value of C.tokenBinding.status
// matches the state of Token Binding for the TLS connection over which the assertion was
// obtained. If Token Binding was used on that TLS connection, also verify that C.tokenBinding.id
// matches the base64url encoding of the Token Binding ID for the connection.

// Not yet fully implemented by the spec, browsers, and me.

// TopOriginVerificationMode determines how the Relying Party validates the topOrigin field in
// [CollectedClientData]. This is relevant for cross-origin iframe scenarios where the top-level browsing context's
// origin differs from the embedded origin making the WebAuthn API call.
//
// WebAuthn Level 3.
type TopOriginVerificationMode int

const (
	// TopOriginDefaultVerificationMode is the zero value of [TopOriginVerificationMode] and has no matching rule in
	// the verifier; passing it directly to [CollectedClientData.Verify] returns an "unknown Top Origin verification
	// mode" error. High-level callers using [webauthn.Config] have this value coerced to
	// [TopOriginExplicitVerificationMode] by config validation, which is the recommended default.
	TopOriginDefaultVerificationMode TopOriginVerificationMode = iota

	// TopOriginAutoVerificationMode accepts the Top Origin if it matches any entry in either the allowed Top Origins
	// list or the allowed Origins list. The two lists are unioned (RPTopOrigins ∪ RPOrigins). This is the most
	// permissive of the three active modes and should only be used when an RP deliberately wants cross-origin and
	// same-origin embeddings to share an allow-list.
	TopOriginAutoVerificationMode

	// TopOriginImplicitVerificationMode accepts the Top Origin only if it matches an entry in the allowed Origins
	// list (RPOrigins). The RPTopOrigins list is ignored in this mode.
	TopOriginImplicitVerificationMode

	// TopOriginExplicitVerificationMode accepts the Top Origin only if it matches an entry in the allowed Top Origins
	// list (RPTopOrigins). The RPOrigins list is ignored in this mode. This is the strictest mode and the one
	// [webauthn.Config] coerces the zero value to.
	TopOriginExplicitVerificationMode
)

// IsOriginInHaystack checks if the needle is in the haystack using the mechanism to determine origin equality defined
// in HTML5 Section 5.3 and RFC3986 Section 6.2.1.
//
// Specifically if the needle value has the 'http://' or 'https://' prefix (case-insensitive) and can be parsed as a
// URL; we check each item in the haystack to see if it matches the same rules, and then if the scheme and host (with
// a normalized port) components match case-insensitively then they're considered a match.
//
// If the needle value does not have the 'http://' or 'https://' prefix (case-insensitive) or can't be parsed as a URL
// equality is determined using simple string comparison.
//
// It is important to note that this function completely ignores Apple Associated Domains entirely as Apple is using
// an unassigned Well-Known URI in breech of Well-Known Uniform Resource Identifiers (RFC8615).
//
// See (Origin Definition): https://www.w3.org/TR/2011/WD-html5-20110525/origin-0.html
//
// See (Simple String Comparison Definition): https://datatracker.ietf.org/doc/html/rfc3986#section-6.2.1
//
// See (Apple Associated Domains): https://developer.apple.com/documentation/xcode/supporting-associated-domains
//
// See (IANA Well Known URI Assignments): https://www.iana.org/assignments/well-known-uris/well-known-uris.xhtml
//
// See (Well-Known Uniform Resource Identifiers): https://datatracker.ietf.org/doc/html/rfc8615
func IsOriginInHaystack(needle string, haystack []string) bool {
	_ = "STUB: not implemented"
	return false
}

func isOriginEqual(a *url.URL, b *url.URL) bool { _ = "STUB: not implemented"; return false }

func parseOriginURI(raw string) *url.URL { _ = "STUB: not implemented"; return nil }

// We can ignore the error here because it's effectively not a FQDN if this fails.

// Normalize the port if necessary.

func isPossibleFQDN(raw string) bool { _ = "STUB: not implemented"; return false }
