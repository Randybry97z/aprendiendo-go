package CodigoFacilito

func (self *Curso) getTitle() string {
	return self.Title
}

func (self *Curso) ObtenerTitulo() string {
	return self.getTitle()
}
