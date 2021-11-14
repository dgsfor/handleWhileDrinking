// pages/qywxlogin/qywxlogin.js
const app = getApp()
import { MSreApiRequest } from "../../request/index.js"
Page({
  qywxlogin(){
    var that = this;
    wx.qy.login({
      success: function(res) {
        that.loginApi(res.code)
      }
    })
  },
  async loginApi(code){
      var data = {
        "qy_login_code":code,
      }
      const res = await MSreApiRequest({url:"/oauth/qywx/login",data:data,method:"POST"})
      console.log(res)
      if(res.status === 200){
        wx.setStorageSync("user_token", res.data.data.user_token);
        wx.setStorageSync("user_id", res.data.data.user_id);
        wx.setStorageSync("user_name",res.data.data.data.user_name);
        wx.setStorageSync("user_memeber", res.data.data.data.member);
        wx.setStorageSync("user_dept", res.data.data.data.dept);
        wx.setStorageSync("user_email", res.data.data.data.email);
        wx.setStorageSync("user_avatar", res.data.data.data.avatar);
        wx.setStorageSync('isLogin', true)
        wx.setStorageSync('isAvatar', true)
        wx.showToast({
          title: "登录成功",
          icon: 'none'
        });
        wx.navigateBack({
          delta: 1
        })
      } else{
        wx.showToast({
          title: res.data.msg,
          icon: 'none'
        }) 
      }
  },
})