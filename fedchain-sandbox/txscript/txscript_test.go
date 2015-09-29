package txscript

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestAddrPkScript(t *testing.T) {
	cases := []struct {
		addr   string
		script string
	}{
		{
			addr:   "3H9gBofbYu4uQXwfMVcFiWjQHXf6vmnVGB",
			script: "a914a994a46855d8f4442b3a6db863628cc020537f4087",
		},
	}

	for _, c := range cases {
		pkScript, err := hex.DecodeString(c.script)
		if err != nil {
			t.Fatal(err)
		}
		got, err := AddrPkScript(c.addr)
		if err != nil {
			t.Error("unexptected error", err)
		}
		if !bytes.Equal(got, pkScript) {
			t.Errorf("got AddrPkScript(%s) = %X want %X", c.addr, got, pkScript)
		}

		got2, err := PkScriptAddr(pkScript)
		if err != nil {
			t.Error("unexptected error", err)
		}
		if got2.String() != c.addr {
			t.Errorf("got PkScriptAddr(%s) = %v want %v", c.script, got2, c.addr)
		}
	}
}

// Taken from PAPI
func TestRedeemToPkScript(t *testing.T) {
	redeem := []byte{
		82, 65, 4, 2, 83, 21, 116, 23, 208, 223, 22, 63, 33, 52, 55, 175, 75,
		119, 114, 250, 19, 22, 177, 255, 206, 20, 137, 199, 197, 174, 244, 194,
		15, 245, 81, 94, 80, 76, 230, 243, 156, 11, 161, 17, 245, 68, 250, 134,
		98, 63, 123, 206, 106, 17, 129, 179, 210, 5, 155, 242, 97, 194, 119,
		175, 122, 32, 45, 65, 4, 219, 47, 252, 31, 82, 125, 34, 225, 107, 200,
		88, 45, 78, 46, 221, 232, 119, 33, 245, 22, 107, 5, 210, 37, 38, 160,
		107, 38, 218, 198, 70, 140, 97, 52, 204, 27, 97, 252, 237, 156, 154,
		175, 86, 193, 177, 245, 210, 222, 244, 235, 8, 179, 15, 187, 126, 249,
		192, 138, 143, 251, 198, 230, 98, 172, 82, 174,
	}

	want := []byte{
		169, 20, 10, 63, 117, 193, 26, 249, 104, 211, 169, 228, 39, 135, 197,
		179, 65, 183, 169, 3, 163, 165, 135,
	}

	got := RedeemToPkScript(redeem)
	if bytes.Compare(got, want) != 0 {
		t.Errorf("got pkscript = %x want %x", got, want)
	}
}
