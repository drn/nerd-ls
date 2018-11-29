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
    "all": boolToInt(opts.All),
    "long": boolToInt(opts.Long),
    "icon": boolToInt(opts.Icon),
  }
}

func boolToInt(boolean bool) int {
  if boolean { return 1 }
  return 0
}
