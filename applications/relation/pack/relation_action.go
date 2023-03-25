package pack

import "github.com/TremblingV5/DouTok/kitex_gen/relation"

func PackageRelationActionResponse(code int32, msg string, err error) (*relation.DouyinRelationActionResponse, error) {
	if err != nil {
		return nil, err
	}

	return &relation.DouyinRelationActionResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}, err
}
