package backend

import (
	"fmt"
	"math"

	m "CurrencyAlertApi/Model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DB variables declared
var DB *gorm.DB
var err error

//DNS points to local MySQL database connection
const DNS = "root:pass@tcp(127.0.0.1:3306)/kryptodb?charset=utf8mb4&parseTime=True&loc=Local"

//InitialMigration function to establish connection
func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}
	DB.AutoMigrate(&m.Currencies{})
	DB.AutoMigrate(&m.Alerts{})
}

//GetAllUserAlerts function to get alerts from database
func GetAllUserAlerts(email string) []m.Alerts {
	var alerts []m.Alerts
	DB.Find(&alerts, "email = ?", email)
	return alerts
}

//GetTriggeredUserAlerts function to get triggered alerts from database
func GetTriggeredUserAlerts(email string) []m.Alerts {
	var alerts []m.Alerts
	DB.Where("email = ?", email).Where("triggered = ?", "true").Find(&alerts)
	return alerts
}

//GetAllOngoingAlerts function to get Alert from database
func GetAllOngoingAlerts() []m.Alerts {
	var alerts []m.Alerts
	DB.Find(&alerts, "triggered = ?", "false")
	return alerts
}

//CreateAlert function to save Alert to database
func CreateAlert(email string, symbol string, target float64) m.Alerts {
	var alert m.Alerts
	alert.Email = email
	alert.Currency = symbol
	alert.Target = target
	alert.Triggered = "false"
	DB.Create(&alert)

	return alert
}

//DeleteAlert function to save Alert to database
func DeleteAlert(id int) {

	DB.Delete(&m.Alerts{}, id)
}

//TriggerAlert function to change trigger status of an alert
func TriggerAlert(id int) {
	var alerts m.Alerts
	DB.Find(&alerts, id)
	alerts.Triggered = "true"
	DB.Save(&alerts)
}

//UpdateCurrencies function to update currencies
func UpdateCurrencies(symbol string, name string, currPrice float64) {
	var currency m.Currencies

	currency.Symbol = symbol
	currency.Name = name
	currency.CurrentPrice = currPrice
	if len(GetCurrency(symbol).Symbol) != 0 {
		DB.Save(&currency)
	} else {
		DB.Create(&currency)
	}
}

//GetCurrency function gets details of the given currency
func GetCurrency(symbol string) m.Currencies {
	var currency m.Currencies
	DB.Find(&currency, "symbol = ?", symbol)
	return currency

}

//PaginatedAlerts function gets the alert details in a paged format
func PaginatedAlerts(limit int, sort string, page int, email string, triggered string) ([]m.Alerts, int, int, int) {
	var alerts []m.Alerts

	sql := "SELECT * FROM alerts"

	if s := triggered; s != "" {
		sql = fmt.Sprintf("%s WHERE triggered LIKE '%%%s%%' AND email like '%%%s%%' ", sql, s, email)
	}

	PerPage := 9
	if limit > 0 {
		PerPage = limit
	}
	if order := sort; sort != "" {
		sql = fmt.Sprintf("%s ORDER BY id %s", sql, order)
	}

	Page := 1
	if page > 1 {
		Page = page
	}
	var total int64
	DB.Raw(sql).Count(&total)
	sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, PerPage, (Page-1)*PerPage)
	DB.Raw(sql).Scan(&alerts)

	return alerts, int(total), Page, int(math.Ceil(float64(total / int64(PerPage))))
}
