package script

const (
	// Push-data opcodes
	OP_0         = 0x00 // an empty array of bytes is pushed onto the stack
	OP_FALSE     = OP_0
	OP_DATA_1    = 0x01 // single-byte push up to
	OP_DATA_75   = 0x4b // 75 - The next opcode bytes is data to be pushed onto the stack
	OP_PUSHDATA1 = 0x4c // 76 - The next byte contains the number of bytes to be pushed onto the stack.
	OP_PUSHDATA2 = 0x4d // 77 - The next two bytes contain the number of bytes to be pushed onto the stack in little endian order.
	OP_PUSHDATA4 = 0x4e // 78 - The next four bytes contain the number of bytes to be pushed onto the stack in little endian order.

	// Constants used to push specific numeric values
	OP_1NEGATE  = 0x4f // pushes -1
	OP_RESERVED = 0x50
	OP_1        = 0x51 // pushes 1
	OP_TRUE     = OP_1
	OP_2        = 0x52
	OP_3        = 0x53
	OP_4        = 0x54
	OP_5        = 0x55
	OP_6        = 0x56
	OP_7        = 0x57
	OP_8        = 0x58
	OP_9        = 0x59
	OP_10       = 0x5a
	OP_11       = 0x5b
	OP_12       = 0x5c
	OP_13       = 0x5d
	OP_14       = 0x5e
	OP_15       = 0x5f
	OP_16       = 0x60
)
