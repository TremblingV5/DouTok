// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package videoDomain

import (
	fmt "fmt"
	entity "github.com/TremblingV5/DouTok/kitex_gen/entity"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *DoutokGetFeedRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DoutokGetFeedRequest[number], err)
}

func (x *DoutokGetFeedRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.LatestTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DoutokGetFeedRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DoutokGetFeedResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DoutokGetFeedResponse[number], err)
}

func (x *DoutokGetFeedResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DoutokGetFeedResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DoutokGetFeedResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v entity.Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.VideoList = append(x.VideoList, &v)
	return offset, nil
}

func (x *DoutokGetFeedResponse) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.NextTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DoutokAddPublishRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DoutokAddPublishRequest[number], err)
}

func (x *DoutokAddPublishRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Data, offset, err = fastpb.ReadBytes(buf, _type)
	return offset, err
}

func (x *DoutokAddPublishRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DoutokAddPublishRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DoutokAddPublishRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DoutokAddPublishResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DoutokAddPublishResponse[number], err)
}

func (x *DoutokAddPublishResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DoutokAddPublishResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DoutokListPublishRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DoutokListPublishRequest[number], err)
}

func (x *DoutokListPublishRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DoutokListPublishResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DoutokListPublishResponse[number], err)
}

func (x *DoutokListPublishResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DoutokListPublishResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DoutokListPublishResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v entity.Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.VideoList = append(x.VideoList, &v)
	return offset, nil
}

func (x *DoutokCountPublishRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DoutokCountPublishRequest[number], err)
}

func (x *DoutokCountPublishRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v int64
			v, offset, err = fastpb.ReadInt64(buf, _type)
			if err != nil {
				return offset, err
			}
			x.UserId = append(x.UserId, v)
			return offset, err
		})
	return offset, err
}

func (x *DoutokCountPublishResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DoutokCountPublishResponse[number], err)
}

func (x *DoutokCountPublishResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DoutokCountPublishResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DoutokCountPublishResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	if x.Result == nil {
		x.Result = make(map[int64]int64)
	}
	var key int64
	var value int64
	offset, err = fastpb.ReadMapEntry(buf, _type,
		func(buf []byte, _type int8) (offset int, err error) {
			key, offset, err = fastpb.ReadInt64(buf, _type)
			return offset, err
		},
		func(buf []byte, _type int8) (offset int, err error) {
			value, offset, err = fastpb.ReadInt64(buf, _type)
			return offset, err
		})
	if err != nil {
		return offset, err
	}
	x.Result[key] = value
	return offset, nil
}

func (x *DoutokGetVideoInfoRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DoutokGetVideoInfoRequest[number], err)
}

func (x *DoutokGetVideoInfoRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.VideoId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DoutokGetFeedRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DoutokGetFeedRequest) fastWriteField1(buf []byte) (offset int) {
	if x.LatestTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.LatestTime)
	return offset
}

func (x *DoutokGetFeedRequest) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.UserId)
	return offset
}

func (x *DoutokGetFeedResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *DoutokGetFeedResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *DoutokGetFeedResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMsg)
	return offset
}

func (x *DoutokGetFeedResponse) fastWriteField3(buf []byte) (offset int) {
	if x.VideoList == nil {
		return offset
	}
	for i := range x.VideoList {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.VideoList[i])
	}
	return offset
}

func (x *DoutokGetFeedResponse) fastWriteField4(buf []byte) (offset int) {
	if x.NextTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.NextTime)
	return offset
}

func (x *DoutokAddPublishRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *DoutokAddPublishRequest) fastWriteField1(buf []byte) (offset int) {
	if len(x.Data) == 0 {
		return offset
	}
	offset += fastpb.WriteBytes(buf[offset:], 1, x.Data)
	return offset
}

func (x *DoutokAddPublishRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.Title)
	return offset
}

func (x *DoutokAddPublishRequest) fastWriteField3(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.UserId)
	return offset
}

func (x *DoutokAddPublishRequest) fastWriteField4(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.Name)
	return offset
}

func (x *DoutokAddPublishResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DoutokAddPublishResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *DoutokAddPublishResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMsg)
	return offset
}

func (x *DoutokListPublishRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DoutokListPublishRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *DoutokListPublishResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *DoutokListPublishResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *DoutokListPublishResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMsg)
	return offset
}

func (x *DoutokListPublishResponse) fastWriteField3(buf []byte) (offset int) {
	if x.VideoList == nil {
		return offset
	}
	for i := range x.VideoList {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.VideoList[i])
	}
	return offset
}

func (x *DoutokCountPublishRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DoutokCountPublishRequest) fastWriteField1(buf []byte) (offset int) {
	if len(x.UserId) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 1, len(x.UserId),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteInt64(buf[offset:], numTagOrKey, x.UserId[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *DoutokCountPublishResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *DoutokCountPublishResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *DoutokCountPublishResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMsg)
	return offset
}

func (x *DoutokCountPublishResponse) fastWriteField3(buf []byte) (offset int) {
	if x.Result == nil {
		return offset
	}
	for k, v := range x.Result {
		offset += fastpb.WriteMapEntry(buf[offset:], 3,
			func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
				offset := 0
				offset += fastpb.WriteInt64(buf[offset:], numTagOrKey, k)
				offset += fastpb.WriteInt64(buf[offset:], numIdxOrVal, v)
				return offset
			})
	}
	return offset
}

func (x *DoutokGetVideoInfoRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DoutokGetVideoInfoRequest) fastWriteField1(buf []byte) (offset int) {
	if x.VideoId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.VideoId)
	return offset
}

func (x *DoutokGetFeedRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DoutokGetFeedRequest) sizeField1() (n int) {
	if x.LatestTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.LatestTime)
	return n
}

func (x *DoutokGetFeedRequest) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.UserId)
	return n
}

func (x *DoutokGetFeedResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *DoutokGetFeedResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *DoutokGetFeedResponse) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMsg)
	return n
}

func (x *DoutokGetFeedResponse) sizeField3() (n int) {
	if x.VideoList == nil {
		return n
	}
	for i := range x.VideoList {
		n += fastpb.SizeMessage(3, x.VideoList[i])
	}
	return n
}

func (x *DoutokGetFeedResponse) sizeField4() (n int) {
	if x.NextTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.NextTime)
	return n
}

func (x *DoutokAddPublishRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *DoutokAddPublishRequest) sizeField1() (n int) {
	if len(x.Data) == 0 {
		return n
	}
	n += fastpb.SizeBytes(1, x.Data)
	return n
}

func (x *DoutokAddPublishRequest) sizeField2() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(2, x.Title)
	return n
}

func (x *DoutokAddPublishRequest) sizeField3() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.UserId)
	return n
}

func (x *DoutokAddPublishRequest) sizeField4() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(4, x.Name)
	return n
}

func (x *DoutokAddPublishResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DoutokAddPublishResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *DoutokAddPublishResponse) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMsg)
	return n
}

func (x *DoutokListPublishRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DoutokListPublishRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *DoutokListPublishResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *DoutokListPublishResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *DoutokListPublishResponse) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMsg)
	return n
}

func (x *DoutokListPublishResponse) sizeField3() (n int) {
	if x.VideoList == nil {
		return n
	}
	for i := range x.VideoList {
		n += fastpb.SizeMessage(3, x.VideoList[i])
	}
	return n
}

func (x *DoutokCountPublishRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DoutokCountPublishRequest) sizeField1() (n int) {
	if len(x.UserId) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(1, len(x.UserId),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeInt64(numTagOrKey, x.UserId[numIdxOrVal])
			return n
		})
	return n
}

func (x *DoutokCountPublishResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *DoutokCountPublishResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *DoutokCountPublishResponse) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMsg)
	return n
}

func (x *DoutokCountPublishResponse) sizeField3() (n int) {
	if x.Result == nil {
		return n
	}
	for k, v := range x.Result {
		n += fastpb.SizeMapEntry(3,
			func(numTagOrKey, numIdxOrVal int32) int {
				n := 0
				n += fastpb.SizeInt64(numTagOrKey, k)
				n += fastpb.SizeInt64(numIdxOrVal, v)
				return n
			})
	}
	return n
}

func (x *DoutokGetVideoInfoRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DoutokGetVideoInfoRequest) sizeField1() (n int) {
	if x.VideoId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.VideoId)
	return n
}

var fieldIDToName_DoutokGetFeedRequest = map[int32]string{
	1: "LatestTime",
	2: "UserId",
}

var fieldIDToName_DoutokGetFeedResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "VideoList",
	4: "NextTime",
}

var fieldIDToName_DoutokAddPublishRequest = map[int32]string{
	1: "Data",
	2: "Title",
	3: "UserId",
	4: "Name",
}

var fieldIDToName_DoutokAddPublishResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
}

var fieldIDToName_DoutokListPublishRequest = map[int32]string{
	1: "UserId",
}

var fieldIDToName_DoutokListPublishResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "VideoList",
}

var fieldIDToName_DoutokCountPublishRequest = map[int32]string{
	1: "UserId",
}

var fieldIDToName_DoutokCountPublishResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "Result",
}

var fieldIDToName_DoutokGetVideoInfoRequest = map[int32]string{
	1: "VideoId",
}

var _ = entity.File_entity_proto
