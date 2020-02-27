package route

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type MyMux struct {
	handlers map[string][]*Handler //用于存储http.HandleFunc("/", SayhelloName) 格式的路由规则
}

type Handler struct {
	path string
	f    http.HandlerFunc
}

//进行路由分配
func (m *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//处理静态文件
	url := strings.Split(strings.TrimLeft(r.URL.Path, "/"), "/")

	//map[get:[0xc0000ae300 0xc0000ae380 0xc00000e060] post:[0xc00000e100]]
	for _, handler := range m.handlers[strings.ToLower(r.Method)] {
		if handler.path == "/"+strings.ToLower(url[0]) {
			//调用的是func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) { f(w, r) }
			handler.f.ServeHTTP(w, r)
			//意思就是调用了自己的函数本身,本身就是写的那个具体的实现
			return
		}
	}
	http.NotFound(w, r)
	return
}

//开启http服务
func (m *MyMux) Conn(port string) {
	s := &http.Server{
		Addr:    port,
		Handler: m,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("开启http服务错误!")
	}
}

func NewMyMux() *MyMux {
	return &MyMux{make(map[string][]*Handler)}
}

//添加路由
func (m *MyMux) AddRoute(mode string, path string, fun http.HandlerFunc) {
	m.add(mode, path, fun)
}

//添加路由
/**
  mode  Post|Get|Put|Delete
  path  前缀
  fun    方法
*/
func (m *MyMux) add(mode, path string, fun http.HandlerFunc) {
	h := &Handler{strings.ToLower(path), fun}

	//下面是存储路由的核心代码，这里的路由m.handlers存储的格式是Get|Post|Put|Delete:String:http.HandlerFunc
	//map[get:[0xc0000ae300 0xc0000ae380 0xc00000e060] post:[0xc00000e100]]
	m.handlers[strings.ToLower(mode)] = append(m.handlers[strings.ToLower(mode)], h)
	fmt.Println(m.handlers)
}
