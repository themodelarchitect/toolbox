package tools

import (
	"fmt"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	if !*hyperConverged {
		t.Skip()
	}

	channelNames := GetLiveChannelNames()
	for _, channelName := range channelNames {
		channelName := channelName
		t.Run(fmt.Sprintf("Latency_between_instances_%s", channelName), func(t *testing.T) {
			teardownTest := setupTest(t)
			defer teardownTest(t)

			redisInstances, err := getRedisContainers()
			ok(t, err)
			Log("", DEBUG, "%v", redisInstances)
			assert(t, len(redisInstances) > 1, fmt.Sprintf("Found %d redis instances.", len(redisInstances)))
			timeout := time.After(60 * time.Second)
			tick := time.Tick(10 * time.Second)
			channelInstance, err := getChannelInstance(channelName)
			ok(t, err)
			Log("", DEBUG, "Channel instance: %s", channelInstance)
			for {
				select {
				case <-timeout:
					return
				case <-tick:
					var etimes []int
					streamInfo, err := getStreamInfo(channelName, redisInstances)
					ok(t, err)
					for k, v := range streamInfo {
						Log("", DEBUG, "%s => %#v", k, v)
						for _, v := range v {
							etimes = append(etimes, v.etime)
						}
					}

					Log("", DEBUG, "%#v", etimes)
					min, max := minMax(etimes)
					Log("", DEBUG, "Max etime: %d", max)
					Log("", DEBUG, "Min etime: %d", min)

					assert(t, (max-min) < 5, fmt.Sprintf("etimes are too far apart...\n Min => %d\n Max => %d\n", min, max))
				}
			}
		})
	}
}
