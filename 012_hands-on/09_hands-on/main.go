package main

func main() {

}

func prs(filePath string) []Record {
	/* src is a pointer to the file */
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	reader := csv.NewReader(src)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
}