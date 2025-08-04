package handler

import (
	"net/http"
	"random_today/internal/application/service"
)

// SetupRoutes 设置路由
func SetupRoutes() http.Handler {
	// 创建应用服务实例（这里需要依赖注入，暂时创建空实例）
	appService := &service.PreferenceAppService{}
	
	// 创建处理器
	preferenceHandler := NewPreferenceHandler(appService)
	
	// 创建多路复用器
	mux := http.NewServeMux()
	
	// 设置路由
	mux.HandleFunc("/api/preferences", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			preferenceHandler.CreatePreference(w, r)
		case http.MethodGet:
			preferenceHandler.GetAllPreferences(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	
	mux.HandleFunc("/api/preferences/random", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		preferenceHandler.GetRandomPreference(w, r)
	})
	
	mux.HandleFunc("/api/preferences/type", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		preferenceHandler.GetPreferencesByType(w, r)
	})
	
	mux.HandleFunc("/api/preferences/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		preferenceHandler.DeletePreference(w, r)
	})
	
	return mux
} 