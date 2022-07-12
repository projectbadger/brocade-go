package utils

import "encoding/base64"

// base64URLEncode BASE64 URL encodes the content without
// any padding.
func Base64URLEncode(content []byte) []byte {
	b64 := make([]byte, base64.RawURLEncoding.EncodedLen(len(content)))
	base64.RawURLEncoding.Encode(b64, content)
	return b64
}

// base64URLDecode BASE64 URL decodes the content without
// any padding.
func Base64URLDecode(content []byte) []byte {
	b64 := make([]byte, base64.RawURLEncoding.DecodedLen(len(content)))
	base64.RawURLEncoding.Decode(b64, content)
	return b64
}
