package pricefiller

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func floatEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < 1e-9
}

func TestCalculateAmountUsd(t *testing.T) {
	// RSR
	usdAmountRSR := calculateAmountUsd("336970435721651800000000", 18, 0.003863)
	assert.True(t, floatEqual(usdAmountRSR, 1301.716793192741), usdAmountRSR)

	// WETH
	usdtAmountWETH := calculateAmountUsd("503568522079108960", 18, 2600)
	assert.True(t, floatEqual(usdtAmountWETH, 1309.2781574056833), usdtAmountWETH)
}
