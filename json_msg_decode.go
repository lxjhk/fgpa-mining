package main

func decodeMiningNotify(s map[string]interface{}) *mining_notify_job {
	t1 := s["params"].([]interface{})
	resp := mining_notify_job{
		job_id:     t1[0].(string),
		prevhash:   t1[1].(string),
		coinb1:     t1[2].(string),
		coinb2:     t1[3].(string),
		version:    t1[5].(string),
		nbits:      t1[6].(string),
		ntime:      t1[7].(string),
		clean_jobs: t1[8].(bool),
	}
	resp.merkle_branch = make([]string, len((t1[4]).([]interface{})))
	for i, v := range (t1[4]).([]interface{}) {
		resp.merkle_branch[i] = v.(string)
	}
	return &resp
}
