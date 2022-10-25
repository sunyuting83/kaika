const CROSUrl = 'https://crossorigin.me/'
const RootUrl = 'http://localhost:13002/api/'
const IndexUrl = `${CROSUrl}${RootUrl}`
const GlobalTitle = '开卡管理'
const ImagesRoot = 'https://pic7.58cdn.com.cn/nowater/webim/'
const images = [
  `${ImagesRoot}big/n_v26a171fb1a3394f2abcfce3e1d0e2b662.jpg`,
  `${ImagesRoot}big/n_v224cd0671e0f4483d95f395494dd3a891.jpg`,
  `${ImagesRoot}small/n_v224cd0671e0f4483d95f395494dd3a891.jpg`,
  `https://image.suning.cn/uimg/ZR/share_order/158501870837440052.jpg`
]
const Api = {
  'login': `${RootUrl}loginadmin`,
  'checklogin': `${RootUrl}checklogin`,
  'repassword': `${RootUrl}repassword`,
  'creatdcard': `${RootUrl}creatdcard`,
  'jt': `https://api.uomg.com/api/rand.qinghua?format=json`,
  'ip': 'https://api.uomg.com/api/visitor.info?skey=774740085',
  'city': 'https://api.map.baidu.com/location/ip?ak=OGNLmlzGl46KE7HU0hblDk2zXPPv0w5v&ip=121.30.139.214',
}
const makeIP = (IP) =>{
  return `https://api.seniverse.com/v3/weather/now.json?key=faua8rceea8uoemu&location=${IP}&language=zh-Hans&unit=c`
}

export default {
  IndexUrl,
  RootUrl,
  CROSUrl,
  GlobalTitle,
  images,
  Api,
  makeIP
}