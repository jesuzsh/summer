package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"sync"

	api "github.com/jesuzsh/proglog/api/v1"
)

// Log consists of a list of segments and a pointer to the active segment to
// append writes to. The directory is where we store the segments.
type Log struct {
	// REVIEW
	mu sync.RWMutex

	Dir    string
	Config Config

	activeSegment *segment
	segments      []*segment
}

// New sets defaults for the configs the caller didn't specify, create a log
// instance, and set up that instance.
func NewLog(dir string, c Config) (*Log, error) {
	if c.Segmetn.MaxStoreBytes == 0 {
		c.Segment.MaxStoreBytes = 1024
	}
	if c.Segment.MaxIndexBytes == 0 {
		c.Segment.MaxIndexBytes = 1024
	}
	l := &Log{
		Dir:    dir,
		Config: c,
	}

	return l, l.setup()
}
