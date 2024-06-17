package ex9_test

import (
	"testing"

	"github.com/StasMerzlyakov/yagoadvanst/cmd/sprint7/task4/les2/ex9"
)

func BenchmarkAlignRight(b *testing.B) {
	testStr := "echo ыва asdasd"
	lenght := 50
	lead := ' '
	const testN = 100000

	b.ResetTimer()

	b.Run("AlignRightAppend", func(b *testing.B) {
		for i := 0; i < testN; i++ {
			ex9.AlignRightAppend(testStr, lenght, lead)
		}
	})

	b.Run("AlignRightBuffer", func(b *testing.B) {
		for i := 0; i < testN; i++ {
			ex9.AlignRightBuffer(testStr, lenght, lead)
		}

	})

	b.Run("AlignRightRepeat", func(b *testing.B) {
		for i := 0; i < testN; i++ {
			ex9.AlignRightRepeat(testStr, lenght, lead)
		}
	})
}
