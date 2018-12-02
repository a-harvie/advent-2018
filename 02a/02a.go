package main

import "fmt"

var tests = []string{
	"abcde",
	"fghij",
	"klmno",
	"pqrst",
	"fguij",
	"axcye",
	"wvxyz",
}
var testResults = "fgij"

var input = []string{"asgwdcmbrkerohqoutfylvzpnx", "asgwjcmbrkejihqoutfylvipne", "asgwjcmbrkedihqoutvylizpnz", "azgsjcmbrkedihqouifylvzpnx", "asgwucmbrktddhqoutfylvzpnx", "asgwocmbrkedihqoutfylvzivx", "aqgwjcmbrkevihqvutfylvzpnx", "tsgljcmbrkedihqourfylvzpnx", "asgpjcmbrkedihqoutfnlvzsnx", "astwjcmbrktdihqrutfylvzpnx", "asgwjcmbrpedhhqoutfylvzynx", "xsgwjcmbrkedieqowtfylvzpnx", "asgwjcmbvkedihfoutnylvzpnx", "asgwjcmtrkedihqouafylvzcnx", "asgwjcmbrkedihqoutfylvxpvm", "usgwjcmbrkedihqortfyuvzpnx", "asgwjcmbrwedihqoutfylizpix", "asgrjcvbrkedixqoutfylvzpnx", "asgwjcmbrogdihqoutfelvzpnx", "aggwjcmbrkesihqoutoylvzpnx", "asgtjccbrkedihqoutfrlvzpnx", "asgcucmbrbedihqoutfylvzpnx", "esgwjcmbrkedihqsutfylvzcnx", "asgwjcmbrkedrhqoutfyobzpnx", "mngwjcbbrkedihqoutfylvzpnx", "asgwjcrbrkeoihqyutfylvzpnx", "apgwjcmbrkednhqogtfylvzpnx", "asgwjcwbrkedihqoutfylplpnx", "asgwjcmbrkfdihqoutfxlvzpyx", "aegwjcmbrkedihqoutfylbxpnx", "asgljcmbrkedixqoutaylvzpnx", "aigwjcmbrkedihqouhfylvzpex", "asgwjbmbrkedihqoutfylfzpnp", "asgwjcmzroedihqoutcylvzinx", "asgwjcwbrieuihqoutfylvzpnx", "aagwjcmbrkedjhqdutfylvzpnx", "ahgwjcmbrkedihqsutfylvzpfx", "asgwjcmbrkedihzosttylvzpnx", "aegwjcmbrkedioqnutfylvzpnx", "asgwjcmbykidihqoutfysvzpnx", "asgwkcxbrkeddhqoutfylvzpnx", "ashwjcmbrkeeihqoutfylvzknx", "acgwjcmbrqedihqoqtfylvzpnx", "asgwjcmtrkedihooutfylszpnx", "asgwjcmbrkmdihqfutrylvzpnx", "asgwjcmbrkedihqoutjylvapnn", "asgwjcmbwkedihqoutkylkzpnx", "asgwjrmbrkedihqoutfycnzpnx", "asgwtcmbrkedihqoqtfylozpnx", "asgajcmbrkedihqoutuylvzpny", "asgwjcmbykedihqoutfylfzpwx", "asgwjcsbrkedihpoutfylvvpnx", "hsdwjcmbrvedihqoutfylvzpnx", "asgwjcmbrkedihqoutfdmszpnx", "adgwjcmbrtidihqoutfylvzpnx", "augwjcmbriedihqoutgylvzpnx", "asgwjvmbreedihqoutfllvzpnx", "asgwjcnbfkedihqoltfylvzpnx", "asgwjcmbykddihqoutqylvzpnx", "ajgwjcmbrkedihqoutfylvpvnx", "asgwjcmbrkydihqoutfylszpnl", "xsgwjcmbrkqdihqoutfylvkpnx", "asgwjcmbrkedimqoutfklvzknx", "csgwjbmbrkedihqoftfylvzpnx", "asgwjcmbjkedihjoutfylvzpnn", "asgwjcmprkedihqoulfalvzpnx", "asgwjcmbrvediqqoutfyuvzpnx", "asgwjambrkedhhqoutkylvzpnx", "asgejcmbrkidihqoutfylvzpnk", "hsiwjcmbrkedihqoutfylvzpnq", "asswjczbrkedihqoutfylczpnx", "asgwjnmbrkedyhzoutfylvzpnx", "asgwscmbrkedihqoutfklvlpnx", "asgwlcmbrktdihqoutfylvzpax", "asfwjcmerkedihqoutfylvipnx", "asgwjcmbrkeditqoeafylvzpnx", "asgwgcmbrkesihqoutfylyzpnx", "fsgwjcmbrkedihqouvfyavzpnx", "asgwjcmbrpedwhqoutfylmzpnx", "asgwjcmbrkzdzhqoucfylvzpnx", "asgwjcmnrketmhqoutfylvzpnx", "asgwjcmbrkedihxoutsylvzpnh", "asgwjcobrkedihqoutfrlvzpox", "asgwjcmbrkedihqootfylxzpox", "asgjjcmcrkedihqoutfylmzpnx", "lsgwjcmbrkedihqoutfyqvzunx", "asgwjcmbrwedihqoutoylvzpnu", "aszwjcmbtkedihqoutfylczpnx", "asgwjcmbykedihqoutfylvgpex", "asgijcmbrkedilqoutkylvzpnx", "astwxcmzrkedihqoutfylvzpnx", "akgwjcmbnkedihqfutfylvzpnx", "asgwjcmbrqndivqoutfylvzpnx", "asgwjrmbrleqihqoutfylvzpnx", "asgwjcmbrkevihqoutfxlvzpvx", "asbwjcmbrkedihqoutfelvwpnx", "asewjcbbrkmdihqoutfylvzpnx", "asgwjcmbrkeaihxoutfylpzpnx", "asgwjzmbrkedihqrotfylvzpnx", "asgwjcmbrkedihqoutgdxvzpnx", "asgwjcwbrkmdihqoutfylvzlnx", "asgwjcmbrkegihqoutfylrzpax", "ajgwjcmbrkegihqhutfylvzpnx", "asgwjcmbrzedihqhutfylvkpnx", "asgwjcmwrkedihqouhfylkzpnx", "aygwjcmbrkedihqoutfdlvzpnr", "asgwjcmbrkednhqoutiylvypnx", "aqgwjcmbrkezihqoutfylvzonx", "bsgwjcmbrkedihqouhfylvzsnx", "asgwjcmcrkedihqokyfylvzpnx", "asgsjcmbrkewiyqoutfylvzpnx", "asgwpcmbrkejihqoutfylzzpnx", "asgwjumbrkedbeqoutfylvzpnx", "asgwjcmbrkedihpoutqylqzpnx", "awgwjcmbrredihqoutfylvzpna", "asgwjsmbraedihqoutfylvzpvx", "asgwncmbrkedihqoutfyljzrnx", "asgwncmbrkedihqohtfylvzonx", "asgwjcmbrkedihqlutfylvypux", "asgwjcmbbkedihooutfylkzpnx", "asghjcmsryedihqoutfylvzpnx", "asgwjcmbrkevihqoulfzlvzpnx", "asggjcmbrkedizqoutfylvzknx", "asbwjcmbriedihqoutfylvmpnx", "asgwjcmbrkedqbqoutfylvzenx", "asgwjcmprkedihqoutfylvzknp", "asgwjcmbrkerihqoutfwlvzpno", "asgwjcmvrkesihqoutrylvzpnx", "asgzjcmbrkedihqoutfnlvbpnx", "asfwjcmbrkhdihqoutfylpzpnx", "asgwjcmbskedihqdutfyyvzpnx", "asgwjcmzrkedihqoutcylvzinx", "asgwjcmbrkedibqoutfylvjonx", "asgwjcmbrbedihqoutfylmzbnx", "asgwjcmbrkedhhqoutmylczpnx", "asgwjcmbrkbgihqoutzylvzpnx", "asgwjcfbrkedihqoupfyxvzpnx", "asiwjcmbzkedihqoutfyluzpnx", "asvwjcmbrkedihqoitfylvzpns", "asgwjcmxikedihqoutfyevzpnx", "asgwjcmbrkedioqoutfylvzwox", "asgwjcmbrkedivqoutjyuvzpnx", "asgwjcmbkkydihqrutfylvzpnx", "asgwjcmbrkxdiuqoutfylvopnx", "asgwjcmbrkedihqouthylvzpra", "asgwjcmbrzedimloutfylvzpnx", "asgwjcmbrkedmhqoulfytvzpnx", "asgwjcmbrkzdihqrutfysvzpnx", "ssgwjcmxrkedihqoutftlvzpnx", "asgwjcmbrkedihqoutfajvzynx", "asgwjcmbrkqdihqxuufylvzpnx", "asmwjcabrkedihqouxfylvzpnx", "asgwjcmbrkeeihqoatfycvzpnx", "asgwjcjbrgedjhqoutfylvzpnx", "asgljcmtrkedihqoutoylvzpnx", "asgwjcmbrkedigqouzfylvzpvx", "ajgvjcmbkkedihqoutfylvzpnx", "asgwjcmbrkedihqtugfygvzpnx", "asgbjcmbrkedihboftfylvzpnx", "asgwjwmbrkedihqontfylhzpnx", "asgwjfmhrkedihqoutfylvqpnx", "asgwjxmbrkedihqoutzylvzpnj", "asgwjcrlrkedihqoutfylvzpsx", "aygwjcmbrkedihqoutsylvzdnx", "zsgwjcmbrkedihjogtfylvzpnx", "asgwjxmbrkegihqoutfylvopnx", "asgwjcmbrkedihqhutfylvzcnr", "asgwicmbrkewihvoutfylvzpnx", "asqwjcmbvkedihqoutfylvzknx", "asgwjcmbrkedihqoktfyevzpnu", "asgwjcmbrkudihqoutfylqzznx", "asgwjdmbrkedihqoutfylvvdnx", "asgwjcmbrkwmihqautfylvzpnx", "asgwjcmbrxedihqoctfyldzpnx", "asgwjdmbrkedjhqoutfyfvzpnx", "asgwjcmtrzedihqoutfylvzpnm", "bpgwjcmbrmedihqoutfylvzpnx", "asgwjctbrkedihqoqtfynvzpnx", "askhjcmbrkedihqoutfylvzrnx", "asgkjcmblkehihqoutfylvzpnx", "asgwjjmbrkedvhqoutfhlvzpnx", "asgwjcmbrkedihqoupzylvzknx", "asgwjcmbukedchqoutfylizpnx", "askwjcmdrkedihqoutwylvzpnx", "asgwjcmbtkcdihloutfylvzpnx", "asgwjcmbrkedwgqoutvylvzpnx", "asmwjcmbrkedihqoutfylozpnc", "asgwjcmbriedibqouofylvzpnx", "asgnjcmcrkedihqoupfylvzpnx", "asgzjcmbrksdihqoutiylvzpnx", "asgwjcrbkkedihqouafylvzpnx", "asgwjcmbkkvdihqqutfylvzpnx", "astwjcqbrkedihqoutfylvzpvx", "asgwjcmhrkehihqoutfylvzvnx", "asgwjcmbraeduhqoutmylvzpnx", "asgwjcmbrkedihquutnylvzptx", "asgwjcmbrkedilqoftfylvzpnz", "akgwjmmbrkedihqoutfylxzpnx", "asgwjcmbrkhxikqoutfylvzpnx", "asgcjcmetkedihqoutfylvzpnx", "fsgwjcmsrkedihooutfylvzpnx", "gsgwjcmbrkedihdoutfylvzdnx", "asgwtccbrkedihqoutfylvwpnx", "asuwjcmbrkedihqcutfylvzpox", "asgwacmbrkodihqlutfylvzpnx", "asgwjcmbrkediuqoutfylqhpnx", "asgwjcmbrkwdrhqoutfylvzpno", "angwjcsblkedihqoutfylvzpnx", "aigwjcmbyoedihqoutfylvzpnx", "adgwjcmbrkedihqtutfylyzpnx", "asgwjzmbrkeeihqputfylvzpnx", "asgwjcmbrkwdihqoutfylvzpwc", "asgpjcmbrkgdihqbutfylvzpnx", "osgwjmmbrkedijqoutfylvzpnx", "asgjjcmbrkkdihqoutfylvzynx", "asgwjcnerwedihqoutfylvzpnx", "azgwhcmbrkedicqoutfylvzpnx", "asnwjcmbrsedihqoutfylvzpnm", "hsgwjcmgrkedihqoutfilvzpnx", "asgwscmbrkjdihqoutfylvzpnm", "asgbjbmbrkedhhqoutfylvzpnx", "aswwjcmtrkedihqjutfylvzpnx", "asgwicmbrbedihqoutfylvzpnm", "asgwjcubrkedihqoutfylvbnnx", "asvwjcmbrkehihqoutfylhzpnx", "gsgwjcmbrkedihqoutsklvzpnx", "asgwjcubikedihqoitfylvzpnx", "asgwjpmbskedilqoutfylvzpnx", "aigwjcmbrxedihqoutyylvzpnx", "asgwjcpbrkedihxoutfynvzpnx", "asgwjcmbrkegihqoutfklvzcnx", "asgwjvubrkedjhqoutfylvzpnx", "asgwjcabrkedihqoutfyivzplx", "asgwjcxbrkedihqgutfylvepnx", "asgwlcmbrkedihqoutqylvwpnx", "asgwjhmbrkydihqhutfylvzpnx", "asgwjcmbrkedihqoutfylwzone", "asgwycmbrkadihqoutuylvzpnx", "asgwjcybrkedihqoftfylvzpne", "asgwjcmnrkedihrodtfylvzpnx", "asgwicmwrkedihqoutfyovzpnx", "aqgwjlmbrkedilqoutfylvzpnx", "asgwjcmzskwdihqoutfylvzpnx", "asgwjcmdrkebihqoutfylvjpnx", "asgwjcmbrkpdihqoutfylxzphx", "asgwjcmbrkedixqoutlylfzpnx", "asgwjcmbrkadihqoutfylvlpdx", "asgejcmqrkedyhqoutfylvzpnx", "asgwjcmvroedihpoutfylvzpnx", "asgwjcmxrkedihqoutfyivzpmx"}

func main() {

	s1, s2 := getMatches(tests)
	s3 := getCommon(s1, s2)
	fmt.Printf("Tests: %v %v %v %v\n", s1, s2, s3, s3 == testResults)

	s1, s2 = getMatches(input)
	a := getCommon(s1, s2)
	fmt.Printf("Answer: %v\n", a)
}

func getMatches(strings []string) (string, string) {
	for i := 0; i < len(strings); i++ {
		s := strings[i]

		for j := 0; j < len(strings); j++ {
			if j == i {
				continue
			}
			s2 := strings[j]

			dc := getDiffCount(s, s2)
			if dc == 1 {
				return s, s2
			}
		}
	}

	return "", ""
}

func getCommon(s1 string, s2 string) string {
	common := ""
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			common += string(s1[i])
		}
	}

	return common
}

func getDiffCount(s1 string, s2 string) int {
	diffs := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffs++
			if diffs > 1 {
				break
			}
		}
	}
	return diffs
}
