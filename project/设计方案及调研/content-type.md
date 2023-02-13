Content-Type
Content-Type（MediaType），即是Internet Media Type，互联网媒体类型，也叫做MIME类型。在互联网中有成百上千中不同的数据类型，HTTP在传输数据对象时会为他们打上称为MIME的数据格式标签，用于区分数据类型。最初MIME是用于电子邮件系统的，后来HTTP也采用了这一方案。

当浏览器在请求资源时，会通过http返回头中的content-type决定如何显示/处理将要加载的数据，如果这个类型浏览器能够支持阅览，浏览器就会直接展示该资源，比如png、jpeg、video等格式。在某些下载文件的场景中，服务端可能会返回文件流，并在返回头中带上Content-Type:application/octet-stream，告知浏览器这是一个字节流，浏览器处理字节流的默认方式就是下载。
Application/octet-stream是应用程序文件的默认值。意思是未知的应用程序文件，浏览器一般不会自动执行或询问执行。浏览器会像对待，设置了HTTP头Content-Disposition值为attachment的文件一样来对待这类文件，即浏览器会触发下载行为。说人话就是，浏览器并不认得这是什么类型，也不知道应该如何展示，只知道这是一种二进制文件，因此遇到content-type为application/octet-stream的文件时，浏览器会直接把它下载下来。这个类型一般会配合另一个响应头Content-Disposition,该响应头指示回复的内容该以何种形式展示，是以内联的形式（即网页或者网页的一部分），还是以附件的形式下载并保存到本地。



Content-Type的格式：
Content-Type：type/subtype ;parameter

type：主类型，任意的字符串，如text，如果是*号代表所有；
subtype：子类型，任意的字符串，如html，如果是*号代表所有，用“/”与主类型隔开；
parameter：可选参数，如charset，boundary等。
例如：
Content-Type: text/html;
Content-Type: application/json;charset:utf-8;

常见Content-Type
常见的Content-Type有数百个，下面例举了一些



1、form-data: 

就是http请求中的multipart/form-data,它会将[表单](https://so.csdn.net/so/search?q=表单&spm=1001.2101.3001.7020)的数据处理为一条消息，以标签为单元，用分隔符分开。既可以上传键值对，也可以上传文件。当上传的字段是文件时，会有Content-Type来说明文件类型；content-disposition，用来说明字段的一些信息；

由于有boundary隔离，既可以上传文件等二进制数据，也可以上传表单键值对，它采用了键值对的方式，所以可以上传多个文件，只是最后会转化为一条信息，在springmvc中可以使用MultipartHttpServletRequest接收通过api根据"name"获取不同的键值，也可以通过MulTipartFile数组接收多个文件。



一、application/x-www-form-urlencoded
最常见的POST提交数据的方式,原生Form表单,如果不设置enctype属性,默认为application/x-www-form-urlencoded方式提交数据。只能上传键值对，并且键值对都是间隔分开的。
首先,Content-Type都指定为application/x-www-form-urlencoded;其次,提交的表单数据会转换为键值对并按照key1=val&key2=val2的方式进行编码,key和val都进行了URL转码。大部分服务端语言都对这种方式有很好的支持。
另外,如利用AJAX提交数据时,也可使用这种方式。例如jQuery,Content-Type默认值都是"application/x-www-form-urlencoded;charset=utf-8"。

二、multipart/form-data
另一个常见的POST数据提交的方式,Form表单的enctype设置为multipart/form-data,它会将表单的数据处理为一条消息,以标签为单元,用分隔符(这就是boundary的作用)分开,类似我们上面Content-Type中的例子。
由于这种方式将数据有很多部分,它既可以上传键值对,也可以上传文件,甚至多个文件。当上传的字段是文件时,会有Content-Type来说明文件类型;Content-disposition,用来说明字段的一些信息。每部分都是以-boundary开始,紧接着是内容描述信息,然后是回车,最后是字段具体内容(字段、文本或二进制等)。如果传输的是文件,还要包含文件名和文件类型信息。消息主体最后以-boundary-标示结束。

三、application/json
Content-Type: application/json作为响应头比较常见。实际上,现在越来越多的人把它作为请求头,用来告诉服务端消息主体是序列化后的JSON字符串,其中一个好处就是JSON格式支持比键值对复杂得多的结构化数据。由于JSON规范的流行,除了低版本IE之外的各大浏览器都原生支持JSON.stringify,服务端语言也都有处理JSON的函数,使用起来没有困难。
Goodle的AngularJS种的Ajax功能,默认就是提交JSON字符串。

四、text/xml
答: XML的作用不言而喻,用于传输和存储数据,它非常适合万维网传输,提供统一的方法来描述和交换独立于应用程序或供应商的结构化数据,在JSON出现之前是业界一大标准(当然现在也是),相比JSON的优缺点大家有兴趣可以上网search。因此,在POST提交数据时,xml类型也是不可缺少的一种,虽然一般场景上使用JSON可能更轻巧、灵活。

五、binary(application/octet-stream)

答: 在谷歌浏览器的Postman工具中,还可以看到"binary"这一类型,指的就是一些二进制文件类型。如application/pdf,指定了特定二进制文件的MIME类型。就像对于text文件类型若没有特定的子类型(subtype),就使用text/plain。类似的,二进制文件没有特定或已知的subtype,即使用application/octet-stream,这是应用程序文件的默认值,一般很少直接使用。对于application/octet-stream,只能提交二进制,而且只能提交一个二进制,如果提交文件的话,只能提交一个文件,后台接收参数只能有一个,而且只能是流(或者字节数组)。很多web服务器使用默认的application/octet-stream来发送未知类型。出于一些安全原因,对于这些资源浏览器不允许设置一些自定义默认操作,导致用户必须存储到本地以使用。一般来说,设置正确的MIME类型很重要。只可以上传二进制数据，通常用来上传文件，由于没有键值，所以，一次只能上传一个文件。
