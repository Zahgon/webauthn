package protocol

import (
	"crypto/x509"

	"github.com/google/go-tpm/tpm2"

	"github.com/go-webauthn/webauthn/metadata"
)

// attestationFormatValidationHandlerTPM is the handler for the TPM Attestation Statement Format.
//
// The syntax of a TPM Attestation statement is as follows:
//
// $$attStmtType // = (
//
//	    fmt: "tpm",
//	    attStmt: tpmStmtFormat
//	)
//
//	tpmStmtFormat = {
//	                    ver: "2.0",
//	                    (
//	                        alg: COSEAlgorithmIdentifier,
//	                        x5c: [ aikCert: bytes, * (caCert: bytes) ]
//	                    )
//	                    sig: bytes,
//	                    certInfo: bytes,
//	                    pubArea: bytes
//	                }
//
// Specification: §8.3. TPM Attestation Statement Format
//
// See: https://www.w3.org/TR/webauthn/#sctn-tpm-attestation
//
//nolint:gocyclo
func attestationFormatValidationHandlerTPM(att AttestationObject, clientDataHash []byte, _ metadata.Provider) (attestationType string, x5cs []any, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Concatenate authenticatorData and clientDataHash to form attToBeSigned.
//nolint:gocritic // This is intentional.

// Validate that certInfo is valid:
// 1/4 Verify that magic is set to TPM_GENERATED_VALUE, handled here.

// 2/4 Verify that type is set to TPM_ST_ATTEST_CERTIFY.

// 3/4 Verify that extraData is set to the hash of attToBeSigned using the hash algorithm employed in "alg".

// Note that the remaining fields in the "Standard Attestation Structure"
// [TPMv2-Part1] section 31.2, i.e., qualifiedSigner, clockInfo and firmwareVersion
// are ignored. These fields MAY be used as an input to risk engines.

// In this case:
// Verify the sig is a valid signature over certInfo using the attestation public key in aikCert with the algorithm specified in alg.

// Verify that aikCert meets the requirements in §8.3.1 TPM Attestation Statement Certificate Requirements.

// 1/6 Version MUST be set to 3.

// 2/6 Subject field MUST be set to empty.

// 3/6 The Subject Alternative Name extension MUST be set as defined in [TPMv2-EK-Profile] section 3.2.9.

// 4/6 The Extended Key Usage extension MUST contain the "joint-iso-itu-t(2) internationalorganizations(23) 133 tcg-kp(8) tcg-kp-AIKCertificate(3)" OID.

// 6/6 An Authority Information Access (AIA) extension with entry id-ad-ocsp and a CRL Distribution Point
// extension [RFC5280] are both OPTIONAL as the status of many attestation certificates is available
// through metadata services. See, for example, the FIDO Metadata Service.

// 4/4 Verify that attested contains a TPMS_CERTIFY_INFO structure as specified in
// [TPMv2-Part2] section 10.12.3, whose name field contains a valid Name for pubArea,
// as computed using the algorithm in the nameAlg field of pubArea
// using the procedure specified in [TPMv2-Part1] section 16.
//
// This needs to move after the x5c check as the QualifiedSigner only gets populated when it can be verified.

func tpm2Exponent(params *tpm2.TPMSRSAParms) (exp uint32) { _ = "STUB: not implemented"; return 0 }

func tpm2NameMatch(certInfo *tpm2.TPMSAttest, pubArea *tpm2.TPMTPublic) (match bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Per the WebAuthn Specification §8.3 step 5:
//
// Note: The remaining fields in the "Standard Attestation Structure" [TPMv2-Part1] section 31.2, i.e.,
// qualifiedSigner, clockInfo and firmwareVersion are ignored. Depending on the properties of the aikCert key used,
// these fields may be obfuscated. If valid, these MAY be used as an input to risk engines.
//
// See: https://w3c.github.io/webauthn/#sctn-tpm-attestation

func tpm2NameDigest(name tpm2.TPM2BName) (alg tpm2.TPMIAlgHash, digest []byte, err error) {
	_ = "STUB: not implemented"
	return *new(tpm2.TPMIAlgHash), nil, nil
}

type tpm2AttStatement struct {
	Version   string
	Algorithm int64
	Signature []byte
	CertInfo  []byte
	PubArea   []byte

	X5C         []any
	HasX5C      bool
	HasValidX5C bool

	HasECDAAKeyID      bool
	HasValidECDAAKeyID bool
	ECDAAKeyID         []byte
}

func newTPM2AttStatement(raw map[string]any) (statement *tpm2AttStatement, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Given the verification procedure inputs attStmt, authenticatorData
// and clientDataHash, the verification procedure is as follows.

// Verify that attStmt is valid CBOR conforming to the syntax defined
// above and perform CBOR decoding on it to extract the contained fields.

// forEachSAN loops through the TPM SAN extension.
//
// RFC 5280, 4.2.1.6
// SubjectAltName ::= GeneralNames
//
// GeneralNames ::= SEQUENCE SIZE (1..MAX) OF GeneralName
//
//	GeneralName ::= CHOICE {
//	     otherName                       [0]     OtherName,
//	     rfc822Name                      [1]     IA5String,
//	     dNSName                         [2]     IA5String,
//	     x400Address                     [3]     ORAddress,
//	     directoryName                   [4]     Name,
//	     ediPartyName                    [5]     EDIPartyName,
//	     uniformResourceIdentifier       [6]     IA5String,
//	     iPAddress                       [7]     OCTET STRING,
//	     registeredID                    [8]     OBJECT IDENTIFIER }
func forEachSAN(extension []byte, callback func(tag int, data []byte) error) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	nameTypeDN = 4
)

func parseSANExtension(value []byte) (manufacturer string, model string, version string, err error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

type tpmManufacturer struct {
	id   string
	name string
	code string
}

// See https://trustedcomputinggroup.org/resource/vendor-id-registry/ for registry contents.
var (
	tpmManufacturers = []tpmManufacturer{
		{"414D4400", "AMD", "AMD"},
		{"414E5400", "Ant Group", "ANT"},
		{"41544D4C", "Atmel", "ATML"},
		{"4252434D", "Broadcom", "BRCM"},
		{"4353434F", "Cisco", "CSCO"},
		{"464C5953", "Flyslice Technologies", "FLYS"},
		{"524F4343", "Fuzhou Rockchip", "ROCC"},
		{"474F4F47", "Google", "GOOG"},
		{"48504900", "HPI", "HPI"},
		{"48504500", "HPE", "HPE"},
		{"48495349", "Huawei", "HISI"},
		{"49424d00", "IBM", "IBM"},
		{"49424D00", "IBM", "IBM"},
		{"49465800", "Infineon", "IFX"},
		{"494E5443", "Intel", "INTC"},
		{"4C454E00", "Lenovo", "LEN"},
		{"4D534654", "Microsoft", "MSFT"},
		{"4E534D20", "National Semiconductor", "NSM"},
		{"4E545A00", "Nationz", "NTZ"},
		{"4E534700", "NSING", "NSG"},
		{"4E544300", "Nuvoton Technology", "NTC"},
		{"51434F4D", "Qualcomm", "QCOM"},
		{"534D534E", "Samsung", "SECE"},
		{"53454345", "SecEdge", "SecEdge"},
		{"534E5300", "Sinosun", "SNS"},
		{"534D5343", "SMSC", "SMSC"},
		{"53544D20", "ST Microelectronics", "STM"},
		{"54584E00", "Texas Instruments", "TXN"},
		{"57454300", "Winbond", "WEC"},
		{"5345414C", "Wisekey", "SEAL"},
		{"FFFFF1D0", "FIDO Alliance Conformance Testing", "FIDO"},
	}
)

func isValidTPMManufacturer(id string) bool { _ = "STUB: not implemented"; return false }

func tpmParseAIKAttCA(x5c *x509.Certificate, x5cis []*x509.Certificate) (err *Error) {
	_ = "STUB: not implemented"
	return nil
}

func tpmParseSANExtension(attestation *x509.Certificate) (protoErr *Error) {
	_ = "STUB: not implemented"
	return nil
}

type tpmBasicConstraints struct {
	IsCA       bool `asn1:"optional"`
	MaxPathLen int  `asn1:"optional,default:-1"`
}

// Remove extension key usage to avoid ExtKeyUsage check failure.
func tpmRemoveEKU(x5c *x509.Certificate) *Error { _ = "STUB: not implemented"; return nil }

func init() {
	RegisterAttestationFormat(AttestationFormatTPM, attestationFormatValidationHandlerTPM)
}
