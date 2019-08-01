/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/4 16:14
* @Description: The file is for
***********************************************************************/

package logger

import (
	//"EChain/enode/conf"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type Logger struct {
	client *mongo.Client
	logCollection *mongo.Collection
	logChan chan *EndLog
	autoCommitChan chan *logBatch
}

var (
	once sync.Once
	E_logger *Logger
)

func InitLogger() (err error) {

	//建立mongoDb连接
	client, err := mongo.NewClient(options.Client().ApplyURI(config().mongodbUri))
	//fmt.Println(config().mongodbUri)
	ctx, _ := context.WithTimeout(
		context.Background(),
		time.Duration(config().mongodbConnectTimeout)*time.Millisecond)
	err = client.Connect(ctx)

	//选择db和collection
	once.Do(func() {
		E_logger = &Logger{
			client:client,
			logCollection:client.Database("enode").Collection("log"),
			logChan:make(chan *EndLog, 1000),
			autoCommitChan: make(chan *logBatch,1000),
		}
	})

	//启动一个mongoDB处理携程
	go E_logger.writeLoop()

	return
}

//日志存储协程
func (logger *Logger) writeLoop() {

	var (
		log *EndLog
		batch *logBatch  //一般批次
		commitTimer *time.Timer
		timeoutBatch *logBatch  //超时批次

		// 配置
		logCommitTimeout = config().logCommitTimeout
		logBatchSize = config().logBatchSize

	)

	for {
		select {
		case log = <-logger.logChan:
			if batch == nil {  //先对logBatch做初始化
				batch = &logBatch{}
				//让这个批次超时自动提交（有的时候可能日志数较少，batch村不满，那么用户久久看不到日志）
				commitTimer = time.AfterFunc(
					time.Duration(logCommitTimeout)*time.Millisecond,
					func(batch *logBatch) func() {
						return func() {
							logger.autoCommitChan <- batch
						}
					}(batch),  //这样是把当前时刻的logBatch当成参数传给回调函数，就立即得到使用了
				)
			}
			//把新的日志追加到数组中
			batch.logs = append(batch.logs, log)

			//如果批次满了，立即发送
			if len(batch.logs) >= logBatchSize {
				//发送日志
				logger.saveLogs(batch)
				//清空logBatch
				batch = nil
				//此时还需要取消定时器，考虑的情形是：在5秒内（超时时间）批次已满（达到100条）,那么需要停掉定时器
				//下一个循环又会重新创建这个定时器
				commitTimer.Stop()
				//这里还要注意一个问题：有可能100条批次满时刚好定时器超时（达到5秒），这种情况下就会导致，
				// 在这里把批次提交了一次，而在自动提交那里也会提交一次，这就会产生冲突
				// 因此，需要在超时提交的case里先做判断
			}
		case timeoutBatch = <- logger.autoCommitChan:
			//先判断超时批次是否仍是当前批次
			if timeoutBatch != batch {
				continue   //timeoutBatch != logBatch只有一种解释就是logBatch，
				// 在前边那个case里已经被清空了，说明已经提交过了。就直接跳过
			}
			//把超时批次写入到mongodb中
			logger.saveLogs(timeoutBatch)
			//把当前批次清空
			batch = nil
		}
	}

}

//批量写入日志
func (logger *Logger) saveLogs(batch *logBatch) {

	logger.logCollection.InsertMany(context.TODO(), batch.logs)
	fmt.Println("日志批次提交到mongoDB")
}

//发送日志的API,给其他地方调用，然后把日志添加到logChan。 logChan的日志会在writeLoop中提交到mongodb
func (logger *Logger) Append(endLog *EndLog) {
	select {
	case logger.logChan <- endLog:  //将日志写入logChan队列
	default:
		//logChan队列满了(可能是因为mongodb写入比较慢)
		// 那么就把日志丢弃（这个考量是因为日志不是特别重要的）。丢弃就啥都不干就行
	}
}

//MONGODB
//打开电脑D:/mongodb/bin，启动mongod（服务器）， 再启动mongo（客户端），use cron， db.log.find()查看日志