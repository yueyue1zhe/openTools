package y

import (
	"openTools/y/compute"
	"openTools/y/file"
	"openTools/y/global"
	"openTools/y/jwt"
	"openTools/y/qiniu"
)

func Global() *global.Global {
	return global.NewGlobal()
}
func Compute() *compute.Compute {
	return compute.NewCompute()
}
func File() *file.File {
	return file.NewFile()
}

func JWT(client string) *jwt.JWT {
	return jwt.NewJWT(client)
}

func QiNiu(conf qiniu.Conf) *qiniu.QiNiu {
	return qiniu.NewQiNiu(conf)
}
