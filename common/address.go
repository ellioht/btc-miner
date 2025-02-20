package common

const (
	AddressLength = 16
)

type Address [AddressLength]byte

func DecodeAddress(addr string) (Address, error) {
	var a Address
	copy(a[:], addr)
	return a, nil
}

func (a Address) String() string {
	return string(a[:])
}

func (a Address) Bytes() []byte {
	return a[:]
}
