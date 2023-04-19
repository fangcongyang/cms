package request

import "cms/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}