package customAtomic

import "sync/atomic"

type SwappableRune struct {
	value *atomic.Int32
}

func (sa *SwappableRune) SwapAtom(other *SwappableRune, condition rune) (success bool) {
	agentVal := sa.Load()
	otherVal := other.Load()

	if otherVal != condition {
		return false
	}

	if other.CompareAndSwap(' ', agentVal) {
		if sa.CompareAndSwap(agentVal, otherVal) {
			return true
		}
	}

	return false
}

func NewSwappableRune(rn rune) SwappableRune {
	var atom atomic.Int32
	atom.Store(int32(rn))
	return SwappableRune{value: &atom}
}

func (sa *SwappableRune) Load() rune {
	return rune(sa.value.Load())
}

func (sa *SwappableRune) Store(rn rune) {
	sa.value.Store(int32(rn))
}

func (sa *SwappableRune) CompareAndSwap(old, new rune) (success bool) {
	return sa.value.CompareAndSwap(int32(old), int32(new))
}

func (sa *SwappableRune) Swap(new rune) (old rune) {
	return sa.value.Swap(int32(new))
}
