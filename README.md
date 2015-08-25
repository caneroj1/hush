<h1>Hush</h1>

This a Go package for reading in a config file and providing access to the
config settings.

Hush reads from a local file called a .hushfile. You can have things in the
.hushfile like:

```
super_secret_key: abcdefghijklmnopqrstuvwxyz
secret_app_number: 42
```

It's probably a good idea to add the .hushfile to your .gitignore.

<h2>Still a WIP!!</h2>

* If you are using Hush without Revel, just put the .hushfile in the same directory from where you are running a Go program.

* If you are using Hush with Revel, put your .hushfile in the app/ directory or in the conf/ directory and be sure to call `revel run yourapp` from the root directory of yourapp.

<h3>How to Use Hush</h3>
<h4>Without Revel</h4>
```
package main

import "github.com/caneroj1/hush"

var secrets hush.Hush

func main() {
  secrets = hush.Hushfile()
  key, ok := secrets.GetString("super_secret_key")
}
```

<h4>With Revel</h4>
* In init.go:
```
package app

import (
	"github.com/caneroj1/hush"
	"github.com/revel/revel"
)

var secrets hush.Hush

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		... // omitted code
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	secrets = hush.Hushfile()}
```
