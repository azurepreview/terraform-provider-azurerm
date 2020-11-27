package client

import (
	"github.com/Azure/azure-sdk-for-go/services/mediaservices/mgmt/2020-05-01/media"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/common"
)

type Client struct {
	ServicesClient           *media.MediaservicesClient
	StreamingEndpointsClient *media.StreamingEndpointsClient
	AssetsClient   *media.AssetsClient
}

func NewClient(o *common.ClientOptions) *Client {
	ServicesClient := media.NewMediaservicesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&ServicesClient.Client, o.ResourceManagerAuthorizer)

	StreamingEndpointsClient := media.NewStreamingEndpointsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&StreamingEndpointsClient.Client, o.ResourceManagerAuthorizer)
  
  AssetsClient := media.NewAssetsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&AssetsClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		ServicesClient:           &ServicesClient,
		StreamingEndpointsClient: &StreamingEndpointsClient,
    AssetsClient:   &AssetsClient,
	}
}
