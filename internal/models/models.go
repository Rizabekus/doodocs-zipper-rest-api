package models

type ZipperService interface{}

type ArchiveInfo struct {
	FileName     string  `json:"filename"`
	Archive_Size float64 `json:"archive_size"`
	Total_Size   float64 `json:"total_size"`
	Total_Files  float64 `json:"total_files"`
	Files        []Files `json:"files"`
}

type Files struct {
	File_Path string  `json:"file_path"`
	Size      float64 `json:"size"`
	MIMEType  string  `json:"mimetype"`
}
