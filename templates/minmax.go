package main

type minmaxType1 int

func lt(a minmaxType1, b minmaxType1) bool {
	return a < b
}

//code starts

func min(a minmaxType1, b minmaxType1) minmaxType1 {
	if lt(a, b) {
		return a
	} else {
		return b
	}
}

func max(a minmaxType1, b minmaxType1) minmaxType1 {
	if lt(a, b) {
		return b
	} else {
		return a
	}
}
