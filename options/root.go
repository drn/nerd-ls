package options

// Options - Parsed input flags schema
type Options struct {
  All bool `short:"a" long:"all" description:"Include directory entries whose names begin with a dot (.)"`
  Long bool `short:"l" long:"long" description:"List in long format"`
  Icon bool `short:"i" long:"icon" description:"Display nerd-font icons"`
}
