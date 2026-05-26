package webauthn

import (
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
)

// LoginOption is a functional option that modifies the [protocol.PublicKeyCredentialRequestOptions] sent to the
// client during a login ceremony. Use the With* functions in this package (i.e. [WithUserVerification],
// [WithAllowedCredentials]) to create login options.
type LoginOption func(*protocol.PublicKeyCredentialRequestOptions)

// DiscoverableUserHandler is a callback function that the Relying Party must provide when performing a discoverable
// (passkey) login. It is called with the rawID of the credential and the userHandle from the authenticator response,
// and must return the [User] who owns the credential. This is necessary because in discoverable login flows, the
// Relying Party does not know which user is authenticating until the authenticator response is received.
type DiscoverableUserHandler func(rawID, userHandle []byte) (user User, err error)

// BeginLogin creates the [*protocol.CredentialAssertion] data payload that should be sent to the user agent for beginning
// the login/assertion process. This function is used to perform a login when the identity of the user is known such as
// multifactor authentications, to specify a conditional mediation requirement use [WebAuthn.BeginMediatedLogin], to
// perform a login when the identity of the user is not known see [WebAuthn.BeginDiscoverableLogin] and
// [WebAuthn.BeginDiscoverableMediatedLogin] instead. The format of this data can be seen in §5.5 of the WebAuthn
// specification. These default values can be amended by providing additional [LoginOption] parameters. This function
// also returns [SessionData], that must be stored by the RP in a secure manner and then provided to the
// [WebAuthn.FinishLogin] function. This data helps us verify the ownership of the credential being retrieved.
//
// Specification: §5.5. Options for Assertion Generation (https://www.w3.org/TR/webauthn/#dictionary-assertion-options)
func (webauthn *WebAuthn) BeginLogin(user User, opts ...LoginOption) (*protocol.CredentialAssertion, *SessionData, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// BeginDiscoverableLogin creates the [*protocol.CredentialAssertion] data payload that should be sent to the user agent
// for beginning the login/assertion process. This function is used to perform a client-side discoverable login when the
// identity of the user is not known such as passwordless or usernameless authentication, to specify a conditional
// mediation requirement use [WebAuthn.BeginDiscoverableMediatedLogin], to perform logins where the identity of the user
// is known such as multifactor authentication see [WebAuthn.BeginLogin] and [WebAuthn.BeginMediatedLogin] instead.
// The format of this data can be seen in §5.5 of the WebAuthn specification. These default values can be amended by
// providing additional [LoginOption] parameters. This function also returns [SessionData], that
// must be stored by the RP in a secure manner and then provided to the [WebAuthn.FinishLogin] function. This data helps
// us verify the ownership of the credential being retrieved.
//
// Specification: §5.5. Options for Assertion Generation (https://www.w3.org/TR/webauthn/#dictionary-assertion-options)
func (webauthn *WebAuthn) BeginDiscoverableLogin(opts ...LoginOption) (*protocol.CredentialAssertion, *SessionData, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// BeginMediatedLogin is similar to [WebAuthn.BeginLogin] however it also allows specifying a credential mediation
// requirement.
func (webauthn *WebAuthn) BeginMediatedLogin(user User, mediation protocol.CredentialMediationRequirement, opts ...LoginOption) (*protocol.CredentialAssertion, *SessionData, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// If the user does not have any credentials, we cannot perform an assertion.

// BeginDiscoverableMediatedLogin is similar to [WebAuthn.BeginDiscoverableLogin] however it also allows specifying a
// credential mediation requirement.
func (webauthn *WebAuthn) BeginDiscoverableMediatedLogin(mediation protocol.CredentialMediationRequirement, opts ...LoginOption) (*protocol.CredentialAssertion, *SessionData, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (webauthn *WebAuthn) beginLogin(userID []byte, allowedCredentials []protocol.CredentialDescriptor, mediation protocol.CredentialMediationRequirement, opts ...LoginOption) (assertion *protocol.CredentialAssertion, session *SessionData, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// FinishLogin takes the response from the client and validates it against the user credentials and stored session data.
//
// As with all Finish functions, this function requires a [*http.Request] but you can perform the same steps with the
// [protocol.ParseCredentialRequestResponseBody] or [protocol.ParseCredentialRequestResponseBytes] which require an
// [io.Reader] or byte array respectively, you can also use an arbitrary [*protocol.ParsedCredentialAssertionData] which is
// returned from all of these functions i.e. by implementing a custom parser. The [*SessionData],
// and [*protocol.ParsedCredentialAssertionData] can then be used with the [WebAuthn.ValidateLogin] function.
//
// This function will return the [protocol.ErrorUnknownCredential] error type when the [User] provided does not contain
// a [Credential] with the same ID byte array provided all [Credential]'s in the [SessionData] exist in the [User]'s
// [Credential] list.
func (webauthn *WebAuthn) FinishLogin(user User, session SessionData, response *http.Request) (credential *Credential, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FinishDiscoverableLogin takes the response from the client and validates it against the handler and stored session data.
// The handler helps to find out which user must be used to validate the response. This is a function defined in your
// business code that will retrieve the user from your persistent data.
//
// As with all Finish functions, this function requires a [*http.Request] but you can perform the same steps with the
// [protocol.ParseCredentialRequestResponseBody] or [protocol.ParseCredentialRequestResponseBytes] which require an
// [io.Reader] or byte array respectively, you can also use an arbitrary [*protocol.ParsedCredentialAssertionData] which is
// returned from all of these functions i.e. by implementing a custom parser. The [DiscoverableUserHandler], [*SessionData],
// and [*protocol.ParsedCredentialAssertionData] can then be used with the [WebAuthn.ValidatePasskeyLogin] function.
//
// This function will return the [protocol.ErrorUnknownCredential] error type when the [User] returned by the
// handler does not contain a [Credential] with the same ID byte array provided all [Credential]'s
// in the [SessionData] exist in the [User]'s [Credential] list.
func (webauthn *WebAuthn) FinishDiscoverableLogin(handler DiscoverableUserHandler, session SessionData, response *http.Request) (credential *Credential, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FinishPasskeyLogin takes the response from the client and validate it against the handler and stored session data.
// The handler helps to find out which user must be used to validate the response. This is a function defined in your
// business code that will retrieve the user from your persistent data.
//
// As with all Finish functions this function requires a [*http.Request] but you can perform the same steps with the
// [protocol.ParseCredentialRequestResponseBody] or [protocol.ParseCredentialRequestResponseBytes] which require an
// io.Reader or byte array respectively, you can also use an arbitrary [*protocol.ParsedCredentialAssertionData] which is
// returned from all of these functions i.e. by implementing a custom parser. The [DiscoverableUserHandler], [*SessionData],
// and [*protocol.ParsedCredentialAssertionData] can then be used with the [WebAuthn.ValidatePasskeyLogin] function.
//
// This function will return the [protocol.ErrorUnknownCredential] error type when the [User] returned by the
// handler does not contain a [Credential] with the same ID byte array provided all [Credential]'s
// in the [SessionData] exist in the [User]'s [Credential] list.
func (webauthn *WebAuthn) FinishPasskeyLogin(handler DiscoverableUserHandler, session SessionData, response *http.Request) (user User, credential *Credential, err error) {
	_ = "STUB: not implemented"
	return *new(User), nil, nil
}

// ValidateLogin takes a parsed response and validates it against the user credentials and session data.
//
// If you wish to skip performing the step required to parse the *protocol.ParsedCredentialAssertionData and
// you're using net/http then you can use [WebAuthn.FinishLogin] instead.
//
// This function will return the [protocol.ErrorUnknownCredential] error type when the [User] provided does not contain
// a [Credential] with the same ID byte array provided all [Credential]'s in the [SessionData] exist in
// the [User]'s [Credential] list.
func (webauthn *WebAuthn) ValidateLogin(user User, session SessionData, parsedResponse *protocol.ParsedCredentialAssertionData) (credential *Credential, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateDiscoverableLogin is similar to [WebAuthn.ValidateLogin] that allows for discoverable credentials. It's
// recommended that [WebAuthn.ValidatePasskeyLogin] is used instead.
//
// If you wish to skip performing the step required to parse the [*protocol.ParsedCredentialAssertionData] and
// you're using net/http then you can use [WebAuthn.FinishDiscoverableLogin] instead.
//
// This function will return the [protocol.ErrorUnknownCredential] error type when the [User] returned by the
// handler does not contain a [Credential] with the same ID byte array provided all [Credential]'s
// in the [SessionData] exist in the [User]'s [Credential] list.
//
// Note: this is just a backwards compatibility layer over [WebAuthn.ValidatePasskeyLogin] which returns more information.
func (webauthn *WebAuthn) ValidateDiscoverableLogin(handler DiscoverableUserHandler, session SessionData, parsedResponse *protocol.ParsedCredentialAssertionData) (credential *Credential, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidatePasskeyLogin is similar to [WebAuthn.ValidateLogin] that allows for discoverable credentials.
//
// If you wish to skip performing the step required to parse the [*protocol.ParsedCredentialAssertionData] and
// you're using net/http then you can use [WebAuthn.FinishPasskeyLogin] instead.
//
// This function will return the [protocol.ErrorUnknownCredential] error type when the [User] returned by the
// handler does not contain a [Credential] with the same ID byte array provided all [Credential]'s
// in the [SessionData] exist in the [User]'s [Credential] list.
func (webauthn *WebAuthn) ValidatePasskeyLogin(handler DiscoverableUserHandler, session SessionData, parsedResponse *protocol.ParsedCredentialAssertionData) (user User, credential *Credential, err error) {
	_ = "STUB: not implemented"
	return *new(User), nil, nil
}

// validateLogin takes a parsed response and validates it against the user credentials and session data.
//
//nolint:gocyclo
func (webauthn *WebAuthn) validateLogin(user User, session SessionData, parsedResponse *protocol.ParsedCredentialAssertionData) (*Credential, error) {
	_ = "STUB: not implemented"
	// Step 1. If the allowCredentials option was given when this authentication ceremony was initiated,
	// verify that credential.id identifies one of the public key credentials that were listed in
	// allowCredentials.
	return nil, nil
}

// NON-NORMATIVE Prior Step: Verify that the allowCredentials for the session are owned by the user provided.

// Step 2. If credential.response.userHandle is present, verify that the user identified by this value is
// the owner of the public key credential identified by credential.id. This is in part handled by our Step 1.

// Step 3. Using credential’s id attribute (or the corresponding rawId, if base64url encoding is inappropriate
// for your use case), look up the corresponding credential public key.

// Ensure authenticators with a bad status are not used.

// Handle steps 4 through 16.

// Check if the BackupEligible flag has changed.

// Check for the invalid combination BE=0 and BS=1.

// Handle step 17.

// Update flags from response data.
