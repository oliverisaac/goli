package goli

// Must is a helper function to unwrap two-type returns into a single return
// This is helpful when writing hacky go scripts just to swallow the error value.
func Must[T any](ret T, err error) T {
	if err != nil {
		panic(err)
	}
	return ret
}

// MustOK is a helper function to swallow the "found" boolean that is used
// For example: `envVal := MustOK(os.LookupEnv("RANDOM_THING"))`
func MustOK[T any](ret T, ok bool) T {
	if !ok {
		panic("unexpected false from ok check")
	}
	return ret
}
