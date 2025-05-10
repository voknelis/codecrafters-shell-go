package redirect

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ParseRedirection(operator string, target string) (*Redirection, error) {
	if len(operator) == 0 {
		return nil, fmt.Errorf("-shell: \"\": Invalid redirect operator")
	}

	redirection := &Redirection{
		FileDescriptor: 1,
	}

	fileDescription, fds, err := parseRedirectionFileDescriptor(operator)
	if err != nil {
		return nil, err
	}
	if fileDescription != -1 {
		redirection.FileDescriptor = fileDescription
	}

	operatorSuffix, _ := strings.CutPrefix(operator, fds)
	switch operatorSuffix {
	case string(RedirectionOutput):
		redirection.Type = RedirectionOutput
	case string(RedirectionOutputAppend):
		redirection.Type = RedirectionOutputAppend
	default:
		return nil, fmt.Errorf("-shell: %s: Invalid redirect operator", operatorSuffix)
	}

	redirection.Target = target

	return redirection, nil
}

func parseRedirectionFileDescriptor(operator string) (int, string, error) {
	re := regexp.MustCompile("^[0-9]*")
	fileDescriptorString := re.FindString(operator)

	if fileDescriptorString == "" {
		return -1, "", nil
	}

	fileDescriptor, _ := strconv.Atoi(fileDescriptorString)
	if fileDescriptor == 1 || fileDescriptor == 2 {
		return fileDescriptor, fileDescriptorString, nil
	}

	if fileDescriptor >= 0 && fileDescriptor <= 9 {
		return -1, "", fmt.Errorf("-shell: %d: Unsupported file description", fileDescriptor)
	}

	return -1, "", fmt.Errorf("-shell: %d: Bad file descriptor", fileDescriptor)
}
