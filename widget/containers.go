package widget

type rowWidget struct {
	element *Element
}

func (rowWidget) Widget() {}

func Row(children ...Widget) *rowWidget {
	return &rowWidget{
		element: NewElement(kindRow, 0, children...),
	}
}

func (r *rowWidget) Align(a ...Alignment) *rowWidget {
	r.element.Align(a...)
	return r
}

func (r *rowWidget) Layer(layer int16) *rowWidget {
	r.element.Props.Layer = layer
	return r
}

func (r *rowWidget) Visible(v bool) *rowWidget {
	r.element.Visible(v)
	return r
}

func (r *rowWidget) Clip(v bool) *rowWidget {
	r.element.Clip(v)
	return r
}

func (r *rowWidget) Opacity(v float32) *rowWidget {
	r.element.Opacity(v)
	return r
}

func (r *rowWidget) Size(width, height int16) *rowWidget {
	r.element.Size(width, height)
	return r
}

type columnWidget struct {
	element *Element
}

func (columnWidget) Widget() {}

func Column(children ...Widget) *columnWidget {
	return &columnWidget{
		element: NewElement(kindColumn, 0, children...),
	}
}

func (c *columnWidget) Align(a ...Alignment) *columnWidget {
	c.element.Align(a...)
	return c
}

func (c *columnWidget) Layer(layer int16) *columnWidget {
	c.element.Props.Layer = layer
	return c
}

func (c *columnWidget) Visible(v bool) *columnWidget {
	c.element.Visible(v)
	return c
}

func (c *columnWidget) Clip(v bool) *columnWidget {
	c.element.Clip(v)
	return c
}

func (c *columnWidget) Opacity(v float32) *columnWidget {
	c.element.Opacity(v)
	return c
}

func (c *columnWidget) Size(width, height int16) *columnWidget {
	c.element.Size(width, height)
	return c
}
