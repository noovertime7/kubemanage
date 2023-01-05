package utils

import (
	"strconv"
	"time"
)

var (
	machineID     int64 //机器id
	sn            int64 //序列号
	lastTimeStamp int64 //记录上次的时间戳(毫秒级)
)

func init() {
	lastTimeStamp = time.Now().UnixNano() / 1e6
	// TODO 机器码硬编码
	SetMachineID(111)
}

func SetMachineID(mid int64) {
	machineID = mid << 12
}

func GetSnowflakeID() string {
	res := getSnowflakeID()
	return strconv.Itoa(int(res))
}

func getSnowflakeID() int64 {
	// 单位为毫秒
	curTimeStamp := time.Now().UnixNano() / 1e6
	if curTimeStamp == lastTimeStamp {
		sn++
		//序列号为12位， 2^12 = 4096个
		if sn > 4095 {
			//序列号超出，则重置序列号。这也意味着每毫秒最多能生成4096个id值
			time.Sleep(time.Millisecond)
			curTimeStamp = time.Now().UnixNano() / 1e6
			lastTimeStamp = curTimeStamp //顺便更新下上次的时间戳
			sn = 0
		}
		//与运算 对应位全为1时，则为1.否则为0
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		rightBinValue <<= 22

		//或运算 对应位全为0时，则为0。否则为1
		id := rightBinValue | machineID | sn
		return id
	} else if curTimeStamp > lastTimeStamp {
		sn = 0
		lastTimeStamp = curTimeStamp
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		rightBinValue <<= 22
		return rightBinValue | machineID | sn
	}
	return 0
}
