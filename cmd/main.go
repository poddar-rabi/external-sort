package main

import (
	"bufio"
	"disksort/public"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter input directory: ")
	inputDir, _ := reader.ReadString('\n')

	fmt.Println("Enter output directory. The same directory would be used for output file: ")
	outputDir, _ := reader.ReadString('\n')

	fmt.Println("Enter memory limit: ")
	var limit int
	if _, err := fmt.Scan(&limit); err != nil {
		log.Print("  invalid limit", err)
		return
	}

	sorter, err := public.NewSorter(strings.TrimSpace(inputDir), strings.TrimSpace(outputDir), limit)
	if err != nil {
		log.Fatal(fmt.Sprintf("error thrown while creating sort instance %v", err))
	}

	fileName, err := sorter.Sort()
	if err != nil {
		log.Fatal(fmt.Sprintf("error sorting %v", err))

	}
	fmt.Println("sorted file: ", *fileName)
}
