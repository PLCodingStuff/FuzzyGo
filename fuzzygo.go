package main

// Membership Fuction type
type MembershipFunction struct {
	from     float64
	to       float64
	funcName string
}

// Fuzzy Set type
type FuzzySet struct {
	mFunc MembershipFunction
	name  string
}
