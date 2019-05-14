package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil" //处理http获取的数据转换用
	"mime/multipart"
	"net"
	"net/http" //http 请求用
	"net/url"  //HttpPost3中定义url.Values中使用到
	"os"
	"strings" //HttpPost2中生成Post数据用
	"time"
)

const baseUrl  = "https://msp-st1.uat.cmrh.com/RH_MSPSERVER"

// Get 请求
func GetTest(){

	resp, err := http.Get(baseUrl + "/bankInsurance/activity/getActivityList?agentNo=B000000000000238&type=01&page=1&rows=10")

	if err != nil {
		// handle error
		fmt.Println(err.Error())
		return
	}

	defer resp.Body.Close()					// error corresponding to variable 'resp' may be not nil --> 与变量“RESP”对应的错误可能不为零。

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(body))
}

//Post 请求
func PostTest(){

	resp, err := http.Post(baseUrl + "/bankInsurance/activity/getActivityList", "application/x-www-form-urlencoded", strings.NewReader("agentNo=B000000000000238&type=01&page=1&rows=10"))
	if err != nil {
		// handle error
		fmt.Println(err.Error())
		return
	}

	defer resp.Body.Close()					// error corresponding to variable 'resp' may be not nil --> 与变量“RESP”对应的错误可能不为零。
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(body))
}

//Post 表单提交	PostForm实现了编码为application/x-www-form-urlencoded的表单提交。		https://www.cnblogs.com/276815076/p/7741870.html
func PostFormTest(){

	postValue := url.Values{
		"agentNo": {"B000000000000238"},
		"type": {"01"},
		"page": {"1"},
		"rows": {"10"},
	}

	resp, err := http.PostForm(baseUrl + "/bankInsurance/activity/getActivityList", postValue)
	if err != nil {
		// handle error
		fmt.Println(err.Error())
		return
	}

	defer resp.Body.Close()					// error corresponding to variable 'resp' may be not nil --> 与变量“RESP”对应的错误可能不为零。

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(body))
}

//Post 请求（资源提交，比如 图片上传）
//func FormDataTest()  {
//
//	resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
//	if err != nil {
//		// handle error
//		fmt.Println(err.Error())
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//
//	if err != nil {
//		// handle error
//		fmt.Println(err.Error())
//	}
//
//	fmt.Println(string(body))
//
//}


//Go post 上传文件
func PostFile(filename string, target_url string) (*http.Response, error) {
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	// use the body_writer to write the Part headers to the buffer
	_, err := body_writer.CreateFormFile("userfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return nil, err
	}

	// the file data will be the second part of the body
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return nil, err
	}
	// need to know the boundary to properly close the part myself.
	boundary := body_writer.Boundary()
	//close_string := fmt.Sprintf("\r\n--%s--\r\n", boundary)
	close_buf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))

	// use multi-reader to defer the reading of the file data until
	// writing to the socket buffer.
	request_reader := io.MultiReader(body_buf, fh, close_buf)
	fi, err := fh.Stat()
	if err != nil {
		fmt.Printf("Error Stating file: %s", filename)
		return nil, err
	}
	req, err := http.NewRequest("POST", target_url, request_reader)
	if err != nil {
		return nil, err
	}

	// Set headers for multipart, and Content Length
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.ContentLength = fi.Size() + int64(body_buf.Len()) + int64(close_buf.Len())

	return http.DefaultClient.Do(req)
}

//扩展 Post 表单提交（包括 Header 设置）
func PostTestAttend()  {

	//可以通过client中transport的Dail函数,在自定义Dail函数里面设置建立连接超时时长和发送接受数据超时
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*2)    //设置建立连接超时
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * 2))    //设置发送接受数据超时
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * 2,
		},
	}

	// 创建http.client
	//client := &http.Client{}

	postValue := url.Values{
		"agentNo": {"B000000000000238"},
		"type": {"01"},
		"page": {"1"},
		"rows": {"10"},
	}

	postString := postValue.Encode()

	req, err := http.NewRequest("POST",baseUrl + "/bankInsurance/activity/getActivityList", strings.NewReader(postString))
	//req, err := http.NewRequest("POST",baseUrl + "/bankInsurance/activity/getActivityList", strings.NewReader("agentNo=B000000000000238&type=01&page=1&rows=10"))		// 方式2，直接使用strings.NewReader("agentNo=B000000000000238&type=01&page=1&rows=10")

	if err != nil {
		// handle error
		fmt.Println(err.Error())
		return
	}

	// 表单方式(必须)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//AJAX 方式请求
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	// json方式
	//req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		// handle error
		fmt.Println(err.Error())
		return
	}

	defer resp.Body.Close()
	stdout := os.Stdout            			//将结果定位到标准输出，也可以直接打印出来，或定位到其他地方进行相应处理
	_,err = io.Copy(stdout,resp.Body)      //将第二个参数拷贝到第一个参数，直到第二参数到达EOF或发生错误，返回拷贝的字节喝遇到的第一个错误.
	status := resp.StatusCode        	   //获取返回状态码，正常是200
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(body))
	fmt.Println(status)
}




func main() {

	fmt.Println("=================GetTest=================")
	GetTest()

	fmt.Println("================PostTest=================")
	PostTest()
	PostTestAttend()

	fmt.Println("================Go post 上传文件=================")
	target_url := "http://localhost:8086/upload"
	filename := "/Users/wei/Downloads/21dian_1.9_10"
	fmt.Println(PostFile(filename, target_url))


}

