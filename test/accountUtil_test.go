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
package test

import (
	"testing"
	"time"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/mind-stack-cn/golang_netease_im_sdk/util"
)

func init()  {
	util.Init("#YOUR APP KEY", "#YOUR APP SEC KEY")
}

func Test_TimeFormat(t *testing.T)  {
	fmt.Println(time.Now().Unix())
}

func Test_CreateAccid(t *testing.T)  {
	res, err := util.CreateAccid("lisi", "", "", "", "")
	assert.Nil(t, err, "error showedup while creating account id")
	if err == nil{
		fmt.Println(res)
	}else{
		fmt.Println(err.Error())
	}
}

func Test_RefreshToken(t *testing.T)  {
	res, err := util.RefreshToken("lisi")
	assert.Nil(t, err, "error showedup while refresh account id")
	if err == nil{
		fmt.Println(res)
	}else{
		fmt.Println(err.Error())
	}
}

func Test_BlockAccid(t *testing.T)  {
	res, err := util.Block("lisi")
	assert.Nil(t, err, "error showedup while blocking account id")
	if err == nil{
		fmt.Println(res)
	}else{
		fmt.Println(err.Error())
	}
}

func Test_UnBlockAccid(t *testing.T)  {
	res, err := util.Unblock("lisi")
	assert.Nil(t, err, "error showedup while unblocking account id")
	if err == nil{
		fmt.Println(res)
	}else{
		fmt.Println(err.Error())
	}
}
