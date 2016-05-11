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
	"net/http"
	"strings"
	"time"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"strconv"
	"net/url"
)

const NETEASE_BASE_URL = "https://api.netease.im/nimserver"

var APPKEY = "94kid09c9ig9k1loimjg012345123456"

var APPSECRETKEY = "123456789012"

type UrlPair struct {
	method, url string
}

func Init(appKey string, appSecretKey string){
	APPKEY = appKey
	APPSECRETKEY = appSecretKey
}

// 执行云信http请求
func DoNeteaseHttpRequest(v url.Values, reqMethod string, reqUrl string)(string, error){
	req, _ := http.NewRequest(reqMethod, reqUrl, strings.NewReader(v.Encode()))

	// add checksum
	AddNeteaseHttpHeader(req)

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err;
	}

	// read response
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err;
	}

	return string(respBody), nil;
}


// 在HttpHeader中添加参数用于校验
func AddNeteaseHttpHeader(req *http.Request)  {
	// nonce & curTime in second
	var nonce string = strings.Replace(uuid.NewV4().String(), "-", "", -1)
	var curTime string = strconv.FormatInt(time.Now().Unix(), 10)

	// add request header
	req.Header.Add("AppKey", APPKEY);
	req.Header.Add("Nonce", nonce);
	req.Header.Add("CurTime", curTime);
	req.Header.Add("CheckSum", GetCheckSum(APPSECRETKEY, nonce, curTime));
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8");
}
