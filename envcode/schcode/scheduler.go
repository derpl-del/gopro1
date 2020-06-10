package schcode

import (
	"github.com/derpl-del/gopro1/envcode/chcode"
	"github.com/robfig/cron"
)

//ClearCache scheduler for clean cache every 4 min
func ClearCache() {
	c := cron.New()
	c.AddFunc("4 * * * * *", chcode.DeleteSche)
	c.Start()
}
