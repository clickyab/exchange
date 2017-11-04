package exchange

// IncShare return price of impression plus profit
// p is price and s is share
func IncShare(p float64, s int) float64 {
	return (float64(s+100) * p) / 100
}

// DecShare return price of impression without profit
// p is price and s is share
func DecShare(p float64, s int) float64 {
	return (p * 100) / float64(s+100)
}
