package table

type Service struct {
	Repository *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repository: repo}
}

func (srv *Service) GetAllTables() ([]Table, error) {
	tables, err := srv.Repository.GetAll()
	return tables, err
}

func (srv *Service) GetByID(id uint) (*Table, error) {
	table, err := srv.Repository.GetByID(id)
	return table, err
}

func (srv *Service) Create(table *Table) (*Table, error) {
	return srv.Repository.Create(table)
}

func (srv *Service) Update(table *Table) error {
	return srv.Repository.Update(table)
}
