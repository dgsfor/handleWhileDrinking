// 定义请求次数，防止一个页面请求三次，调用了三次wx.hideLoading()
let ajaxTimes = 0;
// 请求接口
export const MSreApiRequest=(params)=>{
  let header = {...params.header};
  if(params.method==="POST"){
    header["content-type"] = "application/json"; 
  }
  ajaxTimes++;
  // 显示加载中效果
  wx.showLoading({
    title: "加载中~",
    mask: true,
  });

  const baseUrl = "http://localhost:8081/api/hwd/v1";
  // 如果data为空，那么就自定义一个data，并添加token
  if(params.data === undefined){
    var initparams = {}
    params.data = initparams
  }
  header["3rd_session"] = wx.getStorageSync('3rd_session')
  return new Promise((resolve,reject)=>{
    wx.request({
      data:params.data,
      method:params.method,
      header:header,
      url: baseUrl + params.url,
      success:(result)=>{
        if (result.data.code == 401) {
          var rd_session = wx.getStorageSync('3rd_session')
          if (rd_session.length == 0) {
            wx.showToast({
              title: '请登录',
              icon:'none'
            })
          } else {
            wx.showToast({
              title: '登录失效，请重新登录',
              icon:'none'
            })
          }
        }
        if (result.data.code == 500) {
          wx.showToast({
            title: result.data.data,
            icon:'error'
          })
        }
        resolve(result);
      },
      fail:(err)=>{
        reject(err);
      },
      complete:()=>{
        //关闭加载中
        ajaxTimes--;
        if(ajaxTimes===0){
          wx.hideLoading();
        }
      }
    })
  })
}
