/* 
 * The MIT License (MIT)
 *
 * Copyright (c) 2016 tony<wuhaiyang1213@gmail.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.

 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package nimSdk

import (
	"net/url"
	"fmt"
	"encoding/json"
)

var CREATE_ACCOUNT_ID_URL = UrlPair{"POST", NETEASE_BASE_URL + "/user/create.action"}

var UPDATE_ACCOUNT_ID_URL = UrlPair{"POST", NETEASE_BASE_URL + "/user/update.action"}

var REFRESH_TOKEN_URL = UrlPair{"POST", NETEASE_BASE_URL + "/user/refreshToken.action"}

var BLOCK_URL = UrlPair{"POST", NETEASE_BASE_URL + "/user/block.action"}

var UNBLOCK_URL = UrlPair{"POST", NETEASE_BASE_URL + "/user/unblock.action"}

type NIMUInfo struct {
	Accid   string  //云信ID
	Name    string  //云信ID昵称，最大长度64字节，用来PUSH推送时显示的昵称
	Icon    string  //云信ID头像URL，第三方可选填，最大长度1024
	Token   string  //token
}

// 创建云信ID
// 第三方帐号导入到云信平台
// accId string 云信ID，最大长度32字节，必须保证一个 APP内唯一（只允许字母、数字、半角下划线_、@、半角点以及半角-组成， 不区分大小写， 会统一小写处理，请注意以此接口返回结果中的accid为准）
// name string 云信ID昵称，最大长度64字节，用来PUSH推送 时显示的昵称
// props string json属性，第三方可选填，最大长度1024字节
// icon string 云信ID头像URL，第三方可选填，最大长度1024
// token string 云信ID可以指定登录token值，最大长度128字节， 并更新，如果未指定，会自动生成token，并在创建成功后返回
func CreateAccid(accid string, name string, props string, icon string, token string) (*NIMUInfo, error) {
	// format request body
	v := url.Values{}
	if len(accid) <= 0 {
		return nil, fmt.Errorf("accid is empty")
	}
	v.Add("accid", accid);
	if len(name) > 0 {
		v.Add("name", name)
	}
	if len(props) > 0 {
		v.Add("props", props)
	}
	if len(icon) > 0 {
		v.Add("icon", icon)
	}
	if len(token) > 0 {
		v.Add("token", token)
	}

	resStr,err := DoNeteaseHttpRequest(v, CREATE_ACCOUNT_ID_URL.method, CREATE_ACCOUNT_ID_URL.url)

	if err != nil {
		return nil,err
	}

	var resp NIMUInfo
	err = json.Unmarshal([]byte(resStr), &resp)
	if err != nil {
		return nil,err
	}

	return &resp,nil
}

// 云信ID更新
// 云信ID基本信息更新
// accId string 云信ID，最大长度32字节，必须保证一个 APP内唯一（只允许字母、数字、半角下划线_、@、半角点以及半角-组成， 不区分大小写， 会统一小写处理，请注意以此接口返回结果中的accid为准）
// name string 云信ID昵称，最大长度64字节，用来PUSH推送 时显示的昵称
// props string json属性，第三方可选填，最大长度1024字节
// icon string 云信ID头像URL，第三方可选填，最大长度1024
// token string 云信ID可以指定登录token值，最大长度128字节， 并更新，如果未指定，会自动生成token，并在创建成功后返回
func UpdateAccid(accid string, name string, props string, icon string, token string) (string, error) {
	// format request body
	v := url.Values{}
	if len(accid) <= 0 {
		return "", fmt.Errorf("accid is empty")
	}
	v.Add("accid", accid);
	if len(name) > 0 {
		v.Add("name", name)
	}
	if len(props) > 0 {
		v.Add("props", props)
	}
	if len(icon) > 0 {
		v.Add("icon", icon)
	}
	if len(token) > 0 {
		v.Add("token", token)
	}

	return DoNeteaseHttpRequest(v, UPDATE_ACCOUNT_ID_URL.method, UPDATE_ACCOUNT_ID_URL.url)
}

// 更新并获取新token
// WebServer更新云信ID的token，同时返回新的token
// accId string 云信ID，最大长度32字节，必须保证一个 APP内唯一
func RefreshToken(accid string) (string, error) {
	// format request body
	v := url.Values{}
	if len(accid) <= 0 {
		return "", fmt.Errorf("accid is empty")
	}
	v.Add("accid", accid);
	return DoNeteaseHttpRequest(v, REFRESH_TOKEN_URL.method, REFRESH_TOKEN_URL.url)
}

// 封禁云信ID
// 第三方禁用某个云信ID的IM功能,封禁云信ID后，此ID将不能登陆云信imserver
// accId string 云信ID，最大长度32字节，必须保证一个 APP内唯一
func Block(accid string) (string, error) {
	// format request body
	v := url.Values{}
	if len(accid) <= 0 {
		return "", fmt.Errorf("accid is empty")
	}
	v.Add("accid", accid);
	return DoNeteaseHttpRequest(v, BLOCK_URL.method, BLOCK_URL.url)
}

// 解禁云信ID
// 解禁被封禁的云信ID
// accId string 云信ID，最大长度32字节，必须保证一个 APP内唯一
func Unblock(accid string) (string, error) {
	// format request body
	v := url.Values{}
	if len(accid) <= 0 {
		return "", fmt.Errorf("accid is empty")
	}
	v.Add("accid", accid);
	return DoNeteaseHttpRequest(v, UNBLOCK_URL.method, UNBLOCK_URL.url)
}
