package translator

import (
	"log"
	"nand2tetris/vm-translator/constants"
	"nand2tetris/vm-translator/helpers"
	"nand2tetris/vm-translator/parser"
	"strconv"
)

type Translator struct {
	bufIn  *[]string // Buffer to store source instructions
	bufOut *[]string // Buffer to hold translated instructions

	currentLoc   string // Current @ location
	staticLabel  string // Label for static variables
	variableName string // Name of the temporary variable
}

// Specify the input buffer where the source instructions are stored, the
// output buffer where the translated instructions will be stored, and the
// static label which should be the same as the VM file name ("Xxx.vm" would
// have the static label "Xxx")
func (tr *Translator) Setup(in *[]string, out *[]string, staticLabel string) {
	tr.bufIn = in
	tr.bufOut = out
	tr.staticLabel = staticLabel
}

// Translate all of the instructions stored in the input buffer. Results will be
// stored in the output buffer.
func (tr *Translator) TranslateAll() {
	for _, s := range *tr.bufIn {
		in := parser.ParseIn(s)

		switch in.Operator {
		case "PUSH":
			// Fetch data from target segment and offset, then write data to
			// main stack
			tr.fetchFrom(in.Segment, in.Dest, true)
			tr.writeTo("STACK", -1, true)

		case "POP":
			// Fetch data from main stack and then write to target segment and
			// offset

			if in.Segment != "TEMP" && in.Segment != "STATIC" {
				tr.resolveOffsetAdr(in.Segment, in.Dest) // Move to target adr
				*tr.bufOut = append(*tr.bufOut, "D=A")   // Store the address as data
				tr.createAsVariable()                    // Address is now stored as @tr.variableName for quick access
			}

			tr.fetchFrom("STACK", -1, true)
			tr.writeTo(in.Segment, in.Dest, true)
		default:
			tr.arithmetic(in.Operator)
		}
	}
}

// Move to the designated memory segment and slot number, and store the M value
// into D register. If main stack memory is specified, stack pointer is
// automatically decremented. If loadD is false, it perform only half of the
// operation, i.e. moving to the designated memory without loading any values.
// If the main stack is specified when loadD is off, stack pointer will NOT
// decrement. Offset is ignored when main stack is specified. If segment is
// "CONSTANT" then the A value will be loaded into D, instead of the M value.
func (tr *Translator) fetchFrom(seg string, offset int, loadD bool) {
	offsetString := strconv.Itoa(offset)

	if seg == "STACK" {
		if tr.currentLoc != "SP" {
			tr.currentLoc = "SP"
			*tr.bufOut = append(*tr.bufOut, "@SP")
		}

		tr.currentLoc = "*SP"
		*tr.bufOut = append(*tr.bufOut, "A=M-1") // Resolve pointer

		if loadD {
			*tr.bufOut = append(*tr.bufOut, "D=M")

			tr.decrementSP()
		}

		return
	}

	if seg == "CONSTANT" {
		if tr.currentLoc != offsetString {
			tr.currentLoc = offsetString
			*tr.bufOut = append(*tr.bufOut, "@"+offsetString)
		}

		if loadD {
			*tr.bufOut = append(*tr.bufOut, "D=A")
		}

		return
	}

	if seg == "STATIC" {
		varName := tr.staticLabel + "." + offsetString

		if tr.currentLoc != varName {
			tr.currentLoc = varName
			*tr.bufOut = append(*tr.bufOut, "@"+varName)
		}

		if loadD {
			*tr.bufOut = append(*tr.bufOut, "D=M")
		}

		return
	}

	tr.resolveOffsetAdr(seg, offset)

	if loadD {
		*tr.bufOut = append(*tr.bufOut, "D=M")
	}
}

// Move to the designated memory segment and slot number, and write the D value
// into M. If main stack memory is specified, stack pointer is
// automatically increased. If writeD is false, it perform only half of the
// operation, i.e. moving to the designated memory without writing any values.
// If the main stack is specified when writeD is off, stack pointer will NOT
// increment. Offset is ignored when main stack is specified.
func (tr *Translator) writeTo(seg string, offset int, writeD bool) {
	offsetString := strconv.Itoa(offset)

	if seg == "STACK" {
		if tr.currentLoc != "SP" {
			*tr.bufOut = append(*tr.bufOut, "@SP")
		}

		tr.currentLoc = "*SP"
		*tr.bufOut = append(*tr.bufOut, "A=M") // Resolve pointer

		if writeD {
			*tr.bufOut = append(*tr.bufOut, "M=D")

			// Increment stack pointer
			tr.incrementSP()
		}

		return
	}

	if seg == "STATIC" {
		varName := tr.staticLabel + "." + offsetString

		if tr.currentLoc != varName {
			tr.currentLoc = varName
			*tr.bufOut = append(*tr.bufOut, "@"+varName)
		}

		if writeD {
			*tr.bufOut = append(*tr.bufOut, "M=D")
		}

		return
	}

	if seg == "TEMP" {
		tr.resolveOffsetAdr(seg, offset)

		if writeD {
			*tr.bufOut = append(*tr.bufOut, "M=D")
		}

		return
	}

	// local, arg, this, and that
	if len(tr.variableName) == 0 {
		// The regular user of the VM translator probably won't understand this
		// error
		log.Fatalln("[!] Variable required for pointer access not found, aborting")
	}

	if tr.currentLoc != tr.variableName {
		tr.currentLoc = tr.variableName
		*tr.bufOut = append(*tr.bufOut, "@"+tr.variableName)
	}
	*tr.bufOut = append(*tr.bufOut, "A=M") // Go to the stored resolved address

	if writeD {
		*tr.bufOut = append(*tr.bufOut, "M=D") // Write D to the resolved address
	}
}

// Handle arithmetic operations like "EQ", "LT", etc. It pops the first (or
// first two) elements of the stack, perform the operation, and push the result
// onto the stack.
func (tr *Translator) arithmetic(operator string) {
	// Nested function to generate a single pair of condition-jump branch
	// Nested because it is not used anywhere else, for now
	generateBranchPair := func(conT string, conF string) {
		labelName := helpers.GetRandomString(6)
		labelNameNot := labelName + "_NOT"
		labelNameEnd := labelName + "_END"

		*tr.bufOut = append(*tr.bufOut, "@"+labelName)
		*tr.bufOut = append(*tr.bufOut, "D;"+conT) // If conT is true, JMP
		*tr.bufOut = append(*tr.bufOut, "@"+labelNameNot)
		*tr.bufOut = append(*tr.bufOut, "D;"+conF) // else jump not

		*tr.bufOut = append(*tr.bufOut, "("+labelName+")")
		// Because for some reason true is "-1"...
		// Do you know how many hours I've been trying to debug this?
		*tr.bufOut = append(*tr.bufOut, "D=-1")
		// Jump to end otherwise it continues straight-on and then everything
		// will be wrong and you will be sad and you will question your self
		// worth
		*tr.bufOut = append(*tr.bufOut, "@"+labelNameEnd)
		*tr.bufOut = append(*tr.bufOut, "0;JMP")

		*tr.bufOut = append(*tr.bufOut, "("+labelNameNot+")")
		*tr.bufOut = append(*tr.bufOut, "D=0")
		*tr.bufOut = append(*tr.bufOut, "("+labelNameEnd+")")
	}

	// If only half a pop operation is performed, we need to decrement the SP by
	// hand
	needsDecrementSP := false
	// Fetch first operand
	// I've moved this outside the switch-case because some operations only need
	// one operand
	tr.fetchFrom("STACK", -1, true)

	switch operator {
	case "ADD":
		tr.fetchFrom("STACK", -1, false)         // Go to but don't load main stack
		needsDecrementSP = true                  // But still decrement the SP so to overwrite the slot
		*tr.bufOut = append(*tr.bufOut, "D=M+D") // instruction to add

	case "SUB":
		tr.fetchFrom("STACK", -1, false)
		needsDecrementSP = true
		*tr.bufOut = append(*tr.bufOut, "D=M-D")

	case "NEG":
		*tr.bufOut = append(*tr.bufOut, "D=-D")

	case "EQ":
		tr.fetchFrom("STACK", -1, false)
		needsDecrementSP = true
		*tr.bufOut = append(*tr.bufOut, "D=D-M")
		generateBranchPair("JEQ", "JNE")

	// The course video has a typo "GET"
	// I thought it meant "greater or equal to"
	// Can you believe how many hours I spent on this?
	case "GT":
		tr.fetchFrom("STACK", -1, false)
		needsDecrementSP = true
		*tr.bufOut = append(*tr.bufOut, "D=M-D")
		generateBranchPair("JGT", "JLE")

	case "LT":
		tr.fetchFrom("STACK", -1, false)
		needsDecrementSP = true
		*tr.bufOut = append(*tr.bufOut, "D=M-D")
		generateBranchPair("JLT", "JGE")

	case "AND":
		tr.fetchFrom("STACK", -1, false)
		needsDecrementSP = true
		*tr.bufOut = append(*tr.bufOut, "D=D&M")

	case "OR":
		tr.fetchFrom("STACK", -1, false)
		needsDecrementSP = true
		*tr.bufOut = append(*tr.bufOut, "D=D|M")

	case "NOT":
		*tr.bufOut = append(*tr.bufOut, "D=!D")
	}

	if needsDecrementSP {
		// Treat data as popped even if not loaded onto D
		tr.decrementSP()
	}

	tr.writeTo("STACK", -1, true) // Write result to main stack
}

func (tr *Translator) decrementSP() {
	if tr.currentLoc != "SP" {
		tr.currentLoc = "SP"
		*tr.bufOut = append(*tr.bufOut, "@SP")
	}

	*tr.bufOut = append(*tr.bufOut, "M=M-1")
}

func (tr *Translator) incrementSP() {
	if tr.currentLoc != "SP" {
		tr.currentLoc = "SP"
		*tr.bufOut = append(*tr.bufOut, "@SP")
	}

	*tr.bufOut = append(*tr.bufOut, "M=M+1")
}

// Resolve dynamic addresses that use pointers, i.e. local, argument, this,
// that, temp, and pointer segments. This moves the A pointer to the resolved
// address.
// WARNING: D VALUE WILL BE LOST IF SEG IS NOT TEMP, so store it somewhere else first!
func (tr *Translator) resolveOffsetAdr(seg string, offset int) {
	ptr := constants.PtrWithOffset[seg]
	offsetString := strconv.Itoa(offset)

	if ptr == "@" {
		if seg == "TEMP" {
			ptr += strconv.Itoa(offset + constants.TempBaseAdr)
		} else {
			// This means segment is "POINTER"
			switch offset {
			case 0:
				ptr += strconv.Itoa(constants.ThisPtrAdr)
			case 1:
				ptr += strconv.Itoa(constants.ThatPtrAdr)

			default:
				log.Fatalln("[!] Unrecognised pointer offset: ", offset)
			}
		}

		if tr.currentLoc != ptr[1:] {
			tr.currentLoc = ptr[1:]
			*tr.bufOut = append(*tr.bufOut, ptr)
		}

		return
	}

	tr.fetchFrom("CONSTANT", offset, true) // This sets D == offset

	if tr.currentLoc != ptr[1:]+offsetString {
		tr.currentLoc = ptr[1:] + offsetString

		*tr.bufOut = append(*tr.bufOut, ptr)
	}

	*tr.bufOut = append(*tr.bufOut, "A=M+D")
}

// Create a new variable, and write the D register into its M value. The name of
// the created variable can be fetched from tr.variableName
func (tr *Translator) createAsVariable() {
	tr.variableName = helpers.GetRandomString(6)
	adr := "@" + tr.variableName

	// No need to check for currentLoc because it surely won't be at the
	// location you just created
	*tr.bufOut = append(*tr.bufOut, adr)
	*tr.bufOut = append(*tr.bufOut, "M=D")
}
