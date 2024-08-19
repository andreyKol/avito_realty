package usecase

import (
	"fmt"
	"realty/internal/domain"
	"realty/internal/realty"
	"realty/pkg/sender"
)

//go:generate ifacemaker -f *.go -o ../usecase.go -i UseCase -s UseCase -p realty -y "Controller describes methods, implemented by the usecase package."
type UseCase struct {
	realtyRepo realty.Repository
}

func NewRealtyUseCase(realtyRepo realty.Repository) *UseCase {
	return &UseCase{realtyRepo: realtyRepo}
}

func (u *UseCase) CreateHouse(req *realty.CreateHouseRequest) (*domain.House, error) {
	houseInfo := domain.House{
		Address:   req.Address,
		Year:      req.Year,
		Developer: req.Developer,
	}

	createdHouse, err := u.realtyRepo.CreateHouse(&houseInfo)
	if err != nil {
		return nil, fmt.Errorf("creating house: %w", err)
	}

	return createdHouse, nil
}

func (u *UseCase) GetHouseByID(HouseID int64) (*domain.House, error) {
	house, err := u.realtyRepo.GetHouseByID(HouseID)
	if err != nil {
		return nil, fmt.Errorf("getting house by ID: %w", err)
	}
	return house, nil
}

func (u *UseCase) CreateFlat(req *realty.CreateFlatRequest) (*domain.Flat, error) {
	flatInfo := domain.Flat{
		HouseID: req.HouseID,
		Price:   req.Price,
		Rooms:   req.Rooms,
	}

	createdFlat, err := u.realtyRepo.CreateFlat(&flatInfo)
	if err != nil {
		return nil, fmt.Errorf("creating flat: %w", err)
	}

	err = u.realtyRepo.UpdateHouseLastAdded(createdFlat.HouseID)
	if err != nil {
		return nil, fmt.Errorf("updating house last added date: %w", err)
	}

	go func() {
		err = u.NotifySubscribers(createdFlat.HouseID)
		if err != nil {
			fmt.Errorf("failed to notify subscribers: %w", err)
		}
	}()

	return createdFlat, nil
}

func (u *UseCase) GetFlatByID(flatID int64) (*domain.Flat, error) {
	flat, err := u.realtyRepo.GetFlatByID(flatID)
	if err != nil {
		return nil, fmt.Errorf("getting flat by ID: %w", err)
	}
	return flat, nil
}

func (u *UseCase) UpdateFlatStatus(req *realty.UpdateFlatStatusRequest) (*domain.Flat, error) {
	flatInfo := domain.Flat{
		ID:     req.FlatID,
		Status: req.NewStatus,
	}

	updatedFlat, err := u.realtyRepo.UpdateFlatStatus(&flatInfo)
	if err != nil {
		return nil, fmt.Errorf("updating flat status: %w", err)
	}

	return updatedFlat, nil
}

func (u *UseCase) GetFlatsByHouseID(houseID int64, userType string) ([]domain.Flat, error) {
	flats, err := u.realtyRepo.GetFlatsByHouseID(houseID, userType)
	if err != nil {
		return nil, fmt.Errorf("getting flats by house ID: %w", err)
	}
	return flats, nil
}

func (u *UseCase) SubscribeToHouse(email string, houseID int64) error {
	return u.realtyRepo.SubscribeToHouse(email, houseID)
}

func (u *UseCase) NotifySubscribers(houseID int64) error {
	subscribers, err := u.realtyRepo.GetSubscribersByHouseID(houseID)
	if err != nil {
		return fmt.Errorf("getting subscribers: %w", err)
	}

	send := sender.New()

	for _, subscriber := range subscribers {
		go func(email string) {
			err = send.SendEmail(email, fmt.Sprintf("New flat available in house ID %d", houseID))
			if err != nil {
				fmt.Errorf("failed to send email to %s: %v", email, err)
			}
		}(subscriber)
	}

	return nil
}
