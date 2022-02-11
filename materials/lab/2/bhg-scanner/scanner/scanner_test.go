package scanner

import (
	"testing"
)

func TestOpenPort15to30(t *testing.T){
	
    gotopen,_,_ := PortScanner(15,30,true)
    want := 1
    if gotopen != want {
        t.Errorf("got %d, wanted %d", gotopen, want)
    }
}

func TestTotalPortsScanned15to30(t *testing.T){
	
    gotopen,gotclosed,_ := PortScanner(15,30,false)
    want := 15

    if gotopen+gotclosed != want {
        t.Errorf("got %d, wanted %d", gotopen+gotclosed, want)
    }
}
func TestPortOpen15to30(t *testing.T){
	
    _,_,ports := PortScanner(15,30,false)
    want := true

    if ports[22] != want {
        t.Errorf("got %t, wanted %t", ports[22], want)
    }
}
func TestPortClosed15to30(t *testing.T){
	
    _,_,ports := PortScanner(15,30,false)
    want := false

    if ports[5] != want {
        t.Errorf("got %t, wanted %t", ports[5], want)
    }
}
func TestOpenPort0to1024(t *testing.T){
	
    gotopen,_,_ := PortScanner(0,1024,false)
    want := 1
    if gotopen != want {
        t.Errorf("got %d, wanted %d", gotopen, want)
    }
}
func TestTotalPortsScanned0to1024(t *testing.T){
	
    gotopen,gotclosed,_ := PortScanner(0,1024,false)
    want := 1024

    if gotopen+gotclosed != want {
        t.Errorf("got %d, wanted %d", gotopen+gotclosed, want)
    }
}