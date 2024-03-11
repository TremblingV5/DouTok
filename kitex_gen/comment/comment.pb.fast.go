// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package comment

import (
	fmt "fmt"
	entity "github.com/TremblingV5/DouTok/kitex_gen/entity"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *DouyinCommentActionRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 5:
		offset, err = x.fastReadField5(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinCommentActionRequest[number], err)
}

func (x *DouyinCommentActionRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinCommentActionRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.VideoId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinCommentActionRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ActionType, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinCommentActionRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.CommentText, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DouyinCommentActionRequest) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.CommentId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinCommentActionResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinCommentActionResponse[number], err)
}

func (x *DouyinCommentActionResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinCommentActionResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DouyinCommentActionResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v entity.Comment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Comment = &v
	return offset, nil
}

func (x *DouyinCommentListRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinCommentListRequest[number], err)
}

func (x *DouyinCommentListRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.VideoId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DouyinCommentListResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinCommentListResponse[number], err)
}

func (x *DouyinCommentListResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinCommentListResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DouyinCommentListResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v entity.Comment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.CommentList = append(x.CommentList, &v)
	return offset, nil
}

func (x *DouyinCommentCountRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinCommentCountRequest[number], err)
}

func (x *DouyinCommentCountRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v int64
			v, offset, err = fastpb.ReadInt64(buf, _type)
			if err != nil {
				return offset, err
			}
			x.VideoIdList = append(x.VideoIdList, v)
			return offset, err
		})
	return offset, err
}

func (x *DouyinCommentCountResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DouyinCommentCountResponse[number], err)
}

func (x *DouyinCommentCountResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DouyinCommentCountResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DouyinCommentCountResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
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

func (x *DouyinCommentActionRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *DouyinCommentActionRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *DouyinCommentActionRequest) fastWriteField2(buf []byte) (offset int) {
	if x.VideoId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.VideoId)
	return offset
}

func (x *DouyinCommentActionRequest) fastWriteField3(buf []byte) (offset int) {
	if x.ActionType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.ActionType)
	return offset
}

func (x *DouyinCommentActionRequest) fastWriteField4(buf []byte) (offset int) {
	if x.CommentText == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.CommentText)
	return offset
}

func (x *DouyinCommentActionRequest) fastWriteField5(buf []byte) (offset int) {
	if x.CommentId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.CommentId)
	return offset
}

func (x *DouyinCommentActionResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *DouyinCommentActionResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *DouyinCommentActionResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMsg)
	return offset
}

func (x *DouyinCommentActionResponse) fastWriteField3(buf []byte) (offset int) {
	if x.Comment == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.Comment)
	return offset
}

func (x *DouyinCommentListRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DouyinCommentListRequest) fastWriteField1(buf []byte) (offset int) {
	if x.VideoId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.VideoId)
	return offset
}

func (x *DouyinCommentListResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *DouyinCommentListResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *DouyinCommentListResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMsg)
	return offset
}

func (x *DouyinCommentListResponse) fastWriteField3(buf []byte) (offset int) {
	if x.CommentList == nil {
		return offset
	}
	for i := range x.CommentList {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.CommentList[i])
	}
	return offset
}

func (x *DouyinCommentCountRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DouyinCommentCountRequest) fastWriteField1(buf []byte) (offset int) {
	if len(x.VideoIdList) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 1, len(x.VideoIdList),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteInt64(buf[offset:], numTagOrKey, x.VideoIdList[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *DouyinCommentCountResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *DouyinCommentCountResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *DouyinCommentCountResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMsg)
	return offset
}

func (x *DouyinCommentCountResponse) fastWriteField3(buf []byte) (offset int) {
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

func (x *DouyinCommentActionRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *DouyinCommentActionRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *DouyinCommentActionRequest) sizeField2() (n int) {
	if x.VideoId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.VideoId)
	return n
}

func (x *DouyinCommentActionRequest) sizeField3() (n int) {
	if x.ActionType == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.ActionType)
	return n
}

func (x *DouyinCommentActionRequest) sizeField4() (n int) {
	if x.CommentText == "" {
		return n
	}
	n += fastpb.SizeString(4, x.CommentText)
	return n
}

func (x *DouyinCommentActionRequest) sizeField5() (n int) {
	if x.CommentId == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.CommentId)
	return n
}

func (x *DouyinCommentActionResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *DouyinCommentActionResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *DouyinCommentActionResponse) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMsg)
	return n
}

func (x *DouyinCommentActionResponse) sizeField3() (n int) {
	if x.Comment == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.Comment)
	return n
}

func (x *DouyinCommentListRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DouyinCommentListRequest) sizeField1() (n int) {
	if x.VideoId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.VideoId)
	return n
}

func (x *DouyinCommentListResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *DouyinCommentListResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *DouyinCommentListResponse) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMsg)
	return n
}

func (x *DouyinCommentListResponse) sizeField3() (n int) {
	if x.CommentList == nil {
		return n
	}
	for i := range x.CommentList {
		n += fastpb.SizeMessage(3, x.CommentList[i])
	}
	return n
}

func (x *DouyinCommentCountRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DouyinCommentCountRequest) sizeField1() (n int) {
	if len(x.VideoIdList) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(1, len(x.VideoIdList),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeInt64(numTagOrKey, x.VideoIdList[numIdxOrVal])
			return n
		})
	return n
}

func (x *DouyinCommentCountResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *DouyinCommentCountResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *DouyinCommentCountResponse) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMsg)
	return n
}

func (x *DouyinCommentCountResponse) sizeField3() (n int) {
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

var fieldIDToName_DouyinCommentActionRequest = map[int32]string{
	1: "UserId",
	2: "VideoId",
	3: "ActionType",
	4: "CommentText",
	5: "CommentId",
}

var fieldIDToName_DouyinCommentActionResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "Comment",
}

var fieldIDToName_DouyinCommentListRequest = map[int32]string{
	1: "VideoId",
}

var fieldIDToName_DouyinCommentListResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "CommentList",
}

var fieldIDToName_DouyinCommentCountRequest = map[int32]string{
	1: "VideoIdList",
}

var fieldIDToName_DouyinCommentCountResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "Result",
}

var _ = entity.File_entity_proto