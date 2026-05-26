package cached

import (
	"io"
	"os"

	"github.com/go-webauthn/webauthn/metadata"
)

func doTruncateCopyAndSeekStart(f *os.File, rc io.ReadCloser) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func doOpenOrCreate(name string) (f *os.File, created bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func defaultNew(mds *metadata.Metadata) (provider metadata.Provider, err error) {
	_ = "STUB: not implemented"
	return *new(metadata.Provider), nil
}
