package utils

type ContextKey string

const (
	Connection ContextKey = "connection"
	Commit     ContextKey = "commit"
	PatchName  ContextKey = "patchName"
)
