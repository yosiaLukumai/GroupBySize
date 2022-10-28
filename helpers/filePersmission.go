package helpers

import "os"
// Permssion on the fileMode
var (
	Nopermission        os.FileMode = 0000
	RdWrOwner           os.FileMode = 0700
	RdWrOwnerGroup      os.FileMode = 0770
	RdWrAll             os.FileMode = 0777
	Excute              os.FileMode = 0111
	Write               os.FileMode = 0222
	WriteExcute         os.FileMode = 0333
	ReadExcute          os.FileMode = 0555
	ReadWrite           os.FileMode = 0666
	RdWrExcuteOwnerOnly os.FileMode = 0740
)
