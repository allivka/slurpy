package formatter

import (
	"strings"
	"fmt"
)

func Format(src string) (string, error) {
	
	t := strings.Split(src, "\n")
	
	lines := []string{}
	
	for _, v := range t {
		line := strings.TrimSpace(v)
		if line == "" {continue}
		lines = append(lines, line)
	}
	
	var result string
	
	for i := 0; i < len(lines); i++ {
		
		line := lines[i]
		
		if !(line[len(line) - 1] == '\\') {
			result += strings.Join(strings.Fields(line), " ") + "\n"
			fmt.Printf("No \\ operator at the end of line\nline: %v\nformatted line: %v\ncurrent result: %v\n", line, strings.Join(strings.Fields(line), " "), result)
			continue
			
		}
		
		if i == len(lines) - 1 {
			return "", fmt.Errorf("No line for concatenation with \\ operator at line %d", i + 1)
		}
		
		lines[i + 1] = strings.Join(strings.Fields(strings.TrimSpace(string(line[:len(line) - 1]) + lines[i + 1])), " ")
	}
	
	return result, nil
}