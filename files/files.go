package files

import (
	"os"
    "strings"
)

/*
新建目录
dirs：需要建立的目录文件
返回创建目录是否成功
例子：MakeDir("/data/test",0777)
*/
func MakeDir(dirs string, mode os.FileMode) bool {
    err := os.MkdirAll(dirs,mode)
    if err == nil {
        return true
    }
    
    //创建失败
    return false
}

/*
根据文件和目录创建
bas_dir :目录路径位置
file_name :文件路基
返回整个目录
DirNames("/home/wwwroot","2020/12")，创建/home/wwwroot/2020/12的目录
*/
func DirNames(bas_dir string,file_name string) string {
    //根据文件切分目录
    dirs :=strings.Split(file_name,"/")
    var dir string
    if len(dirs) >1 {
        for i:=0;i<len(dirs)-1;i++ {
            dir=dir+dirs[i]+"/"
        }
        dir = strings.TrimLeft(dir,"/")
        MakeDir(bas_dir+"/"+dir,0777)
    }
    return dir
}