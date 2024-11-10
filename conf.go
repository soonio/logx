package logx

type Conf struct {
	Dir    string `mapping:"dir"`    // 日志目录
	Output bool   `mapping:"output"` // 是否直接输出到控制台
}
