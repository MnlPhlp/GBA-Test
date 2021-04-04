module GBA-Test

go 1.16

require (
	github.com/MnlPhlp/gbaLib v0.0.0-20210401170324-785248fecc7b
	tinygo.org/x/drivers v0.15.1 // indirect
	tinygo.org/x/tinydraw v0.0.0-20200416172542-c30d6d84353c
	tinygo.org/x/tinyfont v0.2.1
)

replace github.com/MnlPhlp/gbaLib v0.0.0-20210401170324-785248fecc7b => ../gbaLib
