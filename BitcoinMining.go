package main

import (
	"encoding/hex"
	"fmt"
	"crypto/sha256"
	"math/big"
	"encoding/json"
	"os"
)

const (
	BASE_1_DIFFICULTY = "0x00000000ffff0000000000000000000000000000000000000000000000000000"
	jd = `{"id": null, "method": "mining.notify", "params": ["8be5", "d7ea32ad603a57b96832465b110d4bbfcc70d6c2001b8f4a0000000000000000", "01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff5c03923708192f5669614254432f4d696e65642062792077346d387374322f2cfabe6d6dbd21d45341c3cbb7a366af31c1fe743d27dbe660e44dcae9275699982333733c040000000000000010e58b370862685029", "ffffffff02b3b0094b000000001976a914536ffa992491508dca0354e52f32a3a7a679a53a88ac0000000000000000266a24aa21a9eddf9626ca545c861be96e1a52fbcea9e27a3c1169d605b031dcdedf64d0461e7700000000", ["3a5328ad196525bc87cc3ad830b0354d8b6e87fc2851f1af5dee89cfd773dacf", "0a19ec9f5cdb4978d03f5722dee4b2f12931706d6e14c08661b13f814796cc9e", "f02b6bb2443c4a7221acba9da1c3f65cf1e14b674fdb8b27a26c6569b1b30ef4", "c28a8897a33c53997b0ba86d83067765573560d3a26059099217af5a213677ee", "52e57da3e2de2120dc435257aa8ed777428397de307adb2d1e0b90fb14440283", "d7c134a4dee1ebe221fb6f9fedc928994464feefb1ec445d139cdcce3f0ef326", "ac7f0d7356bd66660860d2cea338c44664ab1821d9b6b279377082789bfd7e38", "ff887f316e5c772906c50cad1ccd23855744deca61bf8686c6bafc0d3caef097", "1325aa2e9f3d24e1cd38b7db9c4c76505ec31cc1c820a9ecc4cbe8d9b7efc912", "6807882f4b9a754e31fd5b4456f8ce5df828dc27c60b1c8aba660f00164ab609"], "20000000", "1729d72d", "5b825563", true]}`
	extranonce2_size = 8
	extranonce1 = "3a2e2ac3"
	)

func GetHexTarget(packedTarget *big.Int) *big.Int {
	var mantissa big.Int
	var exponent big.Int
	var t1 big.Int
	var result big.Int
	t1.SetString("0x00ffffff", 0)
	mantissa.And(packedTarget, &t1)
	t1.SetString("0xff000000", 0)
	exponent.And(packedTarget, &t1)
	exponent.Rsh(&exponent, 6*4)
	result.Sub(&exponent, new(big.Int).SetUint64(3))
	result.Mul(&result, new(big.Int).SetInt64(8))
	result.Lsh(&mantissa, uint(result.Uint64()))
	return &result
}

func CalculateDifficulty(t *big.Int) *big.Float {
	var difficulty1Target big.Int
	difficulty1Target.SetString(BASE_1_DIFFICULTY, 0)
	var result big.Float
	result.Quo(new(big.Float).SetInt(&difficulty1Target), new(big.Float).SetInt(t))
	return &result
}

func doubleSHA256(s []byte) []byte {
	h := sha256.New()
	h.Write(s)
	h2 := sha256.New()
	h2.Write(h.Sum(nil))
	return h2.Sum(nil)
}

func reverseArray(a []byte) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func difficultyToHexTarget(a uint64) *big.Int {
	var difficulty1Target big.Int
	difficulty1Target.SetString(BASE_1_DIFFICULTY, 0)
	diff := new(big.Int).SetUint64(a)
	var target big.Int
	target.Quo(&difficulty1Target, diff)
	return &target
}

func hashRateApprox(diff uint64, t *big.Int) *big.Int {
	var hashRate big.Int
	diff_big := new(big.Int).SetUint64(diff)
	diff_big.Lsh(diff_big, 32)
	hashRate.Quo(diff_big, t)
	return &hashRate
}

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
}

func decodeMiningNotify(s map[string]interface{}) *mining_notify_job {
	t1 := s["params"].([]interface{})
	resp := mining_notify_job{
		job_id: t1[0].(string),
		prevhash: t1[1].(string),
		coinb1:t1[2].(string),
		coinb2:t1[3].(string),
		version:t1[5].(string),
		nbits:t1[6].(string),
		ntime:t1[7].(string),
		clean_jobs:t1[8].(bool),
	}
	resp.merkle_branch = make([]string, len((t1[4]).([]interface{})))
	for i,v := range (t1[4]).([]interface{}){
		resp.merkle_branch[i] = v.(string)
	}
	return &resp
}

func hexStringToByteArray(s *string) []byte {
	ba, _ := hex.DecodeString(*s)
	return ba
}

func calculateMerkelRoot(job  *mining_notify_job, cb *[]byte) []byte {
	merkle_root := make([]byte, len(*cb))
	copy(merkle_root, *cb)

	for _,v := range (*job).merkle_branch {
		merkle_root = doubleSHA256(append(merkle_root, hexStringToByteArray(&v)...))
	}

	return merkle_root
}

//To produce coinbase, we just concatenate Coinb1 + Extranonce1 + Extranonce2 + Coinb2 together.
func buildCoinBase(job  *mining_notify_job, extranonce1 *string, extranonce2 *string) *[]byte{
	s, _ := hex.DecodeString(
			(*job).coinb1 +
			*extranonce1 +
			*extranonce2 +
			(*job).coinb2)
	//fmt.Println(len((*job).coinb1))
	//fmt.Println(len(s))
	a := doubleSHA256(s)
	//reverseArray(a)
	return &a
}


func generateExtraNounce2() *string {
	n2 := "00000000"
	return &n2
}

func processMiningNotify(s map[string]interface{}){
	stru := *decodeMiningNotify(s)
	//fmt.Fprintf(os.Stdout,"%s\n", spew.Sdump(stru))
	ex1 := extranonce1
	coinbase := buildCoinBase(&stru, &ex1, generateExtraNounce2())
	///fmt.Println("coinbase: ", hex.EncodeToString(*coinbase))
	merkelRoot := calculateMerkelRoot(&stru, coinbase)
	//fmt.Println("Merkle Root is ", hex.EncodeToString(merkelRoot))
	mrs := hex.EncodeToString(merkelRoot)
	fmt.Println(" \t<getFPGAJob>", stru.getFPGAJob(&mrs))
	fmt.Printf(" \t<target> 0x%064x\n", difficultyToHexTarget(1024))

}


// TODO: Merkle Root Reverse Byte Order
func (self mining_notify_job) getFPGAJob(mr *string) string{

	t1 := self.version + self.prevhash + *mr + self.ntime + self.nbits
	return t1

}

func main() {
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jd), &dat); err != nil {
		panic(err)
	}

	if _ , ok := dat["method"]; ok{
		switch dat["method"] {
		case "mining.notify":
			processMiningNotify(dat)
		default:
			fmt.Fprintf(os.Stderr, "No such method %s is implemented!", dat["method"])
		}
	} else {
		fmt.Fprint(os.Stderr, "method is not a key in the json payload!")
	}

//
//
//	start := time.Now()
//
//	fmt.Println(jd)
//	fmt.Println()
//
//	var packedTarget big.Int
//	packedTarget.SetString("0x1b0404cb", 0)
//	fmt.Printf("target is 0x%064x \n", GetHexTarget(&packedTarget))
//	fmt.Printf("difficulty is %.8f \n", CalculateDifficulty(GetHexTarget(&packedTarget)))
//	s, _ := hex.DecodeString("01000000" + "81cd02ab7e569e8bcd9317e2fe99f2de44d49ab2b8851ba4a308000000000000" + "e320b6c2fffc8d750423db8b1eb942ae710e951ed797f7affc8892b0f1fc122b" +
//		"c7f5d74d" +
//		"f2b9441a" +
//		"42a14695")
//	a := doubleSHA256(s)
//	reverseArray(a)
//
//	fmt.Printf("target is : 0x%064x\n", difficultyToHexTarget(16307))
//
//	sha1_hash := hex.EncodeToString(a)
//	fmt.Println("Hash", sha1_hash)
//
//	// Hash Rate Approx
//	k := hashRateApprox(20000, new(big.Int).SetUint64(60*60*23.85))
//	fmt.Printf("Hash Rate Approx: %f G\n", float64(k.Uint64())/math.Pow(10, 9))
//
//	elapsed := time.Since(start)
//	log.Printf("Program took %sd", elapsed)
//
}
