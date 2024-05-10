package main

import (
	"context"
	activity "seckill_system/srv-activity/kitex_gen/activity"
	product "seckill_system/srv-activity/kitex_gen/product"
	"seckill_system/srv-activity/service"
)

// ActivityServiceImpl implements the last service interface defined in the IDL.
type ActivityServiceImpl struct{}

// GetActivities implements the ActivityServiceImpl interface.
func (s *ActivityServiceImpl) GetActivities(ctx context.Context, a *activity.Activity) (resp []*product.Product, err error) {
	// TODO: Your code here...
	resp, err = service.GetProductsService(a)
	return
}

// GetActivityInfo implements the ActivityServiceImpl interface.
func (s *ActivityServiceImpl) GetActivityInfo(ctx context.Context, a *activity.Activity) (resp *activity.Activity, err error) {
	// TODO: Your code here...
	resp, err = service.GetActivity(a)
	return
}
