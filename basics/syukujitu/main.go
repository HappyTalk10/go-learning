// save as extract_holidays_2026.go
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	url := "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "download error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// 内閣府CSVはShift_JIS（MS932）で来ることがあるためUTF-8へ変換
	utf8reader := transform.NewReader(resp.Body, japanese.ShiftJIS.NewDecoder())

	r := csv.NewReader(utf8reader)
	// 行の長さが可変の可能性を考えて FieldsPerRecord = -1
	r.FieldsPerRecord = -1

	outFile, err := os.Create("holidays_2026.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "create file error: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close()
	w := csv.NewWriter(outFile)
	defer w.Flush()

	// ヘッダ行を探してDate列を特定してもよいが、単純化のため
	// 各行の先頭フィールドを日付と見なしてパースする実装にしている
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			// 一部ファイルでは空行や注記があるためスキップ
			continue
		}
		if len(rec) < 2 {
			continue
		}

		// 先頭フィールドを日付としてパースする試み
		dateStr := rec[0] // 例: "2026/1/1" など
		// 受け入れる可能性のあるフォーマットの候補
		formats := []string{"2006/1/2", "2006/01/02", "2006-01-02", "2006/1/2 0:00:00"}
		var dt time.Time
		var perr error
		for _, f := range formats {
			dt, perr = time.Parse(f, dateStr)
			if perr == nil {
				break
			}
		}
		if perr != nil {
			// ヘッダ行や別フォーマットならスキップ（必要ならヘッダ出力）
			// もしヘッダ行を出力したければ条件を追加して出力する
			continue
		}

		if dt.Year() == 2026 {
			// 当該行を出力（そのまま書き出す）
			if err := w.Write(rec); err != nil {
				fmt.Fprintf(os.Stderr, "write error: %v\n", err)
			}
		}
	}

	fmt.Println("finished: holidays_2026.csv")
}
