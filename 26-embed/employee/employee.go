package employee

type Result struct {
	Count int
}

func (r Result) Int() int {
	return r.Count
}

type Rows []struct{}

type Stmt interface {
	Close() error
	NumInput() int
	Exec(stmt string, arg ...string) (Result, error)
	Query(arg []string) (Rows, error)
}

func MaleCount(s Stmt) (int, error) {
	result, err := s.Exec("select count(*) from employee_tab where gender = ?")
	if err != nil {
		return 0, err
	}

	return result.Int(), nil
}
