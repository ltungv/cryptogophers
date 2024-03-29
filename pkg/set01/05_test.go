package set01

import (
	"bytes"
	"testing"
)

func TestRepeatingXORHex(t *testing.T) {
	t.Parallel()
	tc := struct {
		key []byte
		in  []byte
		out []byte
	}{
		[]byte("ICE"),
		[]byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"),
		[]byte("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"),
	}

	out := RepeatingXORHex(tc.in, tc.key)
	if !bytes.Equal(out, tc.out) {
		t.Errorf("unexpected output:\nhave %s\n want %s", out, tc.out)
	}
}
