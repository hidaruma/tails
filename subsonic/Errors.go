/*
Copyright (C) 2014 Matt Layher
Released under the MIT license
http://opensource.org/licenses/mit-license.php
 */

package subsonic

var (
	ErrBadCredentials = func() *Container {
		c := NewContainer()
		c.Status = "failed"

		c.SubError = &Error{Code: 40, Message: "Wrong username or password."}
		return c
	}()

	ErrGeneric = func() *Container {
		c := NewContainer()
		c.Status = "failed"
		c.SubError = &Error{Code: 0, Message: "An error occurred."}
		return c
	}()

	ErrMissingParameter = func() *Container {
		c := NewContainer()
		c.Status = "failed"
		c.SubError = &Error{Code: 10, Message:"Required parameter is missing."}
		return c
	}()
)