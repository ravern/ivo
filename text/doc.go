// Package text implements useful text operations.
package text

// Lines are defined as slices of runes seperated by one '\n' rune.
// The result includes the '\n' rune at the end of each line (except
// the last line, which ends in 0.
//
// Paragraphs are defined as slices of sentences seperated by more
// than one '\n' rune.
//
// Sentences are defined as slices of words seperated runes in the
// set '.;!?'.
//
// Words are defined as slices of runes seperated by whitespace.
