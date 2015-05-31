package dhcp6_test

// Tested using dhcp6_test package to allow use of dhcp6test package.

import (
	"testing"

	"github.com/mdlayher/dhcp6"
	"github.com/mdlayher/dhcp6/dhcp6test"
)

// TestServeMuxHandleNoResponse verifies that no Handler is invoked when a
// ServeMux does not have a Handler registered for a given message type.
func TestServeMuxHandleNoResponse(t *testing.T) {
	mux := dhcp6.NewServeMux()

	p, err := dhcp6.NewPacket(dhcp6.MessageTypeSolicit, []byte{0, 1, 2}, nil)
	if err != nil {
		t.Fatal(err)
	}

	r := dhcp6.ParseRequest(p, nil)
	w := dhcp6test.NewRecorder(r.TransactionID)
	mux.ServeDHCP(w, r)

	if l := len(w.Packet); l > 0 {
		t.Fatalf("reply packet should be empty, but got length: %d", l)
	}
	if mt := w.MessageType; mt != dhcp6.MessageType(0) {
		t.Fatalf("reply packet empty, but got message type: %v", mt)
	}
	if l := len(w.Options()); l > 0 {
		t.Fatalf("reply packet empty, but got %d options", l)
	}
}

// TestServeMuxHandleOK verifies that a Handler is invoked when a ServeMux
// has a Handler registered for a given message type.
func TestServeMuxHandleOK(t *testing.T) {
	mux := dhcp6.NewServeMux()
	mt := dhcp6.MessageTypeSolicit

	mux.Handle(mt, &solicitHandler{})

	p, err := dhcp6.NewPacket(mt, []byte{0, 1, 2}, nil)
	if err != nil {
		t.Fatal(err)
	}

	r := dhcp6.ParseRequest(p, nil)
	w := dhcp6test.NewRecorder(r.TransactionID)
	mux.ServeDHCP(w, r)

	if want, got := dhcp6.MessageTypeAdvertise, w.MessageType; want != got {
		t.Fatalf("unexpected response message type: %v != %v", want, got)
	}
}

// TestServeMuxHandleFuncOK verifies that a normal function which can be used
// as a Handler is invoked when a ServeMux has a HandlerFunc registered for
// a given message type.
func TestServeMuxHandleFuncOK(t *testing.T) {
	mux := dhcp6.NewServeMux()
	mt := dhcp6.MessageTypeSolicit

	mux.HandleFunc(mt, solicit)

	p, err := dhcp6.NewPacket(mt, []byte{0, 1, 2}, nil)
	if err != nil {
		t.Fatal(err)
	}

	r := dhcp6.ParseRequest(p, nil)
	w := dhcp6test.NewRecorder(r.TransactionID)
	mux.ServeDHCP(w, r)

	if want, got := dhcp6.MessageTypeAdvertise, w.MessageType; want != got {
		t.Fatalf("unexpected response message type: %v != %v", want, got)
	}
}

// solicitHandler is a Handler which returns an Advertise in reply
// to a Solicit request.
type solicitHandler struct{}

func (h *solicitHandler) ServeDHCP(w dhcp6.Responser, r *dhcp6.Request) {
	solicit(w, r)
}

// solicit is a function which can be adapted as a HandlerFunc.
func solicit(w dhcp6.Responser, r *dhcp6.Request) {
	w.Send(dhcp6.MessageTypeAdvertise)
}
