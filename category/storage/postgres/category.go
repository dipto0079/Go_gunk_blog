package postgres

import (
	"context"
	"blog/category/storage"
)

const insertCategory=`
	INSERT INTO categorys(
		title
	)VALUES(
		:title
	)RETURNING id;
`

func (s *Storage)Create_sto(ctx context.Context, t storage.Category) (int64, error){
	stmt,err :=s.db.PrepareNamed(insertCategory)
	if err!=nil{
		return 0,err
	}
	var id int64
	if err :=stmt.Get(&id,t);err != nil{
		return 0,err
	}
	return id,nil
}

func (s *Storage) Get_sto(ctx context.Context, id int64) (*storage.Category, error) {
	var c storage.Category

	if err :=s.db.Get(&c,"SELECT * FROM categorys WHERE id=$1",id);err != nil{
		return nil,err
	}
	return &c,nil
}

const updateCategory =`
	UPDATE categorys 
	SET
		title =:title
		
	WHERE
	id =:id
	RETURNING *;
`
func (s *Storage) Update(ctx context.Context,t storage.Category) (*storage.Category, error) {
	stmt,err :=s.db.PrepareNamed(updateCategory)
	if err!=nil{
		return nil,err
	}
	if err :=stmt.Get(&t,t);err != nil{
		return nil,err
	}
	return &t,nil
}

func (s *Storage) Delete(ctx context.Context,id int64) error {
	var b storage.Category
	if err := s.db.Get(&b,"DELETE FROM categorys WHERE id=$1 RETURNING * ",id); err != nil {
		return err;
	}
	return nil
}


func (s *Storage)get_all_Data(ctx context.Context) (storage.Category, error){

	var c storage.Category

	if err :=s.db.Select(&c,"SELECT * FROM categorys");err != nil{

		return c,err
	}
	return c,nil
}