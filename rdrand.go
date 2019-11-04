// rdrand.go - get random values using RDRAND x86 instruction.
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to this module of rdrand, using the creative
// commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

// +build amd64

package rdrand

func rdrand() uint64

func cpuid() uint32

func rdrandSupported() bool {
	return (cpuid() & (1 << 30)) != 0
}

// Assuming that we have RDRAND support while we are up.
var gSupported = rdrandSupported()

// Uint64 returns random uint64 value using RDRAND instruction.
// If RDRAND is not supported, Uint64 returns supported set to false.
func Uint64() (n uint64, supported bool) {
	if !gSupported {
		return 0, false
	}
	ret := rdrand()
	return ret, true
}
