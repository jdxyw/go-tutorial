package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type sessionRecord struct {
	Idx          int
	Type         string
	Module       string
	CorrectRatio float64
}

func main() {
	t := time.Unix(1558026985, 0)
	fmt.Println(t)
	fmt.Println(t.Format("2006-01-02"))

	f2, _ := os.Open("/Users/yongweixing/dev/darwin_seq/darwin_order.tsv")
	defer f2.Close()
	orderScanner := bufio.NewScanner(f2)
	orderMap := make(map[string]map[string]bool)
	for orderScanner.Scan() {
		fields := strings.Split(orderScanner.Text(), "\t")
		uid := fields[0]
		dt := fields[1]

		if _, ok := orderMap[uid]; !ok {
			orderMap[uid] = make(map[string]bool)
		}
		orderMap[uid][dt] = bool
	}

	userData := make(map[string]map[string][]sessionRecord)
	loc := time.FixedZone("UTC+3", int((3 * time.Hour).Seconds()))
	f, _ := os.Open("/Users/yongweixing/dev/darwin_seq/darwin_session_0517_0617.tsv")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), "\t")
		uid := fields[0]
		sessTS, _ := strconv.ParseInt(fields[2], 10, 64)
		consumTS, _ := strconv.ParseInt(fields[11], 10, 64)
		idx, _ := strconv.Atoi(fields[8])
		sessType := fields[9]
		correctRatio, _ := strconv.ParseFloat(fields[12], 64)
		sessModule := fields[13]
		//dateStr := fields[len(fields)-1]
		sessTime := time.Unix(int64(sessTS/1000000), 0).In(loc)
		consumTime := time.Unix(int64(consumTS/1000000), 0).In(loc)
		//t := time.Unix(int64(sessTS/1000000), 0)
		if sessTime.Format("2006-01-02") != consumTime.Format("2006-01-02") {
			continue
		}
		date := sessTime.Format("2006-01-02")
		sessRecord := sessionRecord{
			Idx:          idx,
			Type:         sessType,
			Module:       sessModule,
			CorrectRatio: correctRatio,
		}

		if _, ok := userData[uid]; !ok {
			userData[uid] = make(map[string][]sessionRecord)
		}
		userData[uid][date] = append(userData[uid][date], sessRecord)
	}

	userData2 := make(map[string][][]sessionRecord)
	firstDate := Date(2019, 5, 17)
	for uid, uData := range userData {
		userData2[uid] = make([][]sessionRecord, 32)
		for d, records := range uData {
			days := int(parseToDate(d).Sub(firstDate).Hours() / 24)
			if days < 0 {
				continue
			}
			if days > 31 {
				continue
			}
			userData2[uid][days] = append(userData2[uid][days], records...)
		}
	}

	typeCounterByPos := make(map[string]map[int]int)

	for _, udata := range userData2 {
		for day, records := range udata {
			if len(records) == 0 {
				continue
			}

			if day+3 > len(udata)-1 {
				continue
			}

			if len(udata[day+1]) == 0 && len(udata[day+2]) == 0 && len(udata[day+3]) == 0 {
				continue
			}

			for i, r := range records {
				if _, ok := typeCounterByPos[r.Type]; !ok {
					typeCounterByPos[r.Type] = make(map[int]int)
				}
				typeCounterByPos[r.Type][i]++
			}
		}
	}

	for t, values := range typeCounterByPos {
		for pos, c := range values {
			if pos > 7 {
				continue
			}
			fmt.Printf("type %v  pos %v counter %v\n", t, pos, c)
		}
	}

	// gap := 14
	// for uid, uData := range userData2 {
	// 	for start, records := range uData {
	// 		end := start + 14
	// 		if end > len(records)-5 {
	// 			continue
	// 		}
	// 	}
	// }
	fmt.Println(len(userData))
}

// func getTypeNdaysNumber(srs [][]sessionRecord, n int) map[string]int {
// 	tn := make(map[string]int)
// 	if n < 0 || n > len(srs) {
// 		n = len(srs)
// 	}
// 	for _, ss := range srs[len(srs)-n:] {
// 		for _, s := range ss {
// 			tn[s.Type]++
// 		}
// 	}
// 	return tn
// }

// func getModuleNdaysNumber(srs [][]sessionRecord, n int) map[string]int {
// 	mn := make(map[string]int)
// 	if n < 0 || n > len(srs) {
// 		n = len(srs)
// 	}
// 	for _, ss := range srs[len(srs)-n:] {
// 		for _, s := range ss {
// 			mn[s.Module]++
// 		}
// 	}
// 	return mn
// }

// func defaultSlotNumberByNdaysAndRatio(srs [][]sessionRecord, n int) (int, float64) {
// 	res := 0
// 	total := 0
// 	ratio := 0.0
// 	if n < 0 || n > len(srs) {
// 		n = len(srs)
// 	}
// 	for _, ss := range srs[len(srs)-n:] {
// 		total += len(ss)
// 		for _, s := range ss {
// 			if s.Idx == 0 {
// 				res += 1
// 			}
// 		}
// 	}
// 	if total > 0 {
// 		ratio = float64(res * 1.0 / total)
// 	}
// 	return res, ratio
// }

// func getAverageCorrectRatioLatesNDaysByType(srs [][]sessionRecord, n int) map[string]float64 {
// 	corr := make(map[string]float64)
// 	nums := make(map[string]int)
// 	totalRa := make(map[string]float64)

// 	if n < 0 || n > len(srs) {
// 		n = len(srs)
// 	}
// 	for _, ss := range srs[len(srs)-n:] {
// 		for _, s := range ss {
// 			nums[s.Type] += 1
// 			totalRa[s.Type] += s.CorrectRatio
// 		}
// 	}
// 	for k, v := range nums {
// 		corr[k] = totalRa[k] / float64(v)
// 	}
// 	return corr
// }

func parseToDate(s string) time.Time {
	fields := strings.Split(s, "-")
	year, _ := strconv.Atoi(fields[0])
	month, _ := strconv.Atoi(fields[1])
	day, _ := strconv.Atoi(fields[2])
	return Date(year, month, day)
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
