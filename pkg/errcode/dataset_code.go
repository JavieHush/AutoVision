package errcode

var (
	ErrorGetTrainListFail  = NewError(20020001, "获取训练集失败")
	ErrorGetPredListFail   = NewError(20020002, "获取预测集失败")
	ErrorDeleteDatasetFail = NewError(20020003, "删除数据集失败")
	ErrorCreateDatasetFail = NewError(20020004, "创建数据集失败")
	ErrorSavingFilesFail   = NewError(20020005, "本地保存数据集失败")
)
