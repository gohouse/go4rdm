package config

var ApiHost = "http://47.96.235.53:9393"
var (
	ApiSyncConfig        = ApiHost + "/Api/SyncConfig"
	ApiQaList            = ApiHost + "/Api/QaList"
	ApiQaReplyList       = ApiHost + "/Api/QaReplyList/%s"
	ApiQaAdd             = ApiHost + "/Api/QaAdd"
	ApiQaReply           = ApiHost + "/Api/QaReply"
	ApiUserLogin         = ApiHost + "/Api/UserLogin"
	ApiUserRegister      = ApiHost + "/Api/UserRegister"
	ApiUserPasswordReset = ApiHost + "/Api/UserPasswordReset"
	ApiChatHistory       = ApiHost + "/Api/ChatHistory"
	ApiChatUsers         = ApiHost + "/Api/Users"
	ApiVersion           = ApiHost + "/Api/Version"
	ApiStatistic         = ApiHost + "/Api/Statistic"
)

type ApiReturn struct {
	Code   int
	Msg    string
	Result interface{}
}
