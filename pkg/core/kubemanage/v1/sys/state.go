package sys

import (
	"runtime"
	"time"

	"github.com/noovertime7/kubemanage/dto"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func InitServer() (*dto.Server, error) {
	server := &dto.Server{
		Os:   dto.Os{},
		Cpu:  dto.Cpu{},
		Ram:  dto.Ram{},
		Disk: dto.Disk{},
	}
	var err error
	server.Os = InitOS()
	if server.Cpu, err = InitCPU(); err != nil {
		return nil, err
	}
	if server.Ram, err = InitRAM(); err != nil {
		return nil, err
	}
	if server.Disk, err = InitDisk(); err != nil {
		return nil, err
	}
	return server, nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: InitCPU
//@description: OS信息
//@return: o Os, err error

func InitOS() (o dto.Os) {
	o.GOOS = runtime.GOOS
	o.NumCPU = runtime.NumCPU()
	o.Compiler = runtime.Compiler
	o.GoVersion = runtime.Version()
	o.NumGoroutine = runtime.NumGoroutine()
	return o
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: InitCPU
//@description: CPU信息
//@return: c Cpu, err error

func InitCPU() (c dto.Cpu, err error) {
	if cores, err := cpu.Counts(false); err != nil {
		return c, err
	} else {
		c.Cores = cores
	}
	if cpus, err := cpu.Percent(time.Duration(200)*time.Millisecond, true); err != nil {
		return c, err
	} else {
		c.Cpus = cpus
	}
	return c, nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: InitRAM
//@description: RAM信息
//@return: r Ram, err error

func InitRAM() (r dto.Ram, err error) {
	if u, err := mem.VirtualMemory(); err != nil {
		return r, err
	} else {
		r.UsedMB = int(u.Used) / MB
		r.TotalMB = int(u.Total) / MB
		r.UsedPercent = int(u.UsedPercent)
	}
	return r, nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: InitDisk
//@description: 硬盘信息
//@return: d Disk, err error

func InitDisk() (d dto.Disk, err error) {
	if u, err := disk.Usage("/"); err != nil {
		return d, err
	} else {
		d.UsedMB = int(u.Used) / MB
		d.UsedGB = int(u.Used) / GB
		d.TotalMB = int(u.Total) / MB
		d.TotalGB = int(u.Total) / GB
		d.UsedPercent = int(u.UsedPercent)
	}
	return d, nil
}
