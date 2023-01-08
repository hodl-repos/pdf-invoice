// Package middleware defines common http middlewares.
package middleware

// contextKey is a unique key for a context value to prevent clashing between
// different packages using the context to pass data.
type contextKey string
