export default async (url = '', data = {}, type = 'GET') => {
  type = type.toUpperCase()
  // 此处规定get请求的参数使用时放在data中，如同post请求
  if (type === 'GET') {
    let dataStr = ''
    Object.keys(data).forEach(key => {
      dataStr += key + '=' + data[key] + '&'
    })

    if (dataStr !== '') {
      dataStr = dataStr.substr(0, dataStr.lastIndexOf('&'))
      url = url + '?' + dataStr
    }
  }

  let requestConfig = {
    // fetch默认不会带cookie，需要添加配置项credentials允许携带cookie
    method: type,
    headers: {
      Accept: 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
      'Content-type':'text/html;charset=UTF-8'
    }
  }

  if (type === 'POST' || type === 'PUT' || type === 'DELETE') {
    Object.defineProperty(requestConfig, 'body', {
      value: JSON.stringify(data)
    })
  }
  return new Promise((resolve) => {
    fetch(url, requestConfig)
      .then(res => {
        if(res.ok) {
          return res.blob()
        }else {
          resolve({
            status: 1,
            message: "访问出错"
          })
        }
      })
      .then(blob => {
        var reader = new FileReader();
        reader.onload = function () {
          var text = reader.result;
          resolve(text)
        }
        reader.readAsText(blob, 'UTF-8')
      })
      .catch((err) => {
        resolve({
          status: 1,
          message: err.message
        })
      })
  })
}