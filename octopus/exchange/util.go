package exchange

// IncShare return price of impression plus profit
// p is price and s is share
func IncShare(p float64, s int) float64 {
	return (float64(s+100) * p) / 100
}

// DecShare return price of impression without profit
// p is price and s is share
func DecShare(p int64, s int) int64 {
	return (p * 100) / int64(s+100)
}

// ProfitShare returns the pure profit out of an ad
func ProfitShare(p int64, s int) int64 {
	return p - DecShare(p, s)
}
