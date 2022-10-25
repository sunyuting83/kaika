export default async (url = '', data = {}, type = 'GET', token = '') => {
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
    method: type,
    
  }

  if (type === 'POST' || type === 'PUT' || type === 'DELETE') {
    requestConfig.headers = {
      Accept: '*/*',
      'Content-Type': 'application/json;charset=UTF-8',
    }
    Object.defineProperty(requestConfig, 'body', {
      value: JSON.stringify(data)
    })
  }
  if (token !== null && token.length === 64) {
    requestConfig.headers = new Headers({
      Accept: '*/*',
      'Content-Type': 'application/json;charset=UTF-8',
    })
    requestConfig.headers.append('Authorization',`Bearer ${token}`)
  }
  return new Promise((resolve) => {
    fetch(url, requestConfig)
      .then(res => {
        if(res.ok) {
          return res.json()
        }else {
          resolve({
            status: 1,
            message: "访问出错"
          })
        }
      })
      .then(json => {
        resolve(json)
      })
      .catch((err) => {
        resolve({
          status: 1,
          message: err.message
        })
      })
  })
}