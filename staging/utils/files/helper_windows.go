package files

import (
	"os"
	"syscall"
)

import (
	"time"
)

func fileTimes(fi os.FileInfo) (aTime, mTime, cTime time.Time) {
	stat := fi.Sys().(*syscall.Win32FileAttributeData)
	aTime = time.Unix(0, stat.LastAccessTime.Nanoseconds())
	mTime = time.Unix(0, stat.LastWriteTime.Nanoseconds())
	cTime = time.Unix(0, stat.CreationTime.Nanoseconds())

	return
}

// fileOwner is a NO-OP on windows
func fileOwner(fi os.FileInfo) (uid, gid int) {
	return 0, 0
}
