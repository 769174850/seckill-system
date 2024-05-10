// Code generated by Kitex v0.9.1. DO NOT EDIT.

package activityservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	activity "seckill_system/srv-activity/kitex_gen/activity"
	product "seckill_system/srv-activity/kitex_gen/product"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetActivities(ctx context.Context, a *activity.Activity, callOptions ...callopt.Option) (r []*product.Product, err error)
	GetActivityInfo(ctx context.Context, a *activity.Activity, callOptions ...callopt.Option) (r *activity.Activity, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kActivityServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kActivityServiceClient struct {
	*kClient
}

func (p *kActivityServiceClient) GetActivities(ctx context.Context, a *activity.Activity, callOptions ...callopt.Option) (r []*product.Product, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetActivities(ctx, a)
}

func (p *kActivityServiceClient) GetActivityInfo(ctx context.Context, a *activity.Activity, callOptions ...callopt.Option) (r *activity.Activity, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetActivityInfo(ctx, a)
}