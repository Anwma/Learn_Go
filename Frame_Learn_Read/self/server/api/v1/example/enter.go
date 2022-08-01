package example

import "self/server/service"

type ApiGroup struct {
	ExcelApi
	CustomerApi
	FileUploadAndDownloadApi
}

var (
	excelService                 = service.ServiceGroupApp.ExampleServiceGroup.ExcelService
	customService                = serivce.ServiceGroupApp.ExampleServiceGroup.CustomerService
	FileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)
