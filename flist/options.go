package flist

func WithMaxDepth(depth int) Option {
	return func(opt *options) {
		opt.maxDepth = depth
	}
}

func WithMinDepth(depth int) Option {
	return func(opt *options) {
		opt.minDepth = depth
	}
}

func WithExactDepth(depth int) Option {
	return func(opt *options) {
		opt.maxDepth = depth
		opt.minDepth = depth
	}
}

func WithExtOnly(ext string) Option {
	return func(opt *options) {
		opt.extOnly = ext
	}
}

func WithExtsOnly(exts []string) Option {
	return func(opt *options) {
		opt.extsOnly = exts
	}
}

func WithFileOnly() func(*options) {
	return func(opt *options) {
		opt.fileOnly = true
	}
}

func WithFileName() Option {
	return func(opt *options) {
		opt.filename = true
	}
}

func WithDirOnly() Option {
	return func(opt *options) {
		opt.dirOnly = true
	}
}

func WithRelPath() Option {
	return func(opt *options) {
		opt.relpath = true
	}
}

func WithExecutableOnly() Option {
	return func(opt *options) {
		opt.executableOnly = true
	}
}
