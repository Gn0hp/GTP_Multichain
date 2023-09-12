package system_handlers

import "logur.dev/logur"

type SystemHandler struct {
	logger logur.LoggerFacade
}

func New(logger logur.LoggerFacade) SystemHandler {
	return SystemHandler{
		logger: logger,
	}
}
