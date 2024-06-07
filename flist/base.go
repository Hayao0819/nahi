package flist

type options = struct {
	maxDepth       int
	minDepth       int
	fileOnly       bool
	dirOnly        bool
	extOnly        string
	filename       bool
	extsOnly       []string
	relpath        bool
	executableOnly bool
}

type Option func(*options)
