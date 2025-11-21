package utils

import "encoding/base64"

// base64URLEncode BASE64 URL encodes the content without
// any padding.
func Base64URLEncode(content []byte) []byte {
	b64 := make([]byte, base64.URLEncoding.EncodedLen(len(content)))
	base64.URLEncoding.Encode(b64, content)
	return b64
}

// base64URLDecode BASE64 URL decodes the content without
// any padding.
func Base64URLDecode(content []byte) []byte {
	b64 := make([]byte, base64.URLEncoding.DecodedLen(len(content)))
	base64.URLEncoding.Decode(b64, content)
	return b64
}
