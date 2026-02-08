module golang/modules/hello

go 1.25.6

require (
	golang/modules/greetings v0.0.0
)

replace golang/modules/greetings => ../greetings
