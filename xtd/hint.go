package xtd

// NoCompare is an uncomparable type. Embed this inside another struct to make it uncomparable.
// This DOES NOT:
//  - Disallow shallow copies of structs
//  - Disallow comparison of pointers to uncomparable structs
//
// Example:
//   type Foo struct {
//       _ xtd.NoCompare
//       ...
//   }
type NoCompare [0]func()

// NoCopy may be embedded into structs which must not be copied
// after the first use. NoCopy also implies NoCompare.
// See https://golang.org/issues/8005#issuecomment-190753527 for details.
//
// Example:
//   type Foo struct {
//       _ xtd.NoCopy
//       ...
//   }
type NoCopy struct {
	NoCompare
}

// Lock is a no-op used by -copylocks checker from `go vet`.
func (*NoCopy) Lock()   {}
func (*NoCopy) Unlock() {}
