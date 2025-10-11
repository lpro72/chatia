package utils

/*******************
* Import
*******************/
import (
	"encoding/binary"
	"io"
	"os"
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
	var length uint32
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
	var length uint32
	length = uint32(len(strBuffer))

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
		return 0, err
	}
	return offset + int64(length), nil
}

/*******************
* FileReadData
*******************/
func FileWriteData(file *os.File, offset int64, dataBuffer []byte) (int64, error) {
	var length uint32
	length = uint32(len(dataBuffer))

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
