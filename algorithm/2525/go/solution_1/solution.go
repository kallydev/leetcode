package solution

import "math/big"

func categorizeBox(length int, width int, height int, mass int) string {
	volume := new(big.Int).Mul(new(big.Int).Mul(big.NewInt(int64(length)), big.NewInt(int64(width))), big.NewInt(int64(height)))
	bulky := (length >= 10_000 || width >= 10_000 || height >= 10_000) || volume.Cmp(big.NewInt(1_000_000_000)) >= 0
	heavy := mass >= 100

	switch {
	case bulky && heavy:
		return "Both"
	case !bulky && !heavy:
		return "Neither"
	case bulky:
		return "Bulky"
	case heavy:
		return "Heavy"
	default:
		return "Neither"
	}
}
