package helper

// Coalesce returns the first non-zero parameter value, and if the value is zero, it returns the default value def.
// Suitable for various comparable types such as int, string, boolean, struct, etc.
// When value==zero of def, it will trigger the return of def.
func Coalesce[T comparable](value, def T) T {
	var zero T
	if value == zero {
		return def
	}
	return value
}
