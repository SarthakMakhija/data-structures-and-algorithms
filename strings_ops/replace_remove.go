package stringsops

import "unicode"

//ReplaceAndRemove
//Assume characters has enough size
func ReplaceAndRemove(characters []rune) {
	//replace 'a' by 2 'd's
	//delete 'b'
	var backup []rune
	for _, ch := range characters {
		if ch == 'a' {
			backup = append(backup, []rune{'d', 'd'}...)
		} else if ch != 'b' && unicode.IsPrint(ch) {
			backup = append(backup, ch)
		}
	}
	for index, ch := range backup {
		characters[index] = ch
	}
}
