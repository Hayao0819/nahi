package flist

type options = struct {
	maxDepth       int
	minDepth       int
	fileOnly       bool
	dirOnly        bool
	filename       bool
	filenameMatch  []string
	extsOnly       []string
	relpath        bool
	executableOnly bool
}

type Option func(*options)
