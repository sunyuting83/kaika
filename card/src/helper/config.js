const CROSUrl = 'https://seep.eu.org/'
const RootUrl = 'http://localhost:13002/api/'
const IndexUrl = `${CROSUrl}${RootUrl}`
const GlobalTitle = '开卡管理'
const ImagesRoot = 'https://pic7.58cdn.com.cn/nowater/webim/big/'
const images = [
  `${ImagesRoot}n_v26a171fb1a3394f2abcfce3e1d0e2b662.jpg`,
  `${ImagesRoot}n_v224cd0671e0f4483d95f395494dd3a891.jpg`
]
const Api = {
  'login':`${RootUrl}login`
}

export default {
  IndexUrl,
  RootUrl,
  CROSUrl,
  GlobalTitle,
  images,
  Api
}