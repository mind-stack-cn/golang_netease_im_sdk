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
package util

import (
	"fmt"
	"net/url"
	"encoding/json"
)

var UPDATE_PROFILE_URL = UrlPair{"POST", NETEASE_BASE_URL + "/user/updateUinfo.action"}

var GET_PROFILE_URL = UrlPair{"POST", NETEASE_BASE_URL + "/user/getUinfos.action"}


// 更新用户名片
// accId string 云信ID，最大长度32字节，必须保证一个 APP内唯一
// name string 用户昵称，最大长度64字节
// icon string 用户icon，最大长度1024字节
// sign string 用户签名，最大长度256字节
// email string 用户email，最大长度64字节
// birth string 用户生日，最大长度16字节
// mobile string 用户mobile，最大长度32字节
// gender string 用户性别，0表示未知，1表示男，2女表示女，其它会报参数错误
// ex string 用户名片扩展字段，最大长度1024字节，用户可自行扩展，建议封装成JSON字符串
func UpdateProfile(accid string, name string, icon string, sign string, email string, birth string, mobile string, gender string, ex string) (string, error) {
	// format request body
	v := url.Values{}
	if len(accid) <= 0 {
		return "", fmt.Errorf("accid is empty")
	}
	v.Add("accid", accid);
	if len(name) > 0 {
		v.Add("name", name)
	}
	if len(icon) > 0 {
		v.Add("icon", icon)
	}
	if len(sign) > 0 {
		v.Add("token", sign)
	}
	if len(email) > 0 {
		v.Add("email", email)
	}
	if len(birth) > 0 {
		v.Add("birth", birth)
	}
	if len(mobile) > 0 {
		v.Add("mobile", mobile)
	}
	if len(gender) > 0 {
		v.Add("gender", gender)
	}
	if len(ex) > 0 {
		v.Add("ex", ex)
	}
	return DoNeteaseHttpRequest(v, UPDATE_PROFILE_URL.method, UPDATE_PROFILE_URL.url)
}


// 获取用户名片
// 获取用户名片，可批量
// accId string 云信ID，最大长度32字节，必须保证一个 APP内唯一
func GetProfile(accids []string) (string, error){
	// format request body
	v := url.Values{}
	if len(accids) <= 0 {
		return "", fmt.Errorf("accids is empty")
	}
	jsonVal, _ := json.Marshal(accids)
	v.Add("accids", string(jsonVal));
	return DoNeteaseHttpRequest(v, GET_PROFILE_URL.method, GET_PROFILE_URL.url)
}
