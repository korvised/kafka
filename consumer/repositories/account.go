package repositories

import "gorm.io/gorm"

type BackAccount struct {
	ID            string
	AccountHolder string
	AccountType   int
	Balance       float64
}

type AccountRepository interface {
	Save(bankAccount *BackAccount) error
	Delete(id string) error
	FindAll() (bankAccounts []*BackAccount, err error)
	FindByID(id string) (bankAccount *BackAccount, err error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	_ = db.AutoMigrate(&BackAccount{})
	return &accountRepository{db}
}

func (r *accountRepository) Save(bankAccount *BackAccount) error {
	return r.db.Save(bankAccount).Error
}

func (r *accountRepository) Delete(id string) error {
	return r.db.Where("id=?", id).Delete(&BackAccount{}).Error
}

func (r *accountRepository) FindAll() (bankAccounts []*BackAccount, err error) {
	err = r.db.Find(&bankAccounts).Error
	return bankAccounts, err
}

func (r *accountRepository) FindByID(id string) (bankAccount *BackAccount, err error) {
	err = r.db.Where("id=?", id).First(&bankAccount).Error
	return bankAccount, err
}
