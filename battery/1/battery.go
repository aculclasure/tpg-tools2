package battery

import (
	"fmt"
	"regexp"
	"strconv"
)

type Status struct {
	ChargePercent int
}

var acpiOutput = regexp.MustCompile(`([0-9]+)%`)

func ParseAcpiOutput(text string) (Status, error) {
	matches := acpiOutput.FindStringSubmatch(text)
	if len(matches) < 2 {
		return Status{}, fmt.Errorf("failed to parse acpi output: %q", text)
	}
	charge, err := strconv.Atoi(matches[1])
	if err != nil {
		return Status{}, fmt.Errorf("failed to parse charge percentage %q", matches[1])
	}
	return Status{
		ChargePercent: charge,
	}, nil
}
