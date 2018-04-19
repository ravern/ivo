// Package ivo implements core packages.
//
// The root package provides many of the core interfaces for
// customizing and creating packages for Ivo. It also contains the
// main event loop and the core rendering, which relies on termbox
// to interact with the terminal. The other subpackages provide
// common implementations or utilities that supplement these core
// types to make custom packages easier to build.
//
// Window is the most important interface. It is considered the
// core building block of Ivo, since custom packages are split by
// the Window implementations they provide. The concept is that
// the main loop will send events to a single Window. If more than
// one Window is required, then Window multiplexers can be used to
// split the screen.
package ivo
