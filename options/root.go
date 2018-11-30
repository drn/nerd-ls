package options

// Options - Parsed input flags schema
type Options struct {
  All bool `short:"a" long:"all" description:"Include directory entries whose names begin with a dot (.)"`
  Long bool `short:"l" long:"long" description:"List in long format"`
  Icon bool `short:"i" long:"icon" description:"Display nerd-font icons"`
}

// Parse - Converts flags to string -> interace{} map
func Parse(opts Options) map[string]interface{} {
  return map[string]interface{}{
    "all": opts.All,
    "long": opts.Long,
    "icon": opts.Icon,
  }
}
