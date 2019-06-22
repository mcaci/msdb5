package notify

import "io"

func write(writer io.Writer, msg string) {
	writer.Write([]byte(msg))
}
