package dhcp6

// MessageType represents a DHCP message type, as defined in RFC 3315,
// Section 5.3.  Different DHCP message types are used to perform different
// actions between a client and server.
type MessageType uint8

// MessageType constants which indicate the message types described in
// RFCs 3315, 5007, 5460, 6977, and 7341.
//
// These message types are taken from IANA's DHCPv6 parameters registry:
// http://www.iana.org/assignments/dhcpv6-parameters/dhcpv6-parameters.xhtml.
const (
	// RFC 3315
	MessageTypeSolicit            MessageType = 1
	MessageTypeAdvertise          MessageType = 2
	MessageTypeRequest            MessageType = 3
	MessageTypeConfirm            MessageType = 4
	MessageTypeRenew              MessageType = 5
	MessageTypeRebind             MessageType = 6
	MessageTypeReply              MessageType = 7
	MessageTypeRelease            MessageType = 8
	MessageTypeDecline            MessageType = 9
	MessageTypeReconfigure        MessageType = 10
	MessageTypeInformationRequest MessageType = 11
	MessageTypeRelayForw          MessageType = 12
	MessageTypeRelayRepl          MessageType = 13

	// RFC 5007
	MessageTypeLeasequery      MessageType = 14
	MessageTypeLeasequeryReply MessageType = 15

	// RFC 5460
	MessageTypeLeasequeryDone MessageType = 16
	MessageTypeLeasequeryData MessageType = 17

	// RFC 6977
	MessageTypeReconfigureRequest MessageType = 18
	MessageTypeReconfigureReply   MessageType = 19

	// RFC 7341
	MessageTypeDHCPv4Query    MessageType = 20
	MessageTypeDHCPv4Response MessageType = 21
)

// Status represesents a DHCP status code, as defined in RFC 3315,
// Section 5.4.  Status codes are used to communicate success or failure
// between client and server.
type Status uint16

// Status constants which indicate the status codes described in
// RFC 3315, Section 24.4.  Additional status are defined in IANA's
// DHCPv6 parameters registry:
// http://www.iana.org/assignments/dhcpv6-parameters/dhcpv6-parameters.xhtml.
const (
	StatusSuccess      Status = 0
	StatusUnspecFail   Status = 1
	StatusNoAddrsAvail Status = 2
	StatusNoBinding    Status = 3
	StatusNotOnLink    Status = 4
	StatusUseMulticast Status = 5

	// BUG(mdlayher): add additional status codes defined by IANA
)

// OptionCode represents a DHCP option, as defined in RFC 3315,
// Section 22.  Options are used to carry additional information and
// parameters in DHCP messages between client and server.
type OptionCode uint16

// Status constants which indicate the option codes described in
// RFC 3315, Section 24.3.  Additional option codes are defined in IANA's
// DHCPv6 parameters registry:
// http://www.iana.org/assignments/dhcpv6-parameters/dhcpv6-parameters.xhtml.
const (
	OptionClientID     OptionCode = 1
	OptionServerID     OptionCode = 2
	OptionIANA         OptionCode = 3
	OptionIATA         OptionCode = 4
	OptionIAAddr       OptionCode = 5
	OptionORO          OptionCode = 6
	OptionPreference   OptionCode = 7
	OptionElapsedTime  OptionCode = 8
	OptionRelayMsg     OptionCode = 9
	_                  OptionCode = 10
	OptionAuth         OptionCode = 11
	OptionUnicast      OptionCode = 12
	OptionStatusCode   OptionCode = 13
	OptionRapidCommit  OptionCode = 14
	OptionUserClass    OptionCode = 15
	OptionVendorClass  OptionCode = 16
	OptionVendorOpts   OptionCode = 17
	OptionInterfaceID  OptionCode = 18
	OptionReconfMsg    OptionCode = 19
	OptionReconfAccept OptionCode = 20

	// BUG(mdlayher): add additional message types defined by IANA
)
