package admin

type VisitorService struct {
}

type Visitor struct {
	id   uint32
	name string
}

func (vs *VisitorService) query() *Visitor {
	return nil
}
