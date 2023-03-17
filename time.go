package toolbox

import (
	"strconv"
	"strings"
	"time"
)

func TimeStamp() string {
	ts := time.Now().UTC().Format(time.RFC3339)
	return strings.Replace(strings.Replace(ts, ":", "", -1), "-", "", -1)
}

func EpochTime() int64 {
	return time.Now().Unix()
}

func EpochTimeStr() string {
	now := time.Now().Unix()
	s := strconv.FormatInt(now, 10)
	return s
}

func EpochToZulu(epoch int64) string {
	return time.Unix(epoch, 0).UTC().Format(time.RFC3339)
}
