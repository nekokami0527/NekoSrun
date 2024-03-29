package nekosrun

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
)

var srun_error_info = map[string]string{
	"E0000": "登录成功",
	"E2401": "User-Request",
	"E2402": "Lost-Carrier",
	"E2404": "Idle-Timeout",
	"E2405": "Session-Timeout",
	"E2406": "Admin-Reset",
	"E2407": "Admin-Reboot",
	"E2408": "Port-Error",
	"E2409": "NAS-Error",
	"E2410": "NAS-Request",
	"E2411": "NAS-Reboot",
	"E2412": "Port-Unneeded",
	"E2413": "Port-Preempted",
	"E2414": "Port-Suspended",
	"E2415": "Service-Unavailable",
	"E2416": "Callback",
	"E2417": "User-Error",
	"E2531": "用户不存在",
	"E2532": "您的两次认证的间隔太短,请稍候10秒后再重试登录",
	"E2533": "密码错误次数超过限制，请5分钟后再重试登录",
	"E2534": "有代理行为被暂时禁用",
	"E2535": "认证系统已经被禁用",
	"E2536": "授权已过期",
	"E2553": "帐号或密码错误",
	"E2601": "该区域仅允许手机登录。",
	"E2602": "您还没有绑定手机号或绑定的非联通手机号码",
	"E2606": "用户被禁用",
	"E2607": "接口被禁用",
	"E2611": "您当前使用的设备非该账号绑定设备 请绑定或使用绑定的设备登入",
	"E2613": "NAS PORT绑定错误",
	"E2614": "MAC地址绑定错误",
	"E2615": "IP地址绑定错误",
	"E2616": "用户已欠费",
	"E2620": "已经在线了",
	"E2621": "已经达到授权人数",
	"E2806": "找不到符合条件的产品",
	"E2807": "找不到符合条件的计费策略",
	"E2808": "找不到符合条件的控制策略",
	"E2833": "IP不在DHCP表中，需要重新拿地址。",
	"E2840": "校内地址不允许访问外网",
	"E2841": "IP地址绑定错误",
	"E2842": "IP地址无需认证可直接上网",
	"E2843": "IP地址不在IP表中",
	"E2844": "IP地址在黑名单中",
	"E2901": "密码错误",
	"E6500": "认证程序未启动",
	"E6501": "用户名输入错误",
	"E6502": "注销时发生错误，或没有帐号在线",
	"E6503": "您的账号不在线上",
	"E6504": "注销成功，请等1分钟后登录",
	"E6505": "您的MAC地址不正确",
	"E6506": "用户名或密码错误，请重新输入",
	"E6507": "您无须认证，可直接上网",
	"E6508": "您已欠费，请尽快充值",
	"E6509": "您的资料已被修改正在等待同步，请2钟分后再试。如果您的帐号允许多个用户上线，请到WEB登录页面注销",
	"E6510": "您的帐号已经被删除",
	"E6511": "IP已存在，请稍后再试",
	"E6512": "在线用户已满，请稍后再试",
	"E6513": "正在注销在线账号，请重新连接",
	"E6514": "你的IP地址和认证地址不附，可能是经过小路由器登录的",
	"E6515": "系统已禁止客户端登录，请使用WEB方式登录",
	"E6516": "您的流量已用尽",
	"E6517": "您的时长已用尽",
	"E6518": "您的IP地址不合法，可能是：一、与绑的IP地址附；二、IP不允许在当前区域登录",
	"E6519": "当前时段不允许连接",
	"E6520": "抱歉，您的帐号已禁用",
	"E6521": "您的IPv6地址不正确，请重新配置IPv6地址",
	"E6522": "客户端时间不正确，请先同步时间（或者是调用方传送的时间格式不正确，不是时间戳；客户端和服务器之间时差超过2小时，括号里面内容不要提示给客户）",
	"E6523": "认证服务无响应",
	"E6524": "计费系统尚未授权，目前还不能使用",
	"E6525": "后台服务器无响应;请联系管理员检查后台服务运行状态",
	"E6526": "您的IP已经在线;可以直接上网;或者先注销再重新认证",
	"E6527": "当前设备不在线",
	"E6528": "您已经被服务器强制下线",
	"E6529": "身份验证失败，但不返回错误消息",
	"0":     "本机IP已经使用其他账号登陆在线了",
}

type SrunLoginConfig struct {
	_Alpha     string
	_acid      string
	_enc       string
	_n         string
	_types     string
	_version   string
	_server_ip string
	base64     *base64.Encoding
}

type SrunUserInfo struct {
	username       string
	password       string
	user_ip        string
	real_name      string
	wallet_balance float64
}

type SrunLogin struct {
	_user   SrunUserInfo
	_config SrunLoginConfig
}

func New(server string) SrunLogin {
	config := SrunLoginConfig{
		_Alpha:     "LVoJPiCN2R8G90yg+hmFHuacZ1OWMnrsSTXkYpUq/3dlbfKwv6xztjI7DeBE45QA",
		_enc:       "srun_bx1",
		_n:         "200",
		_types:     "1",
		_server_ip: server,
		_acid:      "4",
	}
	user := SrunUserInfo{user_ip: "218.195.217.49"}
	srun := SrunLogin{_config: config, _user: user}
	return srun
}

func (srun *SrunLogin) _get_challenge() string {
	addr := fmt.Sprintf(
		"http://%s/cgi-bin/get_challenge",
		srun._config._server_ip,
	)
	content := http_get(addr, &map[string]interface{}{
		"callback": "neko",
		"username": srun._user.username,
		"ip":       srun._user.user_ip,
	})
	content_json := loadJsonString(content[5 : len(content)-1])
	srun._config._version = content_json.getString("srun_ver")
	srun._user.user_ip = content_json.getString("online_ip")
	log.Println("登录用户IP", srun._user.user_ip)
	log.Println("认证系统版本", srun._config._version)
	log.Println("获取认证挑战码")
	return content_json.getString("challenge")
}

func (srun *SrunLogin) _info(dataIn map[string]interface{}, key string) string {
	data := json_serialize(dataIn, false, &[]string{"username", "password", "ip", "acid", "enc_ver"})
	return "{SRBX1}" + srun._config.base64.EncodeToString(
		[]byte(xEncode(data, key)))
}

func (srun *SrunLogin) _srunPortal(params map[string]interface{}) *JSONObject {
	url := fmt.Sprintf("http://%s/cgi-bin/srun_portal", srun._config._server_ip)
	response := http_get(url, &params)
	json := loadJsonString(response[5 : len(response)-1])
	return json
}

func (srun *SrunLogin) SetClientIp(clientIp string) {
	srun._user.user_ip = clientIp
	log.Println("设置认证IP:", clientIp)
}

func (srun *SrunLogin) Login(username string, password string) {
	srun._user.username = username
	srun._user.password = password

	srun._config.base64 = base64.NewEncoding(srun._config._Alpha)
	log.Println("登录用户名", username)
	challenge_string := srun._get_challenge()
	info_data := srun._info(
		map[string]interface{}{
			"username": username,
			"password": password,
			"ip":       srun._user.user_ip,
			"acid":     srun._config._acid,
			"enc_ver":  srun._config._enc,
		}, challenge_string,
	)

	log.Println("生成信息摘要")

	password_hmac := HASH_HMAC([]uint8(password), []uint8(challenge_string))

	log.Println("生成密码HMAC")

	var checkStr bytes.Buffer
	checkStr.Write([]uint8(challenge_string))
	checkStr.Write([]uint8(username))
	checkStr.Write([]uint8(challenge_string))
	checkStr.Write([]uint8(password_hmac))
	checkStr.Write([]uint8(challenge_string))
	checkStr.Write([]uint8(srun._config._acid))
	checkStr.Write([]uint8(challenge_string))
	checkStr.Write([]uint8(srun._user.user_ip))
	checkStr.Write([]uint8(challenge_string))
	checkStr.Write([]uint8(srun._config._n))
	checkStr.Write([]uint8(challenge_string))
	checkStr.Write([]uint8(srun._config._types))
	checkStr.Write([]uint8(challenge_string))
	checkStr.Write([]uint8(info_data))
	login_status := srun._srunPortal(map[string]interface{}{
		"callback":     "neko",
		"action":       "login",
		"username":     username,
		"password":     "{MD5}" + password_hmac,
		"ac_id":        srun._config._acid,
		"ip":           srun._user.user_ip,
		"chksum":       SHA1(checkStr.Bytes()),
		"info":         info_data,
		"n":            srun._config._n,
		"type":         srun._config._types,
		"os":           "Windows 10",
		"name":         "Windows",
		"double_stack": "1",
	})
	error_str := login_status.getString("error")

	if error_str != "ok" {
		log.Println(login_status)
		if login_status.exist("ecode") {
			log.Fatalln("登陆失败", srun_error_info[login_status.getString("ecode")])
		} else {
			log.Fatalln("登陆失败", error_str)

		}
	}
	log.Println("登陆成功")

	srun._user.real_name = login_status.getString("real_name")
	srun._user.wallet_balance = login_status.getFloat("wallet_balance:0")
	log.Println("姓名:", srun._user.real_name)
	log.Println("学号:", srun._user.username)
	log.Println("设备IP:", srun._user.user_ip)
	log.Println("余额:", srun._user.wallet_balance)
}

func (srun *SrunLogin) Logout() {
	//?callback=neko&action=logout&ac_id=%s&ip=%s&username=%s"
	url := fmt.Sprintf("http://%s/cgi-bin/srun_portal", srun._config._server_ip)
	response := http_get(url, &map[string]interface{}{
		"callback": "neko",
		"action":   "logout",
		"ac_id":    srun._config._acid,
		"ip":       srun._user.user_ip,
		"username": srun._user.username,
	})
	json := loadJsonString(response[5 : len(response)-1])
	log.Println(json)
}

func (srun *SrunLogin) LogoutTest(username string, ip string) {
	//?callback=neko&action=logout&ac_id=%s&ip=%s&username=%s"
	url := fmt.Sprintf("http://%s/cgi-bin/srun_portal", srun._config._server_ip)
	response := http_get(url, &map[string]interface{}{
		"callback": "neko",
		"action":   "logout",
		"ac_id":    srun._config._acid,
		"ip":       ip,
		"username": username,
	})
	json := loadJsonString(response[5 : len(response)-1])
	log.Println(json)
}
