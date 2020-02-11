package cloudwf

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// OemFlowrateAnalyse invokes the cloudwf.OemFlowrateAnalyse API synchronously
// api document: https://help.aliyun.com/api/cloudwf/oemflowrateanalyse.html
func (client *Client) OemFlowrateAnalyse(request *OemFlowrateAnalyseRequest) (response *OemFlowrateAnalyseResponse, err error) {
	response = CreateOemFlowrateAnalyseResponse()
	err = client.DoAction(request, response)
	return
}

// OemFlowrateAnalyseWithChan invokes the cloudwf.OemFlowrateAnalyse API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/oemflowrateanalyse.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OemFlowrateAnalyseWithChan(request *OemFlowrateAnalyseRequest) (<-chan *OemFlowrateAnalyseResponse, <-chan error) {
	responseChan := make(chan *OemFlowrateAnalyseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OemFlowrateAnalyse(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// OemFlowrateAnalyseWithCallback invokes the cloudwf.OemFlowrateAnalyse API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/oemflowrateanalyse.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OemFlowrateAnalyseWithCallback(request *OemFlowrateAnalyseRequest, callback func(response *OemFlowrateAnalyseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OemFlowrateAnalyseResponse
		var err error
		defer close(result)
		response, err = client.OemFlowrateAnalyse(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// OemFlowrateAnalyseRequest is the request struct for api OemFlowrateAnalyse
type OemFlowrateAnalyseRequest struct {
	*requests.RpcRequest
	Gsid requests.Integer `position:"Query" name:"Gsid"`
}

// OemFlowrateAnalyseResponse is the response struct for api OemFlowrateAnalyse
type OemFlowrateAnalyseResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateOemFlowrateAnalyseRequest creates a request to invoke OemFlowrateAnalyse API
func CreateOemFlowrateAnalyseRequest() (request *OemFlowrateAnalyseRequest) {
	request = &OemFlowrateAnalyseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "OemFlowrateAnalyse", "cloudwf", "openAPI")
	return
}

// CreateOemFlowrateAnalyseResponse creates a response to parse from OemFlowrateAnalyse response
func CreateOemFlowrateAnalyseResponse() (response *OemFlowrateAnalyseResponse) {
	response = &OemFlowrateAnalyseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}