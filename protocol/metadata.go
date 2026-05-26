package protocol

import (
	"context"

	"github.com/google/uuid"

	"github.com/go-webauthn/webauthn/metadata"
)

// ValidateMetadata validates the metadata for the given authenticator.
//
//nolint:gocyclo
func ValidateMetadata(ctx context.Context, mds metadata.Provider, aaguid uuid.UUID, attestationType, attestationFormat string, x5cs []any) (protoErr *Error) {
	_ = "STUB: not implemented"
	return nil
}

func loopOrdinalNumber(n int) string { _ = "STUB: not implemented"; return "" }
