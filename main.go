package main

import (
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"strconv"
	"todolist/model"
)

var items []model.Item
func init()  {//初始化
	items = make([]model.Item,0)//动态切片
}


func main(){
	//fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	fmt.Println("server start ..")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {//为处理器配置响应目录 ,参数是处理函数（处理请求和返回信息的逻辑处理）
	//request ：用户请求的信息。post、get、url等这些信息
	fmt.Println("handle",items)
		t,err :=template.ParseFiles("resource/html/index.html")//创建模板对象，t，评估

		//Parse()方法是解析字符串的，且只解析New()出来的模板对象。如果想要解析文件中的内容，ParseFiles()、ParseGlob()。
		//ParseFiles方法解析filenames指定的文件里的模板定义并将解析结果与 t 关联。如果发生错误，会停止解析并返回nil，否则返回(t, nil)。至少要提供一个文件。

		if err != nil {
			fmt.Println("err:", err)
		}

		data := map[string]interface{}{
			"items":items,
		}

		t.Execute(writer,data)//!替换模板 将data 应用到 writer     ->data.xxx  items.xxx
	})

	http.HandleFunc("/delete", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("delete")
		request.ParseForm()	 //获取 html表单
		fmt.Println("   ",request.Form)
		m := request.Form
		value := m["index"]
		fmt.Println("value:",value)

		if len(request.Form["index"]) >= 0 {
			index := request.Form["index"][0]
			fmt.Println(index)
			t := reflect.TypeOf(index)
			fmt.Println(t.Name())
			index1,err := strconv.Atoi(index)
			if err != nil{

			}

			fmt.Println(index1)
			//items = append(items[:index1], items[index1+1:]...)//删除切片
			fmt.Println("items:",items)
			items[index1].IsDelete= true
			http.Redirect(writer,request, "/", http.StatusFound)

		}
		fmt.Println("del done")
		writer.Write([]byte("done"))

	})

//

	http.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		if len(request.Form["name"]) > 0 {
			name := request.Form["name"][0]//获取表单数据
			fmt.Println(name)
			var item model.Item
			item = model.Item{}
			item.Id = "xxx"
			item.Name = name
			item.IsDelete = false
			items = append(items,item)
			fmt.Println("items:",items)
		}
		http.Redirect(writer,request, "/", http.StatusFound)//【重定向】
	})


	http.HandleFunc("/cha", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()

		if len(request.Form["index"]) > 0 {
			text := request.Form["text"][0]//获取表单数据
			index := request.Form["index"][0]

			fmt.Println(index,text)
			t := reflect.TypeOf(index)
			fmt.Println(t)
			index1,err := strconv.Atoi(index)
			if err != nil{

			}
			items[index1].Name= text

			//var item model.Item
			//item = model.Item{}
			//item.Id = "xxx"
			//item.Name = name
			//item.IsDelete = false
			//items = append(items,item)
			//fmt.Println("items:",items)

		}

		http.Redirect(writer,request, "/", http.StatusFound)//【重定向】

	})

	err := http.ListenAndServe(":8080",nil)//监听端口
	if err != nil {
		fmt.Println("err",err)
	}

}