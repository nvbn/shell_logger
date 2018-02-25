package shell

import "testing"

func testGetFirstCommand(t *testing.T) {
	var commands [6][]byte
	commands[0] = []byte("ls")
	commands[1] = []byte("ls ")
	commands[2] = []byte("ls  -a")
	commands[3] = []byte("ls -la ")
	commands[4] = []byte("")
	commands[5] = []byte(" ls")
	expected := []byte("ls")
	for _, command := range commands[:4] {
		result := GetFirstCommand(command)
		if string(result) != string(expected) {
			t.Fatalf("Expected %s, but got %s", expected, result)
		}
	}
	for _, command := range commands[4:6] {
		result := GetFirstCommand(command)
		if result != nil {
			t.Fatalf("Expected nil, but got %s", result)
		}
	}
}

func testGetTopThreeCommands(t *testing.T) {
	goodCommands := [][]byte{	[]byte("git push origin master"),
								[]byte("git push origin master"),
								[]byte("git push origin master"),
								[]byte("git log"),
								[]byte("git status"),
								[]byte("git log"),
								[]byte("git log"),
								[]byte("git status"),
								[]byte("git reflog"),
								[]byte("git log")}
	for i := 0; i < len(goodCommands); i++ {
		Insert(goodCommands[i], nil)
	}

	topThreeCommands, _ := GetTopThreeCommands([]byte("git"))
	if len(topThreeCommands) != 3 {
		t.Fatalf("Expected %d commands, but got %d", 3, len(topThreeCommands))
	}
	if string(topThreeCommands[0]) != string(goodCommands[3]) {
		t.Fatalf("Expected %s , but got %s", goodCommands[3], topThreeCommands[0])
	}
	if string(topThreeCommands[1]) != string(goodCommands[0]) {
		t.Fatalf("Expected %s , but got %s", goodCommands[0], topThreeCommands[1])
	}
	if string(topThreeCommands[2]) != string(goodCommands[4]) {
		t.Fatalf("Expected %s , but got %s", goodCommands[4], topThreeCommands[2])
	}

	ClearData()
}

func testGetLessThanThreeCommands(t *testing.T) {
	goodCommands := [][]byte{	[]byte("git push origin master"),
								[]byte("git push origin master"),
								[]byte("git push origin master"),
								[]byte("git log")}
	for i := 0; i < len(goodCommands); i++ {
		Insert(goodCommands[i], nil)
	}

	topCommands, _ := GetTopThreeCommands([]byte("git"))
	if len(topCommands) != 2 {
		t.Fatalf("Expected %d commands, but got %d", 2, len(topCommands))
	}
	if string(topCommands[0]) != string(goodCommands[0]) {
		t.Fatalf("Expected %s , but got %s", goodCommands[0], topCommands[0])
	}
	if string(topCommands[1]) != string(goodCommands[3]) {
		t.Fatalf("Expected %s , but got %s", goodCommands[3], topCommands[1])
	}

	ClearData()
}

func testNoCommands(t *testing.T) {
	topCommands, err := GetTopThreeCommands([]byte{})
	if err != nil {
		t.Fatalf("Expected %d commands, but got %d", 0, len(topCommands))
	}
}

func TestQuery(t *testing.T) {
	SetupDatabase("test.db")

	t.Run("TestNoCommands", testNoCommands)
	t.Run("TestGetLessThanThreeCommands", testGetLessThanThreeCommands)
	t.Run("TestGetTopThreeCommands", testGetTopThreeCommands)
	t.Run("TestGetFirstCommand", testGetFirstCommand)
}