package drm

import (
	"fmt"
	"os"
	"testing"
)

type (
	cardDetail struct {
		version      Version
		capabilities map[uint64]uint64
	}
)

var (
	card, errCard = Available()
	cards         = map[string]cardDetail{
		"i915": cardDetail{
			version: Version{
				Major: 1,
				Minor: 6,
				Patch: 1,
				Name:  "i915",
				Desc:  "i915",
				Date:  "20160425",
			},
			capabilities: map[uint64]uint64{
				CapDumbBuffer:         1,
				CapVBlankHighCRTC:     1,
				CapDumbPreferredDepth: 24,
				CapDumbPreferShadow:   1,
				CapPrime:              3,
				CapTimestampMonotonic: 1,
				CapAsyncPageFlip:      0,
				CapCursorWidth:        256,
				CapCursorHeight:       256,

				CapAddFB2Modifiers: 1,
			},
		},
	}
	cardInfo cardDetail
)

func TestMain(m *testing.M) {
	cards[""] = cards["i915"] // i915 bug in 4.8 kernel?

	if errCard != nil {
		fmt.Fprintf(os.Stderr, "No graphics card available to test")
		os.Exit(1)
	}
	if _, ok := cards[card.Name]; !ok {
		fmt.Fprintf(os.Stderr, "No tests for card %s", card.Name)
		os.Exit(1)
	}
	cardInfo = cards[card.Name]
	os.Exit(m.Run())
}
