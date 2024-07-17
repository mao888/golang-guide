本文介绍了使用validator库做参数校验的一些十分实用的使用技巧，包括翻译校验错误提示信息、自定义提示信息的字段名称、自定义校验方法等。

# validator库参数校验若干实用技巧

在web开发中一个不可避免的环节就是对请求参数进行校验，通常我们会在代码中定义与请求参数相对应的模型（结构体），借助模型绑定快捷地解析请求中的参数，例如 gin 框架中的`Bind`和`ShouldBind`系列方法。本文就以 gin 框架的请求参数校验为例，介绍一些`validator`库的实用技巧。

gin框架使用[github.com/go-playground/validator](https://github.com/go-playground/validator)进行参数校验，目前已经支持`github.com/go-playground/validator/v10`了，我们需要在定义结构体时使用 `binding` tag标识相关校验规则，可以查看[validator文档](https://godoc.org/github.com/go-playground/validator#hdr-Baked_In_Validators_and_Tags)查看支持的所有 tag。

### 基本示例

首先来看gin框架内置使用`validator`做参数校验的基本示例。

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUpParam struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

func main() {
	r := gin.Default()

	r.POST("/signup", func(c *gin.Context) {
		var u SignUpParam
		if err := c.ShouldBind(&u); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// 保存入库等业务逻辑代码...

		c.JSON(http.StatusOK, "success")
	})

	_ = r.Run(":8999")
}
```

我们使用curl发送一个POST请求测试下：

```bash
curl -H "Content-type: application/json" -X POST -d '{"name":"q1mi","age":18,"email":"123.com"}' http://127.0.0.1:8999/signup
```

输出结果：

```bash
{"msg":"Key: 'SignUpParam.Email' Error:Field validation for 'Email' failed on the 'email' tag\nKey: 'SignUpParam.Password' Error:Field validation for 'Password' failed on the 'required' tag\nKey: 'SignUpParam.RePassword' Error:Field validation for 'RePassword' failed on the 'required' tag"}
```

从最终的输出结果可以看到 `validator` 的检验生效了，但是错误提示的字段不是特别友好，我们可能需要将它翻译成中文。

### 翻译校验错误提示信息

`validator`库本身是支持国际化的，借助相应的语言包可以实现校验错误提示信息的自动翻译。下面的示例代码演示了如何将错误提示信息翻译成中文，翻译成其他语言的方法类似。

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 定义一个全局翻译器T
var trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

type SignUpParam struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

func main() {
	if err := InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}

	r := gin.Default()

	r.POST("/signup", func(c *gin.Context) {
		var u SignUpParam
		if err := c.ShouldBind(&u); err != nil {
			// 获取validator.ValidationErrors类型的errors
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				// 非validator.ValidationErrors类型错误直接返回
				c.JSON(http.StatusOK, gin.H{
					"msg": err.Error(),
				})
				return
			}
			// validator.ValidationErrors类型错误则进行翻译
			c.JSON(http.StatusOK, gin.H{
				"msg":errs.Translate(trans),
			})
			return
		}
		// 保存入库等具体业务逻辑代码...

		c.JSON(http.StatusOK, "success")
	})

	_ = r.Run(":8999")
}
```

同样的请求再来一次：

```bash
curl -H "Content-type: application/json" -X POST -d '{"name":"q1mi","age":18,"email":"123.com"}' http://127.0.0.1:8999/signup
```

这一次的输出结果如下：

```bash
{"msg":{"SignUpParam.Email":"Email必须是一个有效的邮箱","SignUpParam.Password":"Password为必填字段","SignUpParam.RePassword":"RePassword为必填字段"}}
```

### 自定义错误提示信息的字段名

上面的错误提示看起来是可以了，但是还是差点意思，首先是错误提示中的字段并不是请求中使用的字段，例如：`RePassword`是我们后端定义的结构体中的字段名，而请求中使用的是`re_password`字段。如何是错误提示中的字段使用自定义的名称，例如`json`tag指定的值呢？

只需要在初始化翻译器的时候像下面一样添加一个获取`json` tag的自定义方法即可。

```go
// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// ... liwenzhou.com ...
}
```

再尝试发请求，看一下效果：

```bash
{"msg":{"SignUpParam.email":"email必须是一个有效的邮箱","SignUpParam.password":"password为必填字段","SignUpParam.re_password":"re_password为必填字段"}}
```

可以看到现在错误提示信息中使用的就是我们结构体中`json`tag设置的名称了。

但是还是有点瑕疵，那就是最终的错误提示信息中心还是有我们后端定义的结构体名称——`SignUpParam`，这个名称其实是不需要随错误提示返回给前端的，前端并不需要这个值。我们需要想办法把它去掉。

这里参考https://github.com/go-playground/validator/issues/633#issuecomment-654382345提供的方法，定义一个去掉结构体名称前缀的自定义方法：

```go
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
```

我们在代码中使用上述函数将翻译后的`errors`做一下处理即可：

```go
if err := c.ShouldBind(&u); err != nil {
	// 获取validator.ValidationErrors类型的errors
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		// 非validator.ValidationErrors类型错误直接返回
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// validator.ValidationErrors类型错误则进行翻译
	// 并使用removeTopStruct函数去除字段名中的结构体名称标识
	c.JSON(http.StatusOK, gin.H{
		"msg": removeTopStruct(errs.Translate(trans)),
	})
	return
}
```

看一下最终的效果：

```bash
{"msg":{"email":"email必须是一个有效的邮箱","password":"password为必填字段","re_password":"re_password为必填字段"}}
```

这一次看起来就比较符合我们预期的标准了。

### 自定义结构体校验方法

上面的校验还是有点小问题，就是当涉及到一些复杂的校验规则，比如`re_password`字段需要与`password`字段的值相等这样的校验规则，我们的自定义错误提示字段名称方法就不能很好解决错误提示信息中的其他字段名称了。

```bash
curl -H "Content-type: application/json" -X POST -d '{"name":"q1mi","age":18,"email":"123.com","password":"123","re_password":"321"}' http://127.0.0.1:8999/signup
```

最后输出的错误提示信息如下：

```bash
{"msg":{"email":"email必须是一个有效的邮箱","re_password":"re_password必须等于Password"}}
```

可以看到`re_password`字段的提示信息中还是出现了`Password`这个结构体字段名称。这有点小小的遗憾，毕竟自定义字段名称的方法不能影响被当成param传入的值。

此时如果想要追求更好的提示效果，将上面的Password字段也改为和`json` tag一致的名称，就需要我们自定义结构体校验的方法。

例如，我们为`SignUpParam`自定义一个校验方法如下：

```go
// SignUpParamStructLevelValidation 自定义SignUpParam结构体校验函数
func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(SignUpParam)

	if su.Password != su.RePassword {
		// 输出错误提示信息，最后一个参数就是传递的param
		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
	}
}
```

然后在初始化校验器的函数中注册该自定义校验方法即可：

```go
func InitTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// ... liwenzhou.com ...
    
		// 为SignUpParam注册自定义校验方法
		v.RegisterStructValidation(SignUpParamStructLevelValidation, SignUpParam{})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// ... liwenzhou.com ...
}
```

最终再请求一次，看一下效果：

```bash
{"msg":{"email":"email必须是一个有效的邮箱","re_password":"re_password必须等于password"}}
```

这一次`re_password`字段的错误提示信息就符合我们预期了。

### 自定义字段校验方法

除了上面介绍到的自定义结构体校验方法，`validator`还支持为某个字段自定义校验方法，并使用`RegisterValidation()`注册到校验器实例中。

接下来我们来为`SignUpParam`添加一个需要使用自定义校验方法`checkDate`做参数校验的字段`Date`。

```go
type SignUpParam struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	// 需要使用自定义校验方法checkDate做参数校验的字段Date
	Date       string `json:"date" binding:"required,datetime=2006-01-02,checkDate"`
}
```

其中`datetime=2006-01-02`是内置的用于校验日期类参数是否满足指定格式要求的tag。 如果传入的`date`参数不满足`2006-01-02`这种格式就会提示如下错误：

```bash
{"msg":{"date":"date的格式必须是2006-01-02"}}
```

针对date字段除了内置的`datetime=2006-01-02`提供的格式要求外，假设我们还要求该字段的时间必须是一个未来的时间（晚于当前时间），像这样针对某个字段的特殊校验需求就需要我们使用自定义字段校验方法了。

首先我们要在需要执行自定义校验的字段后面添加自定义tag，这里使用的是`checkDate`，注意使用英文分号分隔开。

```go
// customFunc 自定义字段级别校验方法
func customFunc(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	if date.Before(time.Now()) {
		return false
	}
	return true
}
```

定义好了字段及其自定义校验方法后，就需要将它们联系起来并注册到我们的校验器实例中。

```go
// 在校验器注册自定义的校验方法
if err := v.RegisterValidation("checkDate", customFunc); err != nil {
	return err
}
```

这样，我们就可以对请求参数中`date`字段执行自定义的`checkDate`进行校验了。 我们发送如下请求测试一下：

```bash
curl -H "Content-type: application/json" -X POST -d '{"name":"q1mi","age":18,"email":"123@QQ飞车手游.com","password":"123", "re_password": "123", "date":"2020-01-02"}' http://127.0.0.1:8999/signup
```

此时得到的响应结果是：

```bash
{"msg":{"date":"Key: 'SignUpParam.date' Error:Field validation for 'date' failed on the 'checkDate' tag"}}
```

这…自定义字段级别的校验方法的错误提示信息很“简单粗暴”，和我们上面的中文提示风格有出入，必须想办法搞定它呀！

### 自定义翻译方法

我们现在需要为自定义字段校验方法提供一个自定义的翻译方法，从而实现该字段错误提示信息的自定义显示。

```go
// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}
```

定义好了相关翻译方法之后，我们在`InitTrans`函数中通过调用`RegisterTranslation()`方法来注册我们自定义的翻译方法。

```go
// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// ...liwenzhou.com...
	
		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		if err != nil {
			return err
		}
		// 注意！因为这里会使用到trans实例
		// 所以这一步注册要放到trans初始化的后面
		if err := v.RegisterTranslation(
			"checkDate",
			trans,
			registerTranslator("checkDate", "{0}必须要晚于当前日期"),
			translate,
		); err != nil {
			return err
		}
		return
	}
	return
}
```

这样再次尝试发送请求，就能得到想要的错误提示信息了。

```bash
{"msg":{"date":"date必须要晚于当前日期"}}
```

### 总结

由于本篇博客示例代码较多，我已经把文中示例代码上传到我的github仓库——https://github.com/Q1mi/validator_demo，大家可以查看完整的示例代码。

本文总结的gin框架中`validator`的使用技巧同样也适用于直接使用`validator`库，区别仅仅在于我们配置的是gin框架中的校验器还是由`validator.New()`创建的校验器。同时使用`validator`库确实能够在一定程度上减少我们的编码量，但是它不太可能完美解决我们所有需求，所以你需要找到两者之间的平衡点。

参考链接：

https://github.com/go-playground/validator/blob/master/_examples/simple/main.go

https://github.com/go-playground/validator/blob/master/_examples/translations/main.go

https://github.com/go-playground/validator/issues/567

https://github.com/go-playground/validator/issues/633

https://github.com/go-playground/validator/issues/551