package main

import (
	"errors"
	"fmt"
)

type RowData struct {
	RankWebsite int
	Domain      string
	TLD         string
	IDN_TLD     string
	Valid       bool
	RefIPs      int
}

func GetTLD(domain string) (TLD string, IDN_TLD string) {
	var ListIDN_TLD = map[string]string{
		".com": ".co.id",
		".org": ".org.id",
		".gov": ".go.id",
	}

	for i := len(domain) - 1; i >= 0; i-- {
		if domain[i] == '.' {
			TLD = domain[i:]
			break
		}
	}

	if _, ok := ListIDN_TLD[TLD]; ok {
		return TLD, ListIDN_TLD[TLD]
	} else {
		return TLD, TLD
	}
}

func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
	TLD, IDN_TLD := GetTLD(website.Domain)
	if website.Domain == "" {
		chErr <- errors.New("domain name is empty")
		return
	}
	if website.Valid == false {
		chErr <- errors.New("domain not valid")
		return
	}
	if website.RefIPs < 0 {
		chErr <- errors.New("domain RefIPs not valid")
		return
	}
	website.TLD = TLD
	website.IDN_TLD = IDN_TLD
	ch <- website
	chErr <- nil
}

// Gunakan variable ini sebagai goroutine di fungsi FilterAndGetDomain
var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	ch := make(chan RowData, len(data))
	errCh := make(chan error)
	d := make([]RowData, 0)
	for _, website := range data {
		go FuncProcessGetTLD(website, ch, errCh)
	}
	for i := 0; i < len(data); i++ {
		er := <-errCh
		if er != nil {
			return nil, er
		}
		da := <-ch
		if da.TLD == TLD {
			d = append(d, da)
		}
	}
	return d, nil
}

// gunakan untuk melakukan debugging
func main() {
	rows, err := FilterAndFillData(".com", []RowData{
		{1, "google.com", "", "", true, 100},
		{2, "facebook.com", "", "", true, 100},
		{3, "golang.org", "", "", true, 100},
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rows)
}
