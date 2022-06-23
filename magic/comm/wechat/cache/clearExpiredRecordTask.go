package cache

//每天晚上1点执行清理
func startClearExpiredRecordTask() {
	//now := time.Now()
	//next := now.Add(time.Hour * 24)
	//next = time.Date(next.Year(), next.Month(), next.Day(), 1, 0, 0, 0, next.Location())
	//t := time.NewTimer(next.Sub(now))
	//<-t.C
	//clearExpiredRecord()
	//timer := time.NewTicker(24 * time.Hour)
	//for range timer.C {
	//	clearExpiredRecord()
	//}
}

func clearExpiredRecord() {
	//var err error
	//cli := db.Get()
	//result := cli.Table(counterTableName).
	//	Where("updatetime < ?", time.Now().Add(-time.Hour)).
	//	Delete(model.Counter{})
	//if err = result.Error; err != nil {
	//	log.Error(err)
	//}
	//log.Info("delete expired record: ", result.RowsAffected)
}
