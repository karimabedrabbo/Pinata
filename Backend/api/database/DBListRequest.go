package database

func (e *DbEnv) DBListRequest(orderBy string, afterId int64, limit int64) {

	if afterId != 0 {
		e.SetTx(e.GetTx().Where("id > ?", afterId))
	}

	if limit != 0 {
		e.SetTx(e.GetTx().Limit(limit))
	}

	e.SetTx(e.GetTx().Order(orderBy))
}
