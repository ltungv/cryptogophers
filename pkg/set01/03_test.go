package set01

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/pkg/errors"
)

func TestHexSingleXORDecrypt(t *testing.T) {
	t.Parallel()
	tc := struct {
		inHex []byte
	}{
		[]byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"),
	}

	out, key, _, err := HexSingleXORDecrypt(tc.inHex)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	in := make([]byte, hex.EncodedLen(len(tc.inHex)))
	n, err := hex.Decode(in, tc.inHex)
	in = in[:n]
	if err != nil {
		t.Fatal(errors.Wrapf(err, "could not decode: %s", tc.inHex))
	}

	constructed := crypts.SingleXOR(out, key)
	if !bytes.Equal(constructed, in) {
		t.Errorf("incorrect reconstructed input:\nhave %x\nwant %x", constructed, in)
	}
}
