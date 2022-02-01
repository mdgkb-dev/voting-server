package users

import (
	"fmt"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) getAll() (models.Users, error) {
	items := make(models.Users, 0)
	err := r.db.NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.User, error) {
	item := models.User{}
	err := r.db.NewSelect().
		Model(&item).
		Relation("Human").
		//Relation("Questions").
		Relation("DonorRulesUsers.DonorRule.Image").
		Relation("DonorRulesUsers.DonorRule.DonorRulesUsers").
		Relation("DoctorsUsers.Doctor").
		Relation("Children.Human").
		Where("users.id = ?", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) getByEmail(id string) (*models.User, error) {
	item := models.User{}
	err := r.db.NewSelect().Model(&item).
		//Relation("Human").
		//Relation("Questions").
		//Relation("DonorRulesUsers.DonorRule.Image").
		//Relation("DonorRulesUsers.DonorRule.DonorRulesUsers").
		//Relation("Children.Human").
		//Relation("DoctorsUsers").
		Where("users.email = ?", id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(user *models.User) (err error) {
	_, err = r.db.NewInsert().Model(user).Exec(r.ctx)
	return err
}

func (r *Repository) emailExists(email string) (bool, error) {
	exists, err := r.db.NewSelect().Model((*models.User)(nil)).Where("users.email = ?", email).Exists(r.ctx)
	return exists, err
}

func (r *Repository) update(item *models.User) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.User) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) addToUser(values map[string]interface{}, table string) error {
	_, err := r.db.NewInsert().Model(&values).TableExpr(table).Exec(r.ctx)
	return err
}

func (r *Repository) removeFromUser(values map[string]interface{}, table string) error {
	q := r.db.NewDelete().Table(table)
	for key, value := range values {
		q = q.Where(fmt.Sprintf("%s = ?", key), value)
	}
	_, err := q.Exec(r.ctx)
	return err
}
