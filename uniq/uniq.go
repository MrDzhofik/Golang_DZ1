package uniq

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func countLines(scanner *bufio.Scanner, iFlag bool, sFlag int, fFlag int) ([]string, []int) {
	var prev string
	var curr string
	// Инициализация массивов
	text := make([]string, 0)
	number := make([]int, 0)
	// Считываем первую строку и кладем в массив
	scanner.Scan()
	text = append(text, scanner.Text())
	number = append(number, 1)
	index := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Флаг -i
		if iFlag {
			curr = strings.ToLower(line)
			prev = strings.ToLower(text[index])
		} else {
			curr = line
			prev = text[index]
		}

		// Флаг -f
		if fFlag >= 0 {
			fieldsCurr := strings.Fields(curr)
			fieldsPrev := strings.Fields(prev)
			if fFlag < len(fieldsCurr) && fFlag < len(fieldsPrev) {
				curr = strings.Join(fieldsCurr[fFlag:], " ")
				prev = strings.Join(fieldsPrev[fFlag:], " ")
			}
		}

		// Флаг -s
		if sFlag != 0 && len(curr) > sFlag && len(prev) > sFlag {
			curr = curr[sFlag:]
			prev = prev[sFlag:]
		}

		if curr != prev {
			number = append(number, 1)
		} else {
			tmp := index
			for number[tmp] == 0 {
				tmp--
			}
			number[tmp]++
			number = append(number, 0)
		}
		index++
		text = append(text, line)
	}
	return text, number
}

func parseFlag(args []string) ([]string, []*bool, []*int) {
	// Флаги
	fsb := flag.NewFlagSet("uniqFlags", flag.ExitOnError)
	boolFlags := []*bool{
		fsb.Bool("c", false, "Вывести количество повторений"),
		fsb.Bool("d", false, "Показать только повторяющиеся строки"),
		fsb.Bool("u", false, "Вывести только уникальные строки"),
		fsb.Bool("i", false, "Игнорировать регистр"),
	}
	hFlag := fsb.Bool("h", false, "Помощь")
	intFlags := []*int{
		fsb.Int("f", 0, "не учитывать первые `num_fields`"),
		fsb.Int("s", 0, "не учитывать первые `num_chars`"),
	}

	fsb.Parse(args)

	if *hFlag {
		fmt.Println("Использование: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
		flag.PrintDefaults()
		return nil, nil, nil
	}

	// Проверка флагов -c -d -u
	count := 0
	if *boolFlags[0] {
		count++
	}
	if *boolFlags[1] {
		count++
	}
	if *boolFlags[2] {
		count++
	}
	if count > 1 {
		fmt.Println("Ошибка: Флаги -c, -d, -u взаимозаменяемы. Пожалуйста, выберите только один из них.")
		// flag.PrintDefaults()
		return nil, nil, nil
	}

	filenames := []string{
		fsb.Arg(0),
		fsb.Arg(1),
	}
	return filenames, boolFlags, intFlags
}

func writeIn(writer *bufio.Writer, text []string, number []int, cFlag *bool, dFlag *bool, uFlag *bool) {
	if *cFlag {
		// Запись для флага -с
		for i := 0; i < len(text); i++ {
			if number[i] != 0 {
				fmt.Fprint(writer, number[i], " ")
				fmt.Fprint(writer, text[i], "\n")
			}
		}
	} else if *dFlag {
		// Запись для флага -d
		for i := 0; i < len(text); i++ {
			if number[i] > 1 {
				fmt.Fprint(writer, text[i], "\n")
			}
		}
	} else if *uFlag {
		// Запись для флага -u
		for i := 0; i < len(text); i++ {
			if number[i] == 1 {
				fmt.Fprint(writer, text[i], "\n")
			}
		}
	} else {
		// Запись без флагов
		for i := 0; i < len(text); i++ {
			if number[i] != 0 {
				fmt.Fprint(writer, text[i], "\n")
			}
		}
	}
}

func uniq(args []string) error {
	filenames, boolFlags, intFlags := parseFlag(args)

	if filenames == nil && boolFlags == nil && intFlags == nil {
		return nil
	}

	cFlag, dFlag, uFlag, iFlag := boolFlags[0], boolFlags[1], boolFlags[2], boolFlags[3]
	inputFileName, outputFileName := filenames[0], filenames[1]
	fFlag, sFlag := intFlags[0], intFlags[1]

	// Входной файл
	var inputFile *os.File
	var err error
	if inputFileName == "" {
		inputFile = os.Stdin
	} else {
		inputFile, err = os.Open(inputFileName)
		if err != nil {
			return err
		}
		defer inputFile.Close()
	}
	scanner := bufio.NewScanner(inputFile)

	// Выходной файл
	var writer *bufio.Writer
	var outputFile *os.File
	if outputFileName == "" {
		outputFile = os.Stdout
	} else {
		outputFile, err = os.Create(outputFileName)
		if err != nil {
			outputFile = os.Stdout
		}
		defer outputFile.Close()
	}

	writer = bufio.NewWriter(outputFile)

	// выполняем главную магию (подсчет уникальности)
	text, number := countLines(scanner, *iFlag, *sFlag, *fFlag)

	writeIn(writer, text, number, cFlag, dFlag, uFlag)

	writer.Flush()

	return nil
}
