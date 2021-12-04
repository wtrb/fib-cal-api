package db

import (
	historydb "github.com/wtrb/fib-cal-api/database/history-db"
	"github.com/wtrb/fib-cal-api/database/queue"
)

func init() {

}

func SelectAllValues() *[]string {
	return historydb.SelectAll()
}

func InsertValue(val string) {
	historydb.Insert(val)
}

func SelectValueByIndex(idx string) string {
	return queue.Get(idx)
}

func Queue(idx string) {
	queue.Add(idx, "Processing")
	queue.Publish(idx)
}
