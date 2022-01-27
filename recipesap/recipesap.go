package recipesap

import (
	"encoding/csv"
	"fmt"
	"os"

	//	"github.com/eugenefoxx/test/readIni/"
	"github.com/eugenefoxx/readDGSP1forKATE/setting"
)

type datpatrsData struct {
	IDNUM       string
	NAME        string
	EX          string
	LNAME       string
	FA          string
	FB          string
	FC          string
	FD          string
	FE          string
	FF          string
	FG          string
	FH          string
	FI          string
	FJ          string
	REELS       string
	CHIP        string
	Shape       string
	MCOND       string
	DCOND       string
	DCOND2      string
	ICOND       string
	SKIP        string
	A           string
	PACK        string
	PACKB       string
	SSIZE       string
	RETRY       string
	ESTOP       string
	ThrowAway   string
	LandA       string
	LandB       string
	LandC       string
	LandD       string
	PreSend     string
	NoArrange   string
	Proved      string
	UsePeriod   string
	C           string
	CoverSize   string
	Teach       string
	C2          string
	RETRYCOUNT  string
	ESTOPCOUNT  string
	RETRYCOUNT2 string
	SICOND      string
	HRMCOND     string
	HPICOND     string
	SL          string
	InCOND      string
	FeederEX    string
	MSD         string
	NPI         string
	LCR         string
	REELTYPECHG string
	FEEDLENCHG  string
}

type outdatpatternData struct {
	IDNUM string
	Qty   string
}

func MakeRecipeSAP(file1 string, file2 string, npm string) {
	linesdatpatternData, err := os.Open(file1)
	defer linesdatpatternData.Close()
	if err != nil {
		panic(err)
	}

	cr := csv.NewReader(linesdatpatternData)
	cr.LazyQuotes = true
	cr.Comma = ' '
	record, err := cr.ReadAll()
	if err != nil {
		fmt.Println(err) //(err)
	}

	linesdatpatrsData, err := os.Open(file2)
	defer linesdatpatternData.Close()
	if err != nil {
		panic(err)
	}

	cr2 := csv.NewReader(linesdatpatrsData)
	cr2.LazyQuotes = true
	cr2.Comma = ' '
	record2, err := cr2.ReadAll()
	if err != nil {
		fmt.Println(err) //(err)
	}

	//csvdatafile, err := os.Create("/home/a20272/Code/github.com/eugenefoxx/readDGS/files/Рецепт для SAP.csv")
	//csvdatafile, err := os.Create(setting.Conf.Files + "Рецепт для SAP.csv")
	csvdatafile, err := os.Create(setting.Conf.Files + "recipte.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvdatafile.Close()

	writer := csv.NewWriter(csvdatafile)
	writer.Write([]string{"SAP" + " " + npm, "Количество"})
	if err != nil {
		fmt.Println(err)
	}
	writer.Comma = ','
	writer.Flush()

	//split, err := os.OpenFile("/home/a20272/Code/github.com/eugenefoxx/readDGS/files/Рецепт для SAP.csv", os.O_APPEND|os.O_WRONLY, 0644)
	//split, err := os.OpenFile(setting.Conf.Files+"Рецепт для SAP.csv", os.O_APPEND|os.O_WRONLY, 0644)
	split, err := os.OpenFile(setting.Conf.Files+"recipte.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer split.Close()

	for _, linedatpatternData := range record {
		datapatternData := outdatpatternData{
			IDNUM: linedatpatternData[0],
			Qty:   linedatpatternData[1],
		}

		for _, linedatpatrsData := range record2 {
			data := datpatrsData{
				IDNUM:       linedatpatrsData[0],
				NAME:        linedatpatrsData[1],
				EX:          linedatpatrsData[2],
				LNAME:       linedatpatrsData[3],
				FA:          linedatpatrsData[4],
				FB:          linedatpatrsData[5],
				FC:          linedatpatrsData[6],
				FD:          linedatpatrsData[7],
				FE:          linedatpatrsData[8],
				FF:          linedatpatrsData[9],
				FG:          linedatpatrsData[10],
				FH:          linedatpatrsData[11],
				FI:          linedatpatrsData[12],
				FJ:          linedatpatrsData[13],
				REELS:       linedatpatrsData[14],
				CHIP:        linedatpatrsData[15],
				Shape:       linedatpatrsData[16],
				MCOND:       linedatpatrsData[17],
				DCOND:       linedatpatrsData[18],
				DCOND2:      linedatpatrsData[19],
				ICOND:       linedatpatrsData[20],
				SKIP:        linedatpatrsData[21],
				A:           linedatpatrsData[22],
				PACK:        linedatpatrsData[23],
				PACKB:       linedatpatrsData[24],
				SSIZE:       linedatpatrsData[25],
				RETRY:       linedatpatrsData[26],
				ESTOP:       linedatpatrsData[27],
				ThrowAway:   linedatpatrsData[28],
				LandA:       linedatpatrsData[29],
				LandB:       linedatpatrsData[30],
				LandC:       linedatpatrsData[31],
				LandD:       linedatpatrsData[32],
				PreSend:     linedatpatrsData[33],
				NoArrange:   linedatpatrsData[34],
				Proved:      linedatpatrsData[35],
				UsePeriod:   linedatpatrsData[36],
				C:           linedatpatrsData[37],
				CoverSize:   linedatpatrsData[38],
				Teach:       linedatpatrsData[39],
				C2:          linedatpatrsData[40],
				RETRYCOUNT:  linedatpatrsData[41],
				ESTOPCOUNT:  linedatpatrsData[42],
				RETRYCOUNT2: linedatpatrsData[43],
				SICOND:      linedatpatrsData[44],
				HRMCOND:     linedatpatrsData[45],
				HPICOND:     linedatpatrsData[46],
				SL:          linedatpatrsData[47],
				InCOND:      linedatpatrsData[48],
				FeederEX:    linedatpatrsData[49],
				MSD:         linedatpatrsData[50],
				NPI:         linedatpatrsData[51],
				LCR:         linedatpatrsData[52],
				REELTYPECHG: linedatpatrsData[53],
				FEEDLENCHG:  linedatpatrsData[54],
			}

			if datapatternData.IDNUM == data.IDNUM {
				//	fmt.Println(data.NAME + " " + datapatternData.Qty)
				result := []string{data.NAME + "," + datapatternData.Qty}
				for _, v := range result {
					_, err = fmt.Fprintln(split, v)
					if err != nil {

						fmt.Println(err)
						split.Close()
						return
					}
				}
			}

		}

	}

}
