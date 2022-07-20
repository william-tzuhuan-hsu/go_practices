package main

import ("fmt"
        // "strconv"
        "strings"
)

func main() {
    /*
    fmt.Println("Strings.")

    fmt.Println(string('A'))

    fmt.Println(string(45))
    // Go thinks we want the 45-th symbol according to ASCII chart

    mj := strconv.Itoa(45)
    fmt.Println(mj) // prints 45

    j, err := strconv.Atoi("-37")
    // if this goes well, err = nil
    if err != nil { // problem
        panic("Error: issue converting string to integer")
    }

    fmt.Println(j)

    x, err2 := strconv.ParseFloat("3.14", 64)
    if err2 != nil {// problem
        panic("Error: issue converting string to integer")
    }
    fmt.Println(x)

    //recall: string concatenation
    s := "Hi"
    t := "Lovers"
    u := s+t
    fmt.Println(u)

    fmt.Println(u[0]) // it won't print out what we expected
    fmt.Println(u[len(u)-1])

    fmt.Println(string(u[0])) // convert it to string type first
    fmt.Println(string(u[len(u)-1]))

    if t[2] == 'v' {
        fmt.Println("Success!")
    }

    dna := "ACCGAT"
    fmt.Println(ReverseComplement(dna))
    */

    s := "Hi Lovers!"
    fmt.Println(s[2:5]) // substring of s of length 5-2 = 3 starting at position 2.
    fmt.Println(s[4:]) // If I don't specify second index, Go assumes we go to end of string
    fmt.Println(s[:5]) // prefix substring of length 5

    // note: sub-arrays and subslices have same notation
    a := make([]int, 6)
    for i := range a {
        a[i] = -2*i + 1
    }
    fmt.Println(a)
    fmt.Println(a[3:5])

    //brainteaser: given a slice a, how could we delete an element at a given index?
    index := 2
    a = append(a[:index], a[index+1:]...) // ...any time we want to append multiple items.
    fmt.Println(a)

    a[1] = a[3]
    fmt.Println(a)
    a = append(a[:3], a[4:]...)
    fmt.Println(a)

    pattern := "ATA"
    text := "ATATA"

    fmt.Println(strings.Count(text, pattern))
    // we should write our own function
    // since the definition of orrurence count is different from ours
    fmt.Println(CountPattern(pattern, text))
    fmt.Println(CountPattern2(pattern, text))
    fmt.Println(StartingIndices(pattern, text))
}


// ReverseComplement takes a DNA string and returns its
func ReverseComplement(dna string) string {
    return Reverse(Complement(dna))
}

//Reverse takes a string and reverses its symbols to produce a new string.
func Reverse(dna string) string {
    dna2 := ""
    for i := range dna {
        dna2 = string(dna[i]) + dna2
    }
    return dna2
}

func ReverseCompeau(s string) string {
    s2 := ""
    n := len(s)
    for i := range s {
        s2 += string(s[n-1-i])
    }
    return s2
}

func Complement(dna string) string {
    dna2 := ""
    for i := range dna {
        // would give error: can't change symbols of a string in Go.
        /*
        switch dna[i] {
        case: 'A':
            dna[i] = 'T'
        case: 'T':
            dna[i] = 'A'
        case: 'C':
            dna[i] = 'G'
        case: 'G':
            dna[i] = 'C'
        }
        */
        switch dna[i] {
        case 'A':
            dna2 += "T"
        case 'T':
            dna2 += "A"
        case 'C':
            dna2 += "G"
        case 'G':
            dna2 += "C"
        }
    }
    return dna2
}

// let's count the number of occurences of a pattern in a text as a substring.
func CountPattern(pattern, text string) int {
    count := 0
    k := len(pattern)
    n := len(text)
    // range over all substrings of text, and increment count if we find a match.
    /*
    for i := range text {
        if text[i:i+k] == pattern {
            count++
        }
    }
    */
    for i := 0; i < n-k+1; i++ {
         if text[i:i+k] == pattern {
             count ++
         }
    }

    return count
}

func CountPattern2(pattern, text string) int {
    hits := StartingIndices(pattern, text)
    return len(hits)
}

//StartingIndices finds all starting positions of pattern in text
func StartingIndices(pattern, text string) []int {
    positions := make([]int, 0)
    k := len(pattern)
    n := len(text)
    // range over all substrings of text, adn append to positions if we find a match.
    for i := 0; i < n-k+1; i++ {
        if text[i:i+k] == pattern {
            positions = append(positions, i)
        }
    }
    return positions
}
