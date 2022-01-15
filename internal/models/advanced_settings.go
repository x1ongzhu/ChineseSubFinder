package models

import "gorm.io/gorm"

type AdvancedSettings struct {
	gorm.Model
	DebugMode                  bool     // 是否开启调试模式，这个是写入一个特殊的文件来开启日志的 Debug 输出
	SaveFullSeasonTmpSubtitles bool     // 保存整季的缓存字幕
	SubTypePriority            int      // 字幕下载的优先级，0 是自动，1 是 srt 优先，2 是 ass/ssa 优先
	SubNameFormatter           int      // 字幕命名格式(默认不填写或者超出范围，则为 emby 格式)，0，emby 支持的的格式（AAA.chinese(简英,subhd).ass or AAA.chinese(简英,xunlei).default.ass），1常规格式（兼容性更好，AAA.zh.ass or AAA.zh.default.ass）
	SaveMultiSub               bool     // 保存多个网站的 Top 1 字幕
	CustomVideoExts            []string // 自定义视频扩展名，是在原有基础上新增。
	FixTimeLine                bool     // 开启校正字幕时间轴，默认 false
}