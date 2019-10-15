package main

// tar包实现了文件的打包功能,可以将多个文件或者目录存储到单一的.tar压缩文件中
// tar本身不具有压缩功能,只能打包文件或目录

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 普通文件
	unTarFileName := "pkg/archive/tar/example/tar_example.txt"
	// 压缩文件
	tarFileName := "pkg/archive/tar/example/tar_example.tar"

	// 将普通文件生成压缩文件
	res := Tar(unTarFileName, tarFileName)
	switch res.(type) {
	case error:
		fmt.Println("打包失败:", res)
	case bool:
		fmt.Println("打包成功")
	}

	// 将压缩文件解压
	tarFileName2 := "pkg/archive/tar/example/untar_example.tar"
	unTarFilePath := "pkg/archive/tar/example/untar_example/"
	res2 := UnTar(tarFileName2, unTarFilePath)
	switch res2.(type) {
	case error:
		fmt.Println("解压失败:", res)
	case bool:
		fmt.Println("解压成功")
	}
}

func Tar(unTarFileName string, tarFileName string) interface{} {
	// 打开源文件
	sourceFile, err := os.Open(unTarFileName)
	defer sourceFile.Close()
	if err != nil {
		return err
	}

	/* 向 tar 文件中写入数据是通过 tar.Writer 完成的，所以首先要创建 tar.Writer，
	   可以通过 tar.NewWriter 方法来创建它，该方法要求提供一个 os.Writer 对象，
	   以便将打包后的数据写入该对象中。
	   可以先创建一个文件，然后将该文件提供给 tar.NewWriter 使用。
	*/

	// 创建文件用于储存打包后的数据
	destinationFile, err := os.Create(tarFileName)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	/* 创建tar.Writer对象.此时，我们就拥有了一个 tar.Writer 对象 tw，可以用它来打包文件了。
	   这里要注意一点，使用完 tw 后，一定要执行 tw.Close() 操作，
	   因为 tar.Writer 使用了缓存，tw.Close() 会将缓存中的数据写入到文件中，

	   同时 tw.Close() 还会向 .tar 文件的最后写入结束信息，如果不关闭 tw 而直接退出程序，
	   那么将导致 .tar 文件不完整。
	   存储在 .tar 文件中的每个文件都由两部分组成：文件头信息和文件内容，
	   所以向 .tar 文件中写入每个文件都要分两步：
	   第一步写入文件信息，第二步写入文件数据。
	   对于目录来说，由于没有内容可写，所以只需要写入目录信息即可。
	*/
	tw := tar.NewWriter(destinationFile)
	defer tw.Close()

	// 获取源文件信息
	sourceFileInfo, err := os.Stat(unTarFileName)

	if err != nil {
		return err
	}
	// 根据 os.FileInfo创建tar.Header信息头
	hdr, err := tar.FileInfoHeader(sourceFileInfo, "")
	log.Println(hdr)

	// 第一步写入头文件信息,通过tw.WriteHeader方法将hdr写入.tar文件中
	if err := tw.WriteHeader(hdr); err != nil {
		return err
	}

	// 第二部,写入数据
	_, err = io.Copy(tw, sourceFile)
	if err != nil {
		return err
	}
	return true
}

func UnTar(tarFileName string, unTarFilePath string) interface{} {
	// 将压缩文件生成普通文件
	// 解包的方法，从 .tar 文件中读出数据是通过 tar.Reader 完成的，
	// 所以首先要创建 tar.Reader，可以通过 tar.NewReader 方法来创建它.
	// 该方法要求提供一个 os.Reader 对象，
	// 以便从该对象中读出数据。可以先打开一个 .tar 文件，
	// 然后将该文件提供给 tar.NewReader 使用。
	// 这样就可以将 .tar 文件中的数据读出来了：
	tarFile, err := os.Open(tarFileName)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	// 创建tar.Reader 准备进行解包
	tr := tar.NewReader(tarFile)

	// 重新生成文件夹
	path, err := os.Stat(unTarFilePath)
	if path != nil {
		os.RemoveAll(unTarFilePath)
	}
	os.Mkdir(unTarFilePath, os.ModePerm)

	// 遍历包中的文件
	// 遍历压缩包中的文件
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// 读取到文件目录，跳出循环
			break
		}
		if err != nil {
			return err
		}
		fileName := unTarFilePath + hdr.Name
		_, err = os.Create(fileName)
		if err != nil {
			return err
		}
		fw, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			return err
		}
		_, err = io.Copy(fw, tr)
		if err != nil {
			return err
		}
	}
	return true
}
