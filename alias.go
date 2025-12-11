// Package fileprep re-exports fileparser types for backward compatibility.
package fileprep

import "github.com/nao1215/fileparser"

// FileType is an alias for fileparser.FileType for backward compatibility.
type FileType = fileparser.FileType

// File type constants re-exported from fileparser for backward compatibility.
const (
	FileTypeCSV         = fileparser.CSV
	FileTypeTSV         = fileparser.TSV
	FileTypeLTSV        = fileparser.LTSV
	FileTypeParquet     = fileparser.Parquet
	FileTypeXLSX        = fileparser.XLSX
	FileTypeCSVGZ       = fileparser.CSVGZ
	FileTypeCSVBZ2      = fileparser.CSVBZ2
	FileTypeCSVXZ       = fileparser.CSVXZ
	FileTypeCSVZSTD     = fileparser.CSVZSTD
	FileTypeTSVGZ       = fileparser.TSVGZ
	FileTypeTSVBZ2      = fileparser.TSVBZ2
	FileTypeTSVXZ       = fileparser.TSVXZ
	FileTypeTSVZSTD     = fileparser.TSVZSTD
	FileTypeLTSVGZ      = fileparser.LTSVGZ
	FileTypeLTSVBZ2     = fileparser.LTSVBZ2
	FileTypeLTSVXZ      = fileparser.LTSVXZ
	FileTypeLTSVZSTD    = fileparser.LTSVZSTD
	FileTypeParquetGZ   = fileparser.ParquetGZ
	FileTypeParquetBZ2  = fileparser.ParquetBZ2
	FileTypeParquetXZ   = fileparser.ParquetXZ
	FileTypeParquetZSTD = fileparser.ParquetZSTD
	FileTypeXLSXGZ      = fileparser.XLSXGZ
	FileTypeXLSXBZ2     = fileparser.XLSXBZ2
	FileTypeXLSXXZ      = fileparser.XLSXXZ
	FileTypeXLSXZSTD    = fileparser.XLSXZSTD
	FileTypeUnsupported = fileparser.Unsupported
)

// DetectFileType detects file type from extension.
// This is a convenience wrapper around fileparser.DetectFileType.
func DetectFileType(path string) FileType {
	return fileparser.DetectFileType(path)
}

// IsCompressed returns true if the file type is compressed.
// This is a convenience wrapper around fileparser.IsCompressed.
func IsCompressed(ft FileType) bool {
	return fileparser.IsCompressed(ft)
}
