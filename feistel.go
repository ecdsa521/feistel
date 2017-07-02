package feistel

import "strconv"

//Cipher function returns pseudo-ciphered output for given input
//Direct copy of https://stackoverflow.com/questions/12761346/pseudo-encrypt-function-in-plpgsql-that-takes-bigint/12761795#12761795
func Cipher(input uint64) uint64 {
	var l1, l2, r1, r2 uint64

	l1 = (input >> 32) & 4294967295
	r1 = input & 4294967295
	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ ((((1366.0*r1 + 150889) % 714025) / 714025.0) * 32767 * 32767)
		l1 = l2
		r1 = r2
	}
	return ((l1 << 32) + r1)
}

//CipherStr returns string representation for the Cipher
func CipherStr(input uint64) string {
	return strconv.FormatUint(Cipher(input), 36)
}

func init() {

}
