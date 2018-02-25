package shell

import (
	"sort"
	"fmt"
	"unicode"
)

// Get the first command from an shell input
func GetFirstCommand(str []byte) []byte {
	if len(str) > 0 && unicode.IsSpace(int32(str[0])) {
		return nil
	}
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			return str[:i]
		}
	}
    return str
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type commandCount struct {
	frequency int
	command string
}

type commandCountCollection []*commandCount

func (cc commandCountCollection) Len() int {
	return len(cc)
}

func (cc commandCountCollection) Swap(i, j int) {
	cc[i], cc[j] = cc[j], cc[i]
}

func (cc commandCountCollection) Less(i, j int) bool {
	return cc[i].frequency > cc[j].frequency
}

// Get the top three most frequent commands
func GetTopThreeCommands(firstCommand []byte) ([][]byte, error) {
	goodCommands, err := GetGoodCommands(firstCommand)
	if err != nil {
		return [][]byte{}, err
	}
	fmt.Printf("GOOD COMMAND %v", len(goodCommands))
	if goodCommands == nil || len(goodCommands) == 0 {
		return [][]byte{}, nil
	}

	sortedValues := make(map[string]*commandCount)
	for i := 0; i < len(goodCommands); i++ {
		currCommand := string(goodCommands[i])
		if _, ok := sortedValues[currCommand]; ok {
			sortedValues[currCommand].frequency += 1
		} else {
			sortedValues[currCommand] = &commandCount{1, currCommand}
		}
	}
	commandCountSlice := make(commandCountCollection, len(sortedValues))
	i := 0
	for _, val := range sortedValues {
		commandCountSlice[i] = val
		i += 1
	}
	sort.Sort(commandCountSlice)
	threeOrMin := Min(3, len(commandCountSlice))
	returnCommands := make([][]byte, threeOrMin)
	for i, command := range commandCountSlice[:threeOrMin] {
		returnCommands[i] = []byte(command.command)
	}
	return returnCommands, nil
}
