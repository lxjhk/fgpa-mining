package main

func generateExtraNounce2() *string {
	n2 := "00000000"
	return &n2
}

func main() {

	ms := new(miningSession)
	ms.start()

	//start := time.Now()
	//var dat map[string]interface{}
	//if err := json.Unmarshal([]byte(jd), &dat); err != nil {
	//	panic(err)
	//}
	//
	//if _, ok := dat["method"]; ok {
	//	switch dat["method"] {
	//	case "mining.notify":
	//		processMiningNotify(dat)
	//	default:
	//		fmt.Fprintf(os.Stderr, "No such method %s is implemented!", dat["method"])
	//	}
	//} else {
	//	fmt.Fprint(os.Stderr, "method is not a key in the json payload!")
	//}
	//
	//elapsed := time.Since(start)
	//log.Printf("Program took %s", elapsed)

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
