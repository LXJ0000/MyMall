package service

import (
	"MyMall/pkg/e"
	"MyMall/repository/db/dao"
	"MyMall/repository/db/model"
	"MyMall/serializer"
	"context"
	"fmt"
	"strconv"
)

type AddressService struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

func (service *AddressService) CreateAddress(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	addr := &model.Address{
		UserId:  uId,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	fmt.Println(service)
	if err := addressDao.CreateAddress(addr); err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    "地址添加失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddress(addr),
		Error:  "",
	}
}
func (service *AddressService) GetAddress(ctx context.Context, uId uint, addrIdStr string) serializer.Response {
	code := e.Success
	addrId, err := strconv.Atoi(addrIdStr)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	addressDao := dao.NewAddressDao(ctx)
	addr, err := addressDao.GetAddressByAddressIdAndUserId(uint(addrId), uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddress(addr),
	}

}
func (service *AddressService) GetAddressList(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	addrList, err := addressDao.GetAddressListByUserId(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildAddressList(addrList), uint(len(addrList)))

}
func (service *AddressService) DeleteAddress(ctx context.Context, uId uint, addrIdStr string) serializer.Response {
	code := e.Success
	addrId, err := strconv.Atoi(addrIdStr)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	addressDao := dao.NewAddressDao(ctx)
	_, err = addressDao.GetAddressByAddressIdAndUserId(uint(addrId), uId)
	if err != nil {
		code = e.Error
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "用户不存在",
		}
	}

	if err = addressDao.DeleteAddressByAddressIdAndUserId(uint(addrId), uId); err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}

}
func (service *AddressService) UpdateAddress(ctx context.Context, uId uint, addrIdStr string) serializer.Response {
	code := e.Success
	addrId, err := strconv.Atoi(addrIdStr)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	addressDao := dao.NewAddressDao(ctx)
	addr, _ := addressDao.GetAddressByAddressIdAndUserId(uint(addrId), uId)
	if service.Name != "" {
		addr.Name = service.Name
	}
	if service.Address != "" {
		addr.Address = service.Address
	}
	if service.Phone != "" {
		addr.Phone = service.Phone
	}
	fmt.Println(addr)
	if err = addressDao.UpdateAddressByAddressIdAndUserId(uint(addrId), uId, addr); err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	fmt.Println(addr)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddress(addr),
	}
}
