package tokenizer_test

import (
	"testing"

	"github.com/codecrafters-io/shell-starter-go/app/tokenizer"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	input    string
	expected []string
}

func runTests(t *testing.T, tests []testCase) {
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := tokenizer.Tokenize(tc.input)
			assert.Equal(t, tc.expected, output, "Input: %s", tc.input)
		})
	}
}

func TestTokenize(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		tests := []testCase{
			{"Empty input", "", []string{}},
			{"Single argument", "arg1", []string{"arg1"}},
			{"Multiple arguments", "arg1 arg2 arg3", []string{"arg1", "arg2", "arg3"}},
		}
		runTests(t, tests)
	})

	t.Run("Single quotes", func(t *testing.T) {
		tests := []testCase{
			{"Single arg", `'arg1'`, []string{`arg1`}},
			{"Multiple args", `'arg1' 'arg 2' 'arg 3 4'`, []string{`arg1`, `arg 2`, `arg 3 4`}},
			{"Empty spaces", `'  arg1  arg2  '`, []string{`  arg1  arg2  `}},
			{"Non-breaking quotes", `arg1''arg2`, []string{`arg1arg2`}},
			{"Special chars", `'*?[]{}$\\'`, []string{`*?[]{}$\\`}},
			{"Unclosed quote", `'arg1`, []string{`arg1`}},
			{"Unclosed quote at end", `arg1'`, []string{`arg1`}},
			{"Empty quote", `''`, []string{}}, // TODO: should be []string{""}
		}
		runTests(t, tests)
	})

	t.Run("Double quotes", func(t *testing.T) {
		tests := []testCase{
			{"Single arg", `"arg1"`, []string{`arg1`}},
			{"Multiple args", `"arg1" "arg 2" "arg 3 4"`, []string{`arg1`, `arg 2`, `arg 3 4`}},
			{"Empty spaces", `"  arg1  arg2  "`, []string{`  arg1  arg2  `}},
			{"Non-breaking quotes", `arg1""arg2`, []string{`arg1arg2`}},
			{"Unclosed quote", `"arg1`, []string{`arg1`}},
			{"Unclosed quote at end", `arg1"`, []string{`arg1`}},
			{"Empty quote", `""`, []string{}}, // TODO: should be []string{""}
		}
		runTests(t, tests)
	})

	t.Run("Double quotes escaping", func(t *testing.T) {
		tests := []testCase{
			{"Escape quote", `"arg\""`, []string{`arg"`}},
			{"Escape backslash", `"arg\\"`, []string{`arg\`}},
			{"Escape dollar sign", `"arg\$"`, []string{`arg$`}},
			{"Escape backtick", "\"arg`\"", []string{"arg`"}},
			{"Escape newline", "\"arg1\narg2\"", []string{"arg1\narg2"}},
			{"Preserve non-escapable", `"arg1\a\b\c\ "`, []string{`arg1\a\b\c\ `}},
			{"Consecutive backslashes", `"arg1\\\"arg2"`, []string{`arg1\"arg2`}},
		}
		runTests(t, tests)
	})

	t.Run("Mixed cases", func(t *testing.T) {
		tests := []testCase{
			{"Escape quote", `arg1 'arg2 arg3' "arg4 arg5"`, []string{`arg1`, `arg2 arg3`, `arg4 arg5`}},
			{"Single quotes within double quotes", `"arg1'arg2'"`, []string{`arg1'arg2'`}},
			{"Double quotes after word", `arg1"arg2"`, []string{`arg1arg2`}},
			{"Double quote before word", `"arg1"arg2`, []string{`arg1arg2`}},
			{"Escaped double quotes within double quotes", `"arg1\"arg2\""`, []string{`arg1"arg2"`}},
			{"Quotes and escapes mixed", `arg1\ "arg2" \'arg3\'`, []string{`arg1 arg2`, `'arg3'`}},
			{"Complex nesting", `"arg1\"'arg2'\""`, []string{`arg1"'arg2'"`}},
		}
		runTests(t, tests)
	})

	t.Run("Edge cases", func(t *testing.T) {
		tests := []testCase{
			{"Only spaces", "   ", []string{}},
			{"Trailing space", "arg1 ", []string{"arg1"}},
			{"Leading space", " arg1 ", []string{"arg1"}},
			{"Multiple spaces", "arg1  arg2   arg3", []string{"arg1", "arg2", "arg3"}},
			{"Empty quoted strings", `"" '' ""`, []string{}}, // TODO: should be []string{"", "", ""}
		}
		runTests(t, tests)
	})
}

func TestRealWorldExamples(t *testing.T) {
	tests := []testCase{
		{"Simple command", "ls -la", []string{"ls", "-la"}},
		{"Command with path", "/usr/bin/ls -la /home", []string{"/usr/bin/ls", "-la", "/home"}},
		{"Command with quoted path", `ls -la "/home/user/My Documents"`, []string{"ls", "-la", "/home/user/My Documents"}},
		{"Echo with quotes", `echo "Hello, World!"`, []string{"echo", "Hello, World!"}},
		{"Grep with pattern", `grep -E "^[a-z]+" file.txt`, []string{"grep", "-E", "^[a-z]+", "file.txt"}},
		{"Find with escaped spaces", `find /path/to/dir -name "*.txt" -o -name "*.md"`, []string{"find", "/path/to/dir", "-name", "*.txt", "-o", "-name", "*.md"}},
		{"Command with environment variable", `echo "User: $USER"`, []string{"echo", "User: $USER"}},
		{"Complex glob pattern", `ls *.{png,jpg,gif}`, []string{"ls", "*.{png,jpg,gif}"}},
		{"Command with redirection", `echo "output" > file.txt`, []string{"echo", "output", ">", "file.txt"}},
		{"Command with pipe", `cat file.txt | grep pattern`, []string{"cat", "file.txt", "|", "grep", "pattern"}},
	}
	runTests(t, tests)
}
