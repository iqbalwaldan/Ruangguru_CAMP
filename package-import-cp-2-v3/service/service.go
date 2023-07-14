package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"errors"
)

// Service is package for any logic needed in this program

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Pay(money int) (entity.PaymentInformation, error)
}

type Service struct {
	database database.DatabaseInterface
}

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

func (s *Service) AddCart(productName string, quantity int) error {
	// Mengambil Produk Berdasarkan Nama
	product, err := s.database.GetProductByName(productName)
	if err != nil {
		return err
	}

	// Mengambil semua cart items
	cartItems, err := s.database.GetCartItems()
	if err != nil {
		return err
	}
	// pengecekan jika qty kurang dari 0
	if quantity <= 0 {
		return errors.New("invalid quantity")
	}

	// menambahkan cart item baru
	cartItems = append(cartItems, entity.CartItem{
		ProductName: product.Name,
		Price:       product.Price,
		Quantity:    quantity,
	})

	// melakukan save cart items baru kedalam database
	err = s.database.SaveCartItems(cartItems)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *Service) RemoveCart(productName string) error {
	// melakukan pengecekan ketersediaan produk di database
	_, err := s.database.GetProductByName(productName)
	if err != nil {
		return err
	}
	// melakukan pengambilan semua cart items
	cartItems, err := s.database.GetCartItems()
	if err != nil {
		return err
	}

	newCart := []entity.CartItem{}
	isInCart := false
	// melakukan penghapusan data cart items
	for i, cart := range cartItems {
		if cart.ProductName != productName {
			newCart = append(newCart, cartItems[i])
			isInCart = true
		}
	}
	// pengecekan apakah produk terdaapat dalam cart
	if !isInCart {
		return errors.New("product not found")
	}
	// melakukan save cart items baru kedalam database
	err = s.database.SaveCartItems(newCart)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	carts, err := s.database.GetCartItems()
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (s *Service) ResetCart() error {
	// resetcart dengan melakukan save empty cart items kedalam database
	emptyCartItems := []entity.CartItem{}
	err := s.database.SaveCartItems(emptyCartItems)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {
	// mengembil semua data dari data produk
	return s.database.GetProductData(), nil
	// return nil, nil // TODO: replace this
}

func (s *Service) Pay(money int) (entity.PaymentInformation, error) {
	// mengambil isi dari cart items
	cartItems, err := s.ShowCart()
	if err != nil {
		return entity.PaymentInformation{}, nil
	}

	// menghitung total price
	totalPrice := 0
	for _, carItem := range cartItems {
		totalPrice += carItem.Price * carItem.Quantity
	}

	// menghitung kembalian
	change := money - totalPrice
	if change < 0 {
		return entity.PaymentInformation{}, errors.New("money is not enough")

	}

	// melakukan reset cart
	s.ResetCart()

	return entity.PaymentInformation{
		TotalPrice:  totalPrice,
		Change:      change,
		ProductList: cartItems,
		MoneyPaid:   money,
	}, nil // TODO: replace this
}
