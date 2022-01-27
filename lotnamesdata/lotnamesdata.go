package lotnamesdata

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"strings"

	//	"fmt"
	"os"
)

// FileDoc ...
type FileDoc struct {
	IDNUM        string
	LotNum       string
	Lot          string
	MCFileName   string
	Filter       string
	AutoChg      string
	BaseChg      string
	Lane         string
	ProductionID string
	SimProduct   string
	DGSPCBName   string
	DGSPCBRev    string
	DGSPCBSide   string
	DGSRefPin    string
	C            string
	DataGenMode  string
	MountHead    string
	VSTPath      string
	Order        string
	TargetTact   string
}

// LotNamesData ...
func LotNamesData(file string, key string) (str string) {

	linesPana2, err := os.Open(file) // dat1
	defer linesPana2.Close()
	//	cr := csv.NewReader(linesPana2)
	//	cr.LazyQuotes = true
	//	cr.Comma = ' '
	//	records, err := cr.ReadAll()
	records, err := readSkipRow(linesPana2)
	if err != nil {
		fmt.Println(err) //(err)
	}

	//var str string

	for _, linePana := range records {
		dataPana := FileDoc{
			IDNUM:        linePana[0],
			LotNum:       linePana[1],
			Lot:          linePana[2],
			MCFileName:   linePana[3],
			Filter:       linePana[4],
			AutoChg:      linePana[5],
			BaseChg:      linePana[6],
			Lane:         linePana[7],
			ProductionID: linePana[8],
			SimProduct:   linePana[9],
			DGSPCBName:   linePana[10],
			DGSPCBRev:    linePana[11],
			DGSPCBSide:   linePana[12],
			DGSRefPin:    linePana[13],
			C:            linePana[14],
			DataGenMode:  linePana[15],
			MountHead:    linePana[16],
			VSTPath:      linePana[17],
			Order:        linePana[18],
			TargetTact:   linePana[19],
		}

		fmt.Printf("dataPana.Lot: %s, key: %s\n", dataPana.Lot, key)
		fmt.Printf("LEN dataPana.Lot: %d, LEN key: %d\n", len(dataPana.Lot), len(key))

		//	if len(dataPana.Lot) == len(key) {
		if strings.ToLower(dataPana.Lot) == strings.ToLower(key) { // "NPM_brain-1_p"
			//	result := []string{}
			//	fmt.Println(dataPana.LotNum)
			str := dataPana.LotNum
			return str
		}
		//	} else {
		//		color.Set(color.FgRed)
		//		fmt.Println("Название рецепта соответствует частично!")
		//		fmt.Println("Проверьте правильность ввода названия рецепта!")
		//		color.Unset()
		//		log.Fatal()
		//	}
	}
	return
}

func readSkipRow(rs io.ReadSeeker) ([][]string, error) {
	// Skip first row (line)
	row1, err := bufio.NewReader(rs).ReadSlice('\n')
	if err != nil {
		return nil, err
	}
	_, err = rs.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		return nil, err
	}

	// Read remaining rows
	r := csv.NewReader(rs)
	r.Comma = ' '
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return rows, nil
}
