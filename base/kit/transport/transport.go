package transport

import (
	"context"
)

func Decode(ctx context.Context, input interface{}) (output interface{}, err error) {
	return input, nil
	// switch out := input.(type) {
	// case (*pb.Info):
	// 	return out, nil
	// case (*pb.Resp):
	// 	return out, nil
	// }
	// return nil, fmt.Errorf("unknow type")
}
func Encode(ctx context.Context, input interface{}) (output interface{}, err error) {
	return input, nil
	// switch out := input.(type) {
	// case (*pb.Info):
	// 	return out, nil
	// case (*pb.Resp):
	// 	return out, nil
	// }
	// return nil, fmt.Errorf("unknow type")
}
