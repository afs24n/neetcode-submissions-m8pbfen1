
type Fingerprint [26]int

func groupAnagrams(words []string) [][]string {

    fingerprintCollection := make(map[Fingerprint][]string)

    for i:= 0; i < len(words); i++ {

        fingerprint := Fingerprint{}
        
        word := words[i]

        for j := 0; j < len(word); j++ {
    
            fingerprint[word[j]-'a']++

        }

        fingerprintCollection[fingerprint] = append(fingerprintCollection[fingerprint], word)
        
    }

    var result [][]string

    for _, valueSlice := range fingerprintCollection {
        result = append(result,valueSlice)
    }
    
    return result
    

}
