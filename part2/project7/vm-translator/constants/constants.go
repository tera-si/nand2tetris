package constants

const (
	HelpMsg = `Hack VM Translator
Usage:
	vm-translator [-h/--help] [-v/--verbose] BYTECODE

Flags:
	-h/--help            Shows this help message and exits.

Positional Argument:
	BYTECODE             File containing byte code for the Hack virtual machine.
							File is expected to have ".vm" extension. Required.

Description:
	The Hack virtual machine translator reads an Hack virtual machine bytecode
	file (.vm) and outputs its equivalent Hack assembly code (.asm). Both the
	bytecode and assembly code follow the Hack computer architecture and language
	specification defined in the Nand2Tetris courseware.

	This translator is project #7 of the Nand2Tetris (https://www.nand2tetris.org)
	courseware and book "The Elements of Computing Systems" by Noam Nisan and
	Shimon Schocken. This implementation is written in GO by
	tera-si (https://github.com/tera-si).
`

	CommentMarker = "//" // Afaik, there are no multi-line comment in Hack VM

	TempBaseAdr = 5 // Base address of temp is RAM[5]
	MaxTempNum  = 8 // There are only 8 temp slots

	StackBaseAdr = 256

	ThisPtrAdr = 3 // Address where the "THIS" pointer is stored
	ThatPtrAdr = 4 // Address where the "THAT" pointer is stored
)

var PtrWithOffset = map[string]string{
	"LOCAL":    "@LCL",
	"ARGUMENT": "@ARG",
	"THIS":     "@THIS",
	"THAT":     "@THAT",
	"TEMP":     "@",
	"POINTER":  "@",
}
