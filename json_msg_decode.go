package main

import (
	"fmt"
	"strings"
	"time"
)

func (self *miningSession) decodeMiningNotify(s map[string]interface{}) *mining_notify_job {
	t1 := s["params"].([]interface{})
	for self.incomingDifficulty == 0 {
		time.Sleep(100 * time.Millisecond)
	}
	resp := mining_notify_job{
		job_id:     t1[0].(string),
		prevhash:   t1[1].(string),
		coinb1:     t1[2].(string),
		coinb2:     t1[3].(string),
		version:    t1[5].(string),
		nbits:      t1[6].(string),
		ntime:      t1[7].(string),
		clean_jobs: t1[8].(bool),
		difficulty: self.incomingDifficulty,
	}
	resp.merkle_branch = make([]string, len((t1[4]).([]interface{})))
	for i, v := range (t1[4]).([]interface{}) {
		resp.merkle_branch[i] = v.(string)
	}
	return &resp
}

func (self *miningSession) processMiningNotify(s map[string]interface{}) {
	stru := self.decodeMiningNotify(s)
	ex1 := self.extranonce1
	coinbase := stru.buildCoinBase(&ex1, generateExtraNounce2())
	merkelRoot := stru.calculateMerkelRoot(coinbase)

	var sb strings.Builder
	sb.WriteString("===========[NEW WORK]==============\n")
	sb.WriteString("==  " + fmt.Sprintf("<JobID> 0x%s\n", stru.job_id))
	sb.WriteString("==  " + fmt.Sprintf("<Extranonce2> 0x%s\n", *generateExtraNounce2()))
	sb.WriteString("==  " + fmt.Sprintf("<HeaderInfo Bytes> %d\n", len(stru.getFPGAJob(merkelRoot))/2))
	sb.WriteString("==  " + fmt.Sprintf("<^HeaderInfo> 0x%s\n", stru.getFPGAJob(merkelRoot)))
	sb.WriteString("==  " + fmt.Sprintf("<^Target DIFF %d> 0x%064x\n", stru.difficulty, difficultyToHexTarget(uint64(stru.difficulty))))
	sb.WriteString("+++++++++++++++++++++++++++++++++++\n")
	time.Sleep(100 * time.Millisecond)
	fmt.Print(sb.String())
}
