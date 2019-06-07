package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(buf []byte, n int) {
	for i := 0; i < n; i++ {
		alphabetic := false;
		var new_byte byte
		
		if buf[i] >= 'A' && buf[i] <= 'Z' {
			alphabetic = true;
			new_byte = 'A' + ((buf[i] - 'A') + 13) % ('Z' - 'A' + 1)
		} else if buf[i] >= 'a' && buf[i] <= 'z' {
			alphabetic = true;
			new_byte = 'a' + ((buf[i] - 'a') + 13) % ('z' - 'a' + 1)
		}
		
		if alphabetic {
			buf[i] = new_byte
		}
	}	
}

func (self rot13Reader) Read(buf []byte) (int, error) {
	idx := 0
	temp_buf := make([]byte, 512)
	total := 0
	for {
		n, err := self.r.Read(temp_buf)
		if err == io.EOF {
			break	
		}
		
		total += n
		rot13(temp_buf, n)
		for i := 0; i < n; i++ {
			buf[idx] = temp_buf[i]
			idx += 1
		}
	}
	if total == 0 {
		return 0, io.EOF	
	}
	
	return total, nil	
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

