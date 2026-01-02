package utils

/*******************
* Import
*******************/
import (
	"encoding/binary"
	"io"
	"os"
	"path/filepath"

	"chatia/modules/errcode"
	"chatia/modules/interfaces"
)

/*******************
* FileReadInt32
*******************/
func FileReadUint32(file *os.File, offset int64, intBuffer *uint32) (int64, error) {
	var buffer [4]byte
	readLength, err := file.ReadAt(buffer[:], offset)
	if err != nil || readLength != 4 {
		return 0, err
	}
	*intBuffer = uint32(binary.BigEndian.Uint32(buffer[:]))
	return offset + 4, nil
}

/*******************
* FileWriteInt32
*******************/
func FileWriteUint32(file *os.File, offset int64, intBuffer uint32) (int64, error) {
	var buffer [4]byte
	binary.BigEndian.PutUint32(buffer[:], intBuffer)

	// si offset == -1, écrire à la fin du fichier
	if offset == -1 {
		var err error
		offset, err = file.Seek(0, io.SeekEnd)
		if err != nil {
			return 0, err
		}
	}

	writtenLength, err := file.WriteAt(buffer[:], offset)
	if err != nil || writtenLength != 4 {
		return 0, err
	}
	return offset + 4, nil
}

/*******************
* FileReadString
*******************/
func FileReadString(file *os.File, offset int64, strBuffer *string) (int64, error) {
	length := uint32(0)
	offset, err := FileReadUint32(file, offset, &length)
	if err != nil {
		return 0, err
	}

	buffer := make([]byte, length)
	readLength, err := file.ReadAt(buffer[:], offset)
	if err != nil || uint32(readLength) != length {
		return 0, err
	}
	*strBuffer = string(buffer)
	return offset + int64(length), nil
}

/*******************
* FileWriteString
*******************/
func FileWriteString(file *os.File, offset int64, strBuffer string) (int64, error) {
	var length = uint32(len(strBuffer))

	// si offset == -1, écrire à la fin du fichier
	if offset == -1 {
		var err error
		offset, err = file.Seek(0, io.SeekEnd)
		if err != nil {
			return 0, err
		}
	}

	offset, err := FileWriteUint32(file, offset, length)
	if err != nil {
		return 0, err
	}

	writtenLength, err := file.WriteAt([]byte(strBuffer), offset)
	if err != nil || uint32(writtenLength) != length {
		return 0, err
	}
	return offset + int64(length), nil
}

/*******************
* FileReadData
*******************/
func FileReadData(file *os.File, offset int64, dataBuffer *[]byte) (int64, error) {
	var length uint32
	offset, err := FileReadUint32(file, offset, &length)
	if err != nil {
		return 0, err
	}

	*dataBuffer = make([]byte, length)
	readLength, err := file.ReadAt((*dataBuffer)[:], offset)
	if err != nil || uint32(readLength) != length {
		print(err.Error())
		return 0, err
	}
	return offset + int64(length), nil
}

/*******************
* FileWriteData
*******************/
func FileWriteData(file *os.File, offset int64, dataBuffer []byte) (int64, error) {
	var length = uint32(len(dataBuffer))

	// si offset == -1, écrire à la fin du fichier
	if offset == -1 {
		var err error
		offset, err = file.Seek(0, io.SeekEnd)
		if err != nil {
			return 0, err
		}
	}

	offset, err := FileWriteUint32(file, offset, length)
	if err != nil {
		return 0, err
	}

	writtenLength, err := file.WriteAt(dataBuffer, offset)
	if err != nil || uint32(writtenLength) != length {
		return 0, err
	}
	return offset + int64(length), nil
}

/*******************
* FileGetEndOffset
*******************/
func FileGetEndOffset(file *os.File) (int64, error) {
	offset, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, err
	}
	return offset, nil
}

/*******************
* initConfigFile
*******************/
func initConfigFile(confFile string) *os.File {
	// Create or truncate the configuration file
	fileHandle, err := os.OpenFile(confFile, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_CREATE)
		os.Exit(errcode.ERROR_FATAL_CONFIG_CREATE)
	}

	// Write file version
	_, err = FileWriteUint32(fileHandle, 0, 0x62726e01)
	if err != nil {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_WRITE)
		os.Exit(errcode.ERROR_FATAL_CONFIG_WRITE)
	}

	return fileHandle
}

/*******************
* CreateConfigFile
*******************/
func CreateConfigFile(brainConfig interfaces.I_BrainConfig, configFileName string) *os.File {
	path := brainConfig.GetSaveDirectory()
	err := os.MkdirAll(path, 0700)
	if err != nil {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_CREATE)
		os.Exit(errcode.ERROR_FATAL_CONFIG_OPEN)
	}

	confFile := filepath.Join(path, configFileName)

	return initConfigFile(confFile)
}

/*******************
* ReadConfigFile
*******************/
func ReadConfigFile(brainConfig interfaces.I_BrainConfig, configFileName string, loadFunction func(*os.File, int64, interfaces.I_BrainConfig, uint32)) *os.File {
	path := brainConfig.GetSaveDirectory()
	err := os.MkdirAll(path, 0700)
	if err != nil {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_CREATE)
		os.Exit(errcode.ERROR_FATAL_CONFIG_OPEN)
	}

	confFile := filepath.Join(path, configFileName)

	// Open the configuration file
	fileHandle, err := os.OpenFile(confFile, os.O_RDWR, 0)
	if err != nil {
		return initConfigFile(confFile)
	}

	// Read file version
	var version uint32
	dataOffset, err := FileReadUint32(fileHandle, 0, &version)
	if err != nil {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_FATAL_CONFIG_READ)
		os.Exit(errcode.ERROR_FATAL_CONFIG_READ)
	}
	if version != 0x62726e01 {
		errcode.PrintMsgFromErrorCode(errcode.ERROR_CRITICAL_CONFIG_INVALID_DATA)
		os.Exit(errcode.ERROR_CRITICAL_CONFIG_INVALID_DATA)
	}

	loadFunction(fileHandle, dataOffset, brainConfig, version)

	return fileHandle
}

/*******************
* CloseFile
*******************/
func CloseFile(fileHandle *os.File) {
	if fileHandle != nil {
		fileHandle.Close()
	}
}
