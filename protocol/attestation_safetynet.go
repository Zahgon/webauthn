package protocol

import (
	"github.com/golang-jwt/jwt/v5"

	"github.com/go-webauthn/webauthn/metadata"
)

// attestationFormatValidationHandlerAndroidSafetyNet is the handler for the Android SafetyNet Attestation Statement
// Format.
//
// When the authenticator is a platform authenticator on certain Android platforms, the attestation statement may be
// based on the SafetyNet API. In this case the authenticator data is completely controlled by the caller of the
// SafetyNet API (typically an application running on the Android platform) and the attestation statement provides some
// statements about the health of the platform and the identity of the calling application (see SafetyNet Documentation
// for more details).
//
// The syntax of an Android Attestation statement is defined as follows:
//
//	$$attStmtType //= (
//			fmt: "android-safetynet",
//			attStmt: safetynetStmtFormat
//	)
//
//	safetynetStmtFormat = {
//			ver: text,
//			response: bytes
//	}
//
// Specification: §8.5. Android SafetyNet Attestation Statement Format
//
// See: https://www.w3.org/TR/webauthn/#sctn-android-safetynet-attestation
//
//nolint:gocyclo
func attestationFormatValidationHandlerAndroidSafetyNet(att AttestationObject, clientDataHash []byte, mds metadata.Provider) (attestationType string, x5cs []any, err error) {
	_ = "STUB: not implemented"
	// The syntax of an Android Attestation statement is defined as follows:
	//     $$attStmtType //= (
	//                           fmt: "android-safetynet",
	//                           attStmt: safetynetStmtFormat
	//                       )
	return "", nil, nil
}

//     safetynetStmtFormat = {
//                               ver: text,
//                               response: bytes
//                           }

// §8.5.1 Verify that attStmt is valid CBOR conforming to the syntax defined above and perform CBOR decoding on it to extract
// the contained fields.

// We have done this
// §8.5.2 Verify that response is a valid SafetyNet response of version ver.

// TODO: provide user the ability to designate their supported versions.

// marshall the JWT payload into the safetynet response json.

// §8.5.3 Verify that the nonce in the response is identical to the Base64 encoding of the SHA-256 hash of the concatenation
// of authenticatorData and clientDataHash.

// §8.5.4 Let attestationCert be the attestation certificate (https://www.w3.org/TR/webauthn/#attestation-certificate)

// §8.5.5 Verify that attestationCert is issued to the hostname "attest.android.com".

// §8.5.6 Verify that the ctsProfileMatch attribute in the payload of response is true.

// Zero tolerance for post-dated timestamps.

// Small tolerance for pre-dated timestamps.

// §8.5.7 If successful, return implementation-specific values representing attestation type Basic and attestation
// trust path attestationCert.

func keyFuncSafetyNetJWT(token *jwt.Token) (key any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type SafetyNetResponse struct {
	Nonce                      string `json:"nonce"`
	TimestampMs                int64  `json:"timestampMs"`
	ApkPackageName             string `json:"apkPackageName"`
	ApkDigestSha256            string `json:"apkDigestSha256"`
	CtsProfileMatch            bool   `json:"ctsProfileMatch"`
	ApkCertificateDigestSha256 []any  `json:"apkCertificateDigestSha256"`
	BasicIntegrity             bool   `json:"basicIntegrity"`
}

func init() {
	RegisterAttestationFormat(AttestationFormatAndroidSafetyNet, attestationFormatValidationHandlerAndroidSafetyNet)
}
