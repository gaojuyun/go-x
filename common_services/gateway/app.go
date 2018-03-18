package main

import (
	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common/discovery"
)

var (
	xargs *Args           = NewArgs()
	xlog  common.ILogger  = common.NewGLogger()
	xnode *discovery.Node = &discovery.Node{}
	xapp  *App            = NewApp()
)

type App struct {
	common.App
}

func NewApp() *App {
	this := &App{}
	this.Derived = this
	this.Args = xargs
	this.Logger = xlog
	this.Node = xnode
	return this
}

func (this *App) OnAppReady() {
}

func (this *App) OnAppShutDown() {
}