package r_kvstore

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

// AddShardingNode invokes the r_kvstore.AddShardingNode API synchronously
func (client *Client) AddShardingNode(request *AddShardingNodeRequest) (response *AddShardingNodeResponse, err error) {
	response = CreateAddShardingNodeResponse()
	err = client.DoAction(request, response)
	return
}

// AddShardingNodeWithChan invokes the r_kvstore.AddShardingNode API asynchronously
func (client *Client) AddShardingNodeWithChan(request *AddShardingNodeRequest) (<-chan *AddShardingNodeResponse, <-chan error) {
	responseChan := make(chan *AddShardingNodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddShardingNode(request)
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

// AddShardingNodeWithCallback invokes the r_kvstore.AddShardingNode API asynchronously
func (client *Client) AddShardingNodeWithCallback(request *AddShardingNodeRequest, callback func(response *AddShardingNodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddShardingNodeResponse
		var err error
		defer close(result)
		response, err = client.AddShardingNode(request)
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

// AddShardingNodeRequest is the request struct for api AddShardingNode
type AddShardingNodeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CouponNo             string           `position:"Query" name:"CouponNo"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ShardCount           requests.Integer `position:"Query" name:"ShardCount"`
	BusinessInfo         string           `position:"Query" name:"BusinessInfo"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	ShardClass           string           `position:"Query" name:"ShardClass"`
}

// AddShardingNodeResponse is the response struct for api AddShardingNode
type AddShardingNodeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	NodeId    string `json:"NodeId" xml:"NodeId"`
	OrderId   int64  `json:"OrderId" xml:"OrderId"`
}

// CreateAddShardingNodeRequest creates a request to invoke AddShardingNode API
func CreateAddShardingNodeRequest() (request *AddShardingNodeRequest) {
	request = &AddShardingNodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "AddShardingNode", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddShardingNodeResponse creates a response to parse from AddShardingNode response
func CreateAddShardingNodeResponse() (response *AddShardingNodeResponse) {
	response = &AddShardingNodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}