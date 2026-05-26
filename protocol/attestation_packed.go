package protocol

import (
	"github.com/go-webauthn/webauthn/metadata"
)

func init() {
	RegisterAttestationFormat(AttestationFormatPacked, attestationFormatValidationHandlerPacked)
}

// attestationFormatValidationHandlerPacked is the handler for the Packed Attestation Statement Format.
//
// The syntax of a Packed Attestation statement is defined by the following CDDL:
//
// $$attStmtType //= (
//
//	    fmt: "packed",
//	    attStmt: packedStmtFormat
//	)
//
//	packedStmtFormat = {
//	                       alg: COSEAlgorithmIdentifier,
//	                       sig: bytes,
//	                       x5c: [ attestnCert: bytes, * (caCert: bytes) ]
//	                   } //
//	                   {
//	                       alg: COSEAlgorithmIdentifier
//	                       sig: bytes,
//	                   }
//
// Specification: §8.2. Packed Attestation Statement Format
//
// See: https://www.w3.org/TR/webauthn/#sctn-packed-attestation
func attestationFormatValidationHandlerPacked(att AttestationObject, clientDataHash []byte, mds metadata.Provider) (attestationType string, x5cs []any, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Step 1. Verify that attStmt is valid CBOR conforming to the syntax defined
// above and perform CBOR decoding on it to extract the contained fields.
// Get the alg value - A COSEAlgorithmIdentifier containing the identifier of the algorithm
// used to generate the attestation signature.

// Get the sig value - A byte string containing the attestation signature.

// Step 2. If x5c is present, this indicates that the attestation type is not ECDAA.

// Handle Basic Attestation steps for the x509 Certificate.

// Step 3. If ecdaaKeyId is present, then the attestation type is ECDAA.
// Also make sure the we did not have an x509.

// Handle ECDAA Attestation steps for the x509 Certificate.

// Step 4. If neither x5c nor ecdaaKeyId is present, self attestation is in use.

// Handle the attestation steps laid out in the basic format.
//
//nolint:gocyclo
func handleBasicAttestation(sig, clientDataHash, authData, aaguid []byte, alg int64, x5c []any, _ metadata.Provider) (attestationType string, x5cs []any, err error) {
	_ = "STUB: not implemented"
	// Step 2.1. Verify that sig is a valid signature over the concatenation of authenticatorData
	// and clientDataHash using the attestation public key in attestnCert with the algorithm specified in alg.
	return "", nil, nil
}

//nolint:gocritic // This is intentional.

// Step 2.2 Verify that attestnCert meets the requirements in §8.2.1 Packed attestation statement certificate requirements.
// §8.2.1 can be found here https://www.w3.org/TR/webauthn/#packed-attestation-cert-requirements

// Step 2.2.1 (from §8.2.1) Version MUST be set to 3 (which is indicated by an ASN.1 INTEGER with value 2).

// Step 2.2.2 (from §8.2.1) Subject field MUST be set to:
// 	Subject-C
// 	ISO 3166 code specifying the country where the Authenticator vendor is incorporated (PrintableString).

// 	Subject-O
// 	Legal name of the Authenticator vendor (UTF8String).

// Subject-OU
// Literal string “Authenticator Attestation” (UTF8String).

//  Subject-CN
//  A UTF8String of the vendor’s choosing.

// Step 2.2.3 (from §8.2.1) If the related attestation root certificate is used for multiple authenticator models,
// the Extension OID 1.3.6.1.4.1.45724.1.1.4 (id-fido-gen-ce-aaguid) MUST be present, containing the
// AAGUID as a 16-byte OCTET STRING. The extension MUST NOT be marked as critical.

// We validate the AAGUID as mentioned above
// This is not well defined in§8.2.1 but mentioned in step 2.3: we validate the AAGUID if it is present within the certificate
// and make sure it matches the auth data AAGUID
// Note that an X.509 Extension encodes the DER-encoding of the value in an OCTET STRING. Thus, the
// AAGUID MUST be wrapped in two OCTET STRINGS to be valid.

// Step 2.2.4 The Basic Constraints extension MUST have the CA component set to false.

// Note for 2.2.5 An Authority Information Access (AIA) extension with entry id-ad-ocsp and a CRL
// Distribution Point extension [RFC5280](https://www.w3.org/TR/webauthn/#biblio-rfc5280) are
// both OPTIONAL as the status of many attestation certificates is available through authenticator
// metadata services. See, for example, the FIDO Metadata Service
// [FIDOMetadataService] (https://www.w3.org/TR/webauthn/#biblio-fidometadataservice)

// Step 2.4 If successful, return attestation type Basic and attestation trust path x5c.
// We don't handle trust paths yet but we're done.

func handleECDAAAttestation(sig, clientDataHash, ecdaaKeyID []byte, _ metadata.Provider) (attestationType string, x5cs []any, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func handleSelfAttestation(alg int64, pubKey, authData, clientDataHash, sig []byte, _ metadata.Provider) (attestationType string, x5cs []any, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

//nolint:gocritic // This is intentional.

// §4.1 Validate that alg matches the algorithm of the credentialPublicKey in authenticatorData.

// §4.2 Verify that sig is a valid signature over the concatenation of authenticatorData and
// clientDataHash using the credential public key with alg.

func verifyKeyAlgorithm(keyAlgorithm, attestedAlgorithm int64) error {
	_ = "STUB: not implemented"
	return nil
}
