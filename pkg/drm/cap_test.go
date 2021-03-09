package drm

import (
	"testing"
)

func TestHasDumbBuffer(t *testing.T) {
	file, err := OpenCard(0)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	version, err := GetVersion(file)
	if err != nil {
		t.Error(err)
		return
	}
	if hasDumb := HasDumbBuffer(file); hasDumb != (cardInfo.capabilities[CapDumbBuffer] != 0) {
		t.Errorf("Card '%s' should support dumb buffers...Got %v but %d", version.Name, hasDumb, cardInfo.capabilities[CapDumbBuffer])
		return
	}
}

func TestGetCap(t *testing.T) {
	file, err := OpenCard(0)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	for cap, capval := range cardInfo.capabilities {
		ccap, err := GetCap(file, cap)
		if err != nil {
			t.Error(err)
			return
		}
		if ccap != capval {
			t.Errorf("Capability %d differs: %d != %d", cap, ccap, capval)
			return
		}

	}
}
