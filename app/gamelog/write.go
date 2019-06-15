package gamelog

import "io"

func write(writer io.Writer, msg string) {
	writer.Write([]byte(msg))
}
