// Package ledger provides double ledger accounting.
package ledger

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Entry represents one entry in the ledger.
type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

// Ledger is a series of entries.
type Ledger []Entry

var localeConfig = map[string]struct {
	header                 []string
	dateFormat             string
	thousandsSeparator     string
	positiveCurrencyFormat string
	negativeCurrencyFormat string
}{
	"nl-NL": {
		header:                 []string{"Datum", "Omschrijving", "Verandering"},
		dateFormat:             "02-01-2006",
		thousandsSeparator:     ".",
		positiveCurrencyFormat: "%s %s,%02s ",
		negativeCurrencyFormat: "%s %s,%02s-",
	},
	"en-US": {
		header:                 []string{"Date", "Description", "Change"},
		dateFormat:             "01/02/2006",
		thousandsSeparator:     ",",
		positiveCurrencyFormat: "%s%s.%s ",
		negativeCurrencyFormat: "(%s%s.%s)",
	},
}

var currencies = map[string]string{
	"EUR": "€",
	"USD": "$",
}

// Less orders the ledger by date > description > change
func (a Ledger) Less(i, j int) bool {
	if a[i].Date != a[j].Date {
		return a[i].Date < a[j].Date
	}
	if a[i].Description != a[j].Description {
		return a[i].Description < a[j].Description
	}
	return a[i].Change < a[j].Change
}

type chanMsg struct {
	i int
	s string
	e error
}

func formatCurrency(locale, currency string, cents int) string {
	format := localeConfig[locale].positiveCurrencyFormat
	if cents < 0 {
		cents = -cents
		format = localeConfig[locale].negativeCurrencyFormat
	}
	var parts []string
	dollars := strconv.Itoa(cents / 100)
	centsStr := fmt.Sprintf("%02d", cents%100)

	o := len(dollars) % 3
	if o == 0 {
		o = 3
	}

	for i := 0; i < len(dollars); i += o {
		if i != 0 {
			o = 3
		}
		parts = append(parts, dollars[i:i+o])
	}

	return fmt.Sprintf(format, currencies[currency], strings.Join(parts, localeConfig[locale].thousandsSeparator), centsStr)
}

func truncate(s, cont string, l int) string {
	if len(s) > l {
		return s[:l-len(cont)] + cont
	}
	return s
}

func formatEntry(e Entry, locale, currency string) (string, error) {
	l := localeConfig[locale]
	t, err := time.Parse("2006-01-02", e.Date)
	if err != nil {
		return "", err
	}
	// Truncate long descriptions.
	description := truncate(e.Description, "...", 25)
	date := t.Format(l.dateFormat)
	amount := formatCurrency(locale, currency, e.Change)
	return fmt.Sprintf("%10s | %-25s | %13s\n", date, description, amount), nil
}

// FormatLedger formats a set of ledger entries.
func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	// Validation.
	if _, ok := currencies[currency]; !ok {
		return "", fmt.Errorf("invalid currency: %s", currency)
	}

	l, ok := localeConfig[locale]
	if !ok {
		return "", fmt.Errorf("unknown locale: %s", locale)
	}

	// Make a copy and sort it.
	var entriesCopy Ledger
	for _, e := range entries {
		entriesCopy = append(entriesCopy, e)
	}
	sort.Slice(entriesCopy, entriesCopy.Less)

	ss := make([]string, len(entriesCopy))
	for i, et := range entriesCopy {
		s, err := formatEntry(et, locale, currency)
		if err != nil {
			return "", err
		}
		ss[i] = s
	}

	header := fmt.Sprintf("%-10s | %-25s | %s\n", l.header[0], l.header[1], l.header[2])
	return header + strings.Join(ss, ""), nil
}
