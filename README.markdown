JSON Utility Package
====================

Install it:

	goinstall github.com/crazy2be/jsonutil

Import it:

	import "github.com/crazy2be/jsonutil"

Use it:

	v := someFunction()
	err := jsonutil.EncodeToFile("someFunction.json", v)

	if err != nil {
		fmt.Println("Error encoding value to file!", err)
	}

Intro
-----

Useful utility functions for use with the JSON package.

Functions
---------

Full function reference available at http://gopkgdoc.appspot.com/pkg/github.com/crazy2be/jsonutil.
