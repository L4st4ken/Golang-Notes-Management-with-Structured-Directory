package repositories

import(
	"database/sql"
	"errors"
	"notesmanagement/internal/models"
)

type NoteRepository struct{
	db *sql.DB
}

func NewNoteRepo(db *sql.DB) *NoteRepository{
	return &NoteRepository{db: db}
}

func (repo *NoteRepository) GetAll() ([]models.Note, error){
	query := "SELECT id, title, content, category_id FROM notes"
	rows, err := repo.db.Query(query)
	if err != nil{
		return  nil, err
	}

	defer rows.Close()

	notes := make([]models.Note, 0)
	for rows.Next(){
		var p models.Note
		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.CategoryID)
		if err != nil{
			return nil, err
		}
		notes = append(notes, p)
	}

	return notes, nil
}

func (repo *NoteRepository) Create(note *models.Note) error{
	query := "INSERT INTO notes (title, content,category_id) VALUES ($1, $2, $3) RETURNING id"
	err := repo.db.QueryRow(query, note.Title, note.Content, note.CategoryID).Scan(&note.ID)
	return err
}

//ambil ID Note
func (repo *NoteRepository) GetByID(id int) (*models.Note, error){
	query := "SELECT n.id, n.title, n.content, c.id, c.name FROM notes n JOIN categories c ON n.category_id = c.id WHERE n.id = $1"

	var p models.Note	
	err :=  repo.db.QueryRow(query, id).Scan(&p.ID, &p.Title, &p.Content, &p.CategoryID, &p.CategoryName)
	if err == sql.ErrNoRows{
		return nil, errors.New("note tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (repo *NoteRepository) Update(note *models.Note) error{
	query := "UPDATE notes SET title = $1, content = $2 WHERE id = $3"
	result, err := repo.db.Exec(query, note.Title, note.Content, note.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("note tidak ditemukan")
	}

	return nil
}

func (repo *NoteRepository) Delete(id int) error{
	query := "DELETE FROM notes WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("note tidak ditemukan")
	}

	return err
}