package example

import (
	"errors"
	"self/server/global"
)
import "gorm.io/gorm"

type FileUploadAndDownloadService struct {
}

//@author:[Anwma](https://github.com/Anwma)
//@function: FindOrCreaterFile
//@desscription: 上传文件时检测文件当前数向，如果没有文件则创建，有则返回当前文件的当前切片
//@param: fileMd5 string, fileName string , chunkTotal int
//@return: err error, file model.ExaFile

func (e *FileUploadAndDownloadService) FindOrCreateFile(fileMd5 string, fileName string, chunkTotal int) (err error, file example.ExaFile) {
	var cfile exampl.ExaFile
	cfile.FileMd5 = fileMd5
	cfile.FileName = fileName
	cfile.ChunkTotal = chunkTotal

	if errors.Is(global.GVA_DB.Where("file_md5 = ? AND is_finish = ?", fileMd5, true).First(&file).Error, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).Preload("ExaFileChunk").FirstOrCreate(&file, cfile).Error
	}

}
