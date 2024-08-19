package postgresql

import (
	"fmt"
	"realty/internal/domain"
	"realty/pkg/storage/postgres"
)

//go:generate ifacemaker -f postgres.go -o ../repository.go -i Repository -s PostgresRepository -p realty -y "Controller describes methods, implemented by the repository package."
type PostgresRepository struct {
	db postgres.Postgres
}

func NewPostgresRepository(db postgres.Postgres) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (p *PostgresRepository) CreateHouse(house *domain.House) (*domain.House, error) {
	var createdHouse domain.House
	err := p.db.QueryRow(`
        INSERT INTO houses(address, year, developer, created_at)
        VALUES ($1, $2, $3, NOW())
        RETURNING id, address, year, developer, created_at, last_flat_added_at`,
		house.Address,
		house.Year,
		house.Developer,
	).Scan(&createdHouse.ID, &createdHouse.Address, &createdHouse.Year, &createdHouse.Developer, &createdHouse.CreatedAt, &createdHouse.LastFlatAddedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to create house: %w", err)
	}

	return &createdHouse, nil
}

func (p *PostgresRepository) GetHouseByID(id int64) (*domain.House, error) {
	var house domain.House
	err := p.db.QueryRow(`
		SELECT id, address, year, developer, created_at, last_flat_added_at
		FROM houses
		WHERE id = $1`, id,
	).Scan(&house.ID, &house.Address, &house.Year, &house.Developer, &house.CreatedAt, &house.LastFlatAddedAt)

	if err != nil {
		return nil, fmt.Errorf("querying house: %w", err)
	}

	return &house, nil
}

func (p *PostgresRepository) CreateFlat(flat *domain.Flat) (*domain.Flat, error) {
	var createdFlat domain.Flat
	err := p.db.QueryRow(`
        INSERT INTO flats(house_id, price, rooms, status)
        VALUES ($1, $2, $3, 'created')
        RETURNING id, house_id, price, rooms, status`,
		flat.HouseID,
		flat.Price,
		flat.Rooms,
	).Scan(&createdFlat.ID, &createdFlat.HouseID, &createdFlat.Price, &createdFlat.Rooms, &createdFlat.Status)

	if err != nil {
		return nil, fmt.Errorf("failed to create flat: %w", err)
	}

	return &createdFlat, nil
}

func (p *PostgresRepository) GetFlatByID(id int64) (*domain.Flat, error) {
	var flat domain.Flat
	err := p.db.QueryRow(`
		SELECT id, house_id, price, rooms, status
		FROM flats
		WHERE id = $1`, id,
	).Scan(&flat.ID, &flat.HouseID, &flat.Price, &flat.Rooms, &flat.Status)

	if err != nil {
		return nil, fmt.Errorf("querying flat: %w", err)
	}

	return &flat, nil
}

func (p *PostgresRepository) UpdateFlatStatus(flat *domain.Flat) (*domain.Flat, error) {
	var updatedFlat domain.Flat
	err := p.db.QueryRow(`
		UPDATE flats
		SET status = $1
		WHERE id = $2
		RETURNING id, house_id, price, rooms, status`,
		flat.Status,
		flat.ID,
	).Scan(&updatedFlat.ID, &updatedFlat.HouseID, &updatedFlat.Price, &updatedFlat.Rooms, &updatedFlat.Status)

	if err != nil {
		return nil, fmt.Errorf("failed to update flat status: %w", err)
	}

	return &updatedFlat, nil
}

func (p *PostgresRepository) UpdateHouseLastAdded(houseID int64) error {
	_, err := p.db.Exec(`
		UPDATE houses
		SET last_flat_added_at = NOW()
		WHERE id = $1`, houseID,
	)

	if err != nil {
		return fmt.Errorf("updating house last added date: %w", err)
	}

	return nil
}

func (p *PostgresRepository) GetFlatsByHouseID(houseID int64, userType string) ([]domain.Flat, error) {
	var flats []domain.Flat
	var query string

	if userType == "Moderator" {
		query = `
			SELECT id, house_id, price, rooms, status
			FROM flats
			WHERE house_id = $1`
	} else {
		query = `
			SELECT id, house_id, price, rooms, status
			FROM flats
			WHERE house_id = $1 AND status = 'approved'`
	}

	rows, err := p.db.Query(query, houseID)
	if err != nil {
		return nil, fmt.Errorf("failed to get flats by house ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var flat domain.Flat
		if err = rows.Scan(&flat.ID, &flat.HouseID, &flat.Price, &flat.Rooms, &flat.Status); err != nil {
			return nil, fmt.Errorf("failed to scan flat: %w", err)
		}
		flats = append(flats, flat)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return flats, nil
}

func (p *PostgresRepository) SubscribeToHouse(email string, houseID int64) error {
	_, err := p.db.Exec(`
        INSERT INTO subscriptions(email, house_id, created_at)
        VALUES ($1, $2, NOW())`,
		email, houseID)
	if err != nil {
		return fmt.Errorf("failed to subscribe to house: %w", err)
	}
	return nil
}

func (p *PostgresRepository) GetSubscribersByHouseID(houseID int64) ([]string, error) {
	var subscribers []string

	rows, err := p.db.Query(`
        SELECT email
        FROM subscriptions
        WHERE house_id = $1`, houseID)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscribers: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var email string
		if err = rows.Scan(&email); err != nil {
			return nil, fmt.Errorf("failed to scan email: %w", err)
		}
		subscribers = append(subscribers, email)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return subscribers, nil
}
