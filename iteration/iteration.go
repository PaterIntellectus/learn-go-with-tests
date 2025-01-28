package iteration

func Repeat(character string, times int) string {
	if times < 0 {
		panic("repeat number cannot be negative")
	}

	var repeated string

	for i := 0; i < times; i++ {
		repeated += character
	}

	if len(repeated) != len(character)*times {
		panic("resulting string overflows the correct length of correctly repeated string")
	}

	return repeated
}
