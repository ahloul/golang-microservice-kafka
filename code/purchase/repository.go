package purchase

import "github.com/ahloul/loyalty-reports/models"

type Repository interface {
	Add(purchase models.Purchase) string
}
