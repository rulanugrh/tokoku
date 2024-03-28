package service

import (
	"context"

	"github.com/rulanugrh/order/internal/entity"
	"github.com/rulanugrh/order/internal/grpc/order"
	"github.com/rulanugrh/order/internal/repository"
	"github.com/rulanugrh/order/internal/util"
	"github.com/rulanugrh/order/pkg"
)

type OrderServiceServer struct {
	order.UnimplementedOrderServiceServer
	repository repository.OrderInterface
	product repository.ProductInterface
	xendit pkg.XenditInterface
}

func OrderService(repository repository.OrderInterface, product repository.ProductInterface, xendit pkg.XenditInterface) *OrderServiceServer {
	return &OrderServiceServer{repository: repository, product: product, xendit: xendit}
}

func (o *OrderServiceServer) Receipt(ctx context.Context, req *order.Request) (*order.ResponseProccess, error) {
	input := entity.Order{
		UserID: uint(req.Req.GetUserId()),
		ProductID: uint(req.Req.GetProductId()),
		Quantity: uint(req.Req.GetQuantity()),
		MethodPayment: req.Req.GetMethodPayment(),
		Address: req.Req.GetAddress(),
		RequestCurreny: req.Req.RequstCurrency,
		MobilePhone: req.Req.MobilePhone,
		ChannelCode: req.Req.ChannelCode,
	}

	data, find := o.product.FindID(uint(req.Req.ProductId))
	if find != nil {
		return util.NotFoundOrderCreate(find.Error()), find
	}

	result, err := o.repository.Create(input)
	if err != nil {
		return util.BadRequestOrderCreate(err.Error()), err
	}

	response := order.Data{
		Uuid: result.UUID,
		ProductName: data.Name,
		Price: int64(data.Price),
		MethodPayment: result.MethodPayment,
		Address: result.Address,
	}

	return util.SuccessOrderCreate("success create order", &response), nil
}

func (o *OrderServiceServer) Checkout(ctx context.Context, req *order.UUID) (*order.ResponseCheckout, error) {
	data, err := o.repository.Checkout(req.Uuid)
	if err != nil {
		return util.BadRequestOrderCheckout(err.Error()), err
	}

	product, err_product := o.product.FindID(data.ProductID)
	if err_product != nil {
		return util.BadRequestOrderCheckout(err_product.Error()), err
	}

	payment, err_payment := o.xendit.PaymentRequest(*data)
	if err_payment != nil {
		return util.BadRequestOrderCheckout(err_payment.Error()), err
	}

	response := order.CheckOut{
		ProductName: product.Name,
		Price: int64(product.Price),
		Quantity: int64(data.Quantity),
		Total: (int64(data.Quantity) * int64(product.Price)),
		LinkPayment: payment.GetCreated(),
	}

	return util.SuccessOrderCheckout("success checkout order", &response), nil

}