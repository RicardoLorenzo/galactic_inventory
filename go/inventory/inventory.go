package inventory

import (
	"fmt"

	"github.com/RicardoLorenzo/server-test/go/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Represent a spacecraft
// Ref: https://gorm.io/docs/composite_primary_key.html#content-inner
type Spacecraft struct {
	gorm.Model
	// Autoincrements by default
	Id       uint64 `json:"id,omitempty" gorm:"primaryKey"`
	Name     string `json:"name,omitempty" gorm:"index:idx_name,priority:1;unique_index:idx_u_name"`
	Class    string `json:"class,omitempty" gorm:"index:idx_name,priority:2;index:idx_class,priority:1"`
	Armament string `json:"armament,omitempty"`
	Crew     string `json:"crew,omitempty"`
	Image    string `json:"image,omitempty"`
	Value    string `json:"value,omitempty"`
	Status   string `json:"status,omitempty" gorm:"index:idx_name,priority:3;index:idx_class,priority:3"`
}

type Store struct {
	conf *config.Config
	db   *gorm.DB
}

func NewStore(c *config.Config) (*Store, error) {
	s := &Store{
		conf: c,
	}

	err := s.open()
	defer s.close()

	if err != nil {
		return nil, err
	}

	if s.db.HasTable(&Spacecraft{}) == false {
		s.db.CreateTable(&Spacecraft{})
	}

	return s, nil
}

func (s *Store) paginate(page, pageSize int) *gorm.DB {
	offset := (page - 1) * pageSize
	return s.db.Offset(offset).Limit(pageSize)
}

func (s *Store) List(page, pageSize int, spacecraft *Spacecraft) ([]Spacecraft, error) {
	s.open()
	defer s.close()

	spacecrafts := []Spacecraft{}

	result := s.paginate(page, pageSize).Where(spacecraft).Find(spacecrafts)

	return spacecrafts, result.Error
}

func (s *Store) Create(spacecraft *Spacecraft) error {
	s.open()
	defer s.close()

	result := s.db.Create(spacecraft)

	return result.Error
}

func (s *Store) Update(spacecraft *Spacecraft) error {
	s.open()
	defer s.close()

	// According to the documentation, the update
	// is transactional by default.
	// Ref: https://gorm.io/docs/transactions.html#Disable-Default-Transaction
	result := s.db.Update(spacecraft)

	return result.Error
}

func (s *Store) Delete(spacecraft *Spacecraft) error {
	s.open()
	defer s.close()

	result := s.db.Delete(spacecraft, spacecraft.ID)

	return result.Error
}

func (s *Store) open() error {
	var err error

	// user:password@tcp(host:port)/dbname
	mysqlCredentials := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		s.conf.DBUser,
		s.conf.DBPassword,
		s.conf.DBServer,
		s.conf.DBPort,
		s.conf.DBName,
	)

	s.db, err = gorm.Open("mysql", mysqlCredentials)

	return err
}

func (s *Store) close() error {
	return s.db.Close()
}
