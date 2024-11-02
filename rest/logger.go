package rest

import (
	"CPC_Party_Service_Cloud/rest/utils"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// 日志系统必须高保
// 全局唯一的日志对象(在main函数中初始化)
var LOGGER = &Logger{}

type Logger struct {
	// 加锁，保证原子性
	sync.RWMutex

	//继承log.Logger
	goLogger *log.Logger
	//日志记录所在的文件
	file *os.File
	//每天凌晨定时整理器
	maintainTimer *time.Timer
}

func (l *Logger) Init() {
	l.openFile()

	//日志自我备份，自我维护。每天第1秒触发
	nextTime := utils.FirstSecondOfDay(utils.Tomorrow())
	duration := nextTime.Sub(time.Now())

	l.Info("下一次日志维护时间 %s 距当前 %ds", utils.ConvertTimeToDateTimeString(nextTime), duration/time.Second)
	l.maintainTimer = time.AfterFunc(duration, func() {
		go utils.SafeMethod(l.maintain)
	})
}

// 将日志写入到今天的日期中(该方法内必须使用异步方法记录日志，否则会引发死锁)
func (l *Logger) maintain() {
	l.Info("每日维护日志")

	l.Lock()
	defer l.Unlock()

	// 首先 关闭文件
	l.closeFile()
	// 日志归类为昨天
	destPath := utils.GetLogPath() + "/CPC_PSC-" + utils.Yesterday().Local().Format("2006-01-01"+".log")
	// 直接重命名文件
	err := os.Rename(l.fileName(), destPath)
	if err != nil {
		l.Error("重命名文件出错！", err.Error())
	}
	// 再次打开文件
	l.openFile()
	// 准备好下一次维护日志的时间
	now := time.Now()
	nextTime := utils.FirstSecondOfDay(utils.Tomorrow())
	duration := nextTime.Sub(now)
	l.Info("下一次日志维护时间 %s", utils.ConvertTimeToDateTimeString(nextTime))
	l.maintainTimer = time.AfterFunc(duration, func() {
		go utils.SafeMethod(l.maintain)
	})
}
func (l *Logger) Destroy() {
	l.closeFile()
	if l.maintainTimer != nil {
		l.maintainTimer.Stop()
	}
}

// 获取日志文件名
func (l *Logger) fileName() string {
	return utils.GetLogPath() + "/CPC_PSC.log"
}

// 打开日志文件
func (l *Logger) openFile() {
	//日志输出到文件中 且打开后暂不关闭
	fmt.Printf("使用日志文件 %s\r\n", l.fileName())
	f, err := os.OpenFile(l.fileName(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("日志文件无法正常打开: " + err.Error())
	}
	l.goLogger = log.New(f, "", log.Ltime)
	l.file = f
}

// 关闭日志文件
func (l *Logger) closeFile() {
	if l.file != nil {
		err := l.file.Close()
		if err != nil {
			panic("关闭日志时出错: " + err.Error())
		}
	}
}

// 处理日志的统一方法
func (l *Logger) log(prefix string, format string, v ...interface{}) {
	//控制台打印日志
	var consoleFormat = fmt.Sprintf("%s%s %s\r\n", prefix, utils.ConvertTimeToTimeString(time.Now()), format)
	fmt.Printf(consoleFormat, v...)

	l.goLogger.SetPrefix(prefix)
	//每一行加上换行
	var fileFormat = fmt.Sprintf("%s\r\n", format)
	l.goLogger.Printf(fileFormat, v...)
}
func (l *Logger) Debug(format string, v ...interface{}) {
	l.log("[DEBUG]", format, v...)
}
func (l *Logger) Info(format string, v ...interface{}) {
	l.log("[INFO ]", format, v...)
}
func (l *Logger) Warn(format string, v ...interface{}) {
	l.log("[WARN ]", format, v...)
}
func (l *Logger) Error(format string, v ...interface{}) {
	l.log("[ERROR]", format, v...)
}
func (l *Logger) Panic(format string, v ...interface{}) {
	l.log("[PANIC]", format, v...)
	panic(fmt.Sprintf(format, v...))
}
