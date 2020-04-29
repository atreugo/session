// This file is part of fasthttp/session which is released under MIT license.
// Go to https://github.com/fasthttp/session/blob/master/LICENSE for full license details.

package session

import (
	"encoding/base64"

	"github.com/fasthttp/session/v2"
)

var b64Encoding = base64.StdEncoding

// MSGPEncode MessagePack encode.
func MSGPEncode(src session.Dict) ([]byte, error) {
	if len(src.D) == 0 {
		return nil, nil
	}

	dst, err := src.MarshalMsg(nil)
	if err != nil {
		return nil, err
	}

	return dst, nil
}

// MSGPDecode MessagePack decode.
func MSGPDecode(dst *session.Dict, src []byte) error {
	dst.Reset()

	if len(src) == 0 {
		return nil
	}

	_, err := dst.UnmarshalMsg(src)

	return err
}

// Base64Encode base64 encode.
func Base64Encode(src session.Dict) ([]byte, error) {
	srcBytes, err := MSGPEncode(src)
	if err != nil {
		return nil, err
	}

	dst := make([]byte, b64Encoding.EncodedLen(len(srcBytes)))
	b64Encoding.Encode(dst, srcBytes)

	return dst, nil
}

// Base64Decode base64 decode.
func Base64Decode(dst *session.Dict, src []byte) error {
	tmp := make([]byte, b64Encoding.DecodedLen(len(src)))

	n, err := b64Encoding.Decode(tmp, src)
	if err != nil {
		return err
	}

	return MSGPDecode(dst, tmp[:n])
}
