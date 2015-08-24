= Hush

This a Go package for reading in a config file and providing access to the
config settings.

Hush reads from a local file called a .hushfile. You can have things in the
.hushfile like:

`
super_secret_key: abcdefghijklmnopqrstuvwxyz
secret_app_number: 42
`

It's probably a good idea to add the .hushfile to your .gitignore.
