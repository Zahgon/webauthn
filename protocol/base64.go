package protocol

// URLEncodedBase64 represents a byte slice holding URL-encoded base64 data.
// When fields of this type are unmarshalled from JSON, the data is base64
// decoded into a byte slice.
type URLEncodedBase64 []byte

func (e URLEncodedBase64) String() string { _ = "STUB: not implemented"; return "" }

// UnmarshalJSON base64 decodes a URL-encoded value, storing the result in the
// provided byte slice.
func (e *URLEncodedBase64) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// Trim the leading and trailing quotes from raw JSON data (the whole value part).

// Trim the trailing equal characters.

// MarshalJSON base64 encodes a non URL-encoded value, storing the result in the
// provided byte slice.
func (e URLEncodedBase64) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
