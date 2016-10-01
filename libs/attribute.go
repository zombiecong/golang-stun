package libs

import "fmt"

type Attribute struct {
	Type_  uint16
	Length uint16
	Value  []byte
}

func newAttribute(types uint16, value []byte) *Attribute  {
	attr := new(Attribute)
	attr.Type_ = types
	attr.Value = padding(value)
	attr.Length = uint16(len(attr.Value))
	return attr
}

//func newXorMappedAddress()  {
//	attr := new(Attribute)
//	attr.Type_ = AttributeXorMappedAddress
//
//	return attr;
//}

func (a Attribute) TypeToString() (typeString string)  {
	switch a.Type_ {
	case AttributeMappedAddress:
		typeString = "MappedAddress"
	case AttributeResponseAddress:
		typeString = "ResponseAddress"
	case AttributeChangeRequest:
		typeString = "ChangeRequest"
	case AttributeSourceAddress:
		typeString = "SourceAddress"
	case AttributeChangedAddress:
		typeString = "ChangedAddress"
	case  AttributeUsername:
		typeString = "Username"
	case  AttributePassword:
		typeString = "Password"
	case AttributeMessageIntegrity:
		typeString = "MessageIntegrity"
	case AttributeErrorCode:
		typeString = "ErrorCode"
	case AttributeUnknownAttributes:
		typeString = "UnknownAttributes"
	case AttributeReflectedFrom:
		typeString = "ReflectedFrom"
	case AttributeChannelNumber:
		typeString = "ChannelNumber"
	case AttributeLifetime:
		typeString = "Lifetime"
	case AttributeBandwidth:
		typeString = "Bandwidth"
	case AttributeXorPeerAddress:
		typeString = "XorPeerAddress"
	case AttributeData:
		typeString = "Data"
	case AttributeRealm:
		typeString = "Realm"
	case AttributeNonce:
		typeString = "Nonce"
	case AttributeXorRelayedAddress:
		typeString = "XorRelayedAddress"
	case AttributeRequestedAddressFamily:
		typeString = "RequestedAddressFamily"
	case AttributeEvenPort:
		typeString = "EvenPort"
	case AttributeRequestedTransport:
		typeString = "RequestedTransport"
	case AttributeDontFragment:
		typeString = "DontFragment"
	case AttributeXorMappedAddress:
		typeString = "XorMappedAddress"
	case AttributeTimerVal:
		typeString = "TimerVal"
	case AttributeReservationToken:
		typeString = "ReservationToken"
	case AttributePriority:
		typeString = "Priority"
	case AttributeUseCandidate:
		typeString = "UseCandidate"
	case AttributePadding:
		typeString = "Padding"
	case AttributeResponsePort:
		typeString = "ResponsePort"
	case AttributeConnectionID:
		typeString = "ConnectionID"
	case AttributeXorMappedAddressExp:
		typeString = "XorMappedAddressExp"
	case AttributeSoftware:
		typeString = "Software"
	case AttributeAlternateServer:
		typeString = "AlternateServer"
	case AttributeCacheTimeout:
		typeString = "CacheTimeout"
	case AttributeFingerprint:
		typeString = "Fingerprint"
	case AttributeIceControlled:
		typeString = "IceControlled"
	case AttributeIceControlling:
		typeString = "IceControlling"
	case AttributeResponseOrigin:
		typeString = "ResponseOrigin"
	case AttributeOtherAddress:
		typeString = "OtherAddress"
	case AttributeEcnCheckStun:
		typeString = "EcnCheckStun"
	case AttributeCiscoFlowdata:
		typeString = "CiscoFlowdata"
	case AttributeOrigin:
		typeString = "Origin"
	default:
		typeString = "fuck??"
	}

	return
}
func (a Attribute) String() string {
	return fmt.Sprintf("attribute : type -> %s , length -> %d \n , value -> %s",
				a.TypeToString() ,a.Length, a.Value);
}