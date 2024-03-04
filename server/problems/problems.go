package problems

type Problem struct {
	Text     string
	Solution string
	Answer   string
}

var problems map[string]*Problem
