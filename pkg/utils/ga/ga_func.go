package ga

import (
	"context"
	"fmt"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/empty"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gcache"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gmd5"
	"time"
)

var (
	cache = gcache.New()
)

// Md5 encryption
func Md5(str string) string {
	mdsecret, _ := gmd5.Encrypt(str)
	return mdsecret
}

// md5hex编码加密
func Md5Str(origin string) string {
	return gmd5.Md5StrHex(origin)
}

// 把验证码保存在本地，用GetVerifyCode获取key对应缓存
func SetVerifyCode(key, code string) (err error) {
	ctx := context.Background()
	err = cache.Set(ctx, key, code, time.Second*60)
	return
}

// 获取本地保存的验证码，使用SetVerifyCode保存可以对应数据
func GetVerifyCode(key string) (code int, err error) {
	ctx := context.Background()
	val, err := cache.Get(ctx, key)
	if err == nil {
		code = Int(val)
	}
	return
}

// IsNil checks whether given `value` is nil.
func IsNil(value interface{}, traceSource ...bool) bool {
	return empty.IsNil(value, traceSource...)
}

// IsEmpty checks whether given `value` empty.
// It returns true if `value` is in: 0, nil, false, "", len(slice/map/chan) == 0.
// Or else it returns true.
func IsEmpty(value interface{}, traceSource ...bool) bool {
	return empty.IsEmpty(value, traceSource...)
}

// 数据格式化成【】string
func FormatColumnData(data interface{}) []interface{} {
	if data == nil {
		return []interface{}{}
	}
	// 安全类型转换
	switch v := data.(type) {
	case []interface{}:
		column := make([]interface{}, len(v))
		for i, item := range v {
			// 处理每个元素的类型转换
			switch val := item.(type) {
			case string:
				column[i] = val
			case int, int8, int16, int32, int64:
				column[i] = fmt.Sprintf("%d", val)
			case float32, float64:
				column[i] = fmt.Sprintf("%v", val)
			default:
				column[i] = fmt.Sprintf("%v", val)
			}
		}
		return column
	case []string:
		column := make([]interface{}, len(v))
		for i, val := range v {
			column[i] = val
		}
		return column
	case []int:
		column := make([]interface{}, len(v))
		for i, val := range v {
			column[i] = fmt.Sprintf("%d", val)
		}
		return column
	default:
		// 对于其他类型，尝试转换为字符串数组
		return []interface{}{fmt.Sprintf("%v", data)}
	}
}
