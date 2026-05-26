package protocol

// ChallengeLength - Length of bytes to generate for a challenge.
const ChallengeLength = DefaultChallengeLength

// CreateChallenge creates a new challenge that should be signed and returned by the authenticator. The spec recommends
// using at least 16 bytes with 100 bits of entropy. We use 32 bytes.
func CreateChallenge() (challenge URLEncodedBase64, err error) {
	_ = "STUB: not implemented"
	return *new(URLEncodedBase64), nil
}
