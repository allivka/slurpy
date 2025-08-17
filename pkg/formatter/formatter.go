package formatter

import (
	"strings"
	"fmt"
)

func BasicFormat(src string) (string, error) {
	
	lines := strings.Split(src, "\n")
	
	var result string
	
	for i := 0; i < len(lines); i++ {
		
		line := lines[i]
		
		line = strings.Trim(line, "\n ")
		
		if !(line[len(line) - 1] == '\\') {
			result += strings.Join(strings.Fields(line), " ")
			continue
			
		}
		
		if i == len(lines) - 1 {
			return "", fmt.Errorf("No line fo concatenation with \\ operator at line %d", i + 1)
		}
		
		lines[i] = strings.Join(strings.Fields(string(line[:len(line) - 1]) + lines[i + 1]), " ")
		i--
	}
	
	return result, nil
}