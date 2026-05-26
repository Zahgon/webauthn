package webauthn

import (
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
)

// RegistrationOption is a functional option that modifies the [protocol.PublicKeyCredentialCreationOptions] sent
// to the client during a registration ceremony. Use the With* functions in this package (i.e.
// [WithConveyancePreference], [WithExclusions], [WithAuthenticatorSelection]) to create registration options.
type RegistrationOption func(*protocol.PublicKeyCredentialCreationOptions)

// BeginRegistration generates a new set of registration data to be sent to the client and authenticator. To set a
// conditional mediation requirement for the registration see [WebAuthn.BeginMediatedRegistration].
func (webauthn *WebAuthn) BeginRegistration(user User, opts ...RegistrationOption) (creation *protocol.CredentialCreation, session *SessionData, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// BeginMediatedRegistration is similar to [WebAuthn.BeginRegistration] however it also allows specifying a credential
// mediation requirement.
func (webauthn *WebAuthn) BeginMediatedRegistration(user User, mediation protocol.CredentialMediationRequirement, opts ...RegistrationOption) (creation *protocol.CredentialCreation, session *SessionData, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// FinishRegistration takes the response from the authenticator and client and verify the credential against the user's
// credentials and session data.
//
// As with all Finish functions this function requires a [*http.Request] but you can perform the same steps with the
// [protocol.ParseCredentialCreationResponseBody] or [protocol.ParseCredentialCreationResponseBytes] which require an
// [io.Reader] or byte array respectively, you can also use an arbitrary [*protocol.ParsedCredentialCreationData] which is
// returned from all of these functions i.e. by implementing a custom parser. The [User], [*SessionData], and
// [*protocol.ParsedCredentialCreationData] can then be used with the [WebAuthn.CreateCredential] function.
func (webauthn *WebAuthn) FinishRegistration(user User, session SessionData, request *http.Request) (credential *Credential, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateCredential verifies a parsed response against the user's credentials and session data.
//
// If you wish to skip performing the step required to parse the [*protocol.ParsedCredentialCreationData] and
// you're using net/http then you can use [WebAuthn.FinishRegistration] instead.
func (webauthn *WebAuthn) CreateCredential(user User, session SessionData, parsedResponse *protocol.ParsedCredentialCreationData) (credential *Credential, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateFilteredCredential applies the supplied [FilteringConfig] to a freshly-created [Credential]
// and returns a non-nil error when the credential violates any configured filtering rule (backup-eligibility
// prohibition, permitted-AAGUID allow-list, prohibited-AAGUID deny-list). A nil filtering argument is treated
// as "no filtering" and the function returns nil.
//
// The zero AAGUID ([uuid.Nil]) is never excluded by the permitted list, preserving the documented
// [FilteringConfig] contract for authenticators that report no AAGUID.
//
// This function is invoked automatically by [WebAuthn.CreateCredential] using the [Config.Filtering] value;
// relying parties may also call it directly (e.g. to pre-validate a credential before persistence) with any
// FilteringConfig value of their choosing.
//
// The credential argument must be non-nil.
func ValidateFilteredCredential(credential *Credential, filtering *FilteringConfig) (err error) {
	_ = "STUB: not implemented"
	return nil
}
