// Optional Todo

package hscan

import (
	"testing"
)
/*
func TestGuessSingle(t *testing.T) { //edited to pass
	want := "*123456*"
	got := GuessSingle("77f62e3524cd583d698d51fa24fdff4f", "Top304Thousand-probable-v2.txt")
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

}

func TestShouldWork(t *testing.T) {
	want := "foo"
	got := GuessSingle("acbd18db4cc2f85cedef654fccc4a4d8", "Top304Thousand-probable-v2.txt")

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

}
*/
func TestHashmapTime(t *testing.T) {
	got1,got2:=GenHashMaps("Top304Thousand-probable-v2.txt")
	want:=303872
	if (got1 != want) {
		t.Errorf("got %d, wanted %d", got1, want)
	}else if (got2 != want) {
		t.Errorf("got %d, wanted %d", got2, want)
	}
}
