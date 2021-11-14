// pages/account/account.js
const app = getApp()
import { MSreApiRequest } from "../../request/index.js"
Page({
  /**
   * 页面的初始数据
   */
  data: {
    isLogin: false,
    nickName: '',
    avatarUrl: ''
  },

  onShow: function (options) {
    const isLogin = wx.getStorageSync("isLogin");
    const nickName = wx.getStorageSync("nickName");
    const avatarUrl = wx.getStorageSync("avatarUrl");
    if (isLogin){
      this.setData({
        isLogin,
        nickName,
        avatarUrl
      })
    };
  },
   /**
   * 微信登录接口
   */
  wxLogin: function() {
    var isLogin = wx.getStorageSync('isLogin')
    var that = this;
    if (!isLogin) {
      wx.getUserProfile({
        lang: 'zh_CN',
        desc: 'login',
        success: (res) => {
          console.log(res)
          wx.setStorageSync('nickName', res.userInfo.nickName)
          wx.setStorageSync('avatarUrl',res.userInfo.avatarUrl)
          wx.login({
            success: function(res) {
              that.loginApi(res.code)
            }
          })
        }
      })
    } else {
      wx.showToast({
        title: '你已经登录了',
        icon: 'error'
      }) 
    }
    
  },
  /**
   * 微信退出登录接口
   */
  wxloginout: function(){
    this.logoutApi()
    this.setData({
      isLogin: false
    })
  },
  /**
   * 联系我
   */
  showQrcode() {
    wx.previewImage({
      urls: ['https://blog.itmonkey.icu/img/wechat/dgsfor-wechat.jpg'],
      current: 'https://blog.itmonkey.icu/img/wechat/dgsfor-wechat.jpg' // 当前显示图片的http链接      
    })
  },
  async loginApi(code){
    var data = {
      "wx_login_code":code,
    }
    const result = await MSreApiRequest({url:"/oauth/wx/login",data:data,method:"POST"})
    if(result.data.code == 200){
      wx.setStorageSync('3rd_session', result.data.data)
      wx.setStorageSync('isLogin', true)
      wx.showToast({
        title: "登录成功",
        icon: 'none'
      });
      wx.reLaunch({
        url: '../account/account',
      })
    } else{
      wx.showToast({
        title: res.data.msg,
        icon: 'error'
      }) 
    }
  },
  async logoutApi(){
    const res = await MSreApiRequest({url:"/oauth/wx/logout",method:"GET"})
    if(res.data.code == 200){
      wx.setStorageSync("isLogin", false);
      wx.setStorageSync('nickName', '')
      wx.setStorageSync('avatarUrl','')
      wx.setStorageSync('3rd_session', '')
      wx.showToast({
        title: "退出登录成功",
        icon: 'none'
      });
    } else{
      wx.showToast({
        title: res.data.msg,
        icon: 'error'
      }) 
    }
  },
})