package web

import "context"

type TemplateEngine interface {
	//渲染页面
	Render(ctx context.Context, tplname string, data any) ([]byte, error)
}

//
//func (c *Context) Render(ctx context.Context, tplname string, data any) ([]byte, error) {
//	var err error
//	c.RespData, err = c.tplEngine.Render(c.Req.Context(), tplname, data)
//	return '', err
//}
