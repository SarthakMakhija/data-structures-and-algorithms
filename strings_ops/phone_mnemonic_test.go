package stringsops_test

import (
	stringsops "github.com/SarthakMakhija/data-structures-and-algorithms/strings_ops"
	"reflect"
	"testing"
)

func TestPhoneMnemonicFor1(t *testing.T) {
	phoneNumber := "23"
	allMnemonics := stringsops.PhoneMnemonicFor(phoneNumber)

	expectedMnemonics := []string{"AD", "AE", "AF", "BD", "BE", "BF", "CD", "CE", "CF"}

	if !reflect.DeepEqual(allMnemonics, expectedMnemonics) {
		t.Fatalf("Expected %v, received %v", expectedMnemonics, allMnemonics)
	}
}

func TestPhoneMnemonicFor2(t *testing.T) {
	phoneNumber := "27"
	allMnemonics := stringsops.PhoneMnemonicFor(phoneNumber)

	expectedMnemonics := []string{"AP", "AQ", "AR", "AS", "BP", "BQ", "BR", "BS", "CP", "CQ", "CR", "CS"}

	if !reflect.DeepEqual(allMnemonics, expectedMnemonics) {
		t.Fatalf("Expected %v, received %v", expectedMnemonics, allMnemonics)
	}
}

func TestPhoneMnemonicFor3(t *testing.T) {
	phoneNumber := "1*"
	allMnemonics := stringsops.PhoneMnemonicFor(phoneNumber)

	var expectedMnemonics []string

	if !reflect.DeepEqual(allMnemonics, expectedMnemonics) {
		t.Fatalf("Expected %v, received %v", expectedMnemonics, allMnemonics)
	}
}

func TestPhoneMnemonicFor4(t *testing.T) {
	phoneNumber := "237"
	allMnemonics := stringsops.PhoneMnemonicFor(phoneNumber)

	expectedMnemonics := []string{"ADP", "ADQ", "ADR", "ADS", "AEP", "AEQ", "AER", "AES", "AFP", "AFQ", "AFR", "AFS",
		"BDP", "BDQ", "BDR", "BDS", "BEP", "BEQ", "BER", "BES", "BFP", "BFQ", "BFR", "BFS",
		"CDP", "CDQ", "CDR", "CDS", "CEP", "CEQ", "CER", "CES", "CFP", "CFQ", "CFR", "CFS"}

	if !reflect.DeepEqual(allMnemonics, expectedMnemonics) {
		t.Fatalf("Expected %v, received %v", expectedMnemonics, allMnemonics)
	}
}
