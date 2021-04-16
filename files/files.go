package files

import (
	"os"
    "strings"
	"io/ioutil"
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


/*
递归返回子目录下面所有文件
pathname :要读取的目录地址
vals :接受返回文件目录结果的数组
返回整个目录在slice
var vals []string
例子：GetAllFile("./github.com/netaxcess/util", vals)，读取/github.com/netaxcess/util目录下面所以的文件
*/
func GetAllFile(pathname string, vals []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		return vals, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			vals, err = GetAllFile(fullDir, vals)
			if err != nil {
				return vals, err
			}
		} else {
			fullName := pathname + "/" + fi.Name()
			vals = append(vals, fullName)
		}
	}
	return vals, nil
}

/*
将内容写入文件
filename :要写入的文件地址
data :要写入的文件内容，字符串
mode :文件权限
例子：FilePutContents("./github.com/netaxcess/util/a.txt", "内容", 0666)，读取/github.com/netaxcess/util目录下面所以的文件
*/
func FilePutContents(filename string, data string, mode os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(data), mode)
}

/*
读取文件内容
filename :要读取的文件
例子：FileGetContents("./github.com/netaxcess/util/a.txt")，读取/github.com/netaxcess/util目录下面所以的文件
*/
func FileGetContents(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	return string(data), err
}