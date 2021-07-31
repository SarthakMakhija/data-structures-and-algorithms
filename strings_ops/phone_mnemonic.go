package stringsops

func PhoneMnemonicFor(phoneNumber string) []string {

	if phoneNumber == "" {
		return []string{}
	}
	mnemonicMapping := map[string]string{
		"1": "",
		"2": "ABC",
		"3": "DEF",
		"4": "GHI",
		"5": "JKL",
		"6": "MNO",
		"7": "PQRS",
		"8": "TUV",
		"9": "WXYZ",
		"0": "",
		"*": "",
		"#": "",
	}
	var mnemonicCharacterIndex = make([]int, len(phoneNumber))
	var allMnemonics []string

	var reset func(mnemonicCharacterPosition int)
	reset = func(mnemonicCharacterPosition int) {
		previousPosition := mnemonicCharacterPosition - 1

		if previousPosition >= 0 {
			mnemonicCharacterIndex[mnemonicCharacterPosition] = 0
			mnemonicCharacterIndex[previousPosition] = mnemonicCharacterIndex[previousPosition] + 1

			if mnemonicCharacterIndex[previousPosition] >= len(mnemonicMapping[string(phoneNumber[previousPosition])]) {
				reset(mnemonicCharacterPosition - 1)
			}
		} else {
			return
		}
	}

	var phoneMnemonicInner func(mnemonicCharacterPosition int)
	phoneMnemonicInner = func(mnemonicCharacterPosition int) {
		if mnemonicCharacterIndex[mnemonicCharacterPosition] >= len(mnemonicMapping[string(phoneNumber[mnemonicCharacterPosition])]) {
			reset(mnemonicCharacterPosition)
		}
		if mnemonicCharacterIndex[0] >= len(mnemonicMapping[string(phoneNumber[0])]) {
			return
		} else {
			mnemonic := ""
			for index := 0; index < len(phoneNumber); index++ {
				digit := phoneNumber[index]
				s := mnemonicMapping[string(digit)]
				mnemonic = mnemonic + string(s[(mnemonicCharacterIndex)[index]])
			}
			allMnemonics = append(allMnemonics, mnemonic)
			mnemonicCharacterIndex[mnemonicCharacterPosition] = mnemonicCharacterIndex[mnemonicCharacterPosition] + 1
			phoneMnemonicInner(mnemonicCharacterPosition)
		}
	}
	phoneMnemonicInner(len(mnemonicCharacterIndex) - 1)
	return allMnemonics
}
