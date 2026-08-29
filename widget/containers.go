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
	r.element.Layer = layer
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
	c.element.Layer = layer
	return c
}
