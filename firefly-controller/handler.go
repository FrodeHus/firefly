package main

import (
	log "github.com/Sirupsen/logrus"
)

type Handler interface {
	Init() error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(objOld, objNew interface{})
}

type FireFlyAppHandler struct {
}

// Init handles any handler initialization
func (t *FireFlyAppHandler) Init() error { 
	log.Info("FireFlyAppHandler.Init")
	return nil
}

// ObjectCreated is called when an object is created
func (t *FireFlyAppHandler) ObjectCreated(obj interface{}) {
	log.Info("FireFlyAppHandler.ObjectCreated")
}

// ObjectDeleted is called when an object is deleted
func (t *FireFlyAppHandler) ObjectDeleted(obj interface{}) {
	log.Info("FireFlyAppHandler.ObjectDeleted")
}

// ObjectUpdated is called when an object is updated
func (t *FireFlyAppHandler) ObjectUpdated(objOld, objNew interface{}) {
	log.Info("FireFlyAppHandler.ObjectUpdated")
}
