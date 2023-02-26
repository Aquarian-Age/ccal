package main

import (
	"time"

	"github.com/jacobsa/go-serial/serial"
	"github.com/simpleiot/simpleiot/respreader"
)

func main() {

	timeout()
}
func timeout() {
	options := serial.OpenOptions{
		PortName:        "/dev/ttyUSB0",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 0,
		// with serial ports, you just set
		// InterCharacterTimeout to 100 or larger.
		// Otherwise, the goroutine reading the serial
		// port will never exit when you close the read
		// and will still data the next time you open
		// the port. Be aware it may take 100ms for this
		// to close. The linux kernel only accepts timeouts
		// in increments of 0.1s. When using serial ports it
		// makes sense to set the chunkTimeout to 100ms as well.
		// With Go files, a read is supposed to return when
		// the File is closed, but this does not seem to be
		// working with Linux serial devices.
		InterCharacterTimeout: 100,
		RTSCTSFlowControl:     true,
	}

	port, err := serial.Open(options)

	port = respreader.NewResponseReadWriteCloser(port, time.Second,
		time.Millisecond*50)

	// send out prompt
	port.Write("ATAI")

	// read response
	data := make([]byte, 128)
	count, err := port.Read(data)
	data = data[0:count]

	// now process response ...

	// to close the reader process, you must call Close on the reader.
	// This sets a flag that causes the reader goroutine to exit.
	port.Close()
}
