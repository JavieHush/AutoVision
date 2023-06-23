package errcode

var (
	ErrorGetModelList = NewError(20010001, "获取模型列表失败")
	ErrorCreateModel  = NewError(20010002, "创建模型失败")
	ErrorDeleteModel  = NewError(20010003, "删除模型失败")
	ErrorGetModelByID = NewError(20010004, "通过模型ID获取模型失败")
)
