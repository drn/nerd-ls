package options

// Options - Parsed input flags schema
type Options struct {
  All bool `short:"a" long:"all" description:"Include directory entries whose names begin with a dot (.)"`
  Long bool `short:"l" long:"long" description:"List in long format"`
  Icon bool `short:"i" long:"icon" description:"Display nerd-font icons"`
}

// Parse - Converts flags to string -> int map
func Parse(opts Options) map[string]int {
  return map[string]int{
    "all": btoi(opts.All),
    "long": btoi(opts.Long),
    "icon": btoi(opts.Icon),
  }
}

func btoi(boolean bool) int {
  if boolean { return 1 }
  return 0
}
