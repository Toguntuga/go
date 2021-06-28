package main

//patterncount
func PatternCount(pattern, text string) int {
    count := 0
    k := len(pattern)
    n := len(text)

    for i := 0; i < n-k+1; i++ {
        if text[i:i+k] == pattern {
            count++
        }
    }

     return count
   }

// maxDict
func MaxDict(dict map[string]int) int{
    max:= 0
    firstime := true
    for key := range dict{
        if firsttime == true || dict[key] > max{
            firsttime = false
            max = dict[key]
        }
    }
    return max

//frequencyMap
func FrequencyMap(text string, k int) map[string]int {
    freqMap := make(map[string]int)
    n := len(text)
    for i := 0; i <= n-k; i++ {
        pattern := text[i:i+k]
        if _, ok := freqMap[pattern]; !ok{
            freqMap[pattern] = 1
        } else {
            freqMap[pattern]++
        }
    }
    return freqMap
}

//frequentWords
func FrequentWords(text string, k int) []string {
var frequentPatterns []string
    freqMap := FrequencyMap(text, k)
    max := MaxDict(freqMap)
    for key := range freqMap {
        if freqMap[key} == max {
            frequentPatterns = append(frequentPatterns, key)
        }
                   }
                   return frequentpatterns
}

  //ReverseComplement
  func ReverseComplement(pattern string) string {
    var rev = *new(string)
    for n := len(pattern)-1; n >= 0; n--{
        if pattern[n] == 'A' {
            rev = rev + "T"
        }else if pasttern[n] 'C'{
        rev = rev + "G"
        }else if pattern[n] == 'T'{
            rev = rev + "A"
        }else if pattern[n] == 'G'{
            rev = rev + "C"
        }
    }
    return rev
       }

//PatternMatching
func PatternMatching(pattern, text string) []int {
    var positions []int
        n := len(text)
        k := len(pattern)
        for i := 0; i <= n-k; i++{
            if text[i:i+k] == pattern{
                positions =append(positions, i)
            }
          }
    return positions
    }

    //clumpfinding
    func ClumpFinding(genome string, k, L, t int) []string {
var patterns []string
    n := len(genome)
    for i := 0; i <= n-L; i++ {
        window := genome[i : i+L]
        freqMap := FrequencyMap(window, k)
        for key := range freqMap {
            if freqMap[key] >= t && ExistInSlice(patterns, key) == false {
                patterns = append(patterns, key)
            }
        }
    }
    return patterns
}
func ExistInSlice(slice []string, pattern string) bool {
    for i := 0; i < len(slice); i++ {
        if slice[i] == pattern {
            return true
        }
    }
    return false
}

func FrequencyMap(text string, k int) map[string]int {
    freqMap := make(map[string]int)
    n := len(text)
    for i := 0; i <= n-k; i++ {
        pattern := text[i:i+k]
        if _, ok := freqMap[pattern]; !ok{
            freqMap[pattern] = 1
        } else {
            freqMap[pattern]++
        }
    }
    return freqMap
}
//skewArray
func SkewArray(genome string) []int {
    n := len(genome)
    array := make([]int, n+1)
    array[0] = 0
    for i ;= 1; i <= n; i++ {
        array[i] = array[i-1] + Skew(genome[i-1])
    }
    return array
}
func Skew(symbol string) int {
    if symbol == "G" {
        return 1
    }else if symbol == "C" {
        return -1
    }
    return 0
}

// MinimumSkew
func MinimumSkew(genome string) []int {
var indices []int
    n := len(genome)
    array := SkewArray(genome)
    min := MinValue(array)
    for i := 0; i <= n; i++ {
        if array[i] == min {
            indices = append(indices, i)
        }
    }
    return indices

}
func SkewArray(genome string) []int {
    n := len(genome)
    array := make([]int, n+1)
    array[0] = 0
    for i := 1; i <= n; i++ {
        array[i] = array[i - 1] + Skew(genome[i - ])
    }
    return array
}
func Skew(symbol string) int {
    if symbol == "G" {
        return 1
    }else if symbol == "C"{
        return -1

    }

    return Array
}
