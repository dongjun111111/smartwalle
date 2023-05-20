package alipay

import "fmt"

const (
	kSandboxGateway        = "https://openapi.alipaydev.com/gateway.do"
	kProductionGateway     = "https://openapi.alipay.com/gateway.do"
	kProductionMAPIGateway = "https://mapi.alipay.com/gateway.do"

	kFormat       = "JSON"
	kCharset      = "utf-8"
	kVersion      = "1.0"
	kSignTypeRSA2 = "RSA2"
	kContentType  = "application/x-www-form-urlencoded;charset=utf-8"
	kTimeFormat   = "2006-01-02 15:04:05"
)

const (
	kResponseSuffix    = "_response"
	kErrorResponse     = "error_response"
	kSignFieldName     = "sign"
	kSignTypeFieldName = "sign_type"
	kCertSNFieldName   = "alipay_cert_sn"
	kCertificateEnd    = "-----END CERTIFICATE-----"
)

// Code 支付宝接口响应 code https://doc.open.alipay.com/docs/doc.htm?treeId=291&articleId=105806&docType=1
type Code string

func (c Code) Successful() bool {
	return c == CodeSuccess
}

func (c Code) Failed() bool {
	return c != CodeSuccess
}

const (
	CodeSuccess          Code = "10000" // 接口调用成功
	CodeUnknowError      Code = "20000" // 服务不可用
	CodeInvalidAuthToken Code = "20001" // 授权权限不足
	CodeMissingParam     Code = "40001" // 缺少必选参数
	CodeInvalidParam     Code = "40002" // 非法的参数
	CodeBusinessFailed   Code = "40004" // 业务处理失败
	CodePermissionDenied Code = "40006" // 权限不足
)

type Param interface {
	// APIName 用于提供访问的 method
	APIName() string

	// Params 返回参数列表
	Params() map[string]string
}

type Error struct {
	Code    Code   `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

func (this *Error) Error() string {
	return fmt.Sprintf("%s - %s", this.Code, this.SubMsg)
}

func (this *Error) Successful() bool {
	return this.Code.Successful()
}

func (this *Error) Failed() bool {
	return this.Code.Failed()
}

const (
	kCertDownloadAPI = "alipay.open.app.alipaycert.download"
)

// CertDownload 应用支付宝公钥证书下载 https://opendocs.alipay.com/common/06ue2z
type CertDownload struct {
	AppAuthToken string `json:"-"`              // 可选
	AliPayCertSN string `json:"alipay_cert_sn"` // 支付宝公钥证书序列号
}

func (this CertDownload) APIName() string {
	return kCertDownloadAPI
}

func (this CertDownload) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

type CertDownloadRsp struct {
	Content struct {
		Code              Code   `json:"code"`
		Msg               string `json:"msg"`
		SubCode           string `json:"sub_code"`
		SubMsg            string `json:"sub_msg"`
		AliPayCertContent string `json:"alipay_cert_content"`
	} `json:"alipay_open_app_alipaycert_download_response"`
	Sign string `json:"sign"`
}
