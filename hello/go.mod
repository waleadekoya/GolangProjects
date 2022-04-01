module example.com/hello

go 1.18

require (
	example.com/arrays v0.0.0-00010101000000-000000000000
	example.com/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

require (
	github.com/yuin/goldmark v1.4.1 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3 // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/sys v0.0.0-20211019181941-9d821ace8654 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.10 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	rsc.io/sampler v1.3.0 // indirect
)

// https://go.dev/doc/tutorial/call-module-code
replace example.com/greetings => ../greetings

replace example.com/arrays => ../arrays
