# package cliset

`go get github.com/Qs-F/cliset`

cliset is a helping tool for developing cli in Golang.  
Supporting you on Output/Input.

## Installation

`go get github.com/Qs-F/cliset`

## Design

cliset.Warn(string)

cliset.IsPiped() (bool)

cliset.Arg() // Return arguments from line or line is empty try to return piped value

cliset.RegSubcommand // Register sub command by flag name
