package uniq

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputError(t *testing.T) {
	inputFileName := "input_sa.txt"
	args := []string{inputFileName}

	err := uniq(args)

	assert.NotNil(t, err, "Ошибка прочтения входного файла")
}

// Normal situations
func TestNormal(t *testing.T) {
	expected :=
		`I love music.

I love music of Kartik.
Thanks.
I love music of Kartik.
`
	inputFileName := "input.txt"
	outputFileName := "output.txt"
	args := []string{inputFileName, outputFileName}

	uniq(args)

	// Сравниваем вывод с ожидаемым результатом
	content, err := os.ReadFile(outputFileName)
	if err != nil {
		log.Fatalf("Ошибка чтения файла: %v", err)
	}
	got := string(content)
	assert.Equal(t, got, expected, "Ситуация без флагов работает правильн")
}

func TestCFlag(t *testing.T) {
	expected :=
		`3 I love music.
1 
2 I love music of Kartik.
1 Thanks.
2 I love music of Kartik.
`

	flagC := "-c"
	inputFileName := "input.txt"
	outputFileName := "output_c.txt"
	args := []string{flagC, inputFileName, outputFileName}

	uniq(args)

	// Сравниваем вывод с ожидаемым результатом
	content, _ := os.ReadFile(outputFileName)
	got := string(content)
	assert.Equal(t, got, expected, "Флаг -c работает корректно")
}

func TestDFlag(t *testing.T) {
	expected :=
		`I love music.
I love music of Kartik.
I love music of Kartik.
`

	flagD := "-d"
	inputFileName := "input.txt"
	outputFileName := "output_d.txt"
	args := []string{flagD, inputFileName, outputFileName}

	uniq(args)

	// Сравниваем вывод с ожидаемым результатом
	content, _ := os.ReadFile(outputFileName)
	got := string(content)
	assert.Equal(t, got, expected, "Флаг -d работает правильно")
}

func TestUFlag(t *testing.T) {
	expected :=
		`
Thanks.
`
	flagU := "-u"
	inputFileName := "input.txt"
	outputFileName := "output_u.txt"
	args := []string{flagU, inputFileName, outputFileName}

	uniq(args)

	// Сравниваем вывод с ожидаемым результатом
	content, _ := os.ReadFile(outputFileName)
	got := string(content)
	assert.Equal(t, got, expected, "Флаг -u работает правильно")
}

func TestErrFlag(t *testing.T) {
	flagC := "-c"
	flagD := "-d"
	flagU := "-u"
	args := []string{flagC, flagD, flagU}

	err := uniq(args)

	assert.Nil(t, err, "Флаги -cdu взаимозаменяемы")
}

func TestFFlag(t *testing.T) {
	expected :=
		`We love music.

I love music of Kartik.
Thanks.
`
	flagF := "-f"
	inputFileName := "input_f.txt"
	outputFileName := "output_f.txt"
	args := []string{flagF, "1", inputFileName, outputFileName}

	uniq(args)

	// Сравниваем вывод с ожидаемым результатом
	content, _ := os.ReadFile(outputFileName)
	got := string(content)
	assert.Equal(t, got, expected, "Флаг -f работает правильно")
}

func TestSFlag(t *testing.T) {
	expected :=
		`I love music.

I love music of Kartik.
We love music of Kartik.
Thanks.
`
	flagS := "-s"
	inputFileName := "input_s.txt"
	outputFileName := "output_s.txt"
	args := []string{flagS, "1", inputFileName, outputFileName}

	uniq(args)

	// Сравниваем вывод с ожидаемым результатом
	content, _ := os.ReadFile(outputFileName)
	got := string(content)
	assert.Equal(t, got, expected, "Флаг -s работает правильно")
}
