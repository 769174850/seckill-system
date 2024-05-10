package service

import (
	"context"
	"fmt"
	"log"
	"seckill_system/srv-activity/activity_idl"
	"seckill_system/srv-activity/activity_model"
	"seckill_system/srv-activity/dao"
	"seckill_system/srv-activity/kitex_gen/activity"
	"seckill_system/srv-activity/kitex_gen/product"
	_ "time"
)

func GetTheProducts(a activity_model.Activity) (resp []*product.Product, err error) {

	fmt.Println("GetTheProduct id:", a.ID)

	//服务调用
	resp, err = activity_idl.ActivityClient.GetActivities(context.Background(), &activity.Activity{
		Id:        int64(a.ID),
		Name:      &a.ActivityName,
		Desc:      &a.Description,
		StartTime: &a.StartAt,
		EndTime:   &a.EndAt,
	})

	fmt.Println("GetTheProducts:", resp)
	if len(resp) == 0 {
		log.Println("Get Product Rpc call error", err)
		return
	}
	if err != nil {
		log.Println("GetTheProducts error:", err)
		return
	}
	return
}

func GetProductsService(a *activity.Activity) (resp []*product.Product, err error) {
	a1, err := dao.GetProducts(a.Id)
	fmt.Println("a1:", a1)
	if err != nil {
		log.Println("Service GetProduct call error :", err)
		return
	}

	resp = make([]*product.Product, len(a1))
	for i, p := range a1 {
		resp[i] = &product.Product{
			Id:          int64(p.ID),
			Name:        p.ProductName,
			Description: &p.Description,
			Price:       &p.Price,
			Stock:       &p.Stock,
			Image:       &p.Image,
		}
	}
	fmt.Println("GetProductsService:", resp)

	return resp, nil
}

func GetActivity(a *activity.Activity) (info *activity.Activity, err error) {
	fmt.Println("getActivity id :", a.Id)
	Act, err := dao.SearchActivityByID(a.Id)
	fmt.Println("Act:", Act)
	if err != nil {
		log.Println("GetActivity err:", err)
		return
	}

	return &activity.Activity{
		Id:        int64(Act.ID),
		Name:      &Act.ActivityName,
		Desc:      &Act.Description,
		StartTime: &Act.StartAt,
		EndTime:   &Act.EndAt,
	}, nil
}

func GetActInfo(a activity_model.Activity) (resp *activity.Activity, err error) {
	fmt.Println("GetInfo id:", a.ID)

	// 调用rpc
	resp, err = activity_idl.ActivityClient.GetActivityInfo(context.Background(), &activity.Activity{
		Id:        int64(a.ID),
		Name:      &a.ActivityName,
		Desc:      &a.Description,
		StartTime: &a.StartAt,
		EndTime:   &a.EndAt,
	})
	if err != nil {
		log.Println("GetActivityInfo rpc call err:", err)
	}
	return
}
