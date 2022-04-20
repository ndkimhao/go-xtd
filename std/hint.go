package std

type NoCmpCopy struct {
	NoCmp
	NoCopy
}

// NoCmp is an uncomparable type. Embed this inside another struct to make
// it uncomparable.
//
// This DOES NOT:
//  - Disallow shallow copies of structs
//  - Disallow comparison of pointers to uncomparable structs
type NoCmp [0]func()

// NoCopy may be embedded into structs which must not be copied
// after the first use.
//
// See https://golang.org/issues/8005#issuecomment-190753527
// for details.
type NoCopy struct{}

// Lock is a no-op used by -copylocks checker from `go vet`.
func (*NoCopy) Lock()   {}
func (*NoCopy) Unlock() {}
