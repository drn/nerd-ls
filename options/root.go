package options

// Options - Parsed input flags schema
type Options struct {
	All      bool `short:"a" long:"all" description:"Include directory entries whose names begin with a dot (.)"`
	Long     bool `short:"l" long:"long" description:"List in long format"`
	Icon     bool `short:"i" long:"icon" description:"Display nerd-font icons"`
	Tree     bool `short:"T" long:"tree" description:"Recurse into directories as a tree"`
	SizeSort bool `short:"S" long:"size-sort" description:"Sort files by descending size"`
}

// Parse - Converts flags to string -> interface{} map
func Parse(opts Options) map[string]interface{} {
	return map[string]interface{}{
		"all":       opts.All,
		"long":      opts.Long,
		"icon":      opts.Icon,
		"tree":      opts.Tree,
		"size-sort": opts.SizeSort,
	}
}
