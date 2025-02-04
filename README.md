# Silent Map Key Errors in Go

This repository demonstrates a common, yet subtle, error in Go: the silent handling of missing keys in maps.  Accessing a non-existent key in a Go map will not result in an error; instead, it returns the zero value for the map's value type. This can lead to unexpected behavior and difficult-to-debug issues.

The `bug.go` file shows an example of this behavior.  The `bugSolution.go` file offers a solution illustrating how to handle missing keys safely.
