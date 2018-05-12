# gopool
A generic goroutine pooling library

This library provides low level mechanics for maintaining a dynamic pool of goroutines that process tasks.  It is intentionally simple to allow for higher level functionality to be build into the implementation.  As such, concepts such as context, logging, error handling and metrics are not directly impemented by this library, it simply provides the minimum of support for those concepts to exist.

The examples folder holds simple code implemetations of this library.
