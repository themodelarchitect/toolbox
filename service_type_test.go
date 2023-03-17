package tools

import (
	"fmt"
	"testing"
)

var ccursvctypetestcases = []string{"rb", "rec", "ts"}

func TestServiceTypeParam(t *testing.T) {
	var testCases []TestCase

	for _, tc := range testCases {
		tc := tc
		for _, serviceType := range ccursvctypetestcases {
			t.Run(fmt.Sprintf("Service type param %s %s %s %s %s", tc.RollingBuffer.MediaName, tc.StorageType, tc.PlaybackFormat, serviceType, tc.DRMType), func(t *testing.T) {
				teardownTest := setupTest(t)
				defer teardownTest(t)

				if tc.PlaybackFormat == DASH {
					t.Skip("skipping dash")
				}
				var url string
				if tc.PlaybackFormat == TS {
					if tc.DRMType != Unencrypted {
						t.Skip()
					}
					url = fmt.Sprintf("%s?ccur_svc_type=%s", tc.PlaybackURL, serviceType)
				}
				if tc.PlaybackFormat == FMP4 {
					url = fmt.Sprintf("%s&ccur_svc_type=%s", tc.PlaybackURL, serviceType)
				}
				if serviceType == "ts" && tc.DRMType != Unencrypted {
					t.Skip()
				}

				Log("",

					DEBUG, url)
				SegmentsOK(t, url, false, false)
			})
		}
	}
}
