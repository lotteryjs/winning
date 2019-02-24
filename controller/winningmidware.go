package controller

func notFound(c *gin.Context) {
	t, err := template.ParseFiles(filepath.ToSlash(filepath.Join(model.Conf.StaticRoot, "console/dist/start/index.html")))
	if nil != err {
		logger.Errorf("load 404 page failed: " + err.Error())
		c.String(http.StatusNotFound, "load 404 page failed")

		return
	}

	c.Status(http.StatusNotFound)
	t.Execute(c.Writer, nil)
}
