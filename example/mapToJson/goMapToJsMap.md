### Object转换为Map

后端Go的map[string]string类型 前端转换为json

[Global_Objects](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Object/entries)

### map
```js
                var js = JSON.parse(data);
                var map = new Map(Object.entries(js));
```

### 遍历
```js
                for (let [key,value] of map.entries()){
                    console.log(key+"=="+value);
                }
```