package ess

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

// DetachDBInstances invokes the ess.DetachDBInstances API synchronously
// api document: https://help.aliyun.com/api/ess/detachdbinstances.html
func (client *Client) DetachDBInstances(request *DetachDBInstancesRequest) (response *DetachDBInstancesResponse, err error) {
	response = CreateDetachDBInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DetachDBInstancesWithChan invokes the ess.DetachDBInstances API asynchronously
// api document: https://help.aliyun.com/api/ess/detachdbinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetachDBInstancesWithChan(request *DetachDBInstancesRequest) (<-chan *DetachDBInstancesResponse, <-chan error) {
	responseChan := make(chan *DetachDBInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachDBInstances(request)
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

// DetachDBInstancesWithCallback invokes the ess.DetachDBInstances API asynchronously
// api document: https://help.aliyun.com/api/ess/detachdbinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetachDBInstancesWithCallback(request *DetachDBInstancesRequest, callback func(response *DetachDBInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachDBInstancesResponse
		var err error
		defer close(result)
		response, err = client.DetachDBInstances(request)
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

// DetachDBInstancesRequest is the request struct for api DetachDBInstances
type DetachDBInstancesRequest struct {
	*requests.RpcRequest
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBInstance           *[]string        `position:"Query" name:"DBInstance"  type:"Repeated"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ForceDetach          requests.Boolean `position:"Query" name:"ForceDetach"`
}

// DetachDBInstancesResponse is the response struct for api DetachDBInstances
type DetachDBInstancesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDetachDBInstancesRequest creates a request to invoke DetachDBInstances API
func CreateDetachDBInstancesRequest() (request *DetachDBInstancesRequest) {
	request = &DetachDBInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DetachDBInstances", "ess", "openAPI")
	return
}

// CreateDetachDBInstancesResponse creates a response to parse from DetachDBInstances response
func CreateDetachDBInstancesResponse() (response *DetachDBInstancesResponse) {
	response = &DetachDBInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}