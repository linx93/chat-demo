// Package result 这些内容最好封装成公共的pkg来引入使用
package result

// 自定义业务状态码
const (
	Success = 200000

	ParamErr  = 400000
	ParamMiss = 400001

	Fail = 500000
)

// StatusText 自定义业务状态码描述
func StatusText(code int) string {
	switch code {
	case Success:
		return "请求成功"
	case ParamErr:
		return "参数错误"
	case ParamMiss:
		return "缺失参数"
	case Fail:
		return "服务器内部错误"

	default:
		return ""
	}
}
