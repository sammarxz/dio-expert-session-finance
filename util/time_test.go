package util

import (
	"testing"
)

func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2019-02-03 02:09")

	if convertedTime.Year() != 2019 {
		t.Errorf("Espera que o ano seja 2019")
	}

	if convertedTime.Month() != 02 {
		t.Errorf("Espera que o mês seja 02")
	}

	if convertedTime.Day() != 03 {
		t.Errorf("Espera que o dia seja 03")
	}
}
