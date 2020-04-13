package queryfx_test

import (
	"fmt"
	"github.com/donama/queryfx"
	"testing"
)

func TestQueryFX(t *testing.T){
	phids := []string{
		"PHID-STXN-2f0fea1931a290220777a93143dfdcbfa68406e877073ff08834e197a4034aa4",
		"PHID-ACCT-9cdcc595bcce3c7bd3d8df93fab7e125ddebafe65a31bd5d41e2d2ce9c2b1789",
		"PHID-UTXN-5d8857b799acb18e4affabe3037ffe7fa68aa8af5e39cc416e734d373c5ebebc",
		"PHID-ACCT-7694267aef4ebcea406b32d6108bd68584f57e37caac6e33feaa3263a3994370"}

	q := "SELECT %LC from %T where phid IN (%Ls) AND id = %d"
	q1, err := queryfx.FormatQuery(q, []string{"marketplaceName","channelPHID","metaData"}, "marketplace_lists", phids, uint64(234))
	if err != nil{
		t.Fatal(err)
	}
	fmt.Println(q1)
}

func TestEmptyQueryFX(t *testing.T){

	q := "SELECT * FROM lists"
	q1, err := queryfx.FormatQuery(q)
	if err != nil{
		t.Fatal(err)
	}
	fmt.Println(q1)
}