package session

import (
	"github.com/fasthttp/session/v2"
)

// MSGPEncode MessagePack encode.
func MSGPEncode(src session.Dict) ([]byte, error) {
	return session.MSGPEncode(src)
}

// MSGPDecode MessagePack decode.
func MSGPDecode(dst *session.Dict, src []byte) error {
	return session.MSGPDecode(dst, src)
}

// Base64Encode base64 encode.
func Base64Encode(src session.Dict) ([]byte, error) {
	return session.Base64Encode(src)
}

// Base64Decode base64 decode.
func Base64Decode(dst *session.Dict, src []byte) error {
	return session.MSGPDecode(dst, src)
}
