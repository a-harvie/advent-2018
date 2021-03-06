package main

import "fmt"

var tests = []string{
	"abcdef",
	"bababc",
	"abbcde",
	"abcccd",
	"aabcdd",
	"abcdee",
	"ababab",
}
var testResults = 12

var input = []string{"asgwdcmbrkerohqoutfylvzpnx", "asgwjcmbrkejihqoutfylvipne", "asgwjcmbrkedihqoutvylizpnz", "azgsjcmbrkedihqouifylvzpnx", "asgwucmbrktddhqoutfylvzpnx", "asgwocmbrkedihqoutfylvzivx", "aqgwjcmbrkevihqvutfylvzpnx", "tsgljcmbrkedihqourfylvzpnx", "asgpjcmbrkedihqoutfnlvzsnx", "astwjcmbrktdihqrutfylvzpnx", "asgwjcmbrpedhhqoutfylvzynx", "xsgwjcmbrkedieqowtfylvzpnx", "asgwjcmbvkedihfoutnylvzpnx", "asgwjcmtrkedihqouafylvzcnx", "asgwjcmbrkedihqoutfylvxpvm", "usgwjcmbrkedihqortfyuvzpnx", "asgwjcmbrwedihqoutfylizpix", "asgrjcvbrkedixqoutfylvzpnx", "asgwjcmbrogdihqoutfelvzpnx", "aggwjcmbrkesihqoutoylvzpnx", "asgtjccbrkedihqoutfrlvzpnx", "asgcucmbrbedihqoutfylvzpnx", "esgwjcmbrkedihqsutfylvzcnx", "asgwjcmbrkedrhqoutfyobzpnx", "mngwjcbbrkedihqoutfylvzpnx", "asgwjcrbrkeoihqyutfylvzpnx", "apgwjcmbrkednhqogtfylvzpnx", "asgwjcwbrkedihqoutfylplpnx", "asgwjcmbrkfdihqoutfxlvzpyx", "aegwjcmbrkedihqoutfylbxpnx", "asgljcmbrkedixqoutaylvzpnx", "aigwjcmbrkedihqouhfylvzpex", "asgwjbmbrkedihqoutfylfzpnp", "asgwjcmzroedihqoutcylvzinx", "asgwjcwbrieuihqoutfylvzpnx", "aagwjcmbrkedjhqdutfylvzpnx", "ahgwjcmbrkedihqsutfylvzpfx", "asgwjcmbrkedihzosttylvzpnx", "aegwjcmbrkedioqnutfylvzpnx", "asgwjcmbykidihqoutfysvzpnx", "asgwkcxbrkeddhqoutfylvzpnx", "ashwjcmbrkeeihqoutfylvzknx", "acgwjcmbrqedihqoqtfylvzpnx", "asgwjcmtrkedihooutfylszpnx", "asgwjcmbrkmdihqfutrylvzpnx", "asgwjcmbrkedihqoutjylvapnn", "asgwjcmbwkedihqoutkylkzpnx", "asgwjrmbrkedihqoutfycnzpnx", "asgwtcmbrkedihqoqtfylozpnx", "asgajcmbrkedihqoutuylvzpny", "asgwjcmbykedihqoutfylfzpwx", "asgwjcsbrkedihpoutfylvvpnx", "hsdwjcmbrvedihqoutfylvzpnx", "asgwjcmbrkedihqoutfdmszpnx", "adgwjcmbrtidihqoutfylvzpnx", "augwjcmbriedihqoutgylvzpnx", "asgwjvmbreedihqoutfllvzpnx", "asgwjcnbfkedihqoltfylvzpnx", "asgwjcmbykddihqoutqylvzpnx", "ajgwjcmbrkedihqoutfylvpvnx", "asgwjcmbrkydihqoutfylszpnl", "xsgwjcmbrkqdihqoutfylvkpnx", "asgwjcmbrkedimqoutfklvzknx", "csgwjbmbrkedihqoftfylvzpnx", "asgwjcmbjkedihjoutfylvzpnn", "asgwjcmprkedihqoulfalvzpnx", "asgwjcmbrvediqqoutfyuvzpnx", "asgwjambrkedhhqoutkylvzpnx", "asgejcmbrkidihqoutfylvzpnk", "hsiwjcmbrkedihqoutfylvzpnq", "asswjczbrkedihqoutfylczpnx", "asgwjnmbrkedyhzoutfylvzpnx", "asgwscmbrkedihqoutfklvlpnx", "asgwlcmbrktdihqoutfylvzpax", "asfwjcmerkedihqoutfylvipnx", "asgwjcmbrkeditqoeafylvzpnx", "asgwgcmbrkesihqoutfylyzpnx", "fsgwjcmbrkedihqouvfyavzpnx", "asgwjcmbrpedwhqoutfylmzpnx", "asgwjcmbrkzdzhqoucfylvzpnx", "asgwjcmnrketmhqoutfylvzpnx", "asgwjcmbrkedihxoutsylvzpnh", "asgwjcobrkedihqoutfrlvzpox", "asgwjcmbrkedihqootfylxzpox", "asgjjcmcrkedihqoutfylmzpnx", "lsgwjcmbrkedihqoutfyqvzunx", "asgwjcmbrwedihqoutoylvzpnu", "aszwjcmbtkedihqoutfylczpnx", "asgwjcmbykedihqoutfylvgpex", "asgijcmbrkedilqoutkylvzpnx", "astwxcmzrkedihqoutfylvzpnx", "akgwjcmbnkedihqfutfylvzpnx", "asgwjcmbrqndivqoutfylvzpnx", "asgwjrmbrleqihqoutfylvzpnx", "asgwjcmbrkevihqoutfxlvzpvx", "asbwjcmbrkedihqoutfelvwpnx", "asewjcbbrkmdihqoutfylvzpnx", "asgwjcmbrkeaihxoutfylpzpnx", "asgwjzmbrkedihqrotfylvzpnx", "asgwjcmbrkedihqoutgdxvzpnx", "asgwjcwbrkmdihqoutfylvzlnx", "asgwjcmbrkegihqoutfylrzpax", "ajgwjcmbrkegihqhutfylvzpnx", "asgwjcmbrzedihqhutfylvkpnx", "asgwjcmwrkedihqouhfylkzpnx", "aygwjcmbrkedihqoutfdlvzpnr", "asgwjcmbrkednhqoutiylvypnx", "aqgwjcmbrkezihqoutfylvzonx", "bsgwjcmbrkedihqouhfylvzsnx", "asgwjcmcrkedihqokyfylvzpnx", "asgsjcmbrkewiyqoutfylvzpnx", "asgwpcmbrkejihqoutfylzzpnx", "asgwjumbrkedbeqoutfylvzpnx", "asgwjcmbrkedihpoutqylqzpnx", "awgwjcmbrredihqoutfylvzpna", "asgwjsmbraedihqoutfylvzpvx", "asgwncmbrkedihqoutfyljzrnx", "asgwncmbrkedihqohtfylvzonx", "asgwjcmbrkedihqlutfylvypux", "asgwjcmbbkedihooutfylkzpnx", "asghjcmsryedihqoutfylvzpnx", "asgwjcmbrkevihqoulfzlvzpnx", "asggjcmbrkedizqoutfylvzknx", "asbwjcmbriedihqoutfylvmpnx", "asgwjcmbrkedqbqoutfylvzenx", "asgwjcmprkedihqoutfylvzknp", "asgwjcmbrkerihqoutfwlvzpno", "asgwjcmvrkesihqoutrylvzpnx", "asgzjcmbrkedihqoutfnlvbpnx", "asfwjcmbrkhdihqoutfylpzpnx", "asgwjcmbskedihqdutfyyvzpnx", "asgwjcmzrkedihqoutcylvzinx", "asgwjcmbrkedibqoutfylvjonx", "asgwjcmbrbedihqoutfylmzbnx", "asgwjcmbrkedhhqoutmylczpnx", "asgwjcmbrkbgihqoutzylvzpnx", "asgwjcfbrkedihqoupfyxvzpnx", "asiwjcmbzkedihqoutfyluzpnx", "asvwjcmbrkedihqoitfylvzpns", "asgwjcmxikedihqoutfyevzpnx", "asgwjcmbrkedioqoutfylvzwox", "asgwjcmbrkedivqoutjyuvzpnx", "asgwjcmbkkydihqrutfylvzpnx", "asgwjcmbrkxdiuqoutfylvopnx", "asgwjcmbrkedihqouthylvzpra", "asgwjcmbrzedimloutfylvzpnx", "asgwjcmbrkedmhqoulfytvzpnx", "asgwjcmbrkzdihqrutfysvzpnx", "ssgwjcmxrkedihqoutftlvzpnx", "asgwjcmbrkedihqoutfajvzynx", "asgwjcmbrkqdihqxuufylvzpnx", "asmwjcabrkedihqouxfylvzpnx", "asgwjcmbrkeeihqoatfycvzpnx", "asgwjcjbrgedjhqoutfylvzpnx", "asgljcmtrkedihqoutoylvzpnx", "asgwjcmbrkedigqouzfylvzpvx", "ajgvjcmbkkedihqoutfylvzpnx", "asgwjcmbrkedihqtugfygvzpnx", "asgbjcmbrkedihboftfylvzpnx", "asgwjwmbrkedihqontfylhzpnx", "asgwjfmhrkedihqoutfylvqpnx", "asgwjxmbrkedihqoutzylvzpnj", "asgwjcrlrkedihqoutfylvzpsx", "aygwjcmbrkedihqoutsylvzdnx", "zsgwjcmbrkedihjogtfylvzpnx", "asgwjxmbrkegihqoutfylvopnx", "asgwjcmbrkedihqhutfylvzcnr", "asgwicmbrkewihvoutfylvzpnx", "asqwjcmbvkedihqoutfylvzknx", "asgwjcmbrkedihqoktfyevzpnu", "asgwjcmbrkudihqoutfylqzznx", "asgwjdmbrkedihqoutfylvvdnx", "asgwjcmbrkwmihqautfylvzpnx", "asgwjcmbrxedihqoctfyldzpnx", "asgwjdmbrkedjhqoutfyfvzpnx", "asgwjcmtrzedihqoutfylvzpnm", "bpgwjcmbrmedihqoutfylvzpnx", "asgwjctbrkedihqoqtfynvzpnx", "askhjcmbrkedihqoutfylvzrnx", "asgkjcmblkehihqoutfylvzpnx", "asgwjjmbrkedvhqoutfhlvzpnx", "asgwjcmbrkedihqoupzylvzknx", "asgwjcmbukedchqoutfylizpnx", "askwjcmdrkedihqoutwylvzpnx", "asgwjcmbtkcdihloutfylvzpnx", "asgwjcmbrkedwgqoutvylvzpnx", "asmwjcmbrkedihqoutfylozpnc", "asgwjcmbriedibqouofylvzpnx", "asgnjcmcrkedihqoupfylvzpnx", "asgzjcmbrksdihqoutiylvzpnx", "asgwjcrbkkedihqouafylvzpnx", "asgwjcmbkkvdihqqutfylvzpnx", "astwjcqbrkedihqoutfylvzpvx", "asgwjcmhrkehihqoutfylvzvnx", "asgwjcmbraeduhqoutmylvzpnx", "asgwjcmbrkedihquutnylvzptx", "asgwjcmbrkedilqoftfylvzpnz", "akgwjmmbrkedihqoutfylxzpnx", "asgwjcmbrkhxikqoutfylvzpnx", "asgcjcmetkedihqoutfylvzpnx", "fsgwjcmsrkedihooutfylvzpnx", "gsgwjcmbrkedihdoutfylvzdnx", "asgwtccbrkedihqoutfylvwpnx", "asuwjcmbrkedihqcutfylvzpox", "asgwacmbrkodihqlutfylvzpnx", "asgwjcmbrkediuqoutfylqhpnx", "asgwjcmbrkwdrhqoutfylvzpno", "angwjcsblkedihqoutfylvzpnx", "aigwjcmbyoedihqoutfylvzpnx", "adgwjcmbrkedihqtutfylyzpnx", "asgwjzmbrkeeihqputfylvzpnx", "asgwjcmbrkwdihqoutfylvzpwc", "asgpjcmbrkgdihqbutfylvzpnx", "osgwjmmbrkedijqoutfylvzpnx", "asgjjcmbrkkdihqoutfylvzynx", "asgwjcnerwedihqoutfylvzpnx", "azgwhcmbrkedicqoutfylvzpnx", "asnwjcmbrsedihqoutfylvzpnm", "hsgwjcmgrkedihqoutfilvzpnx", "asgwscmbrkjdihqoutfylvzpnm", "asgbjbmbrkedhhqoutfylvzpnx", "aswwjcmtrkedihqjutfylvzpnx", "asgwicmbrbedihqoutfylvzpnm", "asgwjcubrkedihqoutfylvbnnx", "asvwjcmbrkehihqoutfylhzpnx", "gsgwjcmbrkedihqoutsklvzpnx", "asgwjcubikedihqoitfylvzpnx", "asgwjpmbskedilqoutfylvzpnx", "aigwjcmbrxedihqoutyylvzpnx", "asgwjcpbrkedihxoutfynvzpnx", "asgwjcmbrkegihqoutfklvzcnx", "asgwjvubrkedjhqoutfylvzpnx", "asgwjcabrkedihqoutfyivzplx", "asgwjcxbrkedihqgutfylvepnx", "asgwlcmbrkedihqoutqylvwpnx", "asgwjhmbrkydihqhutfylvzpnx", "asgwjcmbrkedihqoutfylwzone", "asgwycmbrkadihqoutuylvzpnx", "asgwjcybrkedihqoftfylvzpne", "asgwjcmnrkedihrodtfylvzpnx", "asgwicmwrkedihqoutfyovzpnx", "aqgwjlmbrkedilqoutfylvzpnx", "asgwjcmzskwdihqoutfylvzpnx", "asgwjcmdrkebihqoutfylvjpnx", "asgwjcmbrkpdihqoutfylxzphx", "asgwjcmbrkedixqoutlylfzpnx", "asgwjcmbrkadihqoutfylvlpdx", "asgejcmqrkedyhqoutfylvzpnx", "asgwjcmvroedihpoutfylvzpnx", "asgwjcmxrkedihqoutfyivzpmx"}

func main() {

	s := getSum(tests)
	fmt.Printf("Tests: %v %v %v\n", tests, s, s == testResults)

	a := getSum(input)
	fmt.Printf("Answer: %v\n", a)
}

func getSum(strings []string) int {

	var twos int
	var threes int
	for _, s := range strings {
		counts := charCount(s)

		hasTwo := false
		hasThree := false
		for _, c := range counts {
			if c == 2 && !hasTwo {
				hasTwo = true
				twos += 1
			}
			if c == 3 && !hasThree {
				hasThree = true
				threes += 1
			}
			if hasTwo && hasThree {
				break
			}
		}
	}

	return twos * threes
}

func charCount(s string) map[byte]int {
	counts := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		count, found := counts[s[i]]
		if found {
			counts[s[i]] = count + 1
		} else {
			counts[s[i]] = 1
		}
	}

	return counts
}
