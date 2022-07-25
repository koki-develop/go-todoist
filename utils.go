package todoist

// Returns a string as a pointer.
func String(s string) *string { return &s }

// Returns a int as a pointer.
func Int(i int) *int { return &i }

// Returns a bool as a pointer.
func Bool(b bool) *bool { return &b }
