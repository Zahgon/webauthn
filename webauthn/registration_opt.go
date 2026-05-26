package webauthn

import "github.com/go-webauthn/webauthn/protocol"

// WithCredentialParameters adjusts the credential parameters in the registration options.
//
// Specification: §5.4. Parameters for Credential Generation (https://www.w3.org/TR/webauthn/#dom-publickeycredentialcreationoptions-pubkeycredparams)
func WithCredentialParameters(credentialParams []protocol.CredentialParameter) RegistrationOption {
	_ = "STUB: not implemented"
	return *new(RegistrationOption)
}

// WithExclusions adjusts the non-default parameters regarding credentials to exclude from registration.
//
// Specification: §5.4. Parameters for Credential Generation (https://www.w3.org/TR/webauthn/#dom-publickeycredentialcreationoptions-excludecredentials)
func WithExclusions(excludeList []protocol.CredentialDescriptor) RegistrationOption {
	_ = "STUB: not implemented"
	return *new(RegistrationOption)
}

// WithAuthenticatorSelection adjusts the non-default parameters regarding the authenticator to select during
// registration.
//
// Specification: §5.4. Parameters for Credential Generation (https://www.w3.org/TR/webauthn/#dom-publickeycredentialcreationoptions-authenticatorselection)
//
// Specification: §5.4.4. Authenticator Selection Criteria (https://www.w3.org/TR/webauthn/#dictdef-authenticatorselectioncriteria)
func WithAuthenticatorSelection(authenticatorSelection protocol.AuthenticatorSelection) RegistrationOption {
	_ = "STUB: not implemented"
	return *new(RegistrationOption)
}

// WithResidentKeyRequirement sets both the resident key and require resident key protocol options.
//
// Specification: §5.4. Parameters for Credential Generation (https://www.w3.org/TR/webauthn/#dom-publickeycredentialcreationoptions-authenticatorselection)
//
// Specification: §5.4.4. Authenticator Selection Criteria (https://www.w3.org/TR/webauthn/#dictdef-authenticatorselectioncriteria)
func WithResidentKeyRequirement(requirement protocol.ResidentKeyRequirement) RegistrationOption {
	_ = "STUB: not implemented"
	return *new(RegistrationOption)
}

// WithPublicKeyCredentialHints adjusts the non-default hints for credential types to select during registration.
//
// WebAuthn Level 3.
//
// Specification: §5.4. Parameters for Credential Generation (https://www.w3.org/TR/webauthn-3/#dom-publickeycredentialcreationoptions-hints)
func WithPublicKeyCredentialHints(hints []protocol.PublicKeyCredentialHints) RegistrationOption {
	_ = "STUB: not implemented"
	return *new(RegistrationOption)
}

// WithConveyancePreference adjusts the non-default parameters regarding whether the authenticator should attest to the
// credential.
//
// Specification: §5.4. Parameters for Credential Generation (https://www.w3.org/TR/webauthn/#dom-publickeycredentialcreationoptions-attestation)
func WithConveyancePreference(preference protocol.ConveyancePreference) RegistrationOption {
	_ = "STUB: not implemented"
	return *new(RegistrationOption)
}

// WithAttestationFormats adjusts the non-default formats for credential types to select during registration.
//
// WebAuthn Level 3.
//
// Specification: §5.4. Parameters for Credential Generation (https://www.w3.org/TR/webauthn-3/#dom-publickeycredentialcreationoptions-attestationformats)
func WithAttestationFormats(formats []protocol.AttestationFormat) RegistrationOption {
	_ = "STUB: not implemented"
	return *new(RegistrationOption)
}

// WithExtensions adjusts the extension parameter in the registration options.
//
// Specification: §5.4. Parameters for Credential Generation (https://www.w3.org/TR/webauthn-3/#dom-publickeycredentialcreationoptions-extensions)
//
// Specification: §9. Extensions (https://www.w3.org/TR/webauthn/#webauthn-extensions)
func WithExtensions(extension protocol.AuthenticationExtensions) RegistrationOption {
	_ = "STUB: not implemented"
	return *new(RegistrationOption)
}

// WithAppIdExcludeExtension automatically includes the specified appid if the CredentialExcludeList contains a credential
// with the type `fido-u2f`.
//
// Specification: §5.4. Parameters for Credential Generation (https://www.w3.org/TR/webauthn-3/#dom-publickeycredentialcreationoptions-extensions)
//
// Specification: §9. Extensions (https://www.w3.org/TR/webauthn/#webauthn-extensions)
//
// Specification: §10.1.2. FIDO AppID Exclusion Extension (https://www.w3.org/TR/webauthn/#sctn-appid-exclude-extension)
func WithAppIdExcludeExtension(appid string) RegistrationOption {
	_ = "STUB: not implemented"
	return *new(RegistrationOption)
}

// WithRegistrationRelyingPartyID sets the relying party id for the registration.
func WithRegistrationRelyingPartyID(id string) RegistrationOption {
	_ = "STUB: not implemented"
	return *new(RegistrationOption)
}

// WithRegistrationRelyingPartyName sets the relying party name for the registration.
func WithRegistrationRelyingPartyName(name string) RegistrationOption {
	_ = "STUB: not implemented"
	return *new(RegistrationOption)
}
