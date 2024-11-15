package purse

import (
	"math/rand"
	"strings"
	"time"
)

func MakeLines(s string) []string {
	return strings.Split(s, "\n")
}

func JoinLines(lines []string) string {
	return strings.Join(lines, "\n")
}

func ReplaceLastSubStr(s, old, new string) string {
	pos := strings.LastIndex(s, old)
	if pos == -1 {
		return s
	}
	return s[:pos] + new + s[pos+len(old):]
}

func GetFirstLine(s string) string {
	lines := MakeLines(s)
	if len(lines) == 0 {
		return s
	}
	return lines[0]
}

func GetLastLine(s string) string {
	lines := MakeLines(s)
	if len(lines) == 0 {
		return s
	}
	return lines[len(lines)-1]
}

func RemoveAllSubStr(s string, subs ...string) string {
	for _, sub := range subs {
		s = strings.ReplaceAll(s, sub, "")
	}
	return s
}

func CountLeadingSpaces(line string) int {
	count := 0
	for _, char := range line {
		if char != ' ' {
			break
		}
		count++
	}
	return count
}

func PrefixLines(str, prefix string) string {
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		lines[i] = prefix + line
	}
	return strings.Join(lines, "\n")
}

func FlattenLines(lines []string) []string {
	for i, line := range lines {
		lines[i] = strings.TrimLeft(line, " \t")
	}
	return lines
}

func Flatten(str string) string {
	lines := MakeLines(str)
	flat := FlattenLines(lines)
	return strings.Join(flat, "")
}

func TrimLeadingSpaces(str string) string {
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimLeft(line, " ")
	}
	return strings.Join(lines, "\n")
}

func SliceContains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func BackTick() string {
	return "`"
}

func ReplaceFirstLine(input, newLine string) string {
	lines := strings.Split(input, "\n")
	if len(lines) > 0 {
		lines[0] = newLine
	}
	return strings.Join(lines, "\n")
}

func ReplaceLastLine(input, newLine string) string {
	lines := strings.Split(input, "\n")
	if len(lines) > 0 {
		lines[len(lines)-1] = newLine
	}
	return strings.Join(lines, "\n")
}

func Squeeze(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func ScanBetweenSubStrs(s, start, end string) []string {
	var out []string
	inSearch := false
	searchStr := ""

	i := 0
	for i < len(s) {
		// Check for the start delimiter
		if !inSearch && i+len(start) <= len(s) && s[i:i+len(start)] == start {
			inSearch = true
			searchStr = start // Start capturing, including the start delimiter
			i += len(start)
			continue
		}

		// If we're in search mode, start capturing until we find the end delimiter
		if inSearch {
			// Check for the end delimiter
			if i+len(end) <= len(s) && s[i:i+len(end)] == end {
				searchStr += end // Include the end delimiter
				out = append(out, searchStr)
				searchStr = ""
				inSearch = false
				i += len(end)
				continue
			}
			// Append current character to searchStr if still inside the delimiters
			searchStr += string(s[i])
		}

		i++
	}

	return out
}

func RemoveFirstLine(input string) string {
	index := strings.Index(input, "\n")
	if index == -1 {
		return ""
	}
	return input[index+1:]
}

func RemoveTrailingEmptyLines(input string) string {
	// Split the input string into lines
	lines := strings.Split(input, "\n")

	// Remove empty lines from the end
	for len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}

	// Join the cleaned lines back into a single string
	return strings.Join(lines, "\n")
}

func RemoveEmptyLines(input string) string {
	// Split the input string into lines
	lines := strings.Split(input, "\n")

	// Create a slice to hold the non-empty lines
	var cleanedLines []string

	// Loop through the lines and add non-empty lines to cleanedLines
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			cleanedLines = append(cleanedLines, line)
		}
	}

	// Join the non-empty lines back into a single string
	return strings.Join(cleanedLines, "\n")
}

func RemoveDuplicatesInSlice(strSlice []string) []string {
	unique := make(map[string]bool)
	var result []string
	for _, item := range strSlice {
		if _, found := unique[item]; !found {
			unique[item] = true
			result = append(result, item)
		}
	}
	return result
}

func WrapStr(s, prefix, suffix string) string {
	return prefix + s + suffix
}

func MatchLeadingSpaces(str1, str2 string) string {
	leadingSpaces := len(str2) - len(strings.TrimLeft(str2, " "))
	padding := strings.Repeat(" ", leadingSpaces)
	return padding + str1
}

func SnipStrAtIndex(s string, x int) string {
	// Ensure that x doesn't exceed the length of the string
	if x > len(s) {
		x = len(s)
	}
	return s[:x]
}

func TargetSearch(input, primarySearch, secondarySearch string) (string, bool) {
	startIndex := strings.Index(input, primarySearch)
	if startIndex == -1 {
		return "", false
	}
	substring := input[startIndex:]
	endIndex := strings.Index(substring, secondarySearch)
	if endIndex == -1 {
		return "", false
	}
	finalEndIndex := startIndex + endIndex + len(secondarySearch)
	return input[startIndex:finalEndIndex], true
}

func SplitWithTargetInclusion(str, target string) []string {
	var parts []string
	start := 0

	for {
		// Find the next occurrence of the target string
		index := strings.Index(str[start:], target)
		if index == -1 {
			// No more occurrences, add the remaining part of the string
			parts = append(parts, str[start:])
			break
		}

		// Adjust the index relative to the original string
		index += start

		// Add the part up to the target
		parts = append(parts, str[start:index])
		// Add the target itself as a separate item
		parts = append(parts, target)

		// Move the start index past the target
		start = index + len(target)
	}

	return parts
}

func PrefixSliceItems(items []string, prefix string) string {
	var prefixedItems []string
	for _, item := range items {
		prefixedItems = append(prefixedItems, prefix+item)
	}
	return strings.Join(prefixedItems, "")
}

func ReverseSlice[T any](slice []T) []T {
	n := len(slice)
	reversed := make([]T, n)
	for i, v := range slice {
		reversed[n-1-i] = v
	}
	return reversed
}

func RandStr(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
