package uot

import (
	"github.com/sagernet/uot/common/socksaddr"
)

const UOTMagicAddress = "sp.udp-over-tcp.arpa"

var AddrParser = socksaddr.NewSerializer(
	socksaddr.AddressFamilyByte(0x00, socksaddr.AddressFamilyIPv4),
	socksaddr.AddressFamilyByte(0x01, socksaddr.AddressFamilyIPv6),
	socksaddr.AddressFamilyByte(0x02, socksaddr.AddressFamilyFqdn),
)
