package main

import("fmt")

func main() {
    // pattern := "AAA"
    // text := "AAAGAGTGTCTGATAGCAGCTTCTGAACTGGTTACCTGCCGTGAGTAAATTAAATTTTATTGACTTAGGTCACTAAATACTTTAACCAATATAGGCATAGCGCACAG"
    // fmt.Println(PatternCount(pattern, text))

    // pattern := "GCTCAGCCACAACACGAGGGATACTATTATCACGGTCAGTACAACAACGCATTTGTGATCAGCAACGCACTAAGCTTGCCCAGGGTAGAACACGAGACGCACTCTGTAGCCGTTGTTATCCGACCCTTTAGGACCTTGCGCTGGGCTAGGATGGATAAACCTCGTGGTGCGGCTGTCTTTAGATGATGCTTCCAGGCGAG"
    // fmt.Println(pattern)
    // fmt.Println(ReverseComplement(pattern))
    // fmt.Println(Reverse(pattern))
    // fmt.Println(Complement(pattern))

    // text := "mamaliga"
    // k := 2
    // freqMap := FrequencyTable(text, k)
    // max := MaxMap(freqMap)
    // fmt.Println(freqMap)
    // fmt.Println(max)

    // text := "CCAGCGGGGGTTGATGCTCTGGGGGTCACAAGATTGCATTTTTATGGGGTTGCAAAAATGTTTTTTACGGCAGATTCATTTAAAATGCCCACTGGCTGGAGACATAGCCCGGATGCGCGTCTTTTACAACGTATTGCGGGGTAAAATCGTAGATGTTTTAAAATAGGCGTAAC"
    // k := 5
    // fmt.Println(BetterFrequentWords(text, k))

    // text := "AGCGTGCCGAAATATGCCGCCAGACCTGCTGCGGTGGCCTCGCCGACTTCACGGATGCCAAGTGCATAGAGGAAGCGAGCAAAGGTGGTTTCTTTCGCTTTATCCAGCGCGTTAACCACGTTCTGTGCCGACTTT"
    // pattern := "TTT"
    // fmt.Println(StartingPosition(pattern, text))

    // text := "CGGACTCGACAGATGTGAAGAACGACAATGTGAAGACTCGACACGACAGAGTGAAGAGAAGAGGAAACATTGTAA"
    // k := 5
    // L := 50
    // t := 4
    // fmt.Println(FindClumps(text, k, L, t))

    genome := "TAAAGACTGCCGAGAGGCCAACACGAGTGCTAGAACGAGGGGCGTAAACGCGGGTCCGAT"
    fmt.Println(SkewArray(genome))
    fmt.Println(MinimumSkew(genome))
}

func PatternCount(pattern string, text string) int {
    count := 0
    n := len(text)
    k := len(pattern)
    for i := 0; i < n-k+1; i++ {
        if text[i:i+k] == pattern{
            count++
        }
    }
    return count
}

func ReverseComplement(pattern string) string {
    pattern = Reverse(pattern)
    pattern = Complement(pattern)
    return pattern
}

func Reverse(pattern string) string {
    rev := ""
    for i := range pattern {
        rev = string(pattern[i]) + rev
    }
    return rev
}

func Complement(pattern string) string {
    com := ""
    for i := range pattern {
        switch pattern[i] {
        case 'A':
            com += "T"
        case 'T':
            com += "A"
        case 'C':
            com += "G"
        case 'G':
            com += "C"
        }
    }
    return com
}

func FrequencyTable(text string, k int) map[string]int {
    n := len(text)
    table := make(map[string]int)
    for i := 0; i < n-k+1; i++ {
        pattern := text[i:i+k]
        table[pattern]++
    }
    return table
}

func MaxMap(freqTable map[string]int) int {
    max := 0
    firstTimeThrough := true
    for _, val := range freqTable {
        if firstTimeThrough == true || val > max {
            firstTimeThrough = false
            max = val
        }
    }
    return max
}

func BetterFrequentWords(text string, k int) []string {
    freqPatterns := make([]string, 0)
    freqMap := FrequencyTable(text, k)
    max := MaxMap(freqMap)
    for patterns, val := range freqMap {
        if val == max {
            freqPatterns = append(freqPatterns, patterns)
        }
    }
    return freqPatterns
}

func StartingPosition(pattern string, text string) []int {
    n := len(text)
    k := len(pattern)
    positions := make([]int, 0)
    for i := 0; i <= n-k; i++ {
        if text[i:i+k] == pattern {
            positions = append(positions, i)
        }
    }
    return positions
}

func FindClumps(text string, k int, L int, t int) []string {
    patterns := make([]string, 0)
    n := len(text)
    for i := 0; i < n-L+1; i++ {
        windows := text[i:i+L]
        freqTable := FrequencyTable(windows, k)
        for key := range freqTable {
            if freqTable[key] >= t && Contain(patterns, key) == false{
                patterns = append(patterns, key)
            }
        }
    }
    return patterns
}

func Contain(patterns []string, key string) bool {
    for _, val := range patterns {
        if string(val) == key {
            return true
        }
    }
    return false
}

func SkewArray(genome string) []int {
    n := len(genome)
    skArray := make([]int, 0)
    skArray = append(skArray, 0)
    for i := 1; i <= n; i++ {
         skArray = append(skArray, skArray[i-1]+Skew(string(genome[i-1])))
    }
    return skArray
}

func Skew(symbol string) int {
    if symbol == "G" {
        return 1
    }else if symbol == "C" {
        return -1
    }
    return 0
}

func MinimumSkew(genome string) []int {
    MinSkew := make([]int, 0)
    skArray := SkewArray(genome)
    m := 0
    firstTimeThrough := true
    for i := range skArray {
        if firstTimeThrough || skArray[i] <= m {
            m = skArray[i]
            firstTimeThrough = false
        }
    }
    for i := range skArray {
        if skArray[i] == m {
            MinSkew = append(MinSkew, i)
        }
    }
    return MinSkew
}
