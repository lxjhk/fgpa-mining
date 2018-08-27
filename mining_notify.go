package main

import "encoding/hex"

type mining_notify_job struct {
	job_id        string
	prevhash      string
	coinb1        string
	coinb2        string
	merkle_branch []string
	version       string
	nbits         string
	ntime         string
	clean_jobs    bool
	difficulty    int
}

func (self *mining_notify_job) calculateMerkelRoot(cb *[]byte) []byte {
	merkle_root := make([]byte, len(*cb))
	copy(merkle_root, *cb)

	for _, v := range self.merkle_branch {
		merkle_root = doubleSHA256(append(merkle_root, hexStringToByteArray(&v)...))
	}

	return merkle_root
}

//To produce coinbase, we just concatenate Coinb1 + Extranonce1 + Extranonce2 + Coinb2 together.
func (self *mining_notify_job) buildCoinBase(extranonce1 *string, extranonce2 *string) *[]byte {
	s, _ := hex.DecodeString(
		self.coinb1 +
			*extranonce1 +
			*extranonce2 +
			self.coinb2)
	a := doubleSHA256(s)
	return &a
}

// need to reverse the byte order of merkle root
func (self *mining_notify_job) getFPGAJob(mr []byte) string {
	t1 := self.version + self.prevhash +
		hex.EncodeToString(reverseArray(mr)) +
		self.ntime +
		self.nbits
	return t1
}
