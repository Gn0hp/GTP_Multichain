package system_handlers

import (
	"eth_bsc_multichain/pkg/build_info"
	"eth_bsc_multichain/pkg/middlewares"
	"net/http"
)

type SystemHandler struct {
	buildInfo build_info.BuildInfo
}

func New(buildInfo build_info.BuildInfo) SystemHandler {
	return SystemHandler{
		buildInfo: buildInfo,
	}
}

func (s *SystemHandler) Liveness(w http.ResponseWriter, r *http.Request) {
	middlewares.WriteResponse(&w, http.StatusOK, s.buildInfo)
}
