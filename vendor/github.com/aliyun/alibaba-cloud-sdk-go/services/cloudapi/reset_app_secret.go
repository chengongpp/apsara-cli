package cloudapi

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

// ResetAppSecret invokes the cloudapi.ResetAppSecret API synchronously
// api document: https://help.aliyun.com/api/cloudapi/resetappsecret.html
func (client *Client) ResetAppSecret(request *ResetAppSecretRequest) (response *ResetAppSecretResponse, err error) {
	response = CreateResetAppSecretResponse()
	err = client.DoAction(request, response)
	return
}

// ResetAppSecretWithChan invokes the cloudapi.ResetAppSecret API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/resetappsecret.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetAppSecretWithChan(request *ResetAppSecretRequest) (<-chan *ResetAppSecretResponse, <-chan error) {
	responseChan := make(chan *ResetAppSecretResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetAppSecret(request)
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

// ResetAppSecretWithCallback invokes the cloudapi.ResetAppSecret API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/resetappsecret.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetAppSecretWithCallback(request *ResetAppSecretRequest, callback func(response *ResetAppSecretResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetAppSecretResponse
		var err error
		defer close(result)
		response, err = client.ResetAppSecret(request)
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

// ResetAppSecretRequest is the request struct for api ResetAppSecret
type ResetAppSecretRequest struct {
	*requests.RpcRequest
	AppKey string `position:"Query" name:"AppKey"`
}

// ResetAppSecretResponse is the response struct for api ResetAppSecret
type ResetAppSecretResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateResetAppSecretRequest creates a request to invoke ResetAppSecret API
func CreateResetAppSecretRequest() (request *ResetAppSecretRequest) {
	request = &ResetAppSecretRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "ResetAppSecret", "apigateway", "openAPI")
	return
}

// CreateResetAppSecretResponse creates a response to parse from ResetAppSecret response
func CreateResetAppSecretResponse() (response *ResetAppSecretResponse) {
	response = &ResetAppSecretResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
