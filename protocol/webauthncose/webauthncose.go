package webauthncose

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"hash"

	"github.com/google/go-tpm/tpm2"
)

// PublicKeyData The public key portion of a Relying Party-specific credential key pair, generated
// by an authenticator and returned to a Relying Party at registration time. We unpack this object
// using fxamacker's cbor library ("github.com/fxamacker/cbor/v2") which is why there are cbor tags
// included. The tag field values correspond to the IANA COSE keys that give their respective
// values.
//
// Specification: §6.4.1.1. Examples of credentialPublicKey Values Encoded in COSE_Key Format (https://www.w3.org/TR/webauthn/#sctn-encoded-credPubKey-examples)
type PublicKeyData struct {
	// Decode the results to int by default.
	_struct bool `cbor:",keyasint" json:"public_key"` //nolint:govet,staticcheck

	// The type of key created. Should be OKP, EC2, or RSA.
	KeyType int64 `cbor:"1,keyasint" json:"kty"`

	// A COSEAlgorithmIdentifier for the algorithm used to derive the key signature.
	Algorithm int64 `cbor:"3,keyasint" json:"alg"`
}

type EC2PublicKeyData struct {
	PublicKeyData

	// If the key type is EC2, the curve on which we derive the signature from.
	Curve int64 `cbor:"-1,keyasint,omitempty" json:"crv"`

	// A byte string 32 bytes in length that holds the x coordinate of the key.
	XCoord []byte `cbor:"-2,keyasint,omitempty" json:"x"`

	// A byte string 32 bytes in length that holds the y coordinate of the key.
	YCoord []byte `cbor:"-3,keyasint,omitempty" json:"y"`
}

type RSAPublicKeyData struct {
	PublicKeyData

	// Represents the modulus parameter for the RSA algorithm.
	Modulus []byte `cbor:"-1,keyasint,omitempty" json:"n"`

	// Represents the exponent parameter for the RSA algorithm.
	Exponent []byte `cbor:"-2,keyasint,omitempty" json:"e"`
}

type OKPPublicKeyData struct {
	PublicKeyData

	Curve int64

	// A byte string that holds the x coordinate of the key.
	XCoord []byte `cbor:"-2,keyasint,omitempty" json:"x"`
}

// Verify Octet Key Pair (OKP) Public Key Signature.
func (k *OKPPublicKeyData) Verify(data []byte, sig []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Verify Elliptic Curve Public Key Signature.
func (k *EC2PublicKeyData) Verify(data []byte, sig []byte) (valid bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ToECDSA converts the EC2PublicKeyData to an ecdsa.PublicKey.
func (k *EC2PublicKeyData) ToECDSA() (key *ecdsa.PublicKey, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verify RSA Public Key Signature.
func (k *RSAPublicKeyData) Verify(data []byte, sig []byte) (valid bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ParsePublicKey figures out what kind of COSE material was provided and create the data for the new key.
func ParsePublicKey(keyBytes []byte) (publicKey any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// ParseFIDOPublicKey is only used when the appID extension is configured by the assertion response.
func ParseFIDOPublicKey(keyBytes []byte) (data EC2PublicKeyData, err error) {
	_ = "STUB: not implemented"
	return *new(EC2PublicKeyData), nil
}

// Raw bytes for an uncompressed P-256 point: 0x04 || x(32) || y(32).

func VerifySignature(key any, data []byte, sig []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func DisplayPublicKey(cpk []byte) string { _ = "STUB: not implemented"; return "" }

func (k *EC2PublicKeyData) TPMCurveID() tpm2.TPMECCCurve {
	_ = "STUB: not implemented"
	return *new(tpm2.TPMECCCurve)
}

// TPM_ECC_NIST_P256.

// TPM_ECC_NIST_P384.

// TPM_ECC_NIST_P521.

// TPM_ECC_NONE.

func ec2AlgCurve(coseAlg int64) elliptic.Curve {
	_ = "STUB: not implemented"
	return *new(elliptic.Curve)
}

// SigAlgFromCOSEAlg return which signature algorithm is being used from the COSE Key.
func SigAlgFromCOSEAlg(coseAlg COSEAlgorithmIdentifier) x509.SignatureAlgorithm {
	_ = "STUB: not implemented"
	return *new(x509.SignatureAlgorithm)
}

// HasherFromCOSEAlg returns the Hashing interface to be used for a given COSE Algorithm.
func HasherFromCOSEAlg(coseAlg COSEAlgorithmIdentifier) hash.Hash {
	_ = "STUB: not implemented"
	return *new(hash.Hash)
}

// default to SHA256?  Why not.

var COSESignatureAlgorithmDetails = map[COSEAlgorithmIdentifier]struct {
	name   string
	hash   crypto.Hash
	sigAlg x509.SignatureAlgorithm
}{
	AlgRS1:     {"SHA1-RSA", crypto.SHA1, x509.SHA1WithRSA},
	AlgRS256:   {"SHA256-RSA", crypto.SHA256, x509.SHA256WithRSA},
	AlgRS384:   {"SHA384-RSA", crypto.SHA384, x509.SHA384WithRSA},
	AlgRS512:   {"SHA512-RSA", crypto.SHA512, x509.SHA512WithRSA},
	AlgPS256:   {"SHA256-RSAPSS", crypto.SHA256, x509.SHA256WithRSAPSS},
	AlgPS384:   {"SHA384-RSAPSS", crypto.SHA384, x509.SHA384WithRSAPSS},
	AlgPS512:   {"SHA512-RSAPSS", crypto.SHA512, x509.SHA512WithRSAPSS},
	AlgES256:   {"ECDSA-SHA256", crypto.SHA256, x509.ECDSAWithSHA256},
	AlgESP256:  {"ECDSA-SHA256-Prehashed", crypto.SHA256, x509.ECDSAWithSHA256},
	AlgES384:   {"ECDSA-SHA384", crypto.SHA384, x509.ECDSAWithSHA384},
	AlgESP384:  {"ECDSA-SHA384-Prehashed", crypto.SHA384, x509.ECDSAWithSHA384},
	AlgES512:   {"ECDSA-SHA512", crypto.SHA512, x509.ECDSAWithSHA512},
	AlgESP512:  {"ECDSA-SHA512-Prehashed", crypto.SHA512, x509.ECDSAWithSHA512},
	AlgEdDSA:   {"EdDSA", crypto.SHA512, x509.PureEd25519},
	AlgEd25519: {"Ed25519", crypto.SHA512, x509.PureEd25519},
}

type Error struct {
	// Short name for the type of error that has occurred.
	Type string `json:"type"`

	// Additional details about the error.
	Details string `json:"error"`

	// Information to help debug the error.
	DevInfo string `json:"debug"`
}

var (
	ErrUnsupportedKey = &Error{
		Type:    "invalid_key_type",
		Details: "Unsupported Public Key Type",
	}
	ErrUnsupportedAlgorithm = &Error{
		Type:    "unsupported_key_algorithm",
		Details: "Unsupported public key algorithm",
	}
	ErrSigNotProvidedOrInvalid = &Error{
		Type:    "signature_not_provided_or_invalid",
		Details: "Signature invalid or not provided",
	}
)

func (err *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (passedError *Error) WithDetails(details string) *Error { _ = "STUB: not implemented"; return nil }

func validateOKPPublicKey(k *OKPPublicKeyData) error { _ = "STUB: not implemented"; return nil }

func validateEC2PublicKey(k *EC2PublicKeyData) error { _ = "STUB: not implemented"; return nil }

func validateRSAPublicKey(k *RSAPublicKeyData) error { _ = "STUB: not implemented"; return nil }

func parseRSAPublicKeyDataExponent(k *RSAPublicKeyData) (exp int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
