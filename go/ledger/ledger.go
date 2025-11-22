package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if err := validateInputs(currency, locale); err != nil {
		return "", err
	}

	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)

	sort.SliceStable(entriesCopy, func(i, j int) bool {
		if entriesCopy[i].Date != entriesCopy[j].Date {
			return entriesCopy[i].Date < entriesCopy[j].Date
		}
		if entriesCopy[i].Description != entriesCopy[j].Description {
			return entriesCopy[i].Description < entriesCopy[j].Description
		}
		return entriesCopy[i].Change < entriesCopy[j].Change
	})

	var result strings.Builder
	result.WriteString(header(locale))

	for _, entry := range entriesCopy {
		line, err := formatEntry(entry, currency, locale)
		if err != nil {
			return "", err
		}
		result.WriteString(line)
	}

	return result.String(), nil
}

func validateInputs(currency string, locale string) error {
	if currency != "USD" && currency != "EUR" {
		return errors.New("invalid currency")
	}
	if locale != "en-US" && locale != "nl-NL" {
		return errors.New("invalid locale")
	}
	return nil
}

func header(locale string) string {
	if locale == "nl-NL" {
		return "Datum      | Omschrijving              | Verandering\n"
	}
	return "Date       | Description               | Change\n"
}

func formatEntry(entry Entry, currency string, locale string) (string, error) {
	if len(entry.Date) != 10 || entry.Date[4] != '-' || entry.Date[7] != '-' {
		return "", errors.New("invalid date")
	}

	date := formatDate(entry.Date, locale)
	desc := formatDescription(entry.Description)
	amount := formatAmount(entry.Change, currency, locale)

	return date + " | " + desc + " | " + amount + "\n", nil
}

func formatDate(date string, locale string) string {
	year, month, day := date[0:4], date[5:7], date[8:10]

	var formatted string
	if locale == "nl-NL" {
		formatted = day + "-" + month + "-" + year
	} else {
		formatted = month + "/" + day + "/" + year
	}

	return formatted + strings.Repeat(" ", 10-len(formatted))
}

func formatDescription(desc string) string {
	if len(desc) > 25 {
		return desc[:22] + "..."
	}
	return desc + strings.Repeat(" ", 25-len(desc))
}

func formatAmount(cents int, currency string, locale string) string {
	negative := cents < 0
	if negative {
		cents = -cents
	}

	symbol := "$"
	if currency == "EUR" {
		symbol = "€"
	}

	var amount string
	if locale == "nl-NL" {
		amount = fmt.Sprintf("%s %s,%02d", symbol, formatThousands(cents/100, "."), cents%100)
		if negative {
			amount += "-"
		} else {
			amount += " "
		}
	} else {
		if negative {
			amount = fmt.Sprintf("(%s%s.%02d)", symbol, formatThousands(cents/100, ","), cents%100)
		} else {
			amount = fmt.Sprintf("%s%s.%02d ", symbol, formatThousands(cents/100, ","), cents%100)
		}
	}

	return strings.Repeat(" ", 13-len([]rune(amount))) + amount
}

func formatThousands(n int, sep string) string {
	s := fmt.Sprintf("%d", n)

	var parts []string
	for len(s) > 3 {
		parts = append([]string{s[len(s)-3:]}, parts...)
		s = s[:len(s)-3]
	}
	parts = append([]string{s}, parts...)

	return strings.Join(parts, sep)
}