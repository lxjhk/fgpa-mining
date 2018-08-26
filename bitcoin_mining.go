package main

import (
	"fmt"
	"encoding/json"
	"os"
	"time"
	"log"
)

const (
	jd               = `{"id": null, "method": "mining.notify", "params": ["8be5", "d7ea32ad603a57b96832465b110d4bbfcc70d6c2001b8f4a0000000000000000", "01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff5c03923708192f5669614254432f4d696e65642062792077346d387374322f2cfabe6d6dbd21d45341c3cbb7a366af31c1fe743d27dbe660e44dcae9275699982333733c040000000000000010e58b370862685029", "ffffffff02b3b0094b000000001976a914536ffa992491508dca0354e52f32a3a7a679a53a88ac0000000000000000266a24aa21a9eddf9626ca545c861be96e1a52fbcea9e27a3c1169d605b031dcdedf64d0461e7700000000", ["3a5328ad196525bc87cc3ad830b0354d8b6e87fc2851f1af5dee89cfd773dacf", "0a19ec9f5cdb4978d03f5722dee4b2f12931706d6e14c08661b13f814796cc9e", "f02b6bb2443c4a7221acba9da1c3f65cf1e14b674fdb8b27a26c6569b1b30ef4", "c28a8897a33c53997b0ba86d83067765573560d3a26059099217af5a213677ee", "52e57da3e2de2120dc435257aa8ed777428397de307adb2d1e0b90fb14440283", "d7c134a4dee1ebe221fb6f9fedc928994464feefb1ec445d139cdcce3f0ef326", "ac7f0d7356bd66660860d2cea338c44664ab1821d9b6b279377082789bfd7e38", "ff887f316e5c772906c50cad1ccd23855744deca61bf8686c6bafc0d3caef097", "1325aa2e9f3d24e1cd38b7db9c4c76505ec31cc1c820a9ecc4cbe8d9b7efc912", "6807882f4b9a754e31fd5b4456f8ce5df828dc27c60b1c8aba660f00164ab609"], "20000000", "1729d72d", "5b825563", true]}`
	extranonce2_size = 8
	extranonce1      = "3a2e2ac3"
)

func generateExtraNounce2() *string {
	n2 := "00000000"
	return &n2
}

func processMiningNotify(s map[string]interface{}) {
	stru := *decodeMiningNotify(s)
	//fmt.Fprintf(os.Stdout,"%s\n", spew.Sdump(stru))
	ex1 := extranonce1
	coinbase := stru.buildCoinBase(&ex1, generateExtraNounce2())
	///fmt.Println("coinbase: ", hex.EncodeToString(*coinbase))
	merkelRoot := stru.calculateMerkelRoot(coinbase)
	fmt.Println("<headerinfo>", stru.getFPGAJob(merkelRoot))
	fmt.Printf("<target> 0x%064x\n", difficultyToHexTarget(1024))

}



func main() {
	start := time.Now()
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jd), &dat); err != nil {
		panic(err)
	}

	if _, ok := dat["method"]; ok {
		switch dat["method"] {
		case "mining.notify":
			processMiningNotify(dat)
		default:
			fmt.Fprintf(os.Stderr, "No such method %s is implemented!", dat["method"])
		}
	} else {
		fmt.Fprint(os.Stderr, "method is not a key in the json payload!")
	}

	elapsed := time.Since(start)
	log.Printf("Program took %s", elapsed)

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
